// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/strings/replace.go:5
package strings

//line /usr/local/go/src/strings/replace.go:5
import (
//line /usr/local/go/src/strings/replace.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/strings/replace.go:5
)
//line /usr/local/go/src/strings/replace.go:5
import (
//line /usr/local/go/src/strings/replace.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/strings/replace.go:5
)

import (
	"io"
	"sync"
)

// Replacer replaces a list of strings with replacements.
//line /usr/local/go/src/strings/replace.go:12
// It is safe for concurrent use by multiple goroutines.
//line /usr/local/go/src/strings/replace.go:14
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
//line /usr/local/go/src/strings/replace.go:26
// pairs. Replacements are performed in the order they appear in the
//line /usr/local/go/src/strings/replace.go:26
// target string, without overlapping matches. The old string
//line /usr/local/go/src/strings/replace.go:26
// comparisons are done in argument order.
//line /usr/local/go/src/strings/replace.go:26
//
//line /usr/local/go/src/strings/replace.go:26
// NewReplacer panics if given an odd number of arguments.
//line /usr/local/go/src/strings/replace.go:32
func NewReplacer(oldnew ...string) *Replacer {
//line /usr/local/go/src/strings/replace.go:32
	_go_fuzz_dep_.CoverTab[959]++
						if len(oldnew)%2 == 1 {
//line /usr/local/go/src/strings/replace.go:33
		_go_fuzz_dep_.CoverTab[961]++
							panic("strings.NewReplacer: odd argument count")
//line /usr/local/go/src/strings/replace.go:34
		// _ = "end of CoverTab[961]"
	} else {
//line /usr/local/go/src/strings/replace.go:35
		_go_fuzz_dep_.CoverTab[962]++
//line /usr/local/go/src/strings/replace.go:35
		// _ = "end of CoverTab[962]"
//line /usr/local/go/src/strings/replace.go:35
	}
//line /usr/local/go/src/strings/replace.go:35
	// _ = "end of CoverTab[959]"
//line /usr/local/go/src/strings/replace.go:35
	_go_fuzz_dep_.CoverTab[960]++
						return &Replacer{oldnew: append([]string(nil), oldnew...)}
//line /usr/local/go/src/strings/replace.go:36
	// _ = "end of CoverTab[960]"
}

func (r *Replacer) buildOnce() {
//line /usr/local/go/src/strings/replace.go:39
	_go_fuzz_dep_.CoverTab[963]++
						r.r = r.build()
						r.oldnew = nil
//line /usr/local/go/src/strings/replace.go:41
	// _ = "end of CoverTab[963]"
}

func (b *Replacer) build() replacer {
//line /usr/local/go/src/strings/replace.go:44
	_go_fuzz_dep_.CoverTab[964]++
						oldnew := b.oldnew
						if len(oldnew) == 2 && func() bool {
//line /usr/local/go/src/strings/replace.go:46
		_go_fuzz_dep_.CoverTab[969]++
//line /usr/local/go/src/strings/replace.go:46
		return len(oldnew[0]) > 1
//line /usr/local/go/src/strings/replace.go:46
		// _ = "end of CoverTab[969]"
//line /usr/local/go/src/strings/replace.go:46
	}() {
//line /usr/local/go/src/strings/replace.go:46
		_go_fuzz_dep_.CoverTab[970]++
							return makeSingleStringReplacer(oldnew[0], oldnew[1])
//line /usr/local/go/src/strings/replace.go:47
		// _ = "end of CoverTab[970]"
	} else {
//line /usr/local/go/src/strings/replace.go:48
		_go_fuzz_dep_.CoverTab[971]++
//line /usr/local/go/src/strings/replace.go:48
		// _ = "end of CoverTab[971]"
//line /usr/local/go/src/strings/replace.go:48
	}
//line /usr/local/go/src/strings/replace.go:48
	// _ = "end of CoverTab[964]"
//line /usr/local/go/src/strings/replace.go:48
	_go_fuzz_dep_.CoverTab[965]++

						allNewBytes := true
						for i := 0; i < len(oldnew); i += 2 {
//line /usr/local/go/src/strings/replace.go:51
		_go_fuzz_dep_.CoverTab[972]++
							if len(oldnew[i]) != 1 {
//line /usr/local/go/src/strings/replace.go:52
			_go_fuzz_dep_.CoverTab[974]++
								return makeGenericReplacer(oldnew)
//line /usr/local/go/src/strings/replace.go:53
			// _ = "end of CoverTab[974]"
		} else {
//line /usr/local/go/src/strings/replace.go:54
			_go_fuzz_dep_.CoverTab[975]++
//line /usr/local/go/src/strings/replace.go:54
			// _ = "end of CoverTab[975]"
//line /usr/local/go/src/strings/replace.go:54
		}
//line /usr/local/go/src/strings/replace.go:54
		// _ = "end of CoverTab[972]"
//line /usr/local/go/src/strings/replace.go:54
		_go_fuzz_dep_.CoverTab[973]++
							if len(oldnew[i+1]) != 1 {
//line /usr/local/go/src/strings/replace.go:55
			_go_fuzz_dep_.CoverTab[976]++
								allNewBytes = false
//line /usr/local/go/src/strings/replace.go:56
			// _ = "end of CoverTab[976]"
		} else {
//line /usr/local/go/src/strings/replace.go:57
			_go_fuzz_dep_.CoverTab[977]++
//line /usr/local/go/src/strings/replace.go:57
			// _ = "end of CoverTab[977]"
//line /usr/local/go/src/strings/replace.go:57
		}
//line /usr/local/go/src/strings/replace.go:57
		// _ = "end of CoverTab[973]"
	}
//line /usr/local/go/src/strings/replace.go:58
	// _ = "end of CoverTab[965]"
//line /usr/local/go/src/strings/replace.go:58
	_go_fuzz_dep_.CoverTab[966]++

						if allNewBytes {
//line /usr/local/go/src/strings/replace.go:60
		_go_fuzz_dep_.CoverTab[978]++
							r := byteReplacer{}
							for i := range r {
//line /usr/local/go/src/strings/replace.go:62
			_go_fuzz_dep_.CoverTab[981]++
								r[i] = byte(i)
//line /usr/local/go/src/strings/replace.go:63
			// _ = "end of CoverTab[981]"
		}
//line /usr/local/go/src/strings/replace.go:64
		// _ = "end of CoverTab[978]"
//line /usr/local/go/src/strings/replace.go:64
		_go_fuzz_dep_.CoverTab[979]++

//line /usr/local/go/src/strings/replace.go:67
		for i := len(oldnew) - 2; i >= 0; i -= 2 {
//line /usr/local/go/src/strings/replace.go:67
			_go_fuzz_dep_.CoverTab[982]++
								o := oldnew[i][0]
								n := oldnew[i+1][0]
								r[o] = n
//line /usr/local/go/src/strings/replace.go:70
			// _ = "end of CoverTab[982]"
		}
//line /usr/local/go/src/strings/replace.go:71
		// _ = "end of CoverTab[979]"
//line /usr/local/go/src/strings/replace.go:71
		_go_fuzz_dep_.CoverTab[980]++
							return &r
//line /usr/local/go/src/strings/replace.go:72
		// _ = "end of CoverTab[980]"
	} else {
//line /usr/local/go/src/strings/replace.go:73
		_go_fuzz_dep_.CoverTab[983]++
//line /usr/local/go/src/strings/replace.go:73
		// _ = "end of CoverTab[983]"
//line /usr/local/go/src/strings/replace.go:73
	}
//line /usr/local/go/src/strings/replace.go:73
	// _ = "end of CoverTab[966]"
//line /usr/local/go/src/strings/replace.go:73
	_go_fuzz_dep_.CoverTab[967]++

						r := byteStringReplacer{toReplace: make([]string, 0, len(oldnew)/2)}

//line /usr/local/go/src/strings/replace.go:78
	for i := len(oldnew) - 2; i >= 0; i -= 2 {
//line /usr/local/go/src/strings/replace.go:78
		_go_fuzz_dep_.CoverTab[984]++
							o := oldnew[i][0]
							n := oldnew[i+1]

							if r.replacements[o] == nil {
//line /usr/local/go/src/strings/replace.go:82
			_go_fuzz_dep_.CoverTab[986]++

//line /usr/local/go/src/strings/replace.go:86
			r.toReplace = append(r.toReplace, string([]byte{o}))
//line /usr/local/go/src/strings/replace.go:86
			// _ = "end of CoverTab[986]"
		} else {
//line /usr/local/go/src/strings/replace.go:87
			_go_fuzz_dep_.CoverTab[987]++
//line /usr/local/go/src/strings/replace.go:87
			// _ = "end of CoverTab[987]"
//line /usr/local/go/src/strings/replace.go:87
		}
//line /usr/local/go/src/strings/replace.go:87
		// _ = "end of CoverTab[984]"
//line /usr/local/go/src/strings/replace.go:87
		_go_fuzz_dep_.CoverTab[985]++
							r.replacements[o] = []byte(n)
//line /usr/local/go/src/strings/replace.go:88
		// _ = "end of CoverTab[985]"

	}
//line /usr/local/go/src/strings/replace.go:90
	// _ = "end of CoverTab[967]"
//line /usr/local/go/src/strings/replace.go:90
	_go_fuzz_dep_.CoverTab[968]++
						return &r
//line /usr/local/go/src/strings/replace.go:91
	// _ = "end of CoverTab[968]"
}

