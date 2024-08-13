// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /snap/go/10455/src/strings/replace.go:5
package strings

//line /snap/go/10455/src/strings/replace.go:5
import (
//line /snap/go/10455/src/strings/replace.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /snap/go/10455/src/strings/replace.go:5
)
//line /snap/go/10455/src/strings/replace.go:5
import (
//line /snap/go/10455/src/strings/replace.go:5
	_atomic_ "sync/atomic"
//line /snap/go/10455/src/strings/replace.go:5
)

import (
	"io"
	"sync"
)

// Replacer replaces a list of strings with replacements.
//line /snap/go/10455/src/strings/replace.go:12
// It is safe for concurrent use by multiple goroutines.
//line /snap/go/10455/src/strings/replace.go:14
type Replacer struct {
	once	sync.Once	// guards buildOnce method
	r	replacer
	oldnew	[]string
}

// replacer is the interface that a replacement algorithm needs to implement.
type replacer interface {
	Replace(s string) string
	WriteString(w io.Writer, s string) (n int, err error)
}

// NewReplacer returns a new Replacer from a list of old, new string
//line /snap/go/10455/src/strings/replace.go:26
// pairs. Replacements are performed in the order they appear in the
//line /snap/go/10455/src/strings/replace.go:26
// target string, without overlapping matches. The old string
//line /snap/go/10455/src/strings/replace.go:26
// comparisons are done in argument order.
//line /snap/go/10455/src/strings/replace.go:26
//
//line /snap/go/10455/src/strings/replace.go:26
// NewReplacer panics if given an odd number of arguments.
//line /snap/go/10455/src/strings/replace.go:32
func NewReplacer(oldnew ...string) *Replacer {
//line /snap/go/10455/src/strings/replace.go:32
	_go_fuzz_dep_.CoverTab[972]++
							if len(oldnew)%2 == 1 {
//line /snap/go/10455/src/strings/replace.go:33
		_go_fuzz_dep_.CoverTab[524977]++
//line /snap/go/10455/src/strings/replace.go:33
		_go_fuzz_dep_.CoverTab[974]++
								panic("strings.NewReplacer: odd argument count")
//line /snap/go/10455/src/strings/replace.go:34
		// _ = "end of CoverTab[974]"
	} else {
//line /snap/go/10455/src/strings/replace.go:35
		_go_fuzz_dep_.CoverTab[524978]++
//line /snap/go/10455/src/strings/replace.go:35
		_go_fuzz_dep_.CoverTab[975]++
//line /snap/go/10455/src/strings/replace.go:35
		// _ = "end of CoverTab[975]"
//line /snap/go/10455/src/strings/replace.go:35
	}
//line /snap/go/10455/src/strings/replace.go:35
	// _ = "end of CoverTab[972]"
//line /snap/go/10455/src/strings/replace.go:35
	_go_fuzz_dep_.CoverTab[973]++
							return &Replacer{oldnew: append([]string(nil), oldnew...)}
//line /snap/go/10455/src/strings/replace.go:36
	// _ = "end of CoverTab[973]"
}

func (r *Replacer) buildOnce() {
//line /snap/go/10455/src/strings/replace.go:39
	_go_fuzz_dep_.CoverTab[976]++
							r.r = r.build()
							r.oldnew = nil
//line /snap/go/10455/src/strings/replace.go:41
	// _ = "end of CoverTab[976]"
}

func (b *Replacer) build() replacer {
//line /snap/go/10455/src/strings/replace.go:44
	_go_fuzz_dep_.CoverTab[977]++
							oldnew := b.oldnew
							if len(oldnew) == 2 && func() bool {
//line /snap/go/10455/src/strings/replace.go:46
		_go_fuzz_dep_.CoverTab[982]++
//line /snap/go/10455/src/strings/replace.go:46
		return len(oldnew[0]) > 1
//line /snap/go/10455/src/strings/replace.go:46
		// _ = "end of CoverTab[982]"
//line /snap/go/10455/src/strings/replace.go:46
	}() {
//line /snap/go/10455/src/strings/replace.go:46
		_go_fuzz_dep_.CoverTab[524979]++
//line /snap/go/10455/src/strings/replace.go:46
		_go_fuzz_dep_.CoverTab[983]++
								return makeSingleStringReplacer(oldnew[0], oldnew[1])
//line /snap/go/10455/src/strings/replace.go:47
		// _ = "end of CoverTab[983]"
	} else {
//line /snap/go/10455/src/strings/replace.go:48
		_go_fuzz_dep_.CoverTab[524980]++
//line /snap/go/10455/src/strings/replace.go:48
		_go_fuzz_dep_.CoverTab[984]++
//line /snap/go/10455/src/strings/replace.go:48
		// _ = "end of CoverTab[984]"
//line /snap/go/10455/src/strings/replace.go:48
	}
//line /snap/go/10455/src/strings/replace.go:48
	// _ = "end of CoverTab[977]"
//line /snap/go/10455/src/strings/replace.go:48
	_go_fuzz_dep_.CoverTab[978]++

							allNewBytes := true
//line /snap/go/10455/src/strings/replace.go:50
	_go_fuzz_dep_.CoverTab[786483] = 0
							for i := 0; i < len(oldnew); i += 2 {
//line /snap/go/10455/src/strings/replace.go:51
		if _go_fuzz_dep_.CoverTab[786483] == 0 {
//line /snap/go/10455/src/strings/replace.go:51
			_go_fuzz_dep_.CoverTab[525081]++
//line /snap/go/10455/src/strings/replace.go:51
		} else {
//line /snap/go/10455/src/strings/replace.go:51
			_go_fuzz_dep_.CoverTab[525082]++
//line /snap/go/10455/src/strings/replace.go:51
		}
//line /snap/go/10455/src/strings/replace.go:51
		_go_fuzz_dep_.CoverTab[786483] = 1
//line /snap/go/10455/src/strings/replace.go:51
		_go_fuzz_dep_.CoverTab[985]++
								if len(oldnew[i]) != 1 {
//line /snap/go/10455/src/strings/replace.go:52
			_go_fuzz_dep_.CoverTab[524981]++
//line /snap/go/10455/src/strings/replace.go:52
			_go_fuzz_dep_.CoverTab[987]++
									return makeGenericReplacer(oldnew)
//line /snap/go/10455/src/strings/replace.go:53
			// _ = "end of CoverTab[987]"
		} else {
//line /snap/go/10455/src/strings/replace.go:54
			_go_fuzz_dep_.CoverTab[524982]++
//line /snap/go/10455/src/strings/replace.go:54
			_go_fuzz_dep_.CoverTab[988]++
//line /snap/go/10455/src/strings/replace.go:54
			// _ = "end of CoverTab[988]"
//line /snap/go/10455/src/strings/replace.go:54
		}
//line /snap/go/10455/src/strings/replace.go:54
		// _ = "end of CoverTab[985]"
//line /snap/go/10455/src/strings/replace.go:54
		_go_fuzz_dep_.CoverTab[986]++
								if len(oldnew[i+1]) != 1 {
//line /snap/go/10455/src/strings/replace.go:55
			_go_fuzz_dep_.CoverTab[524983]++
//line /snap/go/10455/src/strings/replace.go:55
			_go_fuzz_dep_.CoverTab[989]++
									allNewBytes = false
//line /snap/go/10455/src/strings/replace.go:56
			// _ = "end of CoverTab[989]"
		} else {
//line /snap/go/10455/src/strings/replace.go:57
			_go_fuzz_dep_.CoverTab[524984]++
//line /snap/go/10455/src/strings/replace.go:57
			_go_fuzz_dep_.CoverTab[990]++
//line /snap/go/10455/src/strings/replace.go:57
			// _ = "end of CoverTab[990]"
//line /snap/go/10455/src/strings/replace.go:57
		}
//line /snap/go/10455/src/strings/replace.go:57
		// _ = "end of CoverTab[986]"
	}
//line /snap/go/10455/src/strings/replace.go:58
	if _go_fuzz_dep_.CoverTab[786483] == 0 {
//line /snap/go/10455/src/strings/replace.go:58
		_go_fuzz_dep_.CoverTab[525083]++
//line /snap/go/10455/src/strings/replace.go:58
	} else {
//line /snap/go/10455/src/strings/replace.go:58
		_go_fuzz_dep_.CoverTab[525084]++
//line /snap/go/10455/src/strings/replace.go:58
	}
//line /snap/go/10455/src/strings/replace.go:58
	// _ = "end of CoverTab[978]"
//line /snap/go/10455/src/strings/replace.go:58
	_go_fuzz_dep_.CoverTab[979]++

							if allNewBytes {
//line /snap/go/10455/src/strings/replace.go:60
		_go_fuzz_dep_.CoverTab[524985]++
//line /snap/go/10455/src/strings/replace.go:60
		_go_fuzz_dep_.CoverTab[991]++
								r := byteReplacer{}
//line /snap/go/10455/src/strings/replace.go:61
		_go_fuzz_dep_.CoverTab[786485] = 0
								for i := range r {
//line /snap/go/10455/src/strings/replace.go:62
			if _go_fuzz_dep_.CoverTab[786485] == 0 {
//line /snap/go/10455/src/strings/replace.go:62
				_go_fuzz_dep_.CoverTab[525089]++
//line /snap/go/10455/src/strings/replace.go:62
			} else {
//line /snap/go/10455/src/strings/replace.go:62
				_go_fuzz_dep_.CoverTab[525090]++
//line /snap/go/10455/src/strings/replace.go:62
			}
//line /snap/go/10455/src/strings/replace.go:62
			_go_fuzz_dep_.CoverTab[786485] = 1
//line /snap/go/10455/src/strings/replace.go:62
			_go_fuzz_dep_.CoverTab[994]++
									r[i] = byte(i)
//line /snap/go/10455/src/strings/replace.go:63
			// _ = "end of CoverTab[994]"
		}
//line /snap/go/10455/src/strings/replace.go:64
		if _go_fuzz_dep_.CoverTab[786485] == 0 {
//line /snap/go/10455/src/strings/replace.go:64
			_go_fuzz_dep_.CoverTab[525091]++
//line /snap/go/10455/src/strings/replace.go:64
		} else {
//line /snap/go/10455/src/strings/replace.go:64
			_go_fuzz_dep_.CoverTab[525092]++
//line /snap/go/10455/src/strings/replace.go:64
		}
//line /snap/go/10455/src/strings/replace.go:64
		// _ = "end of CoverTab[991]"
//line /snap/go/10455/src/strings/replace.go:64
		_go_fuzz_dep_.CoverTab[992]++
//line /snap/go/10455/src/strings/replace.go:64
		_go_fuzz_dep_.CoverTab[786486] = 0

//line /snap/go/10455/src/strings/replace.go:67
		for i := len(oldnew) - 2; i >= 0; i -= 2 {
//line /snap/go/10455/src/strings/replace.go:67
			if _go_fuzz_dep_.CoverTab[786486] == 0 {
//line /snap/go/10455/src/strings/replace.go:67
				_go_fuzz_dep_.CoverTab[525093]++
//line /snap/go/10455/src/strings/replace.go:67
			} else {
//line /snap/go/10455/src/strings/replace.go:67
				_go_fuzz_dep_.CoverTab[525094]++
//line /snap/go/10455/src/strings/replace.go:67
			}
//line /snap/go/10455/src/strings/replace.go:67
			_go_fuzz_dep_.CoverTab[786486] = 1
//line /snap/go/10455/src/strings/replace.go:67
			_go_fuzz_dep_.CoverTab[995]++
									o := oldnew[i][0]
									n := oldnew[i+1][0]
									r[o] = n
//line /snap/go/10455/src/strings/replace.go:70
			// _ = "end of CoverTab[995]"
		}
//line /snap/go/10455/src/strings/replace.go:71
		if _go_fuzz_dep_.CoverTab[786486] == 0 {
//line /snap/go/10455/src/strings/replace.go:71
			_go_fuzz_dep_.CoverTab[525095]++
//line /snap/go/10455/src/strings/replace.go:71
		} else {
//line /snap/go/10455/src/strings/replace.go:71
			_go_fuzz_dep_.CoverTab[525096]++
//line /snap/go/10455/src/strings/replace.go:71
		}
//line /snap/go/10455/src/strings/replace.go:71
		// _ = "end of CoverTab[992]"
//line /snap/go/10455/src/strings/replace.go:71
		_go_fuzz_dep_.CoverTab[993]++
								return &r
//line /snap/go/10455/src/strings/replace.go:72
		// _ = "end of CoverTab[993]"
	} else {
//line /snap/go/10455/src/strings/replace.go:73
		_go_fuzz_dep_.CoverTab[524986]++
//line /snap/go/10455/src/strings/replace.go:73
		_go_fuzz_dep_.CoverTab[996]++
//line /snap/go/10455/src/strings/replace.go:73
		// _ = "end of CoverTab[996]"
//line /snap/go/10455/src/strings/replace.go:73
	}
//line /snap/go/10455/src/strings/replace.go:73
	// _ = "end of CoverTab[979]"
//line /snap/go/10455/src/strings/replace.go:73
	_go_fuzz_dep_.CoverTab[980]++

							r := byteStringReplacer{toReplace: make([]string, 0, len(oldnew)/2)}
//line /snap/go/10455/src/strings/replace.go:75
	_go_fuzz_dep_.CoverTab[786484] = 0

//line /snap/go/10455/src/strings/replace.go:78
	for i := len(oldnew) - 2; i >= 0; i -= 2 {
//line /snap/go/10455/src/strings/replace.go:78
		if _go_fuzz_dep_.CoverTab[786484] == 0 {
//line /snap/go/10455/src/strings/replace.go:78
			_go_fuzz_dep_.CoverTab[525085]++
//line /snap/go/10455/src/strings/replace.go:78
		} else {
//line /snap/go/10455/src/strings/replace.go:78
			_go_fuzz_dep_.CoverTab[525086]++
//line /snap/go/10455/src/strings/replace.go:78
		}
//line /snap/go/10455/src/strings/replace.go:78
		_go_fuzz_dep_.CoverTab[786484] = 1
//line /snap/go/10455/src/strings/replace.go:78
		_go_fuzz_dep_.CoverTab[997]++
								o := oldnew[i][0]
								n := oldnew[i+1]

								if r.replacements[o] == nil {
//line /snap/go/10455/src/strings/replace.go:82
			_go_fuzz_dep_.CoverTab[524987]++
//line /snap/go/10455/src/strings/replace.go:82
			_go_fuzz_dep_.CoverTab[999]++

//line /snap/go/10455/src/strings/replace.go:86
			r.toReplace = append(r.toReplace, string([]byte{o}))
//line /snap/go/10455/src/strings/replace.go:86
			// _ = "end of CoverTab[999]"
		} else {
//line /snap/go/10455/src/strings/replace.go:87
			_go_fuzz_dep_.CoverTab[524988]++
//line /snap/go/10455/src/strings/replace.go:87
			_go_fuzz_dep_.CoverTab[1000]++
//line /snap/go/10455/src/strings/replace.go:87
			// _ = "end of CoverTab[1000]"
//line /snap/go/10455/src/strings/replace.go:87
		}
//line /snap/go/10455/src/strings/replace.go:87
		// _ = "end of CoverTab[997]"
//line /snap/go/10455/src/strings/replace.go:87
		_go_fuzz_dep_.CoverTab[998]++
								r.replacements[o] = []byte(n)
//line /snap/go/10455/src/strings/replace.go:88
		// _ = "end of CoverTab[998]"

	}
//line /snap/go/10455/src/strings/replace.go:90
	if _go_fuzz_dep_.CoverTab[786484] == 0 {
//line /snap/go/10455/src/strings/replace.go:90
		_go_fuzz_dep_.CoverTab[525087]++
//line /snap/go/10455/src/strings/replace.go:90
	} else {
//line /snap/go/10455/src/strings/replace.go:90
		_go_fuzz_dep_.CoverTab[525088]++
//line /snap/go/10455/src/strings/replace.go:90
	}
//line /snap/go/10455/src/strings/replace.go:90
	// _ = "end of CoverTab[980]"
//line /snap/go/10455/src/strings/replace.go:90
	_go_fuzz_dep_.CoverTab[981]++
							return &r
//line /snap/go/10455/src/strings/replace.go:91
	// _ = "end of CoverTab[981]"
}

