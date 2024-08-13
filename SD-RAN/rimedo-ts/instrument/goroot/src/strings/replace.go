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
	_go_fuzz_dep_.CoverTab[3214]++
						if len(oldnew)%2 == 1 {
//line /usr/local/go/src/strings/replace.go:33
		_go_fuzz_dep_.CoverTab[3216]++
							panic("strings.NewReplacer: odd argument count")
//line /usr/local/go/src/strings/replace.go:34
		// _ = "end of CoverTab[3216]"
	} else {
//line /usr/local/go/src/strings/replace.go:35
		_go_fuzz_dep_.CoverTab[3217]++
//line /usr/local/go/src/strings/replace.go:35
		// _ = "end of CoverTab[3217]"
//line /usr/local/go/src/strings/replace.go:35
	}
//line /usr/local/go/src/strings/replace.go:35
	// _ = "end of CoverTab[3214]"
//line /usr/local/go/src/strings/replace.go:35
	_go_fuzz_dep_.CoverTab[3215]++
						return &Replacer{oldnew: append([]string(nil), oldnew...)}
//line /usr/local/go/src/strings/replace.go:36
	// _ = "end of CoverTab[3215]"
}

func (r *Replacer) buildOnce() {
//line /usr/local/go/src/strings/replace.go:39
	_go_fuzz_dep_.CoverTab[3218]++
						r.r = r.build()
						r.oldnew = nil
//line /usr/local/go/src/strings/replace.go:41
	// _ = "end of CoverTab[3218]"
}

func (b *Replacer) build() replacer {
//line /usr/local/go/src/strings/replace.go:44
	_go_fuzz_dep_.CoverTab[3219]++
						oldnew := b.oldnew
						if len(oldnew) == 2 && func() bool {
//line /usr/local/go/src/strings/replace.go:46
		_go_fuzz_dep_.CoverTab[3224]++
//line /usr/local/go/src/strings/replace.go:46
		return len(oldnew[0]) > 1
//line /usr/local/go/src/strings/replace.go:46
		// _ = "end of CoverTab[3224]"
//line /usr/local/go/src/strings/replace.go:46
	}() {
//line /usr/local/go/src/strings/replace.go:46
		_go_fuzz_dep_.CoverTab[3225]++
							return makeSingleStringReplacer(oldnew[0], oldnew[1])
//line /usr/local/go/src/strings/replace.go:47
		// _ = "end of CoverTab[3225]"
	} else {
//line /usr/local/go/src/strings/replace.go:48
		_go_fuzz_dep_.CoverTab[3226]++
//line /usr/local/go/src/strings/replace.go:48
		// _ = "end of CoverTab[3226]"
//line /usr/local/go/src/strings/replace.go:48
	}
//line /usr/local/go/src/strings/replace.go:48
	// _ = "end of CoverTab[3219]"
//line /usr/local/go/src/strings/replace.go:48
	_go_fuzz_dep_.CoverTab[3220]++

						allNewBytes := true
						for i := 0; i < len(oldnew); i += 2 {
//line /usr/local/go/src/strings/replace.go:51
		_go_fuzz_dep_.CoverTab[3227]++
							if len(oldnew[i]) != 1 {
//line /usr/local/go/src/strings/replace.go:52
			_go_fuzz_dep_.CoverTab[3229]++
								return makeGenericReplacer(oldnew)
//line /usr/local/go/src/strings/replace.go:53
			// _ = "end of CoverTab[3229]"
		} else {
//line /usr/local/go/src/strings/replace.go:54
			_go_fuzz_dep_.CoverTab[3230]++
//line /usr/local/go/src/strings/replace.go:54
			// _ = "end of CoverTab[3230]"
//line /usr/local/go/src/strings/replace.go:54
		}
//line /usr/local/go/src/strings/replace.go:54
		// _ = "end of CoverTab[3227]"
//line /usr/local/go/src/strings/replace.go:54
		_go_fuzz_dep_.CoverTab[3228]++
							if len(oldnew[i+1]) != 1 {
//line /usr/local/go/src/strings/replace.go:55
			_go_fuzz_dep_.CoverTab[3231]++
								allNewBytes = false
//line /usr/local/go/src/strings/replace.go:56
			// _ = "end of CoverTab[3231]"
		} else {
//line /usr/local/go/src/strings/replace.go:57
			_go_fuzz_dep_.CoverTab[3232]++
//line /usr/local/go/src/strings/replace.go:57
			// _ = "end of CoverTab[3232]"
//line /usr/local/go/src/strings/replace.go:57
		}
//line /usr/local/go/src/strings/replace.go:57
		// _ = "end of CoverTab[3228]"
	}
//line /usr/local/go/src/strings/replace.go:58
	// _ = "end of CoverTab[3220]"
//line /usr/local/go/src/strings/replace.go:58
	_go_fuzz_dep_.CoverTab[3221]++

						if allNewBytes {
//line /usr/local/go/src/strings/replace.go:60
		_go_fuzz_dep_.CoverTab[3233]++
							r := byteReplacer{}
							for i := range r {
//line /usr/local/go/src/strings/replace.go:62
			_go_fuzz_dep_.CoverTab[3236]++
								r[i] = byte(i)
//line /usr/local/go/src/strings/replace.go:63
			// _ = "end of CoverTab[3236]"
		}
//line /usr/local/go/src/strings/replace.go:64
		// _ = "end of CoverTab[3233]"
//line /usr/local/go/src/strings/replace.go:64
		_go_fuzz_dep_.CoverTab[3234]++

//line /usr/local/go/src/strings/replace.go:67
		for i := len(oldnew) - 2; i >= 0; i -= 2 {
//line /usr/local/go/src/strings/replace.go:67
			_go_fuzz_dep_.CoverTab[3237]++
								o := oldnew[i][0]
								n := oldnew[i+1][0]
								r[o] = n
//line /usr/local/go/src/strings/replace.go:70
			// _ = "end of CoverTab[3237]"
		}
//line /usr/local/go/src/strings/replace.go:71
		// _ = "end of CoverTab[3234]"
//line /usr/local/go/src/strings/replace.go:71
		_go_fuzz_dep_.CoverTab[3235]++
							return &r
//line /usr/local/go/src/strings/replace.go:72
		// _ = "end of CoverTab[3235]"
	} else {
//line /usr/local/go/src/strings/replace.go:73
		_go_fuzz_dep_.CoverTab[3238]++
//line /usr/local/go/src/strings/replace.go:73
		// _ = "end of CoverTab[3238]"
//line /usr/local/go/src/strings/replace.go:73
	}
//line /usr/local/go/src/strings/replace.go:73
	// _ = "end of CoverTab[3221]"
//line /usr/local/go/src/strings/replace.go:73
	_go_fuzz_dep_.CoverTab[3222]++

						r := byteStringReplacer{toReplace: make([]string, 0, len(oldnew)/2)}

//line /usr/local/go/src/strings/replace.go:78
	for i := len(oldnew) - 2; i >= 0; i -= 2 {
//line /usr/local/go/src/strings/replace.go:78
		_go_fuzz_dep_.CoverTab[3239]++
							o := oldnew[i][0]
							n := oldnew[i+1]

							if r.replacements[o] == nil {
//line /usr/local/go/src/strings/replace.go:82
			_go_fuzz_dep_.CoverTab[3241]++

//line /usr/local/go/src/strings/replace.go:86
			r.toReplace = append(r.toReplace, string([]byte{o}))
//line /usr/local/go/src/strings/replace.go:86
			// _ = "end of CoverTab[3241]"
		} else {
//line /usr/local/go/src/strings/replace.go:87
			_go_fuzz_dep_.CoverTab[3242]++
//line /usr/local/go/src/strings/replace.go:87
			// _ = "end of CoverTab[3242]"
//line /usr/local/go/src/strings/replace.go:87
		}
//line /usr/local/go/src/strings/replace.go:87
		// _ = "end of CoverTab[3239]"
//line /usr/local/go/src/strings/replace.go:87
		_go_fuzz_dep_.CoverTab[3240]++
							r.replacements[o] = []byte(n)
//line /usr/local/go/src/strings/replace.go:88
		// _ = "end of CoverTab[3240]"

	}
//line /usr/local/go/src/strings/replace.go:90
	// _ = "end of CoverTab[3222]"
//line /usr/local/go/src/strings/replace.go:90
	_go_fuzz_dep_.CoverTab[3223]++
						return &r
//line /usr/local/go/src/strings/replace.go:91
	// _ = "end of CoverTab[3223]"
}