// Replace returns a copy of s with all replacements performed.
func (r *Replacer) Replace(s string) string {
//line /usr/local/go/src/strings/replace.go:95
	_go_fuzz_dep_.CoverTab[988]++
						r.once.Do(r.buildOnce)
						return r.r.Replace(s)
//line /usr/local/go/src/strings/replace.go:97
	// _ = "end of CoverTab[988]"
}

// WriteString writes s to w with all replacements performed.
func (r *Replacer) WriteString(w io.Writer, s string) (n int, err error) {
//line /usr/local/go/src/strings/replace.go:101
	_go_fuzz_dep_.CoverTab[989]++
							r.once.Do(r.buildOnce)
							return r.r.WriteString(w, s)
//line /usr/local/go/src/strings/replace.go:103
	// _ = "end of CoverTab[989]"
}

// trieNode is a node in a lookup trie for prioritized key/value pairs. Keys
//line /usr/local/go/src/strings/replace.go:106
// and values may be empty. For example, the trie containing keys "ax", "ay",
//line /usr/local/go/src/strings/replace.go:106
// "bcbc", "x" and "xy" could have eight nodes:
//line /usr/local/go/src/strings/replace.go:106
//
//line /usr/local/go/src/strings/replace.go:106
//	n0  -
//line /usr/local/go/src/strings/replace.go:106
//	n1  a-
//line /usr/local/go/src/strings/replace.go:106
//	n2  .x+
//line /usr/local/go/src/strings/replace.go:106
//	n3  .y+
//line /usr/local/go/src/strings/replace.go:106
//	n4  b-
//line /usr/local/go/src/strings/replace.go:106
//	n5  .cbc+
//line /usr/local/go/src/strings/replace.go:106
//	n6  x+
//line /usr/local/go/src/strings/replace.go:106
//	n7  .y+
//line /usr/local/go/src/strings/replace.go:106
//
//line /usr/local/go/src/strings/replace.go:106
// n0 is the root node, and its children are n1, n4 and n6; n1's children are
//line /usr/local/go/src/strings/replace.go:106
// n2 and n3; n4's child is n5; n6's child is n7. Nodes n0, n1 and n4 (marked
//line /usr/local/go/src/strings/replace.go:106
// with a trailing "-") are partial keys, and nodes n2, n3, n5, n6 and n7
//line /usr/local/go/src/strings/replace.go:106
// (marked with a trailing "+") are complete keys.
//line /usr/local/go/src/strings/replace.go:123
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

//line /usr/local/go/src/strings/replace.go:142
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
//line /usr/local/go/src/strings/replace.go:158
	_go_fuzz_dep_.CoverTab[990]++
							if key == "" {
//line /usr/local/go/src/strings/replace.go:159
		_go_fuzz_dep_.CoverTab[992]++
								if t.priority == 0 {
//line /usr/local/go/src/strings/replace.go:160
			_go_fuzz_dep_.CoverTab[994]++
									t.value = val
									t.priority = priority
//line /usr/local/go/src/strings/replace.go:162
			// _ = "end of CoverTab[994]"
		} else {
//line /usr/local/go/src/strings/replace.go:163
			_go_fuzz_dep_.CoverTab[995]++
//line /usr/local/go/src/strings/replace.go:163
			// _ = "end of CoverTab[995]"
//line /usr/local/go/src/strings/replace.go:163
		}
//line /usr/local/go/src/strings/replace.go:163
		// _ = "end of CoverTab[992]"
//line /usr/local/go/src/strings/replace.go:163
		_go_fuzz_dep_.CoverTab[993]++
								return
//line /usr/local/go/src/strings/replace.go:164
		// _ = "end of CoverTab[993]"
	} else {
//line /usr/local/go/src/strings/replace.go:165
		_go_fuzz_dep_.CoverTab[996]++
//line /usr/local/go/src/strings/replace.go:165
		// _ = "end of CoverTab[996]"
//line /usr/local/go/src/strings/replace.go:165
	}
//line /usr/local/go/src/strings/replace.go:165
	// _ = "end of CoverTab[990]"