// Replace returns a copy of s with all replacements performed.
func (r *Replacer) Replace(s string) string {
//line /snap/go/10455/src/strings/replace.go:95
	_go_fuzz_dep_.CoverTab[1001]++
							r.once.Do(r.buildOnce)
							return r.r.Replace(s)
//line /snap/go/10455/src/strings/replace.go:97
	// _ = "end of CoverTab[1001]"
}

// WriteString writes s to w with all replacements performed.
func (r *Replacer) WriteString(w io.Writer, s string) (n int, err error) {
//line /snap/go/10455/src/strings/replace.go:101
	_go_fuzz_dep_.CoverTab[1002]++
							r.once.Do(r.buildOnce)
							return r.r.WriteString(w, s)
//line /snap/go/10455/src/strings/replace.go:103
	// _ = "end of CoverTab[1002]"
}

// trieNode is a node in a lookup trie for prioritized key/value pairs. Keys
//line /snap/go/10455/src/strings/replace.go:106
// and values may be empty. For example, the trie containing keys "ax", "ay",
//line /snap/go/10455/src/strings/replace.go:106
// "bcbc", "x" and "xy" could have eight nodes:
//line /snap/go/10455/src/strings/replace.go:106
//
//line /snap/go/10455/src/strings/replace.go:106
//	n0  -
//line /snap/go/10455/src/strings/replace.go:106
//	n1  a-
//line /snap/go/10455/src/strings/replace.go:106
//	n2  .x+
//line /snap/go/10455/src/strings/replace.go:106
//	n3  .y+
//line /snap/go/10455/src/strings/replace.go:106
//	n4  b-
//line /snap/go/10455/src/strings/replace.go:106
//	n5  .cbc+
//line /snap/go/10455/src/strings/replace.go:106
//	n6  x+
//line /snap/go/10455/src/strings/replace.go:106
//	n7  .y+
//line /snap/go/10455/src/strings/replace.go:106
//
//line /snap/go/10455/src/strings/replace.go:106
// n0 is the root node, and its children are n1, n4 and n6; n1's children are
//line /snap/go/10455/src/strings/replace.go:106
// n2 and n3; n4's child is n5; n6's child is n7. Nodes n0, n1 and n4 (marked
//line /snap/go/10455/src/strings/replace.go:106
// with a trailing "-") are partial keys, and nodes n2, n3, n5, n6 and n7
//line /snap/go/10455/src/strings/replace.go:106
// (marked with a trailing "+") are complete keys.
//line /snap/go/10455/src/strings/replace.go:123
type trieNode struct {
	// value is the value of the trie node's key/value pair. It is empty if
	// this node is not a complete key.
	value	string
	// priority is the priority (higher is more important) of the trie node's
	// key/value pair; keys are not necessarily matched shortest- or longest-
	// first. Priority is positive if this node is a complete key, and zero
	// otherwise. In the example above, positive/zero priorities are marked
	// with a trailing "+" or "-".
	priority	int

//line /snap/go/10455/src/strings/replace.go:142
	// prefix is the difference in keys between this trie node and the next.
	// In the example above, node n4 has prefix "cbc" and n4's next node is n5.
	// Node n5 has no children and so has zero prefix, next and table fields.
	prefix	string
	next	*trieNode

	// table is a lookup table indexed by the next byte in the key, after
	// remapping that byte through genericReplacer.mapping to create a dense
	// index. In the example above, the keys only use 'a', 'b', 'c', 'x' and
	// 'y', which remap to 0, 1, 2, 3 and 4. All other bytes remap to 5, and
	// genericReplacer.tableSize will be 5. Node n0's table will be
	// []*trieNode{ 0:n1, 1:n4, 3:n6 }, where the 0, 1 and 3 are the remapped
	// 'a', 'b' and 'x'.
	table	[]*trieNode
}