// Replace returns a copy of s with all replacements performed.
func (r *Replacer) Replace(s string) string {
//line /usr/local/go/src/strings/replace.go:95
	_go_fuzz_dep_.CoverTab[3243]++
						r.once.Do(r.buildOnce)
						return r.r.Replace(s)
//line /usr/local/go/src/strings/replace.go:97
	// _ = "end of CoverTab[3243]"
}

// WriteString writes s to w with all replacements performed.
func (r *Replacer) WriteString(w io.Writer, s string) (n int, err error) {
//line /usr/local/go/src/strings/replace.go:101
	_go_fuzz_dep_.CoverTab[3244]++
							r.once.Do(r.buildOnce)
							return r.r.WriteString(w, s)
//line /usr/local/go/src/strings/replace.go:103
	// _ = "end of CoverTab[3244]"
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
	_go_fuzz_dep_.CoverTab[3245]++
							if key == "" {
//line /usr/local/go/src/strings/replace.go:159
		_go_fuzz_dep_.CoverTab[3247]++
								if t.priority == 0 {
//line /usr/local/go/src/strings/replace.go:160
			_go_fuzz_dep_.CoverTab[3249]++
									t.value = val
									t.priority = priority
//line /usr/local/go/src/strings/replace.go:162
			// _ = "end of CoverTab[3249]"
		} else {
//line /usr/local/go/src/strings/replace.go:163
			_go_fuzz_dep_.CoverTab[3250]++
//line /usr/local/go/src/strings/replace.go:163
			// _ = "end of CoverTab[3250]"
//line /usr/local/go/src/strings/replace.go:163
		}
//line /usr/local/go/src/strings/replace.go:163
		// _ = "end of CoverTab[3247]"
//line /usr/local/go/src/strings/replace.go:163
		_go_fuzz_dep_.CoverTab[3248]++
								return
//line /usr/local/go/src/strings/replace.go:164
		// _ = "end of CoverTab[3248]"
	} else {
//line /usr/local/go/src/strings/replace.go:165
		_go_fuzz_dep_.CoverTab[3251]++
//line /usr/local/go/src/strings/replace.go:165
		// _ = "end of CoverTab[3251]"
//line /usr/local/go/src/strings/replace.go:165
	}
//line /usr/local/go/src/strings/replace.go:165
	// _ = "end of CoverTab[3245]"
//line /usr/local/go/src/strings/replace.go:165
	_go_fuzz_dep_.CoverTab[3246]++

							if t.prefix != "" {
//line /usr/local/go/src/strings/replace.go:167
		_go_fuzz_dep_.CoverTab[3252]++
		// Need to split the prefix among multiple nodes.
		var n int	// length of the longest common prefix
		for ; n < len(t.prefix) && func() bool {
//line /usr/local/go/src/strings/replace.go:170
			_go_fuzz_dep_.CoverTab[3254]++
//line /usr/local/go/src/strings/replace.go:170
			return n < len(key)
//line /usr/local/go/src/strings/replace.go:170
			// _ = "end of CoverTab[3254]"
//line /usr/local/go/src/strings/replace.go:170
		}(); n++ {
//line /usr/local/go/src/strings/replace.go:170
			_go_fuzz_dep_.CoverTab[3255]++
									if t.prefix[n] != key[n] {
//line /usr/local/go/src/strings/replace.go:171
				_go_fuzz_dep_.CoverTab[3256]++
										break
//line /usr/local/go/src/strings/replace.go:172
				// _ = "end of CoverTab[3256]"
			} else {
//line /usr/local/go/src/strings/replace.go:173
				_go_fuzz_dep_.CoverTab[3257]++
//line /usr/local/go/src/strings/replace.go:173
				// _ = "end of CoverTab[3257]"
//line /usr/local/go/src/strings/replace.go:173
			}
//line /usr/local/go/src/strings/replace.go:173
			// _ = "end of CoverTab[3255]"
		}
//line /usr/local/go/src/strings/replace.go:174
		// _ = "end of CoverTab[3252]"
//line /usr/local/go/src/strings/replace.go:174
		_go_fuzz_dep_.CoverTab[3253]++
								if n == len(t.prefix) {
//line /usr/local/go/src/strings/replace.go:175
			_go_fuzz_dep_.CoverTab[3258]++
									t.next.add(key[n:], val, priority, r)
//line /usr/local/go/src/strings/replace.go:176
			// _ = "end of CoverTab[3258]"
		} else {
//line /usr/local/go/src/strings/replace.go:177
			_go_fuzz_dep_.CoverTab[3259]++
//line /usr/local/go/src/strings/replace.go:177
			if n == 0 {
//line /usr/local/go/src/strings/replace.go:177
				_go_fuzz_dep_.CoverTab[3260]++
				// First byte differs, start a new lookup table here. Looking up
				// what is currently t.prefix[0] will lead to prefixNode, and
				// looking up key[0] will lead to keyNode.
				var prefixNode *trieNode
				if len(t.prefix) == 1 {
//line /usr/local/go/src/strings/replace.go:182
					_go_fuzz_dep_.CoverTab[3262]++
											prefixNode = t.next
//line /usr/local/go/src/strings/replace.go:183
					// _ = "end of CoverTab[3262]"
				} else {
//line /usr/local/go/src/strings/replace.go:184
					_go_fuzz_dep_.CoverTab[3263]++
											prefixNode = &trieNode{
						prefix:	t.prefix[1:],
						next:	t.next,
					}
//line /usr/local/go/src/strings/replace.go:188
					// _ = "end of CoverTab[3263]"
				}
//line /usr/local/go/src/strings/replace.go:189
				// _ = "end of CoverTab[3260]"
//line /usr/local/go/src/strings/replace.go:189
				_go_fuzz_dep_.CoverTab[3261]++
										keyNode := new(trieNode)
										t.table = make([]*trieNode, r.tableSize)
										t.table[r.mapping[t.prefix[0]]] = prefixNode
										t.table[r.mapping[key[0]]] = keyNode
										t.prefix = ""
										t.next = nil
										keyNode.add(key[1:], val, priority, r)
//line /usr/local/go/src/strings/replace.go:196
				// _ = "end of CoverTab[3261]"
			} else {
//line /usr/local/go/src/strings/replace.go:197
				_go_fuzz_dep_.CoverTab[3264]++

										next := &trieNode{
					prefix:	t.prefix[n:],
					next:	t.next,
				}
										t.prefix = t.prefix[:n]
										t.next = next
										next.add(key[n:], val, priority, r)
//line /usr/local/go/src/strings/replace.go:205
				// _ = "end of CoverTab[3264]"
			}
//line /usr/local/go/src/strings/replace.go:206
			// _ = "end of CoverTab[3259]"
//line /usr/local/go/src/strings/replace.go:206
		}
//line /usr/local/go/src/strings/replace.go:206
		// _ = "end of CoverTab[3253]"
	} else {
//line /usr/local/go/src/strings/replace.go:207
		_go_fuzz_dep_.CoverTab[3265]++
//line /usr/local/go/src/strings/replace.go:207
		if t.table != nil {
//line /usr/local/go/src/strings/replace.go:207
			_go_fuzz_dep_.CoverTab[3266]++

									m := r.mapping[key[0]]
									if t.table[m] == nil {
//line /usr/local/go/src/strings/replace.go:210
				_go_fuzz_dep_.CoverTab[3268]++
										t.table[m] = new(trieNode)
//line /usr/local/go/src/strings/replace.go:211
				// _ = "end of CoverTab[3268]"
			} else {
//line /usr/local/go/src/strings/replace.go:212
				_go_fuzz_dep_.CoverTab[3269]++
//line /usr/local/go/src/strings/replace.go:212
				// _ = "end of CoverTab[3269]"
//line /usr/local/go/src/strings/replace.go:212
			}
//line /usr/local/go/src/strings/replace.go:212
			// _ = "end of CoverTab[3266]"
//line /usr/local/go/src/strings/replace.go:212
			_go_fuzz_dep_.CoverTab[3267]++
									t.table[m].add(key[1:], val, priority, r)
//line /usr/local/go/src/strings/replace.go:213
			// _ = "end of CoverTab[3267]"
		} else {
//line /usr/local/go/src/strings/replace.go:214
			_go_fuzz_dep_.CoverTab[3270]++
									t.prefix = key
									t.next = new(trieNode)
									t.next.add("", val, priority, r)
//line /usr/local/go/src/strings/replace.go:217
			// _ = "end of CoverTab[3270]"
		}
//line /usr/local/go/src/strings/replace.go:218
		// _ = "end of CoverTab[3265]"
//line /usr/local/go/src/strings/replace.go:218
	}
//line /usr/local/go/src/strings/replace.go:218
	// _ = "end of CoverTab[3246]"
}