//line /usr/local/go/src/strings/replace.go:165
	_go_fuzz_dep_.CoverTab[991]++

							if t.prefix != "" {
//line /usr/local/go/src/strings/replace.go:167
		_go_fuzz_dep_.CoverTab[997]++
		// Need to split the prefix among multiple nodes.
		var n int	// length of the longest common prefix
		for ; n < len(t.prefix) && func() bool {
//line /usr/local/go/src/strings/replace.go:170
			_go_fuzz_dep_.CoverTab[999]++
//line /usr/local/go/src/strings/replace.go:170
			return n < len(key)
//line /usr/local/go/src/strings/replace.go:170
			// _ = "end of CoverTab[999]"
//line /usr/local/go/src/strings/replace.go:170
		}(); n++ {
//line /usr/local/go/src/strings/replace.go:170
			_go_fuzz_dep_.CoverTab[1000]++
									if t.prefix[n] != key[n] {
//line /usr/local/go/src/strings/replace.go:171
				_go_fuzz_dep_.CoverTab[1001]++
										break
//line /usr/local/go/src/strings/replace.go:172
				// _ = "end of CoverTab[1001]"
			} else {
//line /usr/local/go/src/strings/replace.go:173
				_go_fuzz_dep_.CoverTab[1002]++
//line /usr/local/go/src/strings/replace.go:173
				// _ = "end of CoverTab[1002]"
//line /usr/local/go/src/strings/replace.go:173
			}
//line /usr/local/go/src/strings/replace.go:173
			// _ = "end of CoverTab[1000]"
		}
//line /usr/local/go/src/strings/replace.go:174
		// _ = "end of CoverTab[997]"
//line /usr/local/go/src/strings/replace.go:174
		_go_fuzz_dep_.CoverTab[998]++
								if n == len(t.prefix) {
//line /usr/local/go/src/strings/replace.go:175
			_go_fuzz_dep_.CoverTab[1003]++
									t.next.add(key[n:], val, priority, r)
//line /usr/local/go/src/strings/replace.go:176
			// _ = "end of CoverTab[1003]"
		} else {
//line /usr/local/go/src/strings/replace.go:177
			_go_fuzz_dep_.CoverTab[1004]++
//line /usr/local/go/src/strings/replace.go:177
			if n == 0 {
//line /usr/local/go/src/strings/replace.go:177
				_go_fuzz_dep_.CoverTab[1005]++
				// First byte differs, start a new lookup table here. Looking up
				// what is currently t.prefix[0] will lead to prefixNode, and
				// looking up key[0] will lead to keyNode.
				var prefixNode *trieNode
				if len(t.prefix) == 1 {
//line /usr/local/go/src/strings/replace.go:182
					_go_fuzz_dep_.CoverTab[1007]++
											prefixNode = t.next
//line /usr/local/go/src/strings/replace.go:183
					// _ = "end of CoverTab[1007]"
				} else {
//line /usr/local/go/src/strings/replace.go:184
					_go_fuzz_dep_.CoverTab[1008]++
											prefixNode = &trieNode{
						prefix:	t.prefix[1:],
						next:	t.next,
					}
//line /usr/local/go/src/strings/replace.go:188
					// _ = "end of CoverTab[1008]"
				}
//line /usr/local/go/src/strings/replace.go:189
				// _ = "end of CoverTab[1005]"
//line /usr/local/go/src/strings/replace.go:189
				_go_fuzz_dep_.CoverTab[1006]++
										keyNode := new(trieNode)
										t.table = make([]*trieNode, r.tableSize)
										t.table[r.mapping[t.prefix[0]]] = prefixNode
										t.table[r.mapping[key[0]]] = keyNode
										t.prefix = ""
										t.next = nil
										keyNode.add(key[1:], val, priority, r)
//line /usr/local/go/src/strings/replace.go:196
				// _ = "end of CoverTab[1006]"
			} else {
//line /usr/local/go/src/strings/replace.go:197
				_go_fuzz_dep_.CoverTab[1009]++

										next := &trieNode{
					prefix:	t.prefix[n:],
					next:	t.next,
				}
										t.prefix = t.prefix[:n]
										t.next = next
										next.add(key[n:], val, priority, r)
//line /usr/local/go/src/strings/replace.go:205
				// _ = "end of CoverTab[1009]"
			}
//line /usr/local/go/src/strings/replace.go:206
			// _ = "end of CoverTab[1004]"
//line /usr/local/go/src/strings/replace.go:206
		}
//line /usr/local/go/src/strings/replace.go:206
		// _ = "end of CoverTab[998]"
	} else {
//line /usr/local/go/src/strings/replace.go:207
		_go_fuzz_dep_.CoverTab[1010]++
//line /usr/local/go/src/strings/replace.go:207
		if t.table != nil {
//line /usr/local/go/src/strings/replace.go:207
			_go_fuzz_dep_.CoverTab[1011]++

									m := r.mapping[key[0]]
									if t.table[m] == nil {
//line /usr/local/go/src/strings/replace.go:210
				_go_fuzz_dep_.CoverTab[1013]++
										t.table[m] = new(trieNode)
//line /usr/local/go/src/strings/replace.go:211
				// _ = "end of CoverTab[1013]"
			} else {
//line /usr/local/go/src/strings/replace.go:212
				_go_fuzz_dep_.CoverTab[1014]++
//line /usr/local/go/src/strings/replace.go:212
				// _ = "end of CoverTab[1014]"
//line /usr/local/go/src/strings/replace.go:212
			}
//line /usr/local/go/src/strings/replace.go:212
			// _ = "end of CoverTab[1011]"
//line /usr/local/go/src/strings/replace.go:212
			_go_fuzz_dep_.CoverTab[1012]++
									t.table[m].add(key[1:], val, priority, r)
//line /usr/local/go/src/strings/replace.go:213
			// _ = "end of CoverTab[1012]"
		} else {
//line /usr/local/go/src/strings/replace.go:214
			_go_fuzz_dep_.CoverTab[1015]++
									t.prefix = key
									t.next = new(trieNode)
									t.next.add("", val, priority, r)
//line /usr/local/go/src/strings/replace.go:217
			// _ = "end of CoverTab[1015]"
		}
//line /usr/local/go/src/strings/replace.go:218
		// _ = "end of CoverTab[1010]"
//line /usr/local/go/src/strings/replace.go:218
	}
//line /usr/local/go/src/strings/replace.go:218
	// _ = "end of CoverTab[991]"
}

func (r *genericReplacer) lookup(s string, ignoreRoot bool) (val string, keylen int, found bool) {
//line /usr/local/go/src/strings/replace.go:221
	_go_fuzz_dep_.CoverTab[1016]++

//line /usr/local/go/src/strings/replace.go:224
	bestPriority := 0
	node := &r.root
	n := 0
	for node != nil {
//line /usr/local/go/src/strings/replace.go:227
		_go_fuzz_dep_.CoverTab[1018]++
								if node.priority > bestPriority && func() bool {
//line /usr/local/go/src/strings/replace.go:228
			_go_fuzz_dep_.CoverTab[1021]++
//line /usr/local/go/src/strings/replace.go:228
			return !(ignoreRoot && func() bool {
//line /usr/local/go/src/strings/replace.go:228
				_go_fuzz_dep_.CoverTab[1022]++
//line /usr/local/go/src/strings/replace.go:228
				return node == &r.root
//line /usr/local/go/src/strings/replace.go:228
				// _ = "end of CoverTab[1022]"
//line /usr/local/go/src/strings/replace.go:228
			}())
//line /usr/local/go/src/strings/replace.go:228
			// _ = "end of CoverTab[1021]"
//line /usr/local/go/src/strings/replace.go:228
		}() {
//line /usr/local/go/src/strings/replace.go:228
			_go_fuzz_dep_.CoverTab[1023]++
									bestPriority = node.priority
									val = node.value
									keylen = n
									found = true
//line /usr/local/go/src/strings/replace.go:232
			// _ = "end of CoverTab[1023]"
		} else {
//line /usr/local/go/src/strings/replace.go:233
			_go_fuzz_dep_.CoverTab[1024]++
//line /usr/local/go/src/strings/replace.go:233
			// _ = "end of CoverTab[1024]"
//line /usr/local/go/src/strings/replace.go:233
		}
//line /usr/local/go/src/strings/replace.go:233
		// _ = "end of CoverTab[1018]"
//line /usr/local/go/src/strings/replace.go:233
		_go_fuzz_dep_.CoverTab[1019]++

								if s == "" {
//line /usr/local/go/src/strings/replace.go:235
			_go_fuzz_dep_.CoverTab[1025]++
									break
//line /usr/local/go/src/strings/replace.go:236
			// _ = "end of CoverTab[1025]"
		} else {
//line /usr/local/go/src/strings/replace.go:237
			_go_fuzz_dep_.CoverTab[1026]++
//line /usr/local/go/src/strings/replace.go:237
			// _ = "end of CoverTab[1026]"
//line /usr/local/go/src/strings/replace.go:237
		}
//line /usr/local/go/src/strings/replace.go:237
		// _ = "end of CoverTab[1019]"
//line /usr/local/go/src/strings/replace.go:237
		_go_fuzz_dep_.CoverTab[1020]++
								if node.table != nil {
//line /usr/local/go/src/strings/replace.go:238
			_go_fuzz_dep_.CoverTab[1027]++
									index := r.mapping[s[0]]
									if int(index) == r.tableSize {
//line /usr/local/go/src/strings/replace.go:240
				_go_fuzz_dep_.CoverTab[1029]++
										break
//line /usr/local/go/src/strings/replace.go:241
				// _ = "end of CoverTab[1029]"
			} else {
//line /usr/local/go/src/strings/replace.go:242
				_go_fuzz_dep_.CoverTab[1030]++
//line /usr/local/go/src/strings/replace.go:242
				// _ = "end of CoverTab[1030]"
//line /usr/local/go/src/strings/replace.go:242
			}
//line /usr/local/go/src/strings/replace.go:242
			// _ = "end of CoverTab[1027]"
//line /usr/local/go/src/strings/replace.go:242
			_go_fuzz_dep_.CoverTab[1028]++
									node = node.table[index]
									s = s[1:]
									n++
//line /usr/local/go/src/strings/replace.go:245
			// _ = "end of CoverTab[1028]"
		} else {
//line /usr/local/go/src/strings/replace.go:246
			_go_fuzz_dep_.CoverTab[1031]++
//line /usr/local/go/src/strings/replace.go:246
			if node.prefix != "" && func() bool {
//line /usr/local/go/src/strings/replace.go:246
				_go_fuzz_dep_.CoverTab[1032]++
//line /usr/local/go/src/strings/replace.go:246
				return HasPrefix(s, node.prefix)
//line /usr/local/go/src/strings/replace.go:246
				// _ = "end of CoverTab[1032]"
//line /usr/local/go/src/strings/replace.go:246
			}() {
//line /usr/local/go/src/strings/replace.go:246
				_go_fuzz_dep_.CoverTab[1033]++
										n += len(node.prefix)
										s = s[len(node.prefix):]
										node = node.next
//line /usr/local/go/src/strings/replace.go:249
				// _ = "end of CoverTab[1033]"
			} else {
//line /usr/local/go/src/strings/replace.go:250
				_go_fuzz_dep_.CoverTab[1034]++
										break
//line /usr/local/go/src/strings/replace.go:251
				// _ = "end of CoverTab[1034]"
			}
//line /usr/local/go/src/strings/replace.go:252
			// _ = "end of CoverTab[1031]"
//line /usr/local/go/src/strings/replace.go:252
		}
//line /usr/local/go/src/strings/replace.go:252
		// _ = "end of CoverTab[1020]"
	}
//line /usr/local/go/src/strings/replace.go:253
	// _ = "end of CoverTab[1016]"
//line /usr/local/go/src/strings/replace.go:253
	_go_fuzz_dep_.CoverTab[1017]++
							return
//line /usr/local/go/src/strings/replace.go:254
	// _ = "end of CoverTab[1017]"
}