func (t *trieNode) add(key, val string, priority int, r *genericReplacer) {
//line /snap/go/10455/src/strings/replace.go:158
	_go_fuzz_dep_.CoverTab[1003]++
							if key == "" {
//line /snap/go/10455/src/strings/replace.go:159
		_go_fuzz_dep_.CoverTab[524989]++
//line /snap/go/10455/src/strings/replace.go:159
		_go_fuzz_dep_.CoverTab[1005]++
								if t.priority == 0 {
//line /snap/go/10455/src/strings/replace.go:160
			_go_fuzz_dep_.CoverTab[524991]++
//line /snap/go/10455/src/strings/replace.go:160
			_go_fuzz_dep_.CoverTab[1007]++
									t.value = val
									t.priority = priority
//line /snap/go/10455/src/strings/replace.go:162
			// _ = "end of CoverTab[1007]"
		} else {
//line /snap/go/10455/src/strings/replace.go:163
			_go_fuzz_dep_.CoverTab[524992]++
//line /snap/go/10455/src/strings/replace.go:163
			_go_fuzz_dep_.CoverTab[1008]++
//line /snap/go/10455/src/strings/replace.go:163
			// _ = "end of CoverTab[1008]"
//line /snap/go/10455/src/strings/replace.go:163
		}
//line /snap/go/10455/src/strings/replace.go:163
		// _ = "end of CoverTab[1005]"
//line /snap/go/10455/src/strings/replace.go:163
		_go_fuzz_dep_.CoverTab[1006]++
								return
//line /snap/go/10455/src/strings/replace.go:164
		// _ = "end of CoverTab[1006]"
	} else {
//line /snap/go/10455/src/strings/replace.go:165
		_go_fuzz_dep_.CoverTab[524990]++
//line /snap/go/10455/src/strings/replace.go:165
		_go_fuzz_dep_.CoverTab[1009]++
//line /snap/go/10455/src/strings/replace.go:165
		// _ = "end of CoverTab[1009]"
//line /snap/go/10455/src/strings/replace.go:165
	}
//line /snap/go/10455/src/strings/replace.go:165
	// _ = "end of CoverTab[1003]"
//line /snap/go/10455/src/strings/replace.go:165
	_go_fuzz_dep_.CoverTab[1004]++

							if t.prefix != "" {
//line /snap/go/10455/src/strings/replace.go:167
		_go_fuzz_dep_.CoverTab[524993]++
//line /snap/go/10455/src/strings/replace.go:167
		_go_fuzz_dep_.CoverTab[1010]++
								// Need to split the prefix among multiple nodes.
								var n int
//line /snap/go/10455/src/strings/replace.go:169
		_go_fuzz_dep_. // length of the longest common prefix
//line /snap/go/10455/src/strings/replace.go:169
		CoverTab[786487] = 0
								for ; n < len(t.prefix) && func() bool {
//line /snap/go/10455/src/strings/replace.go:170
			_go_fuzz_dep_.CoverTab[1012]++
//line /snap/go/10455/src/strings/replace.go:170
			return n < len(key)
//line /snap/go/10455/src/strings/replace.go:170
			// _ = "end of CoverTab[1012]"
//line /snap/go/10455/src/strings/replace.go:170
		}(); n++ {
//line /snap/go/10455/src/strings/replace.go:170
			if _go_fuzz_dep_.CoverTab[786487] == 0 {
//line /snap/go/10455/src/strings/replace.go:170
				_go_fuzz_dep_.CoverTab[525097]++
//line /snap/go/10455/src/strings/replace.go:170
			} else {
//line /snap/go/10455/src/strings/replace.go:170
				_go_fuzz_dep_.CoverTab[525098]++
//line /snap/go/10455/src/strings/replace.go:170
			}
//line /snap/go/10455/src/strings/replace.go:170
			_go_fuzz_dep_.CoverTab[786487] = 1
//line /snap/go/10455/src/strings/replace.go:170
			_go_fuzz_dep_.CoverTab[1013]++
									if t.prefix[n] != key[n] {
//line /snap/go/10455/src/strings/replace.go:171
				_go_fuzz_dep_.CoverTab[524995]++
//line /snap/go/10455/src/strings/replace.go:171
				_go_fuzz_dep_.CoverTab[1014]++
										break
//line /snap/go/10455/src/strings/replace.go:172
				// _ = "end of CoverTab[1014]"
			} else {
//line /snap/go/10455/src/strings/replace.go:173
				_go_fuzz_dep_.CoverTab[524996]++
//line /snap/go/10455/src/strings/replace.go:173
				_go_fuzz_dep_.CoverTab[1015]++
//line /snap/go/10455/src/strings/replace.go:173
				// _ = "end of CoverTab[1015]"
//line /snap/go/10455/src/strings/replace.go:173
			}
//line /snap/go/10455/src/strings/replace.go:173
			// _ = "end of CoverTab[1013]"
		}
//line /snap/go/10455/src/strings/replace.go:174
		if _go_fuzz_dep_.CoverTab[786487] == 0 {
//line /snap/go/10455/src/strings/replace.go:174
			_go_fuzz_dep_.CoverTab[525099]++
//line /snap/go/10455/src/strings/replace.go:174
		} else {
//line /snap/go/10455/src/strings/replace.go:174
			_go_fuzz_dep_.CoverTab[525100]++
//line /snap/go/10455/src/strings/replace.go:174
		}
//line /snap/go/10455/src/strings/replace.go:174
		// _ = "end of CoverTab[1010]"
//line /snap/go/10455/src/strings/replace.go:174
		_go_fuzz_dep_.CoverTab[1011]++
								if n == len(t.prefix) {
//line /snap/go/10455/src/strings/replace.go:175
			_go_fuzz_dep_.CoverTab[524997]++
//line /snap/go/10455/src/strings/replace.go:175
			_go_fuzz_dep_.CoverTab[1016]++
									t.next.add(key[n:], val, priority, r)
//line /snap/go/10455/src/strings/replace.go:176
			// _ = "end of CoverTab[1016]"
		} else {
//line /snap/go/10455/src/strings/replace.go:177
			_go_fuzz_dep_.CoverTab[524998]++
//line /snap/go/10455/src/strings/replace.go:177
			_go_fuzz_dep_.CoverTab[1017]++
//line /snap/go/10455/src/strings/replace.go:177
			if n == 0 {
//line /snap/go/10455/src/strings/replace.go:177
				_go_fuzz_dep_.CoverTab[524999]++
//line /snap/go/10455/src/strings/replace.go:177
				_go_fuzz_dep_.CoverTab[1018]++
				// First byte differs, start a new lookup table here. Looking up
				// what is currently t.prefix[0] will lead to prefixNode, and
				// looking up key[0] will lead to keyNode.
				var prefixNode *trieNode
				if len(t.prefix) == 1 {
//line /snap/go/10455/src/strings/replace.go:182
					_go_fuzz_dep_.CoverTab[525001]++
//line /snap/go/10455/src/strings/replace.go:182
					_go_fuzz_dep_.CoverTab[1020]++
											prefixNode = t.next
//line /snap/go/10455/src/strings/replace.go:183
					// _ = "end of CoverTab[1020]"
				} else {
//line /snap/go/10455/src/strings/replace.go:184
					_go_fuzz_dep_.CoverTab[525002]++
//line /snap/go/10455/src/strings/replace.go:184
					_go_fuzz_dep_.CoverTab[1021]++
											prefixNode = &trieNode{
						prefix:	t.prefix[1:],
						next:	t.next,
					}
//line /snap/go/10455/src/strings/replace.go:188
					// _ = "end of CoverTab[1021]"
				}
//line /snap/go/10455/src/strings/replace.go:189
				// _ = "end of CoverTab[1018]"
//line /snap/go/10455/src/strings/replace.go:189
				_go_fuzz_dep_.CoverTab[1019]++
										keyNode := new(trieNode)
										t.table = make([]*trieNode, r.tableSize)
										t.table[r.mapping[t.prefix[0]]] = prefixNode
										t.table[r.mapping[key[0]]] = keyNode
										t.prefix = ""
										t.next = nil
										keyNode.add(key[1:], val, priority, r)
//line /snap/go/10455/src/strings/replace.go:196
				// _ = "end of CoverTab[1019]"
			} else {
//line /snap/go/10455/src/strings/replace.go:197
				_go_fuzz_dep_.CoverTab[525000]++
//line /snap/go/10455/src/strings/replace.go:197
				_go_fuzz_dep_.CoverTab[1022]++

										next := &trieNode{
					prefix:	t.prefix[n:],
					next:	t.next,
				}
										t.prefix = t.prefix[:n]
										t.next = next
										next.add(key[n:], val, priority, r)
//line /snap/go/10455/src/strings/replace.go:205
				// _ = "end of CoverTab[1022]"
			}
//line /snap/go/10455/src/strings/replace.go:206
			// _ = "end of CoverTab[1017]"
//line /snap/go/10455/src/strings/replace.go:206
		}
//line /snap/go/10455/src/strings/replace.go:206
		// _ = "end of CoverTab[1011]"
	} else {
//line /snap/go/10455/src/strings/replace.go:207
		_go_fuzz_dep_.CoverTab[524994]++
//line /snap/go/10455/src/strings/replace.go:207
		_go_fuzz_dep_.CoverTab[1023]++
//line /snap/go/10455/src/strings/replace.go:207
		if t.table != nil {
//line /snap/go/10455/src/strings/replace.go:207
			_go_fuzz_dep_.CoverTab[525003]++
//line /snap/go/10455/src/strings/replace.go:207
			_go_fuzz_dep_.CoverTab[1024]++

									m := r.mapping[key[0]]
									if t.table[m] == nil {
//line /snap/go/10455/src/strings/replace.go:210
				_go_fuzz_dep_.CoverTab[525005]++
//line /snap/go/10455/src/strings/replace.go:210
				_go_fuzz_dep_.CoverTab[1026]++
										t.table[m] = new(trieNode)
//line /snap/go/10455/src/strings/replace.go:211
				// _ = "end of CoverTab[1026]"
			} else {
//line /snap/go/10455/src/strings/replace.go:212
				_go_fuzz_dep_.CoverTab[525006]++
//line /snap/go/10455/src/strings/replace.go:212
				_go_fuzz_dep_.CoverTab[1027]++
//line /snap/go/10455/src/strings/replace.go:212
				// _ = "end of CoverTab[1027]"
//line /snap/go/10455/src/strings/replace.go:212
			}
//line /snap/go/10455/src/strings/replace.go:212
			// _ = "end of CoverTab[1024]"
//line /snap/go/10455/src/strings/replace.go:212
			_go_fuzz_dep_.CoverTab[1025]++
									t.table[m].add(key[1:], val, priority, r)
//line /snap/go/10455/src/strings/replace.go:213
			// _ = "end of CoverTab[1025]"
		} else {
//line /snap/go/10455/src/strings/replace.go:214
			_go_fuzz_dep_.CoverTab[525004]++
//line /snap/go/10455/src/strings/replace.go:214
			_go_fuzz_dep_.CoverTab[1028]++
									t.prefix = key
									t.next = new(trieNode)
									t.next.add("", val, priority, r)
//line /snap/go/10455/src/strings/replace.go:217
			// _ = "end of CoverTab[1028]"
		}
//line /snap/go/10455/src/strings/replace.go:218
		// _ = "end of CoverTab[1023]"
//line /snap/go/10455/src/strings/replace.go:218
	}
//line /snap/go/10455/src/strings/replace.go:218
	// _ = "end of CoverTab[1004]"
}

func (r *genericReplacer) lookup(s string, ignoreRoot bool) (val string, keylen int, found bool) {
//line /snap/go/10455/src/strings/replace.go:221
	_go_fuzz_dep_.CoverTab[1029]++

//line /snap/go/10455/src/strings/replace.go:224
	bestPriority := 0
							node := &r.root
							n := 0
//line /snap/go/10455/src/strings/replace.go:226
	_go_fuzz_dep_.CoverTab[786488] = 0
							for node != nil {
//line /snap/go/10455/src/strings/replace.go:227
		if _go_fuzz_dep_.CoverTab[786488] == 0 {
//line /snap/go/10455/src/strings/replace.go:227
			_go_fuzz_dep_.CoverTab[525101]++
//line /snap/go/10455/src/strings/replace.go:227
		} else {
//line /snap/go/10455/src/strings/replace.go:227
			_go_fuzz_dep_.CoverTab[525102]++
//line /snap/go/10455/src/strings/replace.go:227
		}
//line /snap/go/10455/src/strings/replace.go:227
		_go_fuzz_dep_.CoverTab[786488] = 1
//line /snap/go/10455/src/strings/replace.go:227
		_go_fuzz_dep_.CoverTab[1031]++
								if node.priority > bestPriority && func() bool {
//line /snap/go/10455/src/strings/replace.go:228
			_go_fuzz_dep_.CoverTab[1034]++
//line /snap/go/10455/src/strings/replace.go:228
			return !(ignoreRoot && func() bool {
//line /snap/go/10455/src/strings/replace.go:228
				_go_fuzz_dep_.CoverTab[1035]++
//line /snap/go/10455/src/strings/replace.go:228
				return node == &r.root
//line /snap/go/10455/src/strings/replace.go:228
				// _ = "end of CoverTab[1035]"
//line /snap/go/10455/src/strings/replace.go:228
			}())
//line /snap/go/10455/src/strings/replace.go:228
			// _ = "end of CoverTab[1034]"
//line /snap/go/10455/src/strings/replace.go:228
		}() {
//line /snap/go/10455/src/strings/replace.go:228
			_go_fuzz_dep_.CoverTab[525007]++
//line /snap/go/10455/src/strings/replace.go:228
			_go_fuzz_dep_.CoverTab[1036]++
									bestPriority = node.priority
									val = node.value
									keylen = n
									found = true
//line /snap/go/10455/src/strings/replace.go:232
			// _ = "end of CoverTab[1036]"
		} else {
//line /snap/go/10455/src/strings/replace.go:233
			_go_fuzz_dep_.CoverTab[525008]++
//line /snap/go/10455/src/strings/replace.go:233
			_go_fuzz_dep_.CoverTab[1037]++
//line /snap/go/10455/src/strings/replace.go:233
			// _ = "end of CoverTab[1037]"
//line /snap/go/10455/src/strings/replace.go:233
		}
//line /snap/go/10455/src/strings/replace.go:233
		// _ = "end of CoverTab[1031]"
//line /snap/go/10455/src/strings/replace.go:233
		_go_fuzz_dep_.CoverTab[1032]++

								if s == "" {
//line /snap/go/10455/src/strings/replace.go:235
			_go_fuzz_dep_.CoverTab[525009]++
//line /snap/go/10455/src/strings/replace.go:235
			_go_fuzz_dep_.CoverTab[1038]++
									break
//line /snap/go/10455/src/strings/replace.go:236
			// _ = "end of CoverTab[1038]"
		} else {
//line /snap/go/10455/src/strings/replace.go:237
			_go_fuzz_dep_.CoverTab[525010]++
//line /snap/go/10455/src/strings/replace.go:237
			_go_fuzz_dep_.CoverTab[1039]++
//line /snap/go/10455/src/strings/replace.go:237
			// _ = "end of CoverTab[1039]"
//line /snap/go/10455/src/strings/replace.go:237
		}
//line /snap/go/10455/src/strings/replace.go:237
		// _ = "end of CoverTab[1032]"
//line /snap/go/10455/src/strings/replace.go:237
		_go_fuzz_dep_.CoverTab[1033]++
								if node.table != nil {
//line /snap/go/10455/src/strings/replace.go:238
			_go_fuzz_dep_.CoverTab[525011]++
//line /snap/go/10455/src/strings/replace.go:238
			_go_fuzz_dep_.CoverTab[1040]++
									index := r.mapping[s[0]]
									if int(index) == r.tableSize {
//line /snap/go/10455/src/strings/replace.go:240
				_go_fuzz_dep_.CoverTab[525013]++
//line /snap/go/10455/src/strings/replace.go:240
				_go_fuzz_dep_.CoverTab[1042]++
										break
//line /snap/go/10455/src/strings/replace.go:241
				// _ = "end of CoverTab[1042]"
			} else {
//line /snap/go/10455/src/strings/replace.go:242
				_go_fuzz_dep_.CoverTab[525014]++
//line /snap/go/10455/src/strings/replace.go:242
				_go_fuzz_dep_.CoverTab[1043]++
//line /snap/go/10455/src/strings/replace.go:242
				// _ = "end of CoverTab[1043]"
//line /snap/go/10455/src/strings/replace.go:242
			}
//line /snap/go/10455/src/strings/replace.go:242
			// _ = "end of CoverTab[1040]"
//line /snap/go/10455/src/strings/replace.go:242
			_go_fuzz_dep_.CoverTab[1041]++
									node = node.table[index]
									s = s[1:]
									n++
//line /snap/go/10455/src/strings/replace.go:245
			// _ = "end of CoverTab[1041]"
		} else {
//line /snap/go/10455/src/strings/replace.go:246
			_go_fuzz_dep_.CoverTab[525012]++
//line /snap/go/10455/src/strings/replace.go:246
			_go_fuzz_dep_.CoverTab[1044]++
//line /snap/go/10455/src/strings/replace.go:246
			if node.prefix != "" && func() bool {
//line /snap/go/10455/src/strings/replace.go:246
				_go_fuzz_dep_.CoverTab[1045]++
//line /snap/go/10455/src/strings/replace.go:246
				return HasPrefix(s, node.prefix)
//line /snap/go/10455/src/strings/replace.go:246
				// _ = "end of CoverTab[1045]"
//line /snap/go/10455/src/strings/replace.go:246
			}() {
//line /snap/go/10455/src/strings/replace.go:246
				_go_fuzz_dep_.CoverTab[525015]++
//line /snap/go/10455/src/strings/replace.go:246
				_go_fuzz_dep_.CoverTab[1046]++
										n += len(node.prefix)
										s = s[len(node.prefix):]
										node = node.next
//line /snap/go/10455/src/strings/replace.go:249
				// _ = "end of CoverTab[1046]"
			} else {
//line /snap/go/10455/src/strings/replace.go:250
				_go_fuzz_dep_.CoverTab[525016]++
//line /snap/go/10455/src/strings/replace.go:250
				_go_fuzz_dep_.CoverTab[1047]++
										break
//line /snap/go/10455/src/strings/replace.go:251
				// _ = "end of CoverTab[1047]"
			}
//line /snap/go/10455/src/strings/replace.go:252
			// _ = "end of CoverTab[1044]"
//line /snap/go/10455/src/strings/replace.go:252
		}
//line /snap/go/10455/src/strings/replace.go:252
		// _ = "end of CoverTab[1033]"
	}
//line /snap/go/10455/src/strings/replace.go:253
	if _go_fuzz_dep_.CoverTab[786488] == 0 {
//line /snap/go/10455/src/strings/replace.go:253
		_go_fuzz_dep_.CoverTab[525103]++
//line /snap/go/10455/src/strings/replace.go:253
	} else {
//line /snap/go/10455/src/strings/replace.go:253
		_go_fuzz_dep_.CoverTab[525104]++
//line /snap/go/10455/src/strings/replace.go:253
	}
//line /snap/go/10455/src/strings/replace.go:253
	// _ = "end of CoverTab[1029]"
//line /snap/go/10455/src/strings/replace.go:253
	_go_fuzz_dep_.CoverTab[1030]++
							return
//line /snap/go/10455/src/strings/replace.go:254
	// _ = "end of CoverTab[1030]"
}