func (r *genericReplacer) lookup(s string, ignoreRoot bool) (val string, keylen int, found bool) {
//line /usr/local/go/src/strings/replace.go:221
	_go_fuzz_dep_.CoverTab[3271]++

//line /usr/local/go/src/strings/replace.go:224
	bestPriority := 0
	node := &r.root
	n := 0
	for node != nil {
//line /usr/local/go/src/strings/replace.go:227
		_go_fuzz_dep_.CoverTab[3273]++
								if node.priority > bestPriority && func() bool {
//line /usr/local/go/src/strings/replace.go:228
			_go_fuzz_dep_.CoverTab[3276]++
//line /usr/local/go/src/strings/replace.go:228
			return !(ignoreRoot && func() bool {
//line /usr/local/go/src/strings/replace.go:228
				_go_fuzz_dep_.CoverTab[3277]++
//line /usr/local/go/src/strings/replace.go:228
				return node == &r.root
//line /usr/local/go/src/strings/replace.go:228
				// _ = "end of CoverTab[3277]"
//line /usr/local/go/src/strings/replace.go:228
			}())
//line /usr/local/go/src/strings/replace.go:228
			// _ = "end of CoverTab[3276]"
//line /usr/local/go/src/strings/replace.go:228
		}() {
//line /usr/local/go/src/strings/replace.go:228
			_go_fuzz_dep_.CoverTab[3278]++
									bestPriority = node.priority
									val = node.value
									keylen = n
									found = true
//line /usr/local/go/src/strings/replace.go:232
			// _ = "end of CoverTab[3278]"
		} else {
//line /usr/local/go/src/strings/replace.go:233
			_go_fuzz_dep_.CoverTab[3279]++
//line /usr/local/go/src/strings/replace.go:233
			// _ = "end of CoverTab[3279]"
//line /usr/local/go/src/strings/replace.go:233
		}
//line /usr/local/go/src/strings/replace.go:233
		// _ = "end of CoverTab[3273]"
//line /usr/local/go/src/strings/replace.go:233
		_go_fuzz_dep_.CoverTab[3274]++

								if s == "" {
//line /usr/local/go/src/strings/replace.go:235
			_go_fuzz_dep_.CoverTab[3280]++
									break
//line /usr/local/go/src/strings/replace.go:236
			// _ = "end of CoverTab[3280]"
		} else {
//line /usr/local/go/src/strings/replace.go:237
			_go_fuzz_dep_.CoverTab[3281]++
//line /usr/local/go/src/strings/replace.go:237
			// _ = "end of CoverTab[3281]"
//line /usr/local/go/src/strings/replace.go:237
		}
//line /usr/local/go/src/strings/replace.go:237
		// _ = "end of CoverTab[3274]"
//line /usr/local/go/src/strings/replace.go:237
		_go_fuzz_dep_.CoverTab[3275]++
								if node.table != nil {
//line /usr/local/go/src/strings/replace.go:238
			_go_fuzz_dep_.CoverTab[3282]++
									index := r.mapping[s[0]]
									if int(index) == r.tableSize {
//line /usr/local/go/src/strings/replace.go:240
				_go_fuzz_dep_.CoverTab[3284]++
										break
//line /usr/local/go/src/strings/replace.go:241
				// _ = "end of CoverTab[3284]"
			} else {
//line /usr/local/go/src/strings/replace.go:242
				_go_fuzz_dep_.CoverTab[3285]++
//line /usr/local/go/src/strings/replace.go:242
				// _ = "end of CoverTab[3285]"
//line /usr/local/go/src/strings/replace.go:242
			}
//line /usr/local/go/src/strings/replace.go:242
			// _ = "end of CoverTab[3282]"
//line /usr/local/go/src/strings/replace.go:242
			_go_fuzz_dep_.CoverTab[3283]++
									node = node.table[index]
									s = s[1:]
									n++
//line /usr/local/go/src/strings/replace.go:245
			// _ = "end of CoverTab[3283]"
		} else {
//line /usr/local/go/src/strings/replace.go:246
			_go_fuzz_dep_.CoverTab[3286]++
//line /usr/local/go/src/strings/replace.go:246
			if node.prefix != "" && func() bool {
//line /usr/local/go/src/strings/replace.go:246
				_go_fuzz_dep_.CoverTab[3287]++
//line /usr/local/go/src/strings/replace.go:246
				return HasPrefix(s, node.prefix)
//line /usr/local/go/src/strings/replace.go:246
				// _ = "end of CoverTab[3287]"
//line /usr/local/go/src/strings/replace.go:246
			}() {
//line /usr/local/go/src/strings/replace.go:246
				_go_fuzz_dep_.CoverTab[3288]++
										n += len(node.prefix)
										s = s[len(node.prefix):]
										node = node.next
//line /usr/local/go/src/strings/replace.go:249
				// _ = "end of CoverTab[3288]"
			} else {
//line /usr/local/go/src/strings/replace.go:250
				_go_fuzz_dep_.CoverTab[3289]++
										break
//line /usr/local/go/src/strings/replace.go:251
				// _ = "end of CoverTab[3289]"
			}
//line /usr/local/go/src/strings/replace.go:252
			// _ = "end of CoverTab[3286]"
//line /usr/local/go/src/strings/replace.go:252
		}
//line /usr/local/go/src/strings/replace.go:252
		// _ = "end of CoverTab[3275]"
	}
//line /usr/local/go/src/strings/replace.go:253
	// _ = "end of CoverTab[3271]"
//line /usr/local/go/src/strings/replace.go:253
	_go_fuzz_dep_.CoverTab[3272]++
							return
//line /usr/local/go/src/strings/replace.go:254
	// _ = "end of CoverTab[3272]"
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
	_go_fuzz_dep_.CoverTab[3290]++
							r := new(genericReplacer)

							for i := 0; i < len(oldnew); i += 2 {
//line /usr/local/go/src/strings/replace.go:271
		_go_fuzz_dep_.CoverTab[3295]++
								key := oldnew[i]
								for j := 0; j < len(key); j++ {
//line /usr/local/go/src/strings/replace.go:273
			_go_fuzz_dep_.CoverTab[3296]++
									r.mapping[key[j]] = 1
//line /usr/local/go/src/strings/replace.go:274
			// _ = "end of CoverTab[3296]"
		}
//line /usr/local/go/src/strings/replace.go:275
		// _ = "end of CoverTab[3295]"
	}
//line /usr/local/go/src/strings/replace.go:276
	// _ = "end of CoverTab[3290]"
//line /usr/local/go/src/strings/replace.go:276
	_go_fuzz_dep_.CoverTab[3291]++

							for _, b := range r.mapping {
//line /usr/local/go/src/strings/replace.go:278
		_go_fuzz_dep_.CoverTab[3297]++
								r.tableSize += int(b)
//line /usr/local/go/src/strings/replace.go:279
		// _ = "end of CoverTab[3297]"
	}
//line /usr/local/go/src/strings/replace.go:280
	// _ = "end of CoverTab[3291]"
//line /usr/local/go/src/strings/replace.go:280
	_go_fuzz_dep_.CoverTab[3292]++

							var index byte
							for i, b := range r.mapping {
//line /usr/local/go/src/strings/replace.go:283
		_go_fuzz_dep_.CoverTab[3298]++
								if b == 0 {
//line /usr/local/go/src/strings/replace.go:284
			_go_fuzz_dep_.CoverTab[3299]++
									r.mapping[i] = byte(r.tableSize)
//line /usr/local/go/src/strings/replace.go:285
			// _ = "end of CoverTab[3299]"
		} else {
//line /usr/local/go/src/strings/replace.go:286
			_go_fuzz_dep_.CoverTab[3300]++
									r.mapping[i] = index
									index++
//line /usr/local/go/src/strings/replace.go:288
			// _ = "end of CoverTab[3300]"
		}
//line /usr/local/go/src/strings/replace.go:289
		// _ = "end of CoverTab[3298]"
	}
//line /usr/local/go/src/strings/replace.go:290
	// _ = "end of CoverTab[3292]"
//line /usr/local/go/src/strings/replace.go:290
	_go_fuzz_dep_.CoverTab[3293]++

							r.root.table = make([]*trieNode, r.tableSize)

							for i := 0; i < len(oldnew); i += 2 {
//line /usr/local/go/src/strings/replace.go:294
		_go_fuzz_dep_.CoverTab[3301]++
								r.root.add(oldnew[i], oldnew[i+1], len(oldnew)-i, r)
//line /usr/local/go/src/strings/replace.go:295
		// _ = "end of CoverTab[3301]"
	}
//line /usr/local/go/src/strings/replace.go:296
	// _ = "end of CoverTab[3293]"
//line /usr/local/go/src/strings/replace.go:296
	_go_fuzz_dep_.CoverTab[3294]++
							return r
//line /usr/local/go/src/strings/replace.go:297
	// _ = "end of CoverTab[3294]"
}