// genericReplacer is the fully generic algorithm.
//line /usr/local/go/src/strings/replace.go:257
// It's used as a fallback when nothing faster can be used.
//line /usr/local/go/src/strings/replace.go:259
type genericReplacer struct {
	root	trieNode
	// tableSize is the size of a trie node's lookup table. It is the number
	// of unique key bytes.
	tableSize	int
	// mapping maps from key bytes to a dense index for trieNode.table.
	mapping	[256]byte
}

func makeGenericReplacer(oldnew []string) *genericReplacer {
//line /usr/local/go/src/strings/replace.go:268
	_go_fuzz_dep_.CoverTab[1035]++
							r := new(genericReplacer)

							for i := 0; i < len(oldnew); i += 2 {
//line /usr/local/go/src/strings/replace.go:271
		_go_fuzz_dep_.CoverTab[1040]++
								key := oldnew[i]
								for j := 0; j < len(key); j++ {
//line /usr/local/go/src/strings/replace.go:273
			_go_fuzz_dep_.CoverTab[1041]++
									r.mapping[key[j]] = 1
//line /usr/local/go/src/strings/replace.go:274
			// _ = "end of CoverTab[1041]"
		}
//line /usr/local/go/src/strings/replace.go:275
		// _ = "end of CoverTab[1040]"
	}
//line /usr/local/go/src/strings/replace.go:276
	// _ = "end of CoverTab[1035]"
//line /usr/local/go/src/strings/replace.go:276
	_go_fuzz_dep_.CoverTab[1036]++

							for _, b := range r.mapping {
//line /usr/local/go/src/strings/replace.go:278
		_go_fuzz_dep_.CoverTab[1042]++
								r.tableSize += int(b)
//line /usr/local/go/src/strings/replace.go:279
		// _ = "end of CoverTab[1042]"
	}
//line /usr/local/go/src/strings/replace.go:280
	// _ = "end of CoverTab[1036]"
//line /usr/local/go/src/strings/replace.go:280
	_go_fuzz_dep_.CoverTab[1037]++

							var index byte
							for i, b := range r.mapping {
//line /usr/local/go/src/strings/replace.go:283
		_go_fuzz_dep_.CoverTab[1043]++
								if b == 0 {
//line /usr/local/go/src/strings/replace.go:284
			_go_fuzz_dep_.CoverTab[1044]++
									r.mapping[i] = byte(r.tableSize)
//line /usr/local/go/src/strings/replace.go:285
			// _ = "end of CoverTab[1044]"
		} else {
//line /usr/local/go/src/strings/replace.go:286
			_go_fuzz_dep_.CoverTab[1045]++
									r.mapping[i] = index
									index++
//line /usr/local/go/src/strings/replace.go:288
			// _ = "end of CoverTab[1045]"
		}
//line /usr/local/go/src/strings/replace.go:289
		// _ = "end of CoverTab[1043]"
	}
//line /usr/local/go/src/strings/replace.go:290
	// _ = "end of CoverTab[1037]"
//line /usr/local/go/src/strings/replace.go:290
	_go_fuzz_dep_.CoverTab[1038]++

							r.root.table = make([]*trieNode, r.tableSize)

							for i := 0; i < len(oldnew); i += 2 {
//line /usr/local/go/src/strings/replace.go:294
		_go_fuzz_dep_.CoverTab[1046]++
								r.root.add(oldnew[i], oldnew[i+1], len(oldnew)-i, r)
//line /usr/local/go/src/strings/replace.go:295
		// _ = "end of CoverTab[1046]"
	}
//line /usr/local/go/src/strings/replace.go:296
	// _ = "end of CoverTab[1038]"
//line /usr/local/go/src/strings/replace.go:296
	_go_fuzz_dep_.CoverTab[1039]++
							return r
//line /usr/local/go/src/strings/replace.go:297
	// _ = "end of CoverTab[1039]"
}

type appendSliceWriter []byte

// Write writes to the buffer to satisfy io.Writer.
func (w *appendSliceWriter) Write(p []byte) (int, error) {
//line /usr/local/go/src/strings/replace.go:303
	_go_fuzz_dep_.CoverTab[1047]++
							*w = append(*w, p...)
							return len(p), nil
//line /usr/local/go/src/strings/replace.go:305
	// _ = "end of CoverTab[1047]"
}

// WriteString writes to the buffer without string->[]byte->string allocations.
func (w *appendSliceWriter) WriteString(s string) (int, error) {
//line /usr/local/go/src/strings/replace.go:309
	_go_fuzz_dep_.CoverTab[1048]++
							*w = append(*w, s...)
							return len(s), nil
//line /usr/local/go/src/strings/replace.go:311
	// _ = "end of CoverTab[1048]"
}

type stringWriter struct {
	w io.Writer
}

func (w stringWriter) WriteString(s string) (int, error) {
//line /usr/local/go/src/strings/replace.go:318
	_go_fuzz_dep_.CoverTab[1049]++
							return w.w.Write([]byte(s))
//line /usr/local/go/src/strings/replace.go:319
	// _ = "end of CoverTab[1049]"
}

func getStringWriter(w io.Writer) io.StringWriter {
//line /usr/local/go/src/strings/replace.go:322
	_go_fuzz_dep_.CoverTab[1050]++
							sw, ok := w.(io.StringWriter)
							if !ok {
//line /usr/local/go/src/strings/replace.go:324
		_go_fuzz_dep_.CoverTab[1052]++
								sw = stringWriter{w}
//line /usr/local/go/src/strings/replace.go:325
		// _ = "end of CoverTab[1052]"
	} else {
//line /usr/local/go/src/strings/replace.go:326
		_go_fuzz_dep_.CoverTab[1053]++
//line /usr/local/go/src/strings/replace.go:326
		// _ = "end of CoverTab[1053]"
//line /usr/local/go/src/strings/replace.go:326
	}
//line /usr/local/go/src/strings/replace.go:326
	// _ = "end of CoverTab[1050]"
//line /usr/local/go/src/strings/replace.go:326
	_go_fuzz_dep_.CoverTab[1051]++
							return sw
//line /usr/local/go/src/strings/replace.go:327
	// _ = "end of CoverTab[1051]"
}

func (r *genericReplacer) Replace(s string) string {
//line /usr/local/go/src/strings/replace.go:330
	_go_fuzz_dep_.CoverTab[1054]++
							buf := make(appendSliceWriter, 0, len(s))
							r.WriteString(&buf, s)
							return string(buf)
//line /usr/local/go/src/strings/replace.go:333
	// _ = "end of CoverTab[1054]"
}