// genericReplacer is the fully generic algorithm.
//line /snap/go/10455/src/strings/replace.go:257
// It's used as a fallback when nothing faster can be used.
//line /snap/go/10455/src/strings/replace.go:259
type genericReplacer struct {
	root	trieNode
	// tableSize is the size of a trie node's lookup table. It is the number
	// of unique key bytes.
	tableSize	int
	// mapping maps from key bytes to a dense index for trieNode.table.
	mapping	[256]byte
}

func makeGenericReplacer(oldnew []string) *genericReplacer {
//line /snap/go/10455/src/strings/replace.go:268
	_go_fuzz_dep_.CoverTab[1048]++
							r := new(genericReplacer)
//line /snap/go/10455/src/strings/replace.go:269
	_go_fuzz_dep_.CoverTab[786489] = 0

							for i := 0; i < len(oldnew); i += 2 {
//line /snap/go/10455/src/strings/replace.go:271
		if _go_fuzz_dep_.CoverTab[786489] == 0 {
//line /snap/go/10455/src/strings/replace.go:271
			_go_fuzz_dep_.CoverTab[525105]++
//line /snap/go/10455/src/strings/replace.go:271
		} else {
//line /snap/go/10455/src/strings/replace.go:271
			_go_fuzz_dep_.CoverTab[525106]++
//line /snap/go/10455/src/strings/replace.go:271
		}
//line /snap/go/10455/src/strings/replace.go:271
		_go_fuzz_dep_.CoverTab[786489] = 1
//line /snap/go/10455/src/strings/replace.go:271
		_go_fuzz_dep_.CoverTab[1053]++
								key := oldnew[i]
//line /snap/go/10455/src/strings/replace.go:272
		_go_fuzz_dep_.CoverTab[786493] = 0
								for j := 0; j < len(key); j++ {
//line /snap/go/10455/src/strings/replace.go:273
			if _go_fuzz_dep_.CoverTab[786493] == 0 {
//line /snap/go/10455/src/strings/replace.go:273
				_go_fuzz_dep_.CoverTab[525121]++
//line /snap/go/10455/src/strings/replace.go:273
			} else {
//line /snap/go/10455/src/strings/replace.go:273
				_go_fuzz_dep_.CoverTab[525122]++
//line /snap/go/10455/src/strings/replace.go:273
			}
//line /snap/go/10455/src/strings/replace.go:273
			_go_fuzz_dep_.CoverTab[786493] = 1
//line /snap/go/10455/src/strings/replace.go:273
			_go_fuzz_dep_.CoverTab[1054]++
									r.mapping[key[j]] = 1
//line /snap/go/10455/src/strings/replace.go:274
			// _ = "end of CoverTab[1054]"
		}
//line /snap/go/10455/src/strings/replace.go:275
		if _go_fuzz_dep_.CoverTab[786493] == 0 {
//line /snap/go/10455/src/strings/replace.go:275
			_go_fuzz_dep_.CoverTab[525123]++
//line /snap/go/10455/src/strings/replace.go:275
		} else {
//line /snap/go/10455/src/strings/replace.go:275
			_go_fuzz_dep_.CoverTab[525124]++
//line /snap/go/10455/src/strings/replace.go:275
		}
//line /snap/go/10455/src/strings/replace.go:275
		// _ = "end of CoverTab[1053]"
	}
//line /snap/go/10455/src/strings/replace.go:276
	if _go_fuzz_dep_.CoverTab[786489] == 0 {
//line /snap/go/10455/src/strings/replace.go:276
		_go_fuzz_dep_.CoverTab[525107]++
//line /snap/go/10455/src/strings/replace.go:276
	} else {
//line /snap/go/10455/src/strings/replace.go:276
		_go_fuzz_dep_.CoverTab[525108]++
//line /snap/go/10455/src/strings/replace.go:276
	}
//line /snap/go/10455/src/strings/replace.go:276
	// _ = "end of CoverTab[1048]"
//line /snap/go/10455/src/strings/replace.go:276
	_go_fuzz_dep_.CoverTab[1049]++
//line /snap/go/10455/src/strings/replace.go:276
	_go_fuzz_dep_.CoverTab[786490] = 0

							for _, b := range r.mapping {
//line /snap/go/10455/src/strings/replace.go:278
		if _go_fuzz_dep_.CoverTab[786490] == 0 {
//line /snap/go/10455/src/strings/replace.go:278
			_go_fuzz_dep_.CoverTab[525109]++
//line /snap/go/10455/src/strings/replace.go:278
		} else {
//line /snap/go/10455/src/strings/replace.go:278
			_go_fuzz_dep_.CoverTab[525110]++
//line /snap/go/10455/src/strings/replace.go:278
		}
//line /snap/go/10455/src/strings/replace.go:278
		_go_fuzz_dep_.CoverTab[786490] = 1
//line /snap/go/10455/src/strings/replace.go:278
		_go_fuzz_dep_.CoverTab[1055]++
								r.tableSize += int(b)
//line /snap/go/10455/src/strings/replace.go:279
		// _ = "end of CoverTab[1055]"
	}
//line /snap/go/10455/src/strings/replace.go:280
	if _go_fuzz_dep_.CoverTab[786490] == 0 {
//line /snap/go/10455/src/strings/replace.go:280
		_go_fuzz_dep_.CoverTab[525111]++
//line /snap/go/10455/src/strings/replace.go:280
	} else {
//line /snap/go/10455/src/strings/replace.go:280
		_go_fuzz_dep_.CoverTab[525112]++
//line /snap/go/10455/src/strings/replace.go:280
	}
//line /snap/go/10455/src/strings/replace.go:280
	// _ = "end of CoverTab[1049]"
//line /snap/go/10455/src/strings/replace.go:280
	_go_fuzz_dep_.CoverTab[1050]++

							var index byte
//line /snap/go/10455/src/strings/replace.go:282
	_go_fuzz_dep_.CoverTab[786491] = 0
							for i, b := range r.mapping {
//line /snap/go/10455/src/strings/replace.go:283
		if _go_fuzz_dep_.CoverTab[786491] == 0 {
//line /snap/go/10455/src/strings/replace.go:283
			_go_fuzz_dep_.CoverTab[525113]++
//line /snap/go/10455/src/strings/replace.go:283
		} else {
//line /snap/go/10455/src/strings/replace.go:283
			_go_fuzz_dep_.CoverTab[525114]++
//line /snap/go/10455/src/strings/replace.go:283
		}
//line /snap/go/10455/src/strings/replace.go:283
		_go_fuzz_dep_.CoverTab[786491] = 1
//line /snap/go/10455/src/strings/replace.go:283
		_go_fuzz_dep_.CoverTab[1056]++
								if b == 0 {
//line /snap/go/10455/src/strings/replace.go:284
			_go_fuzz_dep_.CoverTab[525017]++
//line /snap/go/10455/src/strings/replace.go:284
			_go_fuzz_dep_.CoverTab[1057]++
									r.mapping[i] = byte(r.tableSize)
//line /snap/go/10455/src/strings/replace.go:285
			// _ = "end of CoverTab[1057]"
		} else {
//line /snap/go/10455/src/strings/replace.go:286
			_go_fuzz_dep_.CoverTab[525018]++
//line /snap/go/10455/src/strings/replace.go:286
			_go_fuzz_dep_.CoverTab[1058]++
									r.mapping[i] = index
									index++
//line /snap/go/10455/src/strings/replace.go:288
			// _ = "end of CoverTab[1058]"
		}
//line /snap/go/10455/src/strings/replace.go:289
		// _ = "end of CoverTab[1056]"
	}
//line /snap/go/10455/src/strings/replace.go:290
	if _go_fuzz_dep_.CoverTab[786491] == 0 {
//line /snap/go/10455/src/strings/replace.go:290
		_go_fuzz_dep_.CoverTab[525115]++
//line /snap/go/10455/src/strings/replace.go:290
	} else {
//line /snap/go/10455/src/strings/replace.go:290
		_go_fuzz_dep_.CoverTab[525116]++
//line /snap/go/10455/src/strings/replace.go:290
	}
//line /snap/go/10455/src/strings/replace.go:290
	// _ = "end of CoverTab[1050]"
//line /snap/go/10455/src/strings/replace.go:290
	_go_fuzz_dep_.CoverTab[1051]++

							r.root.table = make([]*trieNode, r.tableSize)
//line /snap/go/10455/src/strings/replace.go:292
	_go_fuzz_dep_.CoverTab[786492] = 0

							for i := 0; i < len(oldnew); i += 2 {
//line /snap/go/10455/src/strings/replace.go:294
		if _go_fuzz_dep_.CoverTab[786492] == 0 {
//line /snap/go/10455/src/strings/replace.go:294
			_go_fuzz_dep_.CoverTab[525117]++
//line /snap/go/10455/src/strings/replace.go:294
		} else {
//line /snap/go/10455/src/strings/replace.go:294
			_go_fuzz_dep_.CoverTab[525118]++
//line /snap/go/10455/src/strings/replace.go:294
		}
//line /snap/go/10455/src/strings/replace.go:294
		_go_fuzz_dep_.CoverTab[786492] = 1
//line /snap/go/10455/src/strings/replace.go:294
		_go_fuzz_dep_.CoverTab[1059]++
								r.root.add(oldnew[i], oldnew[i+1], len(oldnew)-i, r)
//line /snap/go/10455/src/strings/replace.go:295
		// _ = "end of CoverTab[1059]"
	}
//line /snap/go/10455/src/strings/replace.go:296
	if _go_fuzz_dep_.CoverTab[786492] == 0 {
//line /snap/go/10455/src/strings/replace.go:296
		_go_fuzz_dep_.CoverTab[525119]++
//line /snap/go/10455/src/strings/replace.go:296
	} else {
//line /snap/go/10455/src/strings/replace.go:296
		_go_fuzz_dep_.CoverTab[525120]++
//line /snap/go/10455/src/strings/replace.go:296
	}
//line /snap/go/10455/src/strings/replace.go:296
	// _ = "end of CoverTab[1051]"
//line /snap/go/10455/src/strings/replace.go:296
	_go_fuzz_dep_.CoverTab[1052]++
							return r
//line /snap/go/10455/src/strings/replace.go:297
	// _ = "end of CoverTab[1052]"
}

type appendSliceWriter []byte

// Write writes to the buffer to satisfy io.Writer.
func (w *appendSliceWriter) Write(p []byte) (int, error) {
//line /snap/go/10455/src/strings/replace.go:303
	_go_fuzz_dep_.CoverTab[1060]++
							*w = append(*w, p...)
							return len(p), nil
//line /snap/go/10455/src/strings/replace.go:305
	// _ = "end of CoverTab[1060]"
}

// WriteString writes to the buffer without string->[]byte->string allocations.
func (w *appendSliceWriter) WriteString(s string) (int, error) {
//line /snap/go/10455/src/strings/replace.go:309
	_go_fuzz_dep_.CoverTab[1061]++
							*w = append(*w, s...)
							return len(s), nil
//line /snap/go/10455/src/strings/replace.go:311
	// _ = "end of CoverTab[1061]"
}

type stringWriter struct {
	w io.Writer
}

func (w stringWriter) WriteString(s string) (int, error) {
//line /snap/go/10455/src/strings/replace.go:318
	_go_fuzz_dep_.CoverTab[1062]++
							return w.w.Write([]byte(s))
//line /snap/go/10455/src/strings/replace.go:319
	// _ = "end of CoverTab[1062]"
}