type appendSliceWriter []byte

// Write writes to the buffer to satisfy io.Writer.
func (w *appendSliceWriter) Write(p []byte) (int, error) {
//line /usr/local/go/src/strings/replace.go:303
	_go_fuzz_dep_.CoverTab[3302]++
							*w = append(*w, p...)
							return len(p), nil
//line /usr/local/go/src/strings/replace.go:305
	// _ = "end of CoverTab[3302]"
}

// WriteString writes to the buffer without string->[]byte->string allocations.
func (w *appendSliceWriter) WriteString(s string) (int, error) {
//line /usr/local/go/src/strings/replace.go:309
	_go_fuzz_dep_.CoverTab[3303]++
							*w = append(*w, s...)
							return len(s), nil
//line /usr/local/go/src/strings/replace.go:311
	// _ = "end of CoverTab[3303]"
}

type stringWriter struct {
	w io.Writer
}

func (w stringWriter) WriteString(s string) (int, error) {
//line /usr/local/go/src/strings/replace.go:318
	_go_fuzz_dep_.CoverTab[3304]++
							return w.w.Write([]byte(s))
//line /usr/local/go/src/strings/replace.go:319
	// _ = "end of CoverTab[3304]"
}

func getStringWriter(w io.Writer) io.StringWriter {
//line /usr/local/go/src/strings/replace.go:322
	_go_fuzz_dep_.CoverTab[3305]++
							sw, ok := w.(io.StringWriter)
							if !ok {
//line /usr/local/go/src/strings/replace.go:324
		_go_fuzz_dep_.CoverTab[3307]++
								sw = stringWriter{w}
//line /usr/local/go/src/strings/replace.go:325
		// _ = "end of CoverTab[3307]"
	} else {
//line /usr/local/go/src/strings/replace.go:326
		_go_fuzz_dep_.CoverTab[3308]++
//line /usr/local/go/src/strings/replace.go:326
		// _ = "end of CoverTab[3308]"
//line /usr/local/go/src/strings/replace.go:326
	}
//line /usr/local/go/src/strings/replace.go:326
	// _ = "end of CoverTab[3305]"
//line /usr/local/go/src/strings/replace.go:326
	_go_fuzz_dep_.CoverTab[3306]++
							return sw
//line /usr/local/go/src/strings/replace.go:327
	// _ = "end of CoverTab[3306]"
}

func (r *genericReplacer) Replace(s string) string {
//line /usr/local/go/src/strings/replace.go:330
	_go_fuzz_dep_.CoverTab[3309]++
							buf := make(appendSliceWriter, 0, len(s))
							r.WriteString(&buf, s)
							return string(buf)
//line /usr/local/go/src/strings/replace.go:333
	// _ = "end of CoverTab[3309]"
}