func (r *genericReplacer) WriteString(w io.Writer, s string) (n int, err error) {
//line /usr/local/go/src/strings/replace.go:336
	_go_fuzz_dep_.CoverTab[1055]++
							sw := getStringWriter(w)
							var last, wn int
							var prevMatchEmpty bool
							for i := 0; i <= len(s); {
//line /usr/local/go/src/strings/replace.go:340
		_go_fuzz_dep_.CoverTab[1058]++

								if i != len(s) && func() bool {
//line /usr/local/go/src/strings/replace.go:342
			_go_fuzz_dep_.CoverTab[1061]++
//line /usr/local/go/src/strings/replace.go:342
			return r.root.priority == 0
//line /usr/local/go/src/strings/replace.go:342
			// _ = "end of CoverTab[1061]"
//line /usr/local/go/src/strings/replace.go:342
		}() {
//line /usr/local/go/src/strings/replace.go:342
			_go_fuzz_dep_.CoverTab[1062]++
									index := int(r.mapping[s[i]])
									if index == r.tableSize || func() bool {
//line /usr/local/go/src/strings/replace.go:344
				_go_fuzz_dep_.CoverTab[1063]++
//line /usr/local/go/src/strings/replace.go:344
				return r.root.table[index] == nil
//line /usr/local/go/src/strings/replace.go:344
				// _ = "end of CoverTab[1063]"
//line /usr/local/go/src/strings/replace.go:344
			}() {
//line /usr/local/go/src/strings/replace.go:344
				_go_fuzz_dep_.CoverTab[1064]++
										i++
										continue
//line /usr/local/go/src/strings/replace.go:346
				// _ = "end of CoverTab[1064]"
			} else {
//line /usr/local/go/src/strings/replace.go:347
				_go_fuzz_dep_.CoverTab[1065]++
//line /usr/local/go/src/strings/replace.go:347
				// _ = "end of CoverTab[1065]"
//line /usr/local/go/src/strings/replace.go:347
			}
//line /usr/local/go/src/strings/replace.go:347
			// _ = "end of CoverTab[1062]"
		} else {
//line /usr/local/go/src/strings/replace.go:348
			_go_fuzz_dep_.CoverTab[1066]++
//line /usr/local/go/src/strings/replace.go:348
			// _ = "end of CoverTab[1066]"
//line /usr/local/go/src/strings/replace.go:348
		}
//line /usr/local/go/src/strings/replace.go:348
		// _ = "end of CoverTab[1058]"
//line /usr/local/go/src/strings/replace.go:348
		_go_fuzz_dep_.CoverTab[1059]++

//line /usr/local/go/src/strings/replace.go:351
		val, keylen, match := r.lookup(s[i:], prevMatchEmpty)
		prevMatchEmpty = match && func() bool {
//line /usr/local/go/src/strings/replace.go:352
			_go_fuzz_dep_.CoverTab[1067]++
//line /usr/local/go/src/strings/replace.go:352
			return keylen == 0
//line /usr/local/go/src/strings/replace.go:352
			// _ = "end of CoverTab[1067]"
//line /usr/local/go/src/strings/replace.go:352
		}()
								if match {
//line /usr/local/go/src/strings/replace.go:353
			_go_fuzz_dep_.CoverTab[1068]++
									wn, err = sw.WriteString(s[last:i])
									n += wn
									if err != nil {
//line /usr/local/go/src/strings/replace.go:356
				_go_fuzz_dep_.CoverTab[1071]++
										return
//line /usr/local/go/src/strings/replace.go:357
				// _ = "end of CoverTab[1071]"
			} else {
//line /usr/local/go/src/strings/replace.go:358
				_go_fuzz_dep_.CoverTab[1072]++
//line /usr/local/go/src/strings/replace.go:358
				// _ = "end of CoverTab[1072]"
//line /usr/local/go/src/strings/replace.go:358
			}
//line /usr/local/go/src/strings/replace.go:358
			// _ = "end of CoverTab[1068]"
//line /usr/local/go/src/strings/replace.go:358
			_go_fuzz_dep_.CoverTab[1069]++
									wn, err = sw.WriteString(val)
									n += wn
									if err != nil {
//line /usr/local/go/src/strings/replace.go:361
				_go_fuzz_dep_.CoverTab[1073]++
										return
//line /usr/local/go/src/strings/replace.go:362
				// _ = "end of CoverTab[1073]"
			} else {
//line /usr/local/go/src/strings/replace.go:363
				_go_fuzz_dep_.CoverTab[1074]++
//line /usr/local/go/src/strings/replace.go:363
				// _ = "end of CoverTab[1074]"
//line /usr/local/go/src/strings/replace.go:363
			}
//line /usr/local/go/src/strings/replace.go:363
			// _ = "end of CoverTab[1069]"
//line /usr/local/go/src/strings/replace.go:363
			_go_fuzz_dep_.CoverTab[1070]++
									i += keylen
									last = i
									continue
//line /usr/local/go/src/strings/replace.go:366
			// _ = "end of CoverTab[1070]"
		} else {
//line /usr/local/go/src/strings/replace.go:367
			_go_fuzz_dep_.CoverTab[1075]++
//line /usr/local/go/src/strings/replace.go:367
			// _ = "end of CoverTab[1075]"
//line /usr/local/go/src/strings/replace.go:367
		}
//line /usr/local/go/src/strings/replace.go:367
		// _ = "end of CoverTab[1059]"
//line /usr/local/go/src/strings/replace.go:367
		_go_fuzz_dep_.CoverTab[1060]++
								i++
//line /usr/local/go/src/strings/replace.go:368
		// _ = "end of CoverTab[1060]"
	}
//line /usr/local/go/src/strings/replace.go:369
	// _ = "end of CoverTab[1055]"
//line /usr/local/go/src/strings/replace.go:369
	_go_fuzz_dep_.CoverTab[1056]++
							if last != len(s) {
//line /usr/local/go/src/strings/replace.go:370
		_go_fuzz_dep_.CoverTab[1076]++
								wn, err = sw.WriteString(s[last:])
								n += wn
//line /usr/local/go/src/strings/replace.go:372
		// _ = "end of CoverTab[1076]"
	} else {
//line /usr/local/go/src/strings/replace.go:373
		_go_fuzz_dep_.CoverTab[1077]++
//line /usr/local/go/src/strings/replace.go:373
		// _ = "end of CoverTab[1077]"
//line /usr/local/go/src/strings/replace.go:373
	}
//line /usr/local/go/src/strings/replace.go:373
	// _ = "end of CoverTab[1056]"
//line /usr/local/go/src/strings/replace.go:373
	_go_fuzz_dep_.CoverTab[1057]++
							return
//line /usr/local/go/src/strings/replace.go:374
	// _ = "end of CoverTab[1057]"
}

// singleStringReplacer is the implementation that's used when there is only
//line /usr/local/go/src/strings/replace.go:377
// one string to replace (and that string has more than one byte).
//line /usr/local/go/src/strings/replace.go:379
type singleStringReplacer struct {
	finder	*stringFinder
	// value is the new string that replaces that pattern when it's found.
	value	string
}

func makeSingleStringReplacer(pattern string, value string) *singleStringReplacer {
//line /usr/local/go/src/strings/replace.go:385
	_go_fuzz_dep_.CoverTab[1078]++
							return &singleStringReplacer{finder: makeStringFinder(pattern), value: value}
//line /usr/local/go/src/strings/replace.go:386
	// _ = "end of CoverTab[1078]"
}