func getStringWriter(w io.Writer) io.StringWriter {
//line /snap/go/10455/src/strings/replace.go:322
	_go_fuzz_dep_.CoverTab[1063]++
							sw, ok := w.(io.StringWriter)
							if !ok {
//line /snap/go/10455/src/strings/replace.go:324
		_go_fuzz_dep_.CoverTab[525019]++
//line /snap/go/10455/src/strings/replace.go:324
		_go_fuzz_dep_.CoverTab[1065]++
								sw = stringWriter{w}
//line /snap/go/10455/src/strings/replace.go:325
		// _ = "end of CoverTab[1065]"
	} else {
//line /snap/go/10455/src/strings/replace.go:326
		_go_fuzz_dep_.CoverTab[525020]++
//line /snap/go/10455/src/strings/replace.go:326
		_go_fuzz_dep_.CoverTab[1066]++
//line /snap/go/10455/src/strings/replace.go:326
		// _ = "end of CoverTab[1066]"
//line /snap/go/10455/src/strings/replace.go:326
	}
//line /snap/go/10455/src/strings/replace.go:326
	// _ = "end of CoverTab[1063]"
//line /snap/go/10455/src/strings/replace.go:326
	_go_fuzz_dep_.CoverTab[1064]++
							return sw
//line /snap/go/10455/src/strings/replace.go:327
	// _ = "end of CoverTab[1064]"
}

func (r *genericReplacer) Replace(s string) string {
//line /snap/go/10455/src/strings/replace.go:330
	_go_fuzz_dep_.CoverTab[1067]++
							buf := make(appendSliceWriter, 0, len(s))
							r.WriteString(&buf, s)
							return string(buf)
//line /snap/go/10455/src/strings/replace.go:333
	// _ = "end of CoverTab[1067]"
}

func (r *genericReplacer) WriteString(w io.Writer, s string) (n int, err error) {
//line /snap/go/10455/src/strings/replace.go:336
	_go_fuzz_dep_.CoverTab[1068]++
							sw := getStringWriter(w)
							var last, wn int
							var prevMatchEmpty bool
//line /snap/go/10455/src/strings/replace.go:339
	_go_fuzz_dep_.CoverTab[786494] = 0
							for i := 0; i <= len(s); {
//line /snap/go/10455/src/strings/replace.go:340
		if _go_fuzz_dep_.CoverTab[786494] == 0 {
//line /snap/go/10455/src/strings/replace.go:340
			_go_fuzz_dep_.CoverTab[525125]++
//line /snap/go/10455/src/strings/replace.go:340
		} else {
//line /snap/go/10455/src/strings/replace.go:340
			_go_fuzz_dep_.CoverTab[525126]++
//line /snap/go/10455/src/strings/replace.go:340
		}
//line /snap/go/10455/src/strings/replace.go:340
		_go_fuzz_dep_.CoverTab[786494] = 1
//line /snap/go/10455/src/strings/replace.go:340
		_go_fuzz_dep_.CoverTab[1071]++

								if i != len(s) && func() bool {
//line /snap/go/10455/src/strings/replace.go:342
			_go_fuzz_dep_.CoverTab[1074]++
//line /snap/go/10455/src/strings/replace.go:342
			return r.root.priority == 0
//line /snap/go/10455/src/strings/replace.go:342
			// _ = "end of CoverTab[1074]"
//line /snap/go/10455/src/strings/replace.go:342
		}() {
//line /snap/go/10455/src/strings/replace.go:342
			_go_fuzz_dep_.CoverTab[525021]++
//line /snap/go/10455/src/strings/replace.go:342
			_go_fuzz_dep_.CoverTab[1075]++
									index := int(r.mapping[s[i]])
									if index == r.tableSize || func() bool {
//line /snap/go/10455/src/strings/replace.go:344
				_go_fuzz_dep_.CoverTab[1076]++
//line /snap/go/10455/src/strings/replace.go:344
				return r.root.table[index] == nil
//line /snap/go/10455/src/strings/replace.go:344
				// _ = "end of CoverTab[1076]"
//line /snap/go/10455/src/strings/replace.go:344
			}() {
//line /snap/go/10455/src/strings/replace.go:344
				_go_fuzz_dep_.CoverTab[525023]++
//line /snap/go/10455/src/strings/replace.go:344
				_go_fuzz_dep_.CoverTab[1077]++
										i++
										continue
//line /snap/go/10455/src/strings/replace.go:346
				// _ = "end of CoverTab[1077]"
			} else {
//line /snap/go/10455/src/strings/replace.go:347
				_go_fuzz_dep_.CoverTab[525024]++
//line /snap/go/10455/src/strings/replace.go:347
				_go_fuzz_dep_.CoverTab[1078]++
//line /snap/go/10455/src/strings/replace.go:347
				// _ = "end of CoverTab[1078]"
//line /snap/go/10455/src/strings/replace.go:347
			}
//line /snap/go/10455/src/strings/replace.go:347
			// _ = "end of CoverTab[1075]"
		} else {
//line /snap/go/10455/src/strings/replace.go:348
			_go_fuzz_dep_.CoverTab[525022]++
//line /snap/go/10455/src/strings/replace.go:348
			_go_fuzz_dep_.CoverTab[1079]++
//line /snap/go/10455/src/strings/replace.go:348
			// _ = "end of CoverTab[1079]"
//line /snap/go/10455/src/strings/replace.go:348
		}
//line /snap/go/10455/src/strings/replace.go:348
		// _ = "end of CoverTab[1071]"
//line /snap/go/10455/src/strings/replace.go:348
		_go_fuzz_dep_.CoverTab[1072]++

//line /snap/go/10455/src/strings/replace.go:351
		val, keylen, match := r.lookup(s[i:], prevMatchEmpty)
		prevMatchEmpty = match && func() bool {
//line /snap/go/10455/src/strings/replace.go:352
			_go_fuzz_dep_.CoverTab[1080]++
//line /snap/go/10455/src/strings/replace.go:352
			return keylen == 0
//line /snap/go/10455/src/strings/replace.go:352
			// _ = "end of CoverTab[1080]"
//line /snap/go/10455/src/strings/replace.go:352
		}()
								if match {
//line /snap/go/10455/src/strings/replace.go:353
			_go_fuzz_dep_.CoverTab[525025]++
//line /snap/go/10455/src/strings/replace.go:353
			_go_fuzz_dep_.CoverTab[1081]++
									wn, err = sw.WriteString(s[last:i])
									n += wn
									if err != nil {
//line /snap/go/10455/src/strings/replace.go:356
				_go_fuzz_dep_.CoverTab[525027]++
//line /snap/go/10455/src/strings/replace.go:356
				_go_fuzz_dep_.CoverTab[1084]++
										return
//line /snap/go/10455/src/strings/replace.go:357
				// _ = "end of CoverTab[1084]"
			} else {
//line /snap/go/10455/src/strings/replace.go:358
				_go_fuzz_dep_.CoverTab[525028]++
//line /snap/go/10455/src/strings/replace.go:358
				_go_fuzz_dep_.CoverTab[1085]++
//line /snap/go/10455/src/strings/replace.go:358
				// _ = "end of CoverTab[1085]"
//line /snap/go/10455/src/strings/replace.go:358
			}
//line /snap/go/10455/src/strings/replace.go:358
			// _ = "end of CoverTab[1081]"
//line /snap/go/10455/src/strings/replace.go:358
			_go_fuzz_dep_.CoverTab[1082]++
									wn, err = sw.WriteString(val)
									n += wn
									if err != nil {
//line /snap/go/10455/src/strings/replace.go:361
				_go_fuzz_dep_.CoverTab[525029]++
//line /snap/go/10455/src/strings/replace.go:361
				_go_fuzz_dep_.CoverTab[1086]++
										return
//line /snap/go/10455/src/strings/replace.go:362
				// _ = "end of CoverTab[1086]"
			} else {
//line /snap/go/10455/src/strings/replace.go:363
				_go_fuzz_dep_.CoverTab[525030]++
//line /snap/go/10455/src/strings/replace.go:363
				_go_fuzz_dep_.CoverTab[1087]++
//line /snap/go/10455/src/strings/replace.go:363
				// _ = "end of CoverTab[1087]"
//line /snap/go/10455/src/strings/replace.go:363
			}
//line /snap/go/10455/src/strings/replace.go:363
			// _ = "end of CoverTab[1082]"
//line /snap/go/10455/src/strings/replace.go:363
			_go_fuzz_dep_.CoverTab[1083]++
									i += keylen
									last = i
									continue
//line /snap/go/10455/src/strings/replace.go:366
			// _ = "end of CoverTab[1083]"
		} else {
//line /snap/go/10455/src/strings/replace.go:367
			_go_fuzz_dep_.CoverTab[525026]++
//line /snap/go/10455/src/strings/replace.go:367
			_go_fuzz_dep_.CoverTab[1088]++
//line /snap/go/10455/src/strings/replace.go:367
			// _ = "end of CoverTab[1088]"
//line /snap/go/10455/src/strings/replace.go:367
		}
//line /snap/go/10455/src/strings/replace.go:367
		// _ = "end of CoverTab[1072]"
//line /snap/go/10455/src/strings/replace.go:367
		_go_fuzz_dep_.CoverTab[1073]++
								i++
//line /snap/go/10455/src/strings/replace.go:368
		// _ = "end of CoverTab[1073]"
	}
//line /snap/go/10455/src/strings/replace.go:369
	if _go_fuzz_dep_.CoverTab[786494] == 0 {
//line /snap/go/10455/src/strings/replace.go:369
		_go_fuzz_dep_.CoverTab[525127]++
//line /snap/go/10455/src/strings/replace.go:369
	} else {
//line /snap/go/10455/src/strings/replace.go:369
		_go_fuzz_dep_.CoverTab[525128]++
//line /snap/go/10455/src/strings/replace.go:369
	}
//line /snap/go/10455/src/strings/replace.go:369
	// _ = "end of CoverTab[1068]"
//line /snap/go/10455/src/strings/replace.go:369
	_go_fuzz_dep_.CoverTab[1069]++
							if last != len(s) {
//line /snap/go/10455/src/strings/replace.go:370
		_go_fuzz_dep_.CoverTab[525031]++
//line /snap/go/10455/src/strings/replace.go:370
		_go_fuzz_dep_.CoverTab[1089]++
								wn, err = sw.WriteString(s[last:])
								n += wn
//line /snap/go/10455/src/strings/replace.go:372
		// _ = "end of CoverTab[1089]"
	} else {
//line /snap/go/10455/src/strings/replace.go:373
		_go_fuzz_dep_.CoverTab[525032]++
//line /snap/go/10455/src/strings/replace.go:373
		_go_fuzz_dep_.CoverTab[1090]++
//line /snap/go/10455/src/strings/replace.go:373
		// _ = "end of CoverTab[1090]"
//line /snap/go/10455/src/strings/replace.go:373
	}
//line /snap/go/10455/src/strings/replace.go:373
	// _ = "end of CoverTab[1069]"
//line /snap/go/10455/src/strings/replace.go:373
	_go_fuzz_dep_.CoverTab[1070]++
							return
//line /snap/go/10455/src/strings/replace.go:374
	// _ = "end of CoverTab[1070]"
}

// singleStringReplacer is the implementation that's used when there is only
//line /snap/go/10455/src/strings/replace.go:377
// one string to replace (and that string has more than one byte).
//line /snap/go/10455/src/strings/replace.go:379
type singleStringReplacer struct {
	finder	*stringFinder
	// value is the new string that replaces that pattern when it's found.
	value	string
}

func makeSingleStringReplacer(pattern string, value string) *singleStringReplacer {
//line /snap/go/10455/src/strings/replace.go:385
	_go_fuzz_dep_.CoverTab[1091]++
							return &singleStringReplacer{finder: makeStringFinder(pattern), value: value}
//line /snap/go/10455/src/strings/replace.go:386
	// _ = "end of CoverTab[1091]"
}