func (r *genericReplacer) WriteString(w io.Writer, s string) (n int, err error) {
//line /usr/local/go/src/strings/replace.go:336
	_go_fuzz_dep_.CoverTab[3310]++
							sw := getStringWriter(w)
							var last, wn int
							var prevMatchEmpty bool
							for i := 0; i <= len(s); {
//line /usr/local/go/src/strings/replace.go:340
		_go_fuzz_dep_.CoverTab[3313]++

								if i != len(s) && func() bool {
//line /usr/local/go/src/strings/replace.go:342
			_go_fuzz_dep_.CoverTab[3316]++
//line /usr/local/go/src/strings/replace.go:342
			return r.root.priority == 0
//line /usr/local/go/src/strings/replace.go:342
			// _ = "end of CoverTab[3316]"
//line /usr/local/go/src/strings/replace.go:342
		}() {
//line /usr/local/go/src/strings/replace.go:342
			_go_fuzz_dep_.CoverTab[3317]++
									index := int(r.mapping[s[i]])
									if index == r.tableSize || func() bool {
//line /usr/local/go/src/strings/replace.go:344
				_go_fuzz_dep_.CoverTab[3318]++
//line /usr/local/go/src/strings/replace.go:344
				return r.root.table[index] == nil
//line /usr/local/go/src/strings/replace.go:344
				// _ = "end of CoverTab[3318]"
//line /usr/local/go/src/strings/replace.go:344
			}() {
//line /usr/local/go/src/strings/replace.go:344
				_go_fuzz_dep_.CoverTab[3319]++
										i++
										continue
//line /usr/local/go/src/strings/replace.go:346
				// _ = "end of CoverTab[3319]"
			} else {
//line /usr/local/go/src/strings/replace.go:347
				_go_fuzz_dep_.CoverTab[3320]++
//line /usr/local/go/src/strings/replace.go:347
				// _ = "end of CoverTab[3320]"
//line /usr/local/go/src/strings/replace.go:347
			}
//line /usr/local/go/src/strings/replace.go:347
			// _ = "end of CoverTab[3317]"
		} else {
//line /usr/local/go/src/strings/replace.go:348
			_go_fuzz_dep_.CoverTab[3321]++
//line /usr/local/go/src/strings/replace.go:348
			// _ = "end of CoverTab[3321]"
//line /usr/local/go/src/strings/replace.go:348
		}
//line /usr/local/go/src/strings/replace.go:348
		// _ = "end of CoverTab[3313]"
//line /usr/local/go/src/strings/replace.go:348
		_go_fuzz_dep_.CoverTab[3314]++

//line /usr/local/go/src/strings/replace.go:351
		val, keylen, match := r.lookup(s[i:], prevMatchEmpty)
		prevMatchEmpty = match && func() bool {
//line /usr/local/go/src/strings/replace.go:352
			_go_fuzz_dep_.CoverTab[3322]++
//line /usr/local/go/src/strings/replace.go:352
			return keylen == 0
//line /usr/local/go/src/strings/replace.go:352
			// _ = "end of CoverTab[3322]"
//line /usr/local/go/src/strings/replace.go:352
		}()
								if match {
//line /usr/local/go/src/strings/replace.go:353
			_go_fuzz_dep_.CoverTab[3323]++
									wn, err = sw.WriteString(s[last:i])
									n += wn
									if err != nil {
//line /usr/local/go/src/strings/replace.go:356
				_go_fuzz_dep_.CoverTab[3326]++
										return
//line /usr/local/go/src/strings/replace.go:357
				// _ = "end of CoverTab[3326]"
			} else {
//line /usr/local/go/src/strings/replace.go:358
				_go_fuzz_dep_.CoverTab[3327]++
//line /usr/local/go/src/strings/replace.go:358
				// _ = "end of CoverTab[3327]"
//line /usr/local/go/src/strings/replace.go:358
			}
//line /usr/local/go/src/strings/replace.go:358
			// _ = "end of CoverTab[3323]"
//line /usr/local/go/src/strings/replace.go:358
			_go_fuzz_dep_.CoverTab[3324]++
									wn, err = sw.WriteString(val)
									n += wn
									if err != nil {
//line /usr/local/go/src/strings/replace.go:361
				_go_fuzz_dep_.CoverTab[3328]++
										return
//line /usr/local/go/src/strings/replace.go:362
				// _ = "end of CoverTab[3328]"
			} else {
//line /usr/local/go/src/strings/replace.go:363
				_go_fuzz_dep_.CoverTab[3329]++
//line /usr/local/go/src/strings/replace.go:363
				// _ = "end of CoverTab[3329]"
//line /usr/local/go/src/strings/replace.go:363
			}
//line /usr/local/go/src/strings/replace.go:363
			// _ = "end of CoverTab[3324]"
//line /usr/local/go/src/strings/replace.go:363
			_go_fuzz_dep_.CoverTab[3325]++
									i += keylen
									last = i
									continue
//line /usr/local/go/src/strings/replace.go:366
			// _ = "end of CoverTab[3325]"
		} else {
//line /usr/local/go/src/strings/replace.go:367
			_go_fuzz_dep_.CoverTab[3330]++
//line /usr/local/go/src/strings/replace.go:367
			// _ = "end of CoverTab[3330]"
//line /usr/local/go/src/strings/replace.go:367
		}
//line /usr/local/go/src/strings/replace.go:367
		// _ = "end of CoverTab[3314]"
//line /usr/local/go/src/strings/replace.go:367
		_go_fuzz_dep_.CoverTab[3315]++
								i++
//line /usr/local/go/src/strings/replace.go:368
		// _ = "end of CoverTab[3315]"
	}
//line /usr/local/go/src/strings/replace.go:369
	// _ = "end of CoverTab[3310]"
//line /usr/local/go/src/strings/replace.go:369
	_go_fuzz_dep_.CoverTab[3311]++
							if last != len(s) {
//line /usr/local/go/src/strings/replace.go:370
		_go_fuzz_dep_.CoverTab[3331]++
								wn, err = sw.WriteString(s[last:])
								n += wn
//line /usr/local/go/src/strings/replace.go:372
		// _ = "end of CoverTab[3331]"
	} else {
//line /usr/local/go/src/strings/replace.go:373
		_go_fuzz_dep_.CoverTab[3332]++
//line /usr/local/go/src/strings/replace.go:373
		// _ = "end of CoverTab[3332]"
//line /usr/local/go/src/strings/replace.go:373
	}
//line /usr/local/go/src/strings/replace.go:373
	// _ = "end of CoverTab[3311]"
//line /usr/local/go/src/strings/replace.go:373
	_go_fuzz_dep_.CoverTab[3312]++
							return
//line /usr/local/go/src/strings/replace.go:374
	// _ = "end of CoverTab[3312]"
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
	_go_fuzz_dep_.CoverTab[3333]++
							return &singleStringReplacer{finder: makeStringFinder(pattern), value: value}
//line /usr/local/go/src/strings/replace.go:386
	// _ = "end of CoverTab[3333]"
}