func (r *singleStringReplacer) Replace(s string) string {
//line /usr/local/go/src/strings/replace.go:389
	_go_fuzz_dep_.CoverTab[1079]++
							var buf Builder
							i, matched := 0, false
							for {
//line /usr/local/go/src/strings/replace.go:392
		_go_fuzz_dep_.CoverTab[1082]++
								match := r.finder.next(s[i:])
								if match == -1 {
//line /usr/local/go/src/strings/replace.go:394
			_go_fuzz_dep_.CoverTab[1084]++
									break
//line /usr/local/go/src/strings/replace.go:395
			// _ = "end of CoverTab[1084]"
		} else {
//line /usr/local/go/src/strings/replace.go:396
			_go_fuzz_dep_.CoverTab[1085]++
//line /usr/local/go/src/strings/replace.go:396
			// _ = "end of CoverTab[1085]"
//line /usr/local/go/src/strings/replace.go:396
		}
//line /usr/local/go/src/strings/replace.go:396
		// _ = "end of CoverTab[1082]"
//line /usr/local/go/src/strings/replace.go:396
		_go_fuzz_dep_.CoverTab[1083]++
								matched = true
								buf.Grow(match + len(r.value))
								buf.WriteString(s[i : i+match])
								buf.WriteString(r.value)
								i += match + len(r.finder.pattern)
//line /usr/local/go/src/strings/replace.go:401
		// _ = "end of CoverTab[1083]"
	}
//line /usr/local/go/src/strings/replace.go:402
	// _ = "end of CoverTab[1079]"
//line /usr/local/go/src/strings/replace.go:402
	_go_fuzz_dep_.CoverTab[1080]++
							if !matched {
//line /usr/local/go/src/strings/replace.go:403
		_go_fuzz_dep_.CoverTab[1086]++
								return s
//line /usr/local/go/src/strings/replace.go:404
		// _ = "end of CoverTab[1086]"
	} else {
//line /usr/local/go/src/strings/replace.go:405
		_go_fuzz_dep_.CoverTab[1087]++
//line /usr/local/go/src/strings/replace.go:405
		// _ = "end of CoverTab[1087]"
//line /usr/local/go/src/strings/replace.go:405
	}
//line /usr/local/go/src/strings/replace.go:405
	// _ = "end of CoverTab[1080]"
//line /usr/local/go/src/strings/replace.go:405
	_go_fuzz_dep_.CoverTab[1081]++
							buf.WriteString(s[i:])
							return buf.String()
//line /usr/local/go/src/strings/replace.go:407
	// _ = "end of CoverTab[1081]"
}

func (r *singleStringReplacer) WriteString(w io.Writer, s string) (n int, err error) {
//line /usr/local/go/src/strings/replace.go:410
	_go_fuzz_dep_.CoverTab[1088]++
							sw := getStringWriter(w)
							var i, wn int
							for {
//line /usr/local/go/src/strings/replace.go:413
		_go_fuzz_dep_.CoverTab[1090]++
								match := r.finder.next(s[i:])
								if match == -1 {
//line /usr/local/go/src/strings/replace.go:415
			_go_fuzz_dep_.CoverTab[1094]++
									break
//line /usr/local/go/src/strings/replace.go:416
			// _ = "end of CoverTab[1094]"
		} else {
//line /usr/local/go/src/strings/replace.go:417
			_go_fuzz_dep_.CoverTab[1095]++
//line /usr/local/go/src/strings/replace.go:417
			// _ = "end of CoverTab[1095]"
//line /usr/local/go/src/strings/replace.go:417
		}
//line /usr/local/go/src/strings/replace.go:417
		// _ = "end of CoverTab[1090]"
//line /usr/local/go/src/strings/replace.go:417
		_go_fuzz_dep_.CoverTab[1091]++
								wn, err = sw.WriteString(s[i : i+match])
								n += wn
								if err != nil {
//line /usr/local/go/src/strings/replace.go:420
			_go_fuzz_dep_.CoverTab[1096]++
									return
//line /usr/local/go/src/strings/replace.go:421
			// _ = "end of CoverTab[1096]"
		} else {
//line /usr/local/go/src/strings/replace.go:422
			_go_fuzz_dep_.CoverTab[1097]++
//line /usr/local/go/src/strings/replace.go:422
			// _ = "end of CoverTab[1097]"
//line /usr/local/go/src/strings/replace.go:422
		}
//line /usr/local/go/src/strings/replace.go:422
		// _ = "end of CoverTab[1091]"
//line /usr/local/go/src/strings/replace.go:422
		_go_fuzz_dep_.CoverTab[1092]++
								wn, err = sw.WriteString(r.value)
								n += wn
								if err != nil {
//line /usr/local/go/src/strings/replace.go:425
			_go_fuzz_dep_.CoverTab[1098]++
									return
//line /usr/local/go/src/strings/replace.go:426
			// _ = "end of CoverTab[1098]"
		} else {
//line /usr/local/go/src/strings/replace.go:427
			_go_fuzz_dep_.CoverTab[1099]++
//line /usr/local/go/src/strings/replace.go:427
			// _ = "end of CoverTab[1099]"
//line /usr/local/go/src/strings/replace.go:427
		}
//line /usr/local/go/src/strings/replace.go:427
		// _ = "end of CoverTab[1092]"
//line /usr/local/go/src/strings/replace.go:427
		_go_fuzz_dep_.CoverTab[1093]++
								i += match + len(r.finder.pattern)
//line /usr/local/go/src/strings/replace.go:428
		// _ = "end of CoverTab[1093]"
	}
//line /usr/local/go/src/strings/replace.go:429
	// _ = "end of CoverTab[1088]"
//line /usr/local/go/src/strings/replace.go:429
	_go_fuzz_dep_.CoverTab[1089]++
							wn, err = sw.WriteString(s[i:])
							n += wn
							return
//line /usr/local/go/src/strings/replace.go:432
	// _ = "end of CoverTab[1089]"
}

// byteReplacer is the implementation that's used when all the "old"
//line /usr/local/go/src/strings/replace.go:435
// and "new" values are single ASCII bytes.
//line /usr/local/go/src/strings/replace.go:435
// The array contains replacement bytes indexed by old byte.
//line /usr/local/go/src/strings/replace.go:438
type byteReplacer [256]byte

func (r *byteReplacer) Replace(s string) string {
//line /usr/local/go/src/strings/replace.go:440
	_go_fuzz_dep_.CoverTab[1100]++
							var buf []byte	// lazily allocated
							for i := 0; i < len(s); i++ {
//line /usr/local/go/src/strings/replace.go:442
		_go_fuzz_dep_.CoverTab[1103]++
								b := s[i]
								if r[b] != b {
//line /usr/local/go/src/strings/replace.go:444
			_go_fuzz_dep_.CoverTab[1104]++
									if buf == nil {
//line /usr/local/go/src/strings/replace.go:445
				_go_fuzz_dep_.CoverTab[1106]++
										buf = []byte(s)
//line /usr/local/go/src/strings/replace.go:446
				// _ = "end of CoverTab[1106]"
			} else {
//line /usr/local/go/src/strings/replace.go:447
				_go_fuzz_dep_.CoverTab[1107]++
//line /usr/local/go/src/strings/replace.go:447
				// _ = "end of CoverTab[1107]"
//line /usr/local/go/src/strings/replace.go:447
			}
//line /usr/local/go/src/strings/replace.go:447
			// _ = "end of CoverTab[1104]"
//line /usr/local/go/src/strings/replace.go:447
			_go_fuzz_dep_.CoverTab[1105]++
									buf[i] = r[b]
//line /usr/local/go/src/strings/replace.go:448
			// _ = "end of CoverTab[1105]"
		} else {
//line /usr/local/go/src/strings/replace.go:449
			_go_fuzz_dep_.CoverTab[1108]++
//line /usr/local/go/src/strings/replace.go:449
			// _ = "end of CoverTab[1108]"
//line /usr/local/go/src/strings/replace.go:449
		}
//line /usr/local/go/src/strings/replace.go:449
		// _ = "end of CoverTab[1103]"
	}
//line /usr/local/go/src/strings/replace.go:450
	// _ = "end of CoverTab[1100]"
//line /usr/local/go/src/strings/replace.go:450
	_go_fuzz_dep_.CoverTab[1101]++
							if buf == nil {
//line /usr/local/go/src/strings/replace.go:451
		_go_fuzz_dep_.CoverTab[1109]++
								return s
//line /usr/local/go/src/strings/replace.go:452
		// _ = "end of CoverTab[1109]"
	} else {
//line /usr/local/go/src/strings/replace.go:453
		_go_fuzz_dep_.CoverTab[1110]++
//line /usr/local/go/src/strings/replace.go:453
		// _ = "end of CoverTab[1110]"
//line /usr/local/go/src/strings/replace.go:453
	}
//line /usr/local/go/src/strings/replace.go:453
	// _ = "end of CoverTab[1101]"
//line /usr/local/go/src/strings/replace.go:453
	_go_fuzz_dep_.CoverTab[1102]++
							return string(buf)
//line /usr/local/go/src/strings/replace.go:454
	// _ = "end of CoverTab[1102]"
}