func (r *singleStringReplacer) Replace(s string) string {
//line /snap/go/10455/src/strings/replace.go:389
	_go_fuzz_dep_.CoverTab[1092]++
							var buf Builder
							i, matched := 0, false
//line /snap/go/10455/src/strings/replace.go:391
	_go_fuzz_dep_.CoverTab[786495] = 0
							for {
//line /snap/go/10455/src/strings/replace.go:392
		if _go_fuzz_dep_.CoverTab[786495] == 0 {
//line /snap/go/10455/src/strings/replace.go:392
			_go_fuzz_dep_.CoverTab[525129]++
//line /snap/go/10455/src/strings/replace.go:392
		} else {
//line /snap/go/10455/src/strings/replace.go:392
			_go_fuzz_dep_.CoverTab[525130]++
//line /snap/go/10455/src/strings/replace.go:392
		}
//line /snap/go/10455/src/strings/replace.go:392
		_go_fuzz_dep_.CoverTab[786495] = 1
//line /snap/go/10455/src/strings/replace.go:392
		_go_fuzz_dep_.CoverTab[1095]++
								match := r.finder.next(s[i:])
								if match == -1 {
//line /snap/go/10455/src/strings/replace.go:394
			_go_fuzz_dep_.CoverTab[525033]++
//line /snap/go/10455/src/strings/replace.go:394
			_go_fuzz_dep_.CoverTab[1097]++
									break
//line /snap/go/10455/src/strings/replace.go:395
			// _ = "end of CoverTab[1097]"
		} else {
//line /snap/go/10455/src/strings/replace.go:396
			_go_fuzz_dep_.CoverTab[525034]++
//line /snap/go/10455/src/strings/replace.go:396
			_go_fuzz_dep_.CoverTab[1098]++
//line /snap/go/10455/src/strings/replace.go:396
			// _ = "end of CoverTab[1098]"
//line /snap/go/10455/src/strings/replace.go:396
		}
//line /snap/go/10455/src/strings/replace.go:396
		// _ = "end of CoverTab[1095]"
//line /snap/go/10455/src/strings/replace.go:396
		_go_fuzz_dep_.CoverTab[1096]++
								matched = true
								buf.Grow(match + len(r.value))
								buf.WriteString(s[i : i+match])
								buf.WriteString(r.value)
								i += match + len(r.finder.pattern)
//line /snap/go/10455/src/strings/replace.go:401
		// _ = "end of CoverTab[1096]"
	}
//line /snap/go/10455/src/strings/replace.go:402
	// _ = "end of CoverTab[1092]"
//line /snap/go/10455/src/strings/replace.go:402
	_go_fuzz_dep_.CoverTab[1093]++
							if !matched {
//line /snap/go/10455/src/strings/replace.go:403
		_go_fuzz_dep_.CoverTab[525035]++
//line /snap/go/10455/src/strings/replace.go:403
		_go_fuzz_dep_.CoverTab[1099]++
								return s
//line /snap/go/10455/src/strings/replace.go:404
		// _ = "end of CoverTab[1099]"
	} else {
//line /snap/go/10455/src/strings/replace.go:405
		_go_fuzz_dep_.CoverTab[525036]++
//line /snap/go/10455/src/strings/replace.go:405
		_go_fuzz_dep_.CoverTab[1100]++
//line /snap/go/10455/src/strings/replace.go:405
		// _ = "end of CoverTab[1100]"
//line /snap/go/10455/src/strings/replace.go:405
	}
//line /snap/go/10455/src/strings/replace.go:405
	// _ = "end of CoverTab[1093]"
//line /snap/go/10455/src/strings/replace.go:405
	_go_fuzz_dep_.CoverTab[1094]++
							buf.WriteString(s[i:])
							return buf.String()
//line /snap/go/10455/src/strings/replace.go:407
	// _ = "end of CoverTab[1094]"
}

func (r *singleStringReplacer) WriteString(w io.Writer, s string) (n int, err error) {
//line /snap/go/10455/src/strings/replace.go:410
	_go_fuzz_dep_.CoverTab[1101]++
							sw := getStringWriter(w)
							var i, wn int
//line /snap/go/10455/src/strings/replace.go:412
	_go_fuzz_dep_.CoverTab[786496] = 0
							for {
//line /snap/go/10455/src/strings/replace.go:413
		if _go_fuzz_dep_.CoverTab[786496] == 0 {
//line /snap/go/10455/src/strings/replace.go:413
			_go_fuzz_dep_.CoverTab[525133]++
//line /snap/go/10455/src/strings/replace.go:413
		} else {
//line /snap/go/10455/src/strings/replace.go:413
			_go_fuzz_dep_.CoverTab[525134]++
//line /snap/go/10455/src/strings/replace.go:413
		}
//line /snap/go/10455/src/strings/replace.go:413
		_go_fuzz_dep_.CoverTab[786496] = 1
//line /snap/go/10455/src/strings/replace.go:413
		_go_fuzz_dep_.CoverTab[1103]++
								match := r.finder.next(s[i:])
								if match == -1 {
//line /snap/go/10455/src/strings/replace.go:415
			_go_fuzz_dep_.CoverTab[525037]++
//line /snap/go/10455/src/strings/replace.go:415
			_go_fuzz_dep_.CoverTab[1107]++
									break
//line /snap/go/10455/src/strings/replace.go:416
			// _ = "end of CoverTab[1107]"
		} else {
//line /snap/go/10455/src/strings/replace.go:417
			_go_fuzz_dep_.CoverTab[525038]++
//line /snap/go/10455/src/strings/replace.go:417
			_go_fuzz_dep_.CoverTab[1108]++
//line /snap/go/10455/src/strings/replace.go:417
			// _ = "end of CoverTab[1108]"
//line /snap/go/10455/src/strings/replace.go:417
		}
//line /snap/go/10455/src/strings/replace.go:417
		// _ = "end of CoverTab[1103]"
//line /snap/go/10455/src/strings/replace.go:417
		_go_fuzz_dep_.CoverTab[1104]++
								wn, err = sw.WriteString(s[i : i+match])
								n += wn
								if err != nil {
//line /snap/go/10455/src/strings/replace.go:420
			_go_fuzz_dep_.CoverTab[525039]++
//line /snap/go/10455/src/strings/replace.go:420
			_go_fuzz_dep_.CoverTab[1109]++
									return
//line /snap/go/10455/src/strings/replace.go:421
			// _ = "end of CoverTab[1109]"
		} else {
//line /snap/go/10455/src/strings/replace.go:422
			_go_fuzz_dep_.CoverTab[525040]++
//line /snap/go/10455/src/strings/replace.go:422
			_go_fuzz_dep_.CoverTab[1110]++
//line /snap/go/10455/src/strings/replace.go:422
			// _ = "end of CoverTab[1110]"
//line /snap/go/10455/src/strings/replace.go:422
		}
//line /snap/go/10455/src/strings/replace.go:422
		// _ = "end of CoverTab[1104]"
//line /snap/go/10455/src/strings/replace.go:422
		_go_fuzz_dep_.CoverTab[1105]++
								wn, err = sw.WriteString(r.value)
								n += wn
								if err != nil {
//line /snap/go/10455/src/strings/replace.go:425
			_go_fuzz_dep_.CoverTab[525041]++
//line /snap/go/10455/src/strings/replace.go:425
			_go_fuzz_dep_.CoverTab[1111]++
									return
//line /snap/go/10455/src/strings/replace.go:426
			// _ = "end of CoverTab[1111]"
		} else {
//line /snap/go/10455/src/strings/replace.go:427
			_go_fuzz_dep_.CoverTab[525042]++
//line /snap/go/10455/src/strings/replace.go:427
			_go_fuzz_dep_.CoverTab[1112]++
//line /snap/go/10455/src/strings/replace.go:427
			// _ = "end of CoverTab[1112]"
//line /snap/go/10455/src/strings/replace.go:427
		}
//line /snap/go/10455/src/strings/replace.go:427
		// _ = "end of CoverTab[1105]"
//line /snap/go/10455/src/strings/replace.go:427
		_go_fuzz_dep_.CoverTab[1106]++
								i += match + len(r.finder.pattern)
//line /snap/go/10455/src/strings/replace.go:428
		// _ = "end of CoverTab[1106]"
	}
//line /snap/go/10455/src/strings/replace.go:429
	// _ = "end of CoverTab[1101]"
//line /snap/go/10455/src/strings/replace.go:429
	_go_fuzz_dep_.CoverTab[1102]++
							wn, err = sw.WriteString(s[i:])
							n += wn
							return
//line /snap/go/10455/src/strings/replace.go:432
	// _ = "end of CoverTab[1102]"
}

// byteReplacer is the implementation that's used when all the "old"
//line /snap/go/10455/src/strings/replace.go:435
// and "new" values are single ASCII bytes.
//line /snap/go/10455/src/strings/replace.go:435
// The array contains replacement bytes indexed by old byte.
//line /snap/go/10455/src/strings/replace.go:438
type byteReplacer [256]byte

func (r *byteReplacer) Replace(s string) string {
//line /snap/go/10455/src/strings/replace.go:440
	_go_fuzz_dep_.CoverTab[1113]++
							var buf []byte
//line /snap/go/10455/src/strings/replace.go:441
	_go_fuzz_dep_. // lazily allocated
//line /snap/go/10455/src/strings/replace.go:441
	CoverTab[786497] = 0
							for i := 0; i < len(s); i++ {
//line /snap/go/10455/src/strings/replace.go:442
		if _go_fuzz_dep_.CoverTab[786497] == 0 {
//line /snap/go/10455/src/strings/replace.go:442
			_go_fuzz_dep_.CoverTab[525137]++
//line /snap/go/10455/src/strings/replace.go:442
		} else {
//line /snap/go/10455/src/strings/replace.go:442
			_go_fuzz_dep_.CoverTab[525138]++
//line /snap/go/10455/src/strings/replace.go:442
		}
//line /snap/go/10455/src/strings/replace.go:442
		_go_fuzz_dep_.CoverTab[786497] = 1
//line /snap/go/10455/src/strings/replace.go:442
		_go_fuzz_dep_.CoverTab[1116]++
								b := s[i]
								if r[b] != b {
//line /snap/go/10455/src/strings/replace.go:444
			_go_fuzz_dep_.CoverTab[525043]++
//line /snap/go/10455/src/strings/replace.go:444
			_go_fuzz_dep_.CoverTab[1117]++
									if buf == nil {
//line /snap/go/10455/src/strings/replace.go:445
				_go_fuzz_dep_.CoverTab[525045]++
//line /snap/go/10455/src/strings/replace.go:445
				_go_fuzz_dep_.CoverTab[1119]++
										buf = []byte(s)
//line /snap/go/10455/src/strings/replace.go:446
				// _ = "end of CoverTab[1119]"
			} else {
//line /snap/go/10455/src/strings/replace.go:447
				_go_fuzz_dep_.CoverTab[525046]++
//line /snap/go/10455/src/strings/replace.go:447
				_go_fuzz_dep_.CoverTab[1120]++
//line /snap/go/10455/src/strings/replace.go:447
				// _ = "end of CoverTab[1120]"
//line /snap/go/10455/src/strings/replace.go:447
			}
//line /snap/go/10455/src/strings/replace.go:447
			// _ = "end of CoverTab[1117]"
//line /snap/go/10455/src/strings/replace.go:447
			_go_fuzz_dep_.CoverTab[1118]++
									buf[i] = r[b]
//line /snap/go/10455/src/strings/replace.go:448
			// _ = "end of CoverTab[1118]"
		} else {
//line /snap/go/10455/src/strings/replace.go:449
			_go_fuzz_dep_.CoverTab[525044]++
//line /snap/go/10455/src/strings/replace.go:449
			_go_fuzz_dep_.CoverTab[1121]++
//line /snap/go/10455/src/strings/replace.go:449
			// _ = "end of CoverTab[1121]"
//line /snap/go/10455/src/strings/replace.go:449
		}
//line /snap/go/10455/src/strings/replace.go:449
		// _ = "end of CoverTab[1116]"
	}
//line /snap/go/10455/src/strings/replace.go:450
	if _go_fuzz_dep_.CoverTab[786497] == 0 {
//line /snap/go/10455/src/strings/replace.go:450
		_go_fuzz_dep_.CoverTab[525139]++
//line /snap/go/10455/src/strings/replace.go:450
	} else {
//line /snap/go/10455/src/strings/replace.go:450
		_go_fuzz_dep_.CoverTab[525140]++
//line /snap/go/10455/src/strings/replace.go:450
	}
//line /snap/go/10455/src/strings/replace.go:450
	// _ = "end of CoverTab[1113]"
//line /snap/go/10455/src/strings/replace.go:450
	_go_fuzz_dep_.CoverTab[1114]++
							if buf == nil {
//line /snap/go/10455/src/strings/replace.go:451
		_go_fuzz_dep_.CoverTab[525047]++
//line /snap/go/10455/src/strings/replace.go:451
		_go_fuzz_dep_.CoverTab[1122]++
								return s
//line /snap/go/10455/src/strings/replace.go:452
		// _ = "end of CoverTab[1122]"
	} else {
//line /snap/go/10455/src/strings/replace.go:453
		_go_fuzz_dep_.CoverTab[525048]++
//line /snap/go/10455/src/strings/replace.go:453
		_go_fuzz_dep_.CoverTab[1123]++
//line /snap/go/10455/src/strings/replace.go:453
		// _ = "end of CoverTab[1123]"
//line /snap/go/10455/src/strings/replace.go:453
	}
//line /snap/go/10455/src/strings/replace.go:453
	// _ = "end of CoverTab[1114]"
//line /snap/go/10455/src/strings/replace.go:453
	_go_fuzz_dep_.CoverTab[1115]++
							return string(buf)
//line /snap/go/10455/src/strings/replace.go:454
	// _ = "end of CoverTab[1115]"
}