func (r *singleStringReplacer) Replace(s string) string {
//line /usr/local/go/src/strings/replace.go:389
	_go_fuzz_dep_.CoverTab[3334]++
							var buf Builder
							i, matched := 0, false
							for {
//line /usr/local/go/src/strings/replace.go:392
		_go_fuzz_dep_.CoverTab[3337]++
								match := r.finder.next(s[i:])
								if match == -1 {
//line /usr/local/go/src/strings/replace.go:394
			_go_fuzz_dep_.CoverTab[3339]++
									break
//line /usr/local/go/src/strings/replace.go:395
			// _ = "end of CoverTab[3339]"
		} else {
//line /usr/local/go/src/strings/replace.go:396
			_go_fuzz_dep_.CoverTab[3340]++
//line /usr/local/go/src/strings/replace.go:396
			// _ = "end of CoverTab[3340]"
//line /usr/local/go/src/strings/replace.go:396
		}
//line /usr/local/go/src/strings/replace.go:396
		// _ = "end of CoverTab[3337]"
//line /usr/local/go/src/strings/replace.go:396
		_go_fuzz_dep_.CoverTab[3338]++
								matched = true
								buf.Grow(match + len(r.value))
								buf.WriteString(s[i : i+match])
								buf.WriteString(r.value)
								i += match + len(r.finder.pattern)
//line /usr/local/go/src/strings/replace.go:401
		// _ = "end of CoverTab[3338]"
	}
//line /usr/local/go/src/strings/replace.go:402
	// _ = "end of CoverTab[3334]"
//line /usr/local/go/src/strings/replace.go:402
	_go_fuzz_dep_.CoverTab[3335]++
							if !matched {
//line /usr/local/go/src/strings/replace.go:403
		_go_fuzz_dep_.CoverTab[3341]++
								return s
//line /usr/local/go/src/strings/replace.go:404
		// _ = "end of CoverTab[3341]"
	} else {
//line /usr/local/go/src/strings/replace.go:405
		_go_fuzz_dep_.CoverTab[3342]++
//line /usr/local/go/src/strings/replace.go:405
		// _ = "end of CoverTab[3342]"
//line /usr/local/go/src/strings/replace.go:405
	}
//line /usr/local/go/src/strings/replace.go:405
	// _ = "end of CoverTab[3335]"
//line /usr/local/go/src/strings/replace.go:405
	_go_fuzz_dep_.CoverTab[3336]++
							buf.WriteString(s[i:])
							return buf.String()
//line /usr/local/go/src/strings/replace.go:407
	// _ = "end of CoverTab[3336]"
}

func (r *singleStringReplacer) WriteString(w io.Writer, s string) (n int, err error) {
//line /usr/local/go/src/strings/replace.go:410
	_go_fuzz_dep_.CoverTab[3343]++
							sw := getStringWriter(w)
							var i, wn int
							for {
//line /usr/local/go/src/strings/replace.go:413
		_go_fuzz_dep_.CoverTab[3345]++
								match := r.finder.next(s[i:])
								if match == -1 {
//line /usr/local/go/src/strings/replace.go:415
			_go_fuzz_dep_.CoverTab[3349]++
									break
//line /usr/local/go/src/strings/replace.go:416
			// _ = "end of CoverTab[3349]"
		} else {
//line /usr/local/go/src/strings/replace.go:417
			_go_fuzz_dep_.CoverTab[3350]++
//line /usr/local/go/src/strings/replace.go:417
			// _ = "end of CoverTab[3350]"
//line /usr/local/go/src/strings/replace.go:417
		}
//line /usr/local/go/src/strings/replace.go:417
		// _ = "end of CoverTab[3345]"
//line /usr/local/go/src/strings/replace.go:417
		_go_fuzz_dep_.CoverTab[3346]++
								wn, err = sw.WriteString(s[i : i+match])
								n += wn
								if err != nil {
//line /usr/local/go/src/strings/replace.go:420
			_go_fuzz_dep_.CoverTab[3351]++
									return
//line /usr/local/go/src/strings/replace.go:421
			// _ = "end of CoverTab[3351]"
		} else {
//line /usr/local/go/src/strings/replace.go:422
			_go_fuzz_dep_.CoverTab[3352]++
//line /usr/local/go/src/strings/replace.go:422
			// _ = "end of CoverTab[3352]"
//line /usr/local/go/src/strings/replace.go:422
		}
//line /usr/local/go/src/strings/replace.go:422
		// _ = "end of CoverTab[3346]"
//line /usr/local/go/src/strings/replace.go:422
		_go_fuzz_dep_.CoverTab[3347]++
								wn, err = sw.WriteString(r.value)
								n += wn
								if err != nil {
//line /usr/local/go/src/strings/replace.go:425
			_go_fuzz_dep_.CoverTab[3353]++
									return
//line /usr/local/go/src/strings/replace.go:426
			// _ = "end of CoverTab[3353]"
		} else {
//line /usr/local/go/src/strings/replace.go:427
			_go_fuzz_dep_.CoverTab[3354]++
//line /usr/local/go/src/strings/replace.go:427
			// _ = "end of CoverTab[3354]"
//line /usr/local/go/src/strings/replace.go:427
		}
//line /usr/local/go/src/strings/replace.go:427
		// _ = "end of CoverTab[3347]"
//line /usr/local/go/src/strings/replace.go:427
		_go_fuzz_dep_.CoverTab[3348]++
								i += match + len(r.finder.pattern)
//line /usr/local/go/src/strings/replace.go:428
		// _ = "end of CoverTab[3348]"
	}
//line /usr/local/go/src/strings/replace.go:429
	// _ = "end of CoverTab[3343]"
//line /usr/local/go/src/strings/replace.go:429
	_go_fuzz_dep_.CoverTab[3344]++
							wn, err = sw.WriteString(s[i:])
							n += wn
							return
//line /usr/local/go/src/strings/replace.go:432
	// _ = "end of CoverTab[3344]"
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
	_go_fuzz_dep_.CoverTab[3355]++
							var buf []byte	// lazily allocated
							for i := 0; i < len(s); i++ {
//line /usr/local/go/src/strings/replace.go:442
		_go_fuzz_dep_.CoverTab[3358]++
								b := s[i]
								if r[b] != b {
//line /usr/local/go/src/strings/replace.go:444
			_go_fuzz_dep_.CoverTab[3359]++
									if buf == nil {
//line /usr/local/go/src/strings/replace.go:445
				_go_fuzz_dep_.CoverTab[3361]++
										buf = []byte(s)
//line /usr/local/go/src/strings/replace.go:446
				// _ = "end of CoverTab[3361]"
			} else {
//line /usr/local/go/src/strings/replace.go:447
				_go_fuzz_dep_.CoverTab[3362]++
//line /usr/local/go/src/strings/replace.go:447
				// _ = "end of CoverTab[3362]"
//line /usr/local/go/src/strings/replace.go:447
			}
//line /usr/local/go/src/strings/replace.go:447
			// _ = "end of CoverTab[3359]"
//line /usr/local/go/src/strings/replace.go:447
			_go_fuzz_dep_.CoverTab[3360]++
									buf[i] = r[b]
//line /usr/local/go/src/strings/replace.go:448
			// _ = "end of CoverTab[3360]"
		} else {
//line /usr/local/go/src/strings/replace.go:449
			_go_fuzz_dep_.CoverTab[3363]++
//line /usr/local/go/src/strings/replace.go:449
			// _ = "end of CoverTab[3363]"
//line /usr/local/go/src/strings/replace.go:449
		}
//line /usr/local/go/src/strings/replace.go:449
		// _ = "end of CoverTab[3358]"
	}
//line /usr/local/go/src/strings/replace.go:450
	// _ = "end of CoverTab[3355]"
//line /usr/local/go/src/strings/replace.go:450
	_go_fuzz_dep_.CoverTab[3356]++
							if buf == nil {
//line /usr/local/go/src/strings/replace.go:451
		_go_fuzz_dep_.CoverTab[3364]++
								return s
//line /usr/local/go/src/strings/replace.go:452
		// _ = "end of CoverTab[3364]"
	} else {
//line /usr/local/go/src/strings/replace.go:453
		_go_fuzz_dep_.CoverTab[3365]++
//line /usr/local/go/src/strings/replace.go:453
		// _ = "end of CoverTab[3365]"
//line /usr/local/go/src/strings/replace.go:453
	}
//line /usr/local/go/src/strings/replace.go:453
	// _ = "end of CoverTab[3356]"
//line /usr/local/go/src/strings/replace.go:453
	_go_fuzz_dep_.CoverTab[3357]++
							return string(buf)
//line /usr/local/go/src/strings/replace.go:454
	// _ = "end of CoverTab[3357]"
}