func (r *byteReplacer) WriteString(w io.Writer, s string) (n int, err error) {
//line /usr/local/go/src/strings/replace.go:457
	_go_fuzz_dep_.CoverTab[1111]++
							sw := getStringWriter(w)
							last := 0
							for i := 0; i < len(s); i++ {
//line /usr/local/go/src/strings/replace.go:460
		_go_fuzz_dep_.CoverTab[1114]++
								b := s[i]
								if r[b] == b {
//line /usr/local/go/src/strings/replace.go:462
			_go_fuzz_dep_.CoverTab[1117]++
									continue
//line /usr/local/go/src/strings/replace.go:463
			// _ = "end of CoverTab[1117]"
		} else {
//line /usr/local/go/src/strings/replace.go:464
			_go_fuzz_dep_.CoverTab[1118]++
//line /usr/local/go/src/strings/replace.go:464
			// _ = "end of CoverTab[1118]"
//line /usr/local/go/src/strings/replace.go:464
		}
//line /usr/local/go/src/strings/replace.go:464
		// _ = "end of CoverTab[1114]"
//line /usr/local/go/src/strings/replace.go:464
		_go_fuzz_dep_.CoverTab[1115]++
								if last != i {
//line /usr/local/go/src/strings/replace.go:465
			_go_fuzz_dep_.CoverTab[1119]++
									wn, err := sw.WriteString(s[last:i])
									n += wn
									if err != nil {
//line /usr/local/go/src/strings/replace.go:468
				_go_fuzz_dep_.CoverTab[1120]++
										return n, err
//line /usr/local/go/src/strings/replace.go:469
				// _ = "end of CoverTab[1120]"
			} else {
//line /usr/local/go/src/strings/replace.go:470
				_go_fuzz_dep_.CoverTab[1121]++
//line /usr/local/go/src/strings/replace.go:470
				// _ = "end of CoverTab[1121]"
//line /usr/local/go/src/strings/replace.go:470
			}
//line /usr/local/go/src/strings/replace.go:470
			// _ = "end of CoverTab[1119]"
		} else {
//line /usr/local/go/src/strings/replace.go:471
			_go_fuzz_dep_.CoverTab[1122]++
//line /usr/local/go/src/strings/replace.go:471
			// _ = "end of CoverTab[1122]"
//line /usr/local/go/src/strings/replace.go:471
		}
//line /usr/local/go/src/strings/replace.go:471
		// _ = "end of CoverTab[1115]"
//line /usr/local/go/src/strings/replace.go:471
		_go_fuzz_dep_.CoverTab[1116]++
								last = i + 1
								nw, err := w.Write(r[b : int(b)+1])
								n += nw
								if err != nil {
//line /usr/local/go/src/strings/replace.go:475
			_go_fuzz_dep_.CoverTab[1123]++
									return n, err
//line /usr/local/go/src/strings/replace.go:476
			// _ = "end of CoverTab[1123]"
		} else {
//line /usr/local/go/src/strings/replace.go:477
			_go_fuzz_dep_.CoverTab[1124]++
//line /usr/local/go/src/strings/replace.go:477
			// _ = "end of CoverTab[1124]"
//line /usr/local/go/src/strings/replace.go:477
		}
//line /usr/local/go/src/strings/replace.go:477
		// _ = "end of CoverTab[1116]"
	}
//line /usr/local/go/src/strings/replace.go:478
	// _ = "end of CoverTab[1111]"
//line /usr/local/go/src/strings/replace.go:478
	_go_fuzz_dep_.CoverTab[1112]++
							if last != len(s) {
//line /usr/local/go/src/strings/replace.go:479
		_go_fuzz_dep_.CoverTab[1125]++
								nw, err := sw.WriteString(s[last:])
								n += nw
								if err != nil {
//line /usr/local/go/src/strings/replace.go:482
			_go_fuzz_dep_.CoverTab[1126]++
									return n, err
//line /usr/local/go/src/strings/replace.go:483
			// _ = "end of CoverTab[1126]"
		} else {
//line /usr/local/go/src/strings/replace.go:484
			_go_fuzz_dep_.CoverTab[1127]++
//line /usr/local/go/src/strings/replace.go:484
			// _ = "end of CoverTab[1127]"
//line /usr/local/go/src/strings/replace.go:484
		}
//line /usr/local/go/src/strings/replace.go:484
		// _ = "end of CoverTab[1125]"
	} else {
//line /usr/local/go/src/strings/replace.go:485
		_go_fuzz_dep_.CoverTab[1128]++
//line /usr/local/go/src/strings/replace.go:485
		// _ = "end of CoverTab[1128]"
//line /usr/local/go/src/strings/replace.go:485
	}
//line /usr/local/go/src/strings/replace.go:485
	// _ = "end of CoverTab[1112]"
//line /usr/local/go/src/strings/replace.go:485
	_go_fuzz_dep_.CoverTab[1113]++
							return n, nil
//line /usr/local/go/src/strings/replace.go:486
	// _ = "end of CoverTab[1113]"
}

// byteStringReplacer is the implementation that's used when all the
//line /usr/local/go/src/strings/replace.go:489
// "old" values are single ASCII bytes but the "new" values vary in size.
//line /usr/local/go/src/strings/replace.go:491
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
//line /usr/local/go/src/strings/replace.go:501
// at which (*byteStringReplacer).Replace switches algorithms.
//line /usr/local/go/src/strings/replace.go:501
// For strings with higher ration of length to replacements than that value,
//line /usr/local/go/src/strings/replace.go:501
// we call Count, for each replacement from toReplace.
//line /usr/local/go/src/strings/replace.go:501
// For strings, with a lower ratio we use simple loop, because of Count overhead.
//line /usr/local/go/src/strings/replace.go:501
// countCutOff is an empirically determined overhead multiplier.
//line /usr/local/go/src/strings/replace.go:501
// TODO(tocarip) revisit once we have register-based abi/mid-stack inlining.
//line /usr/local/go/src/strings/replace.go:508
const countCutOff = 8