func (r *byteReplacer) WriteString(w io.Writer, s string) (n int, err error) {
//line /snap/go/10455/src/strings/replace.go:457
	_go_fuzz_dep_.CoverTab[1124]++
							sw := getStringWriter(w)
							last := 0
//line /snap/go/10455/src/strings/replace.go:459
	_go_fuzz_dep_.CoverTab[786498] = 0
							for i := 0; i < len(s); i++ {
//line /snap/go/10455/src/strings/replace.go:460
		if _go_fuzz_dep_.CoverTab[786498] == 0 {
//line /snap/go/10455/src/strings/replace.go:460
			_go_fuzz_dep_.CoverTab[525141]++
//line /snap/go/10455/src/strings/replace.go:460
		} else {
//line /snap/go/10455/src/strings/replace.go:460
			_go_fuzz_dep_.CoverTab[525142]++
//line /snap/go/10455/src/strings/replace.go:460
		}
//line /snap/go/10455/src/strings/replace.go:460
		_go_fuzz_dep_.CoverTab[786498] = 1
//line /snap/go/10455/src/strings/replace.go:460
		_go_fuzz_dep_.CoverTab[1127]++
								b := s[i]
								if r[b] == b {
//line /snap/go/10455/src/strings/replace.go:462
			_go_fuzz_dep_.CoverTab[525049]++
//line /snap/go/10455/src/strings/replace.go:462
			_go_fuzz_dep_.CoverTab[1130]++
									continue
//line /snap/go/10455/src/strings/replace.go:463
			// _ = "end of CoverTab[1130]"
		} else {
//line /snap/go/10455/src/strings/replace.go:464
			_go_fuzz_dep_.CoverTab[525050]++
//line /snap/go/10455/src/strings/replace.go:464
			_go_fuzz_dep_.CoverTab[1131]++
//line /snap/go/10455/src/strings/replace.go:464
			// _ = "end of CoverTab[1131]"
//line /snap/go/10455/src/strings/replace.go:464
		}
//line /snap/go/10455/src/strings/replace.go:464
		// _ = "end of CoverTab[1127]"
//line /snap/go/10455/src/strings/replace.go:464
		_go_fuzz_dep_.CoverTab[1128]++
								if last != i {
//line /snap/go/10455/src/strings/replace.go:465
			_go_fuzz_dep_.CoverTab[525051]++
//line /snap/go/10455/src/strings/replace.go:465
			_go_fuzz_dep_.CoverTab[1132]++
									wn, err := sw.WriteString(s[last:i])
									n += wn
									if err != nil {
//line /snap/go/10455/src/strings/replace.go:468
				_go_fuzz_dep_.CoverTab[525053]++
//line /snap/go/10455/src/strings/replace.go:468
				_go_fuzz_dep_.CoverTab[1133]++
										return n, err
//line /snap/go/10455/src/strings/replace.go:469
				// _ = "end of CoverTab[1133]"
			} else {
//line /snap/go/10455/src/strings/replace.go:470
				_go_fuzz_dep_.CoverTab[525054]++
//line /snap/go/10455/src/strings/replace.go:470
				_go_fuzz_dep_.CoverTab[1134]++
//line /snap/go/10455/src/strings/replace.go:470
				// _ = "end of CoverTab[1134]"
//line /snap/go/10455/src/strings/replace.go:470
			}
//line /snap/go/10455/src/strings/replace.go:470
			// _ = "end of CoverTab[1132]"
		} else {
//line /snap/go/10455/src/strings/replace.go:471
			_go_fuzz_dep_.CoverTab[525052]++
//line /snap/go/10455/src/strings/replace.go:471
			_go_fuzz_dep_.CoverTab[1135]++
//line /snap/go/10455/src/strings/replace.go:471
			// _ = "end of CoverTab[1135]"
//line /snap/go/10455/src/strings/replace.go:471
		}
//line /snap/go/10455/src/strings/replace.go:471
		// _ = "end of CoverTab[1128]"
//line /snap/go/10455/src/strings/replace.go:471
		_go_fuzz_dep_.CoverTab[1129]++
								last = i + 1
								nw, err := w.Write(r[b : int(b)+1])
								n += nw
								if err != nil {
//line /snap/go/10455/src/strings/replace.go:475
			_go_fuzz_dep_.CoverTab[525055]++
//line /snap/go/10455/src/strings/replace.go:475
			_go_fuzz_dep_.CoverTab[1136]++
									return n, err
//line /snap/go/10455/src/strings/replace.go:476
			// _ = "end of CoverTab[1136]"
		} else {
//line /snap/go/10455/src/strings/replace.go:477
			_go_fuzz_dep_.CoverTab[525056]++
//line /snap/go/10455/src/strings/replace.go:477
			_go_fuzz_dep_.CoverTab[1137]++
//line /snap/go/10455/src/strings/replace.go:477
			// _ = "end of CoverTab[1137]"
//line /snap/go/10455/src/strings/replace.go:477
		}
//line /snap/go/10455/src/strings/replace.go:477
		// _ = "end of CoverTab[1129]"
	}
//line /snap/go/10455/src/strings/replace.go:478
	if _go_fuzz_dep_.CoverTab[786498] == 0 {
//line /snap/go/10455/src/strings/replace.go:478
		_go_fuzz_dep_.CoverTab[525143]++
//line /snap/go/10455/src/strings/replace.go:478
	} else {
//line /snap/go/10455/src/strings/replace.go:478
		_go_fuzz_dep_.CoverTab[525144]++
//line /snap/go/10455/src/strings/replace.go:478
	}
//line /snap/go/10455/src/strings/replace.go:478
	// _ = "end of CoverTab[1124]"
//line /snap/go/10455/src/strings/replace.go:478
	_go_fuzz_dep_.CoverTab[1125]++
							if last != len(s) {
//line /snap/go/10455/src/strings/replace.go:479
		_go_fuzz_dep_.CoverTab[525057]++
//line /snap/go/10455/src/strings/replace.go:479
		_go_fuzz_dep_.CoverTab[1138]++
								nw, err := sw.WriteString(s[last:])
								n += nw
								if err != nil {
//line /snap/go/10455/src/strings/replace.go:482
			_go_fuzz_dep_.CoverTab[525059]++
//line /snap/go/10455/src/strings/replace.go:482
			_go_fuzz_dep_.CoverTab[1139]++
									return n, err
//line /snap/go/10455/src/strings/replace.go:483
			// _ = "end of CoverTab[1139]"
		} else {
//line /snap/go/10455/src/strings/replace.go:484
			_go_fuzz_dep_.CoverTab[525060]++
//line /snap/go/10455/src/strings/replace.go:484
			_go_fuzz_dep_.CoverTab[1140]++
//line /snap/go/10455/src/strings/replace.go:484
			// _ = "end of CoverTab[1140]"
//line /snap/go/10455/src/strings/replace.go:484
		}
//line /snap/go/10455/src/strings/replace.go:484
		// _ = "end of CoverTab[1138]"
	} else {
//line /snap/go/10455/src/strings/replace.go:485
		_go_fuzz_dep_.CoverTab[525058]++
//line /snap/go/10455/src/strings/replace.go:485
		_go_fuzz_dep_.CoverTab[1141]++
//line /snap/go/10455/src/strings/replace.go:485
		// _ = "end of CoverTab[1141]"
//line /snap/go/10455/src/strings/replace.go:485
	}
//line /snap/go/10455/src/strings/replace.go:485
	// _ = "end of CoverTab[1125]"
//line /snap/go/10455/src/strings/replace.go:485
	_go_fuzz_dep_.CoverTab[1126]++
							return n, nil
//line /snap/go/10455/src/strings/replace.go:486
	// _ = "end of CoverTab[1126]"
}

// byteStringReplacer is the implementation that's used when all the
//line /snap/go/10455/src/strings/replace.go:489
// "old" values are single ASCII bytes but the "new" values vary in size.
//line /snap/go/10455/src/strings/replace.go:491
type byteStringReplacer struct {
	// replacements contains replacement byte slices indexed by old byte.
	// A nil []byte means that the old byte should not be replaced.
	replacements	[256][]byte
	// toReplace keeps a list of bytes to replace. Depending on length of toReplace
	// and length of target string it may be faster to use Count, or a plain loop.
	// We store single byte as a string, because Count takes a string.
	toReplace	[]string
}

// countCutOff controls the ratio of a string length to a number of replacements
//line /snap/go/10455/src/strings/replace.go:501
// at which (*byteStringReplacer).Replace switches algorithms.
//line /snap/go/10455/src/strings/replace.go:501
// For strings with higher ration of length to replacements than that value,
//line /snap/go/10455/src/strings/replace.go:501
// we call Count, for each replacement from toReplace.
//line /snap/go/10455/src/strings/replace.go:501
// For strings, with a lower ratio we use simple loop, because of Count overhead.
//line /snap/go/10455/src/strings/replace.go:501
// countCutOff is an empirically determined overhead multiplier.
//line /snap/go/10455/src/strings/replace.go:501
// TODO(tocarip) revisit once we have register-based abi/mid-stack inlining.
//line /snap/go/10455/src/strings/replace.go:508
const countCutOff = 8