func (r *byteReplacer) WriteString(w io.Writer, s string) (n int, err error) {
//line /usr/local/go/src/strings/replace.go:457
	_go_fuzz_dep_.CoverTab[3366]++
							sw := getStringWriter(w)
							last := 0
							for i := 0; i < len(s); i++ {
//line /usr/local/go/src/strings/replace.go:460
		_go_fuzz_dep_.CoverTab[3369]++
								b := s[i]
								if r[b] == b {
//line /usr/local/go/src/strings/replace.go:462
			_go_fuzz_dep_.CoverTab[3372]++
									continue
//line /usr/local/go/src/strings/replace.go:463
			// _ = "end of CoverTab[3372]"
		} else {
//line /usr/local/go/src/strings/replace.go:464
			_go_fuzz_dep_.CoverTab[3373]++
//line /usr/local/go/src/strings/replace.go:464
			// _ = "end of CoverTab[3373]"
//line /usr/local/go/src/strings/replace.go:464
		}
//line /usr/local/go/src/strings/replace.go:464
		// _ = "end of CoverTab[3369]"
//line /usr/local/go/src/strings/replace.go:464
		_go_fuzz_dep_.CoverTab[3370]++
								if last != i {
//line /usr/local/go/src/strings/replace.go:465
			_go_fuzz_dep_.CoverTab[3374]++
									wn, err := sw.WriteString(s[last:i])
									n += wn
									if err != nil {
//line /usr/local/go/src/strings/replace.go:468
				_go_fuzz_dep_.CoverTab[3375]++
										return n, err
//line /usr/local/go/src/strings/replace.go:469
				// _ = "end of CoverTab[3375]"
			} else {
//line /usr/local/go/src/strings/replace.go:470
				_go_fuzz_dep_.CoverTab[3376]++
//line /usr/local/go/src/strings/replace.go:470
				// _ = "end of CoverTab[3376]"
//line /usr/local/go/src/strings/replace.go:470
			}
//line /usr/local/go/src/strings/replace.go:470
			// _ = "end of CoverTab[3374]"
		} else {
//line /usr/local/go/src/strings/replace.go:471
			_go_fuzz_dep_.CoverTab[3377]++
//line /usr/local/go/src/strings/replace.go:471
			// _ = "end of CoverTab[3377]"
//line /usr/local/go/src/strings/replace.go:471
		}
//line /usr/local/go/src/strings/replace.go:471
		// _ = "end of CoverTab[3370]"
//line /usr/local/go/src/strings/replace.go:471
		_go_fuzz_dep_.CoverTab[3371]++
								last = i + 1
								nw, err := w.Write(r[b : int(b)+1])
								n += nw
								if err != nil {
//line /usr/local/go/src/strings/replace.go:475
			_go_fuzz_dep_.CoverTab[3378]++
									return n, err
//line /usr/local/go/src/strings/replace.go:476
			// _ = "end of CoverTab[3378]"
		} else {
//line /usr/local/go/src/strings/replace.go:477
			_go_fuzz_dep_.CoverTab[3379]++
//line /usr/local/go/src/strings/replace.go:477
			// _ = "end of CoverTab[3379]"
//line /usr/local/go/src/strings/replace.go:477
		}
//line /usr/local/go/src/strings/replace.go:477
		// _ = "end of CoverTab[3371]"
	}
//line /usr/local/go/src/strings/replace.go:478
	// _ = "end of CoverTab[3366]"
//line /usr/local/go/src/strings/replace.go:478
	_go_fuzz_dep_.CoverTab[3367]++
							if last != len(s) {
//line /usr/local/go/src/strings/replace.go:479
		_go_fuzz_dep_.CoverTab[3380]++
								nw, err := sw.WriteString(s[last:])
								n += nw
								if err != nil {
//line /usr/local/go/src/strings/replace.go:482
			_go_fuzz_dep_.CoverTab[3381]++
									return n, err
//line /usr/local/go/src/strings/replace.go:483
			// _ = "end of CoverTab[3381]"
		} else {
//line /usr/local/go/src/strings/replace.go:484
			_go_fuzz_dep_.CoverTab[3382]++
//line /usr/local/go/src/strings/replace.go:484
			// _ = "end of CoverTab[3382]"
//line /usr/local/go/src/strings/replace.go:484
		}
//line /usr/local/go/src/strings/replace.go:484
		// _ = "end of CoverTab[3380]"
	} else {
//line /usr/local/go/src/strings/replace.go:485
		_go_fuzz_dep_.CoverTab[3383]++
//line /usr/local/go/src/strings/replace.go:485
		// _ = "end of CoverTab[3383]"
//line /usr/local/go/src/strings/replace.go:485
	}
//line /usr/local/go/src/strings/replace.go:485
	// _ = "end of CoverTab[3367]"
//line /usr/local/go/src/strings/replace.go:485
	_go_fuzz_dep_.CoverTab[3368]++
							return n, nil
//line /usr/local/go/src/strings/replace.go:486
	// _ = "end of CoverTab[3368]"
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
	_go_fuzz_dep_.CoverTab[3384]++
							newSize := len(s)
							anyChanges := false

							if len(r.toReplace)*countCutOff <= len(s) {
//line /usr/local/go/src/strings/replace.go:514
		_go_fuzz_dep_.CoverTab[3388]++
								for _, x := range r.toReplace {
//line /usr/local/go/src/strings/replace.go:515
			_go_fuzz_dep_.CoverTab[3389]++
									if c := Count(s, x); c != 0 {
//line /usr/local/go/src/strings/replace.go:516
				_go_fuzz_dep_.CoverTab[3390]++

										newSize += c * (len(r.replacements[x[0]]) - 1)
										anyChanges = true
//line /usr/local/go/src/strings/replace.go:519
				// _ = "end of CoverTab[3390]"
			} else {
//line /usr/local/go/src/strings/replace.go:520
				_go_fuzz_dep_.CoverTab[3391]++
//line /usr/local/go/src/strings/replace.go:520
				// _ = "end of CoverTab[3391]"
//line /usr/local/go/src/strings/replace.go:520
			}
//line /usr/local/go/src/strings/replace.go:520
			// _ = "end of CoverTab[3389]"

		}
//line /usr/local/go/src/strings/replace.go:522
		// _ = "end of CoverTab[3388]"
	} else {
//line /usr/local/go/src/strings/replace.go:523
		_go_fuzz_dep_.CoverTab[3392]++
								for i := 0; i < len(s); i++ {
//line /usr/local/go/src/strings/replace.go:524
			_go_fuzz_dep_.CoverTab[3393]++
									b := s[i]
									if r.replacements[b] != nil {
//line /usr/local/go/src/strings/replace.go:526
				_go_fuzz_dep_.CoverTab[3394]++

										newSize += len(r.replacements[b]) - 1
										anyChanges = true
//line /usr/local/go/src/strings/replace.go:529
				// _ = "end of CoverTab[3394]"
			} else {
//line /usr/local/go/src/strings/replace.go:530
				_go_fuzz_dep_.CoverTab[3395]++
//line /usr/local/go/src/strings/replace.go:530
				// _ = "end of CoverTab[3395]"
//line /usr/local/go/src/strings/replace.go:530
			}
//line /usr/local/go/src/strings/replace.go:530
			// _ = "end of CoverTab[3393]"
		}
//line /usr/local/go/src/strings/replace.go:531
		// _ = "end of CoverTab[3392]"
	}
//line /usr/local/go/src/strings/replace.go:532
	// _ = "end of CoverTab[3384]"
//line /usr/local/go/src/strings/replace.go:532
	_go_fuzz_dep_.CoverTab[3385]++
							if !anyChanges {
//line /usr/local/go/src/strings/replace.go:533
		_go_fuzz_dep_.CoverTab[3396]++
								return s
//line /usr/local/go/src/strings/replace.go:534
		// _ = "end of CoverTab[3396]"
	} else {
//line /usr/local/go/src/strings/replace.go:535
		_go_fuzz_dep_.CoverTab[3397]++
//line /usr/local/go/src/strings/replace.go:535
		// _ = "end of CoverTab[3397]"
//line /usr/local/go/src/strings/replace.go:535
	}
//line /usr/local/go/src/strings/replace.go:535
	// _ = "end of CoverTab[3385]"
//line /usr/local/go/src/strings/replace.go:535
	_go_fuzz_dep_.CoverTab[3386]++
							buf := make([]byte, newSize)
							j := 0
							for i := 0; i < len(s); i++ {
//line /usr/local/go/src/strings/replace.go:538
		_go_fuzz_dep_.CoverTab[3398]++
								b := s[i]
								if r.replacements[b] != nil {
//line /usr/local/go/src/strings/replace.go:540
			_go_fuzz_dep_.CoverTab[3399]++
									j += copy(buf[j:], r.replacements[b])
//line /usr/local/go/src/strings/replace.go:541
			// _ = "end of CoverTab[3399]"
		} else {
//line /usr/local/go/src/strings/replace.go:542
			_go_fuzz_dep_.CoverTab[3400]++
									buf[j] = b
									j++
//line /usr/local/go/src/strings/replace.go:544
			// _ = "end of CoverTab[3400]"
		}
//line /usr/local/go/src/strings/replace.go:545
		// _ = "end of CoverTab[3398]"
	}
//line /usr/local/go/src/strings/replace.go:546
	// _ = "end of CoverTab[3386]"
//line /usr/local/go/src/strings/replace.go:546
	_go_fuzz_dep_.CoverTab[3387]++
							return string(buf)
//line /usr/local/go/src/strings/replace.go:547
	// _ = "end of CoverTab[3387]"
}