func (r *byteStringReplacer) Replace(s string) string {
//line /usr/local/go/src/strings/replace.go:510
	_go_fuzz_dep_.CoverTab[1129]++
							newSize := len(s)
							anyChanges := false

							if len(r.toReplace)*countCutOff <= len(s) {
//line /usr/local/go/src/strings/replace.go:514
		_go_fuzz_dep_.CoverTab[1133]++
								for _, x := range r.toReplace {
//line /usr/local/go/src/strings/replace.go:515
			_go_fuzz_dep_.CoverTab[1134]++
									if c := Count(s, x); c != 0 {
//line /usr/local/go/src/strings/replace.go:516
				_go_fuzz_dep_.CoverTab[1135]++

										newSize += c * (len(r.replacements[x[0]]) - 1)
										anyChanges = true
//line /usr/local/go/src/strings/replace.go:519
				// _ = "end of CoverTab[1135]"
			} else {
//line /usr/local/go/src/strings/replace.go:520
				_go_fuzz_dep_.CoverTab[1136]++
//line /usr/local/go/src/strings/replace.go:520
				// _ = "end of CoverTab[1136]"
//line /usr/local/go/src/strings/replace.go:520
			}
//line /usr/local/go/src/strings/replace.go:520
			// _ = "end of CoverTab[1134]"

		}
//line /usr/local/go/src/strings/replace.go:522
		// _ = "end of CoverTab[1133]"
	} else {
//line /usr/local/go/src/strings/replace.go:523
		_go_fuzz_dep_.CoverTab[1137]++
								for i := 0; i < len(s); i++ {
//line /usr/local/go/src/strings/replace.go:524
			_go_fuzz_dep_.CoverTab[1138]++
									b := s[i]
									if r.replacements[b] != nil {
//line /usr/local/go/src/strings/replace.go:526
				_go_fuzz_dep_.CoverTab[1139]++

										newSize += len(r.replacements[b]) - 1
										anyChanges = true
//line /usr/local/go/src/strings/replace.go:529
				// _ = "end of CoverTab[1139]"
			} else {
//line /usr/local/go/src/strings/replace.go:530
				_go_fuzz_dep_.CoverTab[1140]++
//line /usr/local/go/src/strings/replace.go:530
				// _ = "end of CoverTab[1140]"
//line /usr/local/go/src/strings/replace.go:530
			}
//line /usr/local/go/src/strings/replace.go:530
			// _ = "end of CoverTab[1138]"
		}
//line /usr/local/go/src/strings/replace.go:531
		// _ = "end of CoverTab[1137]"
	}
//line /usr/local/go/src/strings/replace.go:532
	// _ = "end of CoverTab[1129]"
//line /usr/local/go/src/strings/replace.go:532
	_go_fuzz_dep_.CoverTab[1130]++
							if !anyChanges {
//line /usr/local/go/src/strings/replace.go:533
		_go_fuzz_dep_.CoverTab[1141]++
								return s
//line /usr/local/go/src/strings/replace.go:534
		// _ = "end of CoverTab[1141]"
	} else {
//line /usr/local/go/src/strings/replace.go:535
		_go_fuzz_dep_.CoverTab[1142]++
//line /usr/local/go/src/strings/replace.go:535
		// _ = "end of CoverTab[1142]"
//line /usr/local/go/src/strings/replace.go:535
	}
//line /usr/local/go/src/strings/replace.go:535
	// _ = "end of CoverTab[1130]"
//line /usr/local/go/src/strings/replace.go:535
	_go_fuzz_dep_.CoverTab[1131]++
							buf := make([]byte, newSize)
							j := 0
							for i := 0; i < len(s); i++ {
//line /usr/local/go/src/strings/replace.go:538
		_go_fuzz_dep_.CoverTab[1143]++
								b := s[i]
								if r.replacements[b] != nil {
//line /usr/local/go/src/strings/replace.go:540
			_go_fuzz_dep_.CoverTab[1144]++
									j += copy(buf[j:], r.replacements[b])
//line /usr/local/go/src/strings/replace.go:541
			// _ = "end of CoverTab[1144]"
		} else {
//line /usr/local/go/src/strings/replace.go:542
			_go_fuzz_dep_.CoverTab[1145]++
									buf[j] = b
									j++
//line /usr/local/go/src/strings/replace.go:544
			// _ = "end of CoverTab[1145]"
		}
//line /usr/local/go/src/strings/replace.go:545
		// _ = "end of CoverTab[1143]"
	}
//line /usr/local/go/src/strings/replace.go:546
	// _ = "end of CoverTab[1131]"
//line /usr/local/go/src/strings/replace.go:546
	_go_fuzz_dep_.CoverTab[1132]++
							return string(buf)
//line /usr/local/go/src/strings/replace.go:547
	// _ = "end of CoverTab[1132]"
}

func (r *byteStringReplacer) WriteString(w io.Writer, s string) (n int, err error) {
//line /usr/local/go/src/strings/replace.go:550
	_go_fuzz_dep_.CoverTab[1146]++
							sw := getStringWriter(w)
							last := 0
							for i := 0; i < len(s); i++ {
//line /usr/local/go/src/strings/replace.go:553
		_go_fuzz_dep_.CoverTab[1149]++
								b := s[i]
								if r.replacements[b] == nil {
//line /usr/local/go/src/strings/replace.go:555
			_go_fuzz_dep_.CoverTab[1152]++
									continue
//line /usr/local/go/src/strings/replace.go:556
			// _ = "end of CoverTab[1152]"
		} else {
//line /usr/local/go/src/strings/replace.go:557
			_go_fuzz_dep_.CoverTab[1153]++
//line /usr/local/go/src/strings/replace.go:557
			// _ = "end of CoverTab[1153]"
//line /usr/local/go/src/strings/replace.go:557
		}
//line /usr/local/go/src/strings/replace.go:557
		// _ = "end of CoverTab[1149]"
//line /usr/local/go/src/strings/replace.go:557
		_go_fuzz_dep_.CoverTab[1150]++
								if last != i {
//line /usr/local/go/src/strings/replace.go:558
			_go_fuzz_dep_.CoverTab[1154]++
									nw, err := sw.WriteString(s[last:i])
									n += nw
									if err != nil {
//line /usr/local/go/src/strings/replace.go:561
				_go_fuzz_dep_.CoverTab[1155]++
										return n, err
//line /usr/local/go/src/strings/replace.go:562
				// _ = "end of CoverTab[1155]"
			} else {
//line /usr/local/go/src/strings/replace.go:563
				_go_fuzz_dep_.CoverTab[1156]++
//line /usr/local/go/src/strings/replace.go:563
				// _ = "end of CoverTab[1156]"
//line /usr/local/go/src/strings/replace.go:563
			}
//line /usr/local/go/src/strings/replace.go:563
			// _ = "end of CoverTab[1154]"
		} else {
//line /usr/local/go/src/strings/replace.go:564
			_go_fuzz_dep_.CoverTab[1157]++
//line /usr/local/go/src/strings/replace.go:564
			// _ = "end of CoverTab[1157]"
//line /usr/local/go/src/strings/replace.go:564
		}
//line /usr/local/go/src/strings/replace.go:564
		// _ = "end of CoverTab[1150]"
//line /usr/local/go/src/strings/replace.go:564
		_go_fuzz_dep_.CoverTab[1151]++
								last = i + 1
								nw, err := w.Write(r.replacements[b])
								n += nw
								if err != nil {
//line /usr/local/go/src/strings/replace.go:568
			_go_fuzz_dep_.CoverTab[1158]++
									return n, err
//line /usr/local/go/src/strings/replace.go:569
			// _ = "end of CoverTab[1158]"
		} else {
//line /usr/local/go/src/strings/replace.go:570
			_go_fuzz_dep_.CoverTab[1159]++
//line /usr/local/go/src/strings/replace.go:570
			// _ = "end of CoverTab[1159]"
//line /usr/local/go/src/strings/replace.go:570
		}
//line /usr/local/go/src/strings/replace.go:570
		// _ = "end of CoverTab[1151]"
	}
//line /usr/local/go/src/strings/replace.go:571
	// _ = "end of CoverTab[1146]"
//line /usr/local/go/src/strings/replace.go:571
	_go_fuzz_dep_.CoverTab[1147]++
							if last != len(s) {
//line /usr/local/go/src/strings/replace.go:572
		_go_fuzz_dep_.CoverTab[1160]++
								var nw int
								nw, err = sw.WriteString(s[last:])
								n += nw
//line /usr/local/go/src/strings/replace.go:575
		// _ = "end of CoverTab[1160]"
	} else {
//line /usr/local/go/src/strings/replace.go:576
		_go_fuzz_dep_.CoverTab[1161]++
//line /usr/local/go/src/strings/replace.go:576
		// _ = "end of CoverTab[1161]"
//line /usr/local/go/src/strings/replace.go:576
	}
//line /usr/local/go/src/strings/replace.go:576
	// _ = "end of CoverTab[1147]"
//line /usr/local/go/src/strings/replace.go:576
	_go_fuzz_dep_.CoverTab[1148]++
							return
//line /usr/local/go/src/strings/replace.go:577
	// _ = "end of CoverTab[1148]"
}

//line /usr/local/go/src/strings/replace.go:578
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/strings/replace.go:578
var _ = _go_fuzz_dep_.CoverTab