func (r *byteStringReplacer) Replace(s string) string {
//line /snap/go/10455/src/strings/replace.go:510
	_go_fuzz_dep_.CoverTab[1142]++
							newSize := len(s)
							anyChanges := false

							if len(r.toReplace)*countCutOff <= len(s) {
//line /snap/go/10455/src/strings/replace.go:514
		_go_fuzz_dep_.CoverTab[525061]++
//line /snap/go/10455/src/strings/replace.go:514
		_go_fuzz_dep_.CoverTab[1146]++
//line /snap/go/10455/src/strings/replace.go:514
		_go_fuzz_dep_.CoverTab[786500] = 0
								for _, x := range r.toReplace {
//line /snap/go/10455/src/strings/replace.go:515
			if _go_fuzz_dep_.CoverTab[786500] == 0 {
//line /snap/go/10455/src/strings/replace.go:515
				_go_fuzz_dep_.CoverTab[525149]++
//line /snap/go/10455/src/strings/replace.go:515
			} else {
//line /snap/go/10455/src/strings/replace.go:515
				_go_fuzz_dep_.CoverTab[525150]++
//line /snap/go/10455/src/strings/replace.go:515
			}
//line /snap/go/10455/src/strings/replace.go:515
			_go_fuzz_dep_.CoverTab[786500] = 1
//line /snap/go/10455/src/strings/replace.go:515
			_go_fuzz_dep_.CoverTab[1147]++
									if c := Count(s, x); c != 0 {
//line /snap/go/10455/src/strings/replace.go:516
				_go_fuzz_dep_.CoverTab[525063]++
//line /snap/go/10455/src/strings/replace.go:516
				_go_fuzz_dep_.CoverTab[1148]++

										newSize += c * (len(r.replacements[x[0]]) - 1)
										anyChanges = true
//line /snap/go/10455/src/strings/replace.go:519
				// _ = "end of CoverTab[1148]"
			} else {
//line /snap/go/10455/src/strings/replace.go:520
				_go_fuzz_dep_.CoverTab[525064]++
//line /snap/go/10455/src/strings/replace.go:520
				_go_fuzz_dep_.CoverTab[1149]++
//line /snap/go/10455/src/strings/replace.go:520
				// _ = "end of CoverTab[1149]"
//line /snap/go/10455/src/strings/replace.go:520
			}
//line /snap/go/10455/src/strings/replace.go:520
			// _ = "end of CoverTab[1147]"

		}
//line /snap/go/10455/src/strings/replace.go:522
		if _go_fuzz_dep_.CoverTab[786500] == 0 {
//line /snap/go/10455/src/strings/replace.go:522
			_go_fuzz_dep_.CoverTab[525151]++
//line /snap/go/10455/src/strings/replace.go:522
		} else {
//line /snap/go/10455/src/strings/replace.go:522
			_go_fuzz_dep_.CoverTab[525152]++
//line /snap/go/10455/src/strings/replace.go:522
		}
//line /snap/go/10455/src/strings/replace.go:522
		// _ = "end of CoverTab[1146]"
	} else {
//line /snap/go/10455/src/strings/replace.go:523
		_go_fuzz_dep_.CoverTab[525062]++
//line /snap/go/10455/src/strings/replace.go:523
		_go_fuzz_dep_.CoverTab[1150]++
//line /snap/go/10455/src/strings/replace.go:523
		_go_fuzz_dep_.CoverTab[786501] = 0
								for i := 0; i < len(s); i++ {
//line /snap/go/10455/src/strings/replace.go:524
			if _go_fuzz_dep_.CoverTab[786501] == 0 {
//line /snap/go/10455/src/strings/replace.go:524
				_go_fuzz_dep_.CoverTab[525153]++
//line /snap/go/10455/src/strings/replace.go:524
			} else {
//line /snap/go/10455/src/strings/replace.go:524
				_go_fuzz_dep_.CoverTab[525154]++
//line /snap/go/10455/src/strings/replace.go:524
			}
//line /snap/go/10455/src/strings/replace.go:524
			_go_fuzz_dep_.CoverTab[786501] = 1
//line /snap/go/10455/src/strings/replace.go:524
			_go_fuzz_dep_.CoverTab[1151]++
									b := s[i]
									if r.replacements[b] != nil {
//line /snap/go/10455/src/strings/replace.go:526
				_go_fuzz_dep_.CoverTab[525065]++
//line /snap/go/10455/src/strings/replace.go:526
				_go_fuzz_dep_.CoverTab[1152]++

										newSize += len(r.replacements[b]) - 1
										anyChanges = true
//line /snap/go/10455/src/strings/replace.go:529
				// _ = "end of CoverTab[1152]"
			} else {
//line /snap/go/10455/src/strings/replace.go:530
				_go_fuzz_dep_.CoverTab[525066]++
//line /snap/go/10455/src/strings/replace.go:530
				_go_fuzz_dep_.CoverTab[1153]++
//line /snap/go/10455/src/strings/replace.go:530
				// _ = "end of CoverTab[1153]"
//line /snap/go/10455/src/strings/replace.go:530
			}
//line /snap/go/10455/src/strings/replace.go:530
			// _ = "end of CoverTab[1151]"
		}
//line /snap/go/10455/src/strings/replace.go:531
		if _go_fuzz_dep_.CoverTab[786501] == 0 {
//line /snap/go/10455/src/strings/replace.go:531
			_go_fuzz_dep_.CoverTab[525155]++
//line /snap/go/10455/src/strings/replace.go:531
		} else {
//line /snap/go/10455/src/strings/replace.go:531
			_go_fuzz_dep_.CoverTab[525156]++
//line /snap/go/10455/src/strings/replace.go:531
		}
//line /snap/go/10455/src/strings/replace.go:531
		// _ = "end of CoverTab[1150]"
	}
//line /snap/go/10455/src/strings/replace.go:532
	// _ = "end of CoverTab[1142]"
//line /snap/go/10455/src/strings/replace.go:532
	_go_fuzz_dep_.CoverTab[1143]++
							if !anyChanges {
//line /snap/go/10455/src/strings/replace.go:533
		_go_fuzz_dep_.CoverTab[525067]++
//line /snap/go/10455/src/strings/replace.go:533
		_go_fuzz_dep_.CoverTab[1154]++
								return s
//line /snap/go/10455/src/strings/replace.go:534
		// _ = "end of CoverTab[1154]"
	} else {
//line /snap/go/10455/src/strings/replace.go:535
		_go_fuzz_dep_.CoverTab[525068]++
//line /snap/go/10455/src/strings/replace.go:535
		_go_fuzz_dep_.CoverTab[1155]++
//line /snap/go/10455/src/strings/replace.go:535
		// _ = "end of CoverTab[1155]"
//line /snap/go/10455/src/strings/replace.go:535
	}
//line /snap/go/10455/src/strings/replace.go:535
	// _ = "end of CoverTab[1143]"
//line /snap/go/10455/src/strings/replace.go:535
	_go_fuzz_dep_.CoverTab[1144]++
							buf := make([]byte, newSize)
							j := 0
//line /snap/go/10455/src/strings/replace.go:537
	_go_fuzz_dep_.CoverTab[786499] = 0
							for i := 0; i < len(s); i++ {
//line /snap/go/10455/src/strings/replace.go:538
		if _go_fuzz_dep_.CoverTab[786499] == 0 {
//line /snap/go/10455/src/strings/replace.go:538
			_go_fuzz_dep_.CoverTab[525145]++
//line /snap/go/10455/src/strings/replace.go:538
		} else {
//line /snap/go/10455/src/strings/replace.go:538
			_go_fuzz_dep_.CoverTab[525146]++
//line /snap/go/10455/src/strings/replace.go:538
		}
//line /snap/go/10455/src/strings/replace.go:538
		_go_fuzz_dep_.CoverTab[786499] = 1
//line /snap/go/10455/src/strings/replace.go:538
		_go_fuzz_dep_.CoverTab[1156]++
								b := s[i]
								if r.replacements[b] != nil {
//line /snap/go/10455/src/strings/replace.go:540
			_go_fuzz_dep_.CoverTab[525069]++
//line /snap/go/10455/src/strings/replace.go:540
			_go_fuzz_dep_.CoverTab[1157]++
									j += copy(buf[j:], r.replacements[b])
//line /snap/go/10455/src/strings/replace.go:541
			// _ = "end of CoverTab[1157]"
		} else {
//line /snap/go/10455/src/strings/replace.go:542
			_go_fuzz_dep_.CoverTab[525070]++
//line /snap/go/10455/src/strings/replace.go:542
			_go_fuzz_dep_.CoverTab[1158]++
									buf[j] = b
									j++
//line /snap/go/10455/src/strings/replace.go:544
			// _ = "end of CoverTab[1158]"
		}
//line /snap/go/10455/src/strings/replace.go:545
		// _ = "end of CoverTab[1156]"
	}
//line /snap/go/10455/src/strings/replace.go:546
	if _go_fuzz_dep_.CoverTab[786499] == 0 {
//line /snap/go/10455/src/strings/replace.go:546
		_go_fuzz_dep_.CoverTab[525147]++
//line /snap/go/10455/src/strings/replace.go:546
	} else {
//line /snap/go/10455/src/strings/replace.go:546
		_go_fuzz_dep_.CoverTab[525148]++
//line /snap/go/10455/src/strings/replace.go:546
	}
//line /snap/go/10455/src/strings/replace.go:546
	// _ = "end of CoverTab[1144]"
//line /snap/go/10455/src/strings/replace.go:546
	_go_fuzz_dep_.CoverTab[1145]++
							return string(buf)
//line /snap/go/10455/src/strings/replace.go:547
	// _ = "end of CoverTab[1145]"
}

func (r *byteStringReplacer) WriteString(w io.Writer, s string) (n int, err error) {
//line /snap/go/10455/src/strings/replace.go:550
	_go_fuzz_dep_.CoverTab[1159]++
							sw := getStringWriter(w)
							last := 0
//line /snap/go/10455/src/strings/replace.go:552
	_go_fuzz_dep_.CoverTab[786502] = 0
							for i := 0; i < len(s); i++ {
//line /snap/go/10455/src/strings/replace.go:553
		if _go_fuzz_dep_.CoverTab[786502] == 0 {
//line /snap/go/10455/src/strings/replace.go:553
			_go_fuzz_dep_.CoverTab[525157]++
//line /snap/go/10455/src/strings/replace.go:553
		} else {
//line /snap/go/10455/src/strings/replace.go:553
			_go_fuzz_dep_.CoverTab[525158]++
//line /snap/go/10455/src/strings/replace.go:553
		}
//line /snap/go/10455/src/strings/replace.go:553
		_go_fuzz_dep_.CoverTab[786502] = 1
//line /snap/go/10455/src/strings/replace.go:553
		_go_fuzz_dep_.CoverTab[1162]++
								b := s[i]
								if r.replacements[b] == nil {
//line /snap/go/10455/src/strings/replace.go:555
			_go_fuzz_dep_.CoverTab[525071]++
//line /snap/go/10455/src/strings/replace.go:555
			_go_fuzz_dep_.CoverTab[1165]++
									continue
//line /snap/go/10455/src/strings/replace.go:556
			// _ = "end of CoverTab[1165]"
		} else {
//line /snap/go/10455/src/strings/replace.go:557
			_go_fuzz_dep_.CoverTab[525072]++
//line /snap/go/10455/src/strings/replace.go:557
			_go_fuzz_dep_.CoverTab[1166]++
//line /snap/go/10455/src/strings/replace.go:557
			// _ = "end of CoverTab[1166]"
//line /snap/go/10455/src/strings/replace.go:557
		}
//line /snap/go/10455/src/strings/replace.go:557
		// _ = "end of CoverTab[1162]"
//line /snap/go/10455/src/strings/replace.go:557
		_go_fuzz_dep_.CoverTab[1163]++
								if last != i {
//line /snap/go/10455/src/strings/replace.go:558
			_go_fuzz_dep_.CoverTab[525073]++
//line /snap/go/10455/src/strings/replace.go:558
			_go_fuzz_dep_.CoverTab[1167]++
									nw, err := sw.WriteString(s[last:i])
									n += nw
									if err != nil {
//line /snap/go/10455/src/strings/replace.go:561
				_go_fuzz_dep_.CoverTab[525075]++
//line /snap/go/10455/src/strings/replace.go:561
				_go_fuzz_dep_.CoverTab[1168]++
										return n, err
//line /snap/go/10455/src/strings/replace.go:562
				// _ = "end of CoverTab[1168]"
			} else {
//line /snap/go/10455/src/strings/replace.go:563
				_go_fuzz_dep_.CoverTab[525076]++
//line /snap/go/10455/src/strings/replace.go:563
				_go_fuzz_dep_.CoverTab[1169]++
//line /snap/go/10455/src/strings/replace.go:563
				// _ = "end of CoverTab[1169]"
//line /snap/go/10455/src/strings/replace.go:563
			}
//line /snap/go/10455/src/strings/replace.go:563
			// _ = "end of CoverTab[1167]"
		} else {
//line /snap/go/10455/src/strings/replace.go:564
			_go_fuzz_dep_.CoverTab[525074]++
//line /snap/go/10455/src/strings/replace.go:564
			_go_fuzz_dep_.CoverTab[1170]++
//line /snap/go/10455/src/strings/replace.go:564
			// _ = "end of CoverTab[1170]"
//line /snap/go/10455/src/strings/replace.go:564
		}
//line /snap/go/10455/src/strings/replace.go:564
		// _ = "end of CoverTab[1163]"
//line /snap/go/10455/src/strings/replace.go:564
		_go_fuzz_dep_.CoverTab[1164]++
								last = i + 1
								nw, err := w.Write(r.replacements[b])
								n += nw
								if err != nil {
//line /snap/go/10455/src/strings/replace.go:568
			_go_fuzz_dep_.CoverTab[525077]++
//line /snap/go/10455/src/strings/replace.go:568
			_go_fuzz_dep_.CoverTab[1171]++
									return n, err
//line /snap/go/10455/src/strings/replace.go:569
			// _ = "end of CoverTab[1171]"
		} else {
//line /snap/go/10455/src/strings/replace.go:570
			_go_fuzz_dep_.CoverTab[525078]++
//line /snap/go/10455/src/strings/replace.go:570
			_go_fuzz_dep_.CoverTab[1172]++
//line /snap/go/10455/src/strings/replace.go:570
			// _ = "end of CoverTab[1172]"
//line /snap/go/10455/src/strings/replace.go:570
		}
//line /snap/go/10455/src/strings/replace.go:570
		// _ = "end of CoverTab[1164]"
	}
//line /snap/go/10455/src/strings/replace.go:571
	if _go_fuzz_dep_.CoverTab[786502] == 0 {
//line /snap/go/10455/src/strings/replace.go:571
		_go_fuzz_dep_.CoverTab[525159]++
//line /snap/go/10455/src/strings/replace.go:571
	} else {
//line /snap/go/10455/src/strings/replace.go:571
		_go_fuzz_dep_.CoverTab[525160]++
//line /snap/go/10455/src/strings/replace.go:571
	}
//line /snap/go/10455/src/strings/replace.go:571
	// _ = "end of CoverTab[1159]"
//line /snap/go/10455/src/strings/replace.go:571
	_go_fuzz_dep_.CoverTab[1160]++
							if last != len(s) {
//line /snap/go/10455/src/strings/replace.go:572
		_go_fuzz_dep_.CoverTab[525079]++
//line /snap/go/10455/src/strings/replace.go:572
		_go_fuzz_dep_.CoverTab[1173]++
								var nw int
								nw, err = sw.WriteString(s[last:])
								n += nw
//line /snap/go/10455/src/strings/replace.go:575
		// _ = "end of CoverTab[1173]"
	} else {
//line /snap/go/10455/src/strings/replace.go:576
		_go_fuzz_dep_.CoverTab[525080]++
//line /snap/go/10455/src/strings/replace.go:576
		_go_fuzz_dep_.CoverTab[1174]++
//line /snap/go/10455/src/strings/replace.go:576
		// _ = "end of CoverTab[1174]"
//line /snap/go/10455/src/strings/replace.go:576
	}
//line /snap/go/10455/src/strings/replace.go:576
	// _ = "end of CoverTab[1160]"
//line /snap/go/10455/src/strings/replace.go:576
	_go_fuzz_dep_.CoverTab[1161]++
							return
//line /snap/go/10455/src/strings/replace.go:577
	// _ = "end of CoverTab[1161]"
}

//line /snap/go/10455/src/strings/replace.go:578
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /snap/go/10455/src/strings/replace.go:578
var _ = _go_fuzz_dep_.CoverTab