func (r *byteStringReplacer) WriteString(w io.Writer, s string) (n int, err error) {
//line /usr/local/go/src/strings/replace.go:550
	_go_fuzz_dep_.CoverTab[3401]++
							sw := getStringWriter(w)
							last := 0
							for i := 0; i < len(s); i++ {
//line /usr/local/go/src/strings/replace.go:553
		_go_fuzz_dep_.CoverTab[3404]++
								b := s[i]
								if r.replacements[b] == nil {
//line /usr/local/go/src/strings/replace.go:555
			_go_fuzz_dep_.CoverTab[3407]++
									continue
//line /usr/local/go/src/strings/replace.go:556
			// _ = "end of CoverTab[3407]"
		} else {
//line /usr/local/go/src/strings/replace.go:557
			_go_fuzz_dep_.CoverTab[3408]++
//line /usr/local/go/src/strings/replace.go:557
			// _ = "end of CoverTab[3408]"
//line /usr/local/go/src/strings/replace.go:557
		}
//line /usr/local/go/src/strings/replace.go:557
		// _ = "end of CoverTab[3404]"
//line /usr/local/go/src/strings/replace.go:557
		_go_fuzz_dep_.CoverTab[3405]++
								if last != i {
//line /usr/local/go/src/strings/replace.go:558
			_go_fuzz_dep_.CoverTab[3409]++
									nw, err := sw.WriteString(s[last:i])
									n += nw
									if err != nil {
//line /usr/local/go/src/strings/replace.go:561
				_go_fuzz_dep_.CoverTab[3410]++
										return n, err
//line /usr/local/go/src/strings/replace.go:562
				// _ = "end of CoverTab[3410]"
			} else {
//line /usr/local/go/src/strings/replace.go:563
				_go_fuzz_dep_.CoverTab[3411]++
//line /usr/local/go/src/strings/replace.go:563
				// _ = "end of CoverTab[3411]"
//line /usr/local/go/src/strings/replace.go:563
			}
//line /usr/local/go/src/strings/replace.go:563
			// _ = "end of CoverTab[3409]"
		} else {
//line /usr/local/go/src/strings/replace.go:564
			_go_fuzz_dep_.CoverTab[3412]++
//line /usr/local/go/src/strings/replace.go:564
			// _ = "end of CoverTab[3412]"
//line /usr/local/go/src/strings/replace.go:564
		}
//line /usr/local/go/src/strings/replace.go:564
		// _ = "end of CoverTab[3405]"
//line /usr/local/go/src/strings/replace.go:564
		_go_fuzz_dep_.CoverTab[3406]++
								last = i + 1
								nw, err := w.Write(r.replacements[b])
								n += nw
								if err != nil {
//line /usr/local/go/src/strings/replace.go:568
			_go_fuzz_dep_.CoverTab[3413]++
									return n, err
//line /usr/local/go/src/strings/replace.go:569
			// _ = "end of CoverTab[3413]"
		} else {
//line /usr/local/go/src/strings/replace.go:570
			_go_fuzz_dep_.CoverTab[3414]++
//line /usr/local/go/src/strings/replace.go:570
			// _ = "end of CoverTab[3414]"
//line /usr/local/go/src/strings/replace.go:570
		}
//line /usr/local/go/src/strings/replace.go:570
		// _ = "end of CoverTab[3406]"
	}
//line /usr/local/go/src/strings/replace.go:571
	// _ = "end of CoverTab[3401]"
//line /usr/local/go/src/strings/replace.go:571
	_go_fuzz_dep_.CoverTab[3402]++
							if last != len(s) {
//line /usr/local/go/src/strings/replace.go:572
		_go_fuzz_dep_.CoverTab[3415]++
								var nw int
								nw, err = sw.WriteString(s[last:])
								n += nw
//line /usr/local/go/src/strings/replace.go:575
		// _ = "end of CoverTab[3415]"
	} else {
//line /usr/local/go/src/strings/replace.go:576
		_go_fuzz_dep_.CoverTab[3416]++
//line /usr/local/go/src/strings/replace.go:576
		// _ = "end of CoverTab[3416]"
//line /usr/local/go/src/strings/replace.go:576
	}
//line /usr/local/go/src/strings/replace.go:576
	// _ = "end of CoverTab[3402]"
//line /usr/local/go/src/strings/replace.go:576
	_go_fuzz_dep_.CoverTab[3403]++
							return
//line /usr/local/go/src/strings/replace.go:577
	// _ = "end of CoverTab[3403]"
}

//line /usr/local/go/src/strings/replace.go:578
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/strings/replace.go:578
var _ = _go_fuzz_dep_.CoverTab
