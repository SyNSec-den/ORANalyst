//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:1
package ndr

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:1
)

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// intFromTag returns an int that is a value in a struct tag key/value pair
func intFromTag(tag reflect.StructTag, key string) (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:11
	_go_fuzz_dep_.CoverTab[86748]++
											ndrTag := parseTags(tag)
											d := 1
											if n, ok := ndrTag.Map[key]; ok {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:14
		_go_fuzz_dep_.CoverTab[86750]++
												i, err := strconv.Atoi(n)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:16
			_go_fuzz_dep_.CoverTab[86752]++
													return d, fmt.Errorf("invalid dimensions tag [%s]: %v", n, err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:17
			// _ = "end of CoverTab[86752]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:18
			_go_fuzz_dep_.CoverTab[86753]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:18
			// _ = "end of CoverTab[86753]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:18
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:18
		// _ = "end of CoverTab[86750]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:18
		_go_fuzz_dep_.CoverTab[86751]++
												d = i
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:19
		// _ = "end of CoverTab[86751]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:20
		_go_fuzz_dep_.CoverTab[86754]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:20
		// _ = "end of CoverTab[86754]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:20
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:20
	// _ = "end of CoverTab[86748]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:20
	_go_fuzz_dep_.CoverTab[86749]++
											return d, nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:21
	// _ = "end of CoverTab[86749]"
}

// parseDimensions returns the a slice of the size of each dimension and type of the member at the deepest level.
func parseDimensions(v reflect.Value) (l []int, tb reflect.Type) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:25
	_go_fuzz_dep_.CoverTab[86755]++
											if v.Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:26
		_go_fuzz_dep_.CoverTab[86760]++
												v = v.Elem()
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:27
		// _ = "end of CoverTab[86760]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:28
		_go_fuzz_dep_.CoverTab[86761]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:28
		// _ = "end of CoverTab[86761]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:28
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:28
	// _ = "end of CoverTab[86755]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:28
	_go_fuzz_dep_.CoverTab[86756]++
											t := v.Type()
											if t.Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:30
		_go_fuzz_dep_.CoverTab[86762]++
												t = t.Elem()
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:31
		// _ = "end of CoverTab[86762]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:32
		_go_fuzz_dep_.CoverTab[86763]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:32
		// _ = "end of CoverTab[86763]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:32
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:32
	// _ = "end of CoverTab[86756]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:32
	_go_fuzz_dep_.CoverTab[86757]++
											if t.Kind() != reflect.Array && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:33
		_go_fuzz_dep_.CoverTab[86764]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:33
		return t.Kind() != reflect.Slice
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:33
		// _ = "end of CoverTab[86764]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:33
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:33
		_go_fuzz_dep_.CoverTab[86765]++
												return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:34
		// _ = "end of CoverTab[86765]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:35
		_go_fuzz_dep_.CoverTab[86766]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:35
		// _ = "end of CoverTab[86766]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:35
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:35
	// _ = "end of CoverTab[86757]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:35
	_go_fuzz_dep_.CoverTab[86758]++
											l = append(l, v.Len())
											if t.Elem().Kind() == reflect.Array || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:37
		_go_fuzz_dep_.CoverTab[86767]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:37
		return t.Elem().Kind() == reflect.Slice
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:37
		// _ = "end of CoverTab[86767]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:37
	}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:37
		_go_fuzz_dep_.CoverTab[86768]++
												// contains array or slice
												var m []int
												m, tb = parseDimensions(v.Index(0))
												l = append(l, m...)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:41
		// _ = "end of CoverTab[86768]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:42
		_go_fuzz_dep_.CoverTab[86769]++
												tb = t.Elem()
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:43
		// _ = "end of CoverTab[86769]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:44
	// _ = "end of CoverTab[86758]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:44
	_go_fuzz_dep_.CoverTab[86759]++
											return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:45
	// _ = "end of CoverTab[86759]"
}

// sliceDimensions returns the count of dimensions a slice has.
func sliceDimensions(t reflect.Type) (d int, tb reflect.Type) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:49
	_go_fuzz_dep_.CoverTab[86770]++
											if t.Kind() == reflect.Ptr {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:50
		_go_fuzz_dep_.CoverTab[86773]++
												t = t.Elem()
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:51
		// _ = "end of CoverTab[86773]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:52
		_go_fuzz_dep_.CoverTab[86774]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:52
		// _ = "end of CoverTab[86774]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:52
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:52
	// _ = "end of CoverTab[86770]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:52
	_go_fuzz_dep_.CoverTab[86771]++
											if t.Kind() == reflect.Slice {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:53
		_go_fuzz_dep_.CoverTab[86775]++
												d++
												var n int
												n, tb = sliceDimensions(t.Elem())
												d += n
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:57
		// _ = "end of CoverTab[86775]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:58
		_go_fuzz_dep_.CoverTab[86776]++
												tb = t
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:59
		// _ = "end of CoverTab[86776]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:60
	// _ = "end of CoverTab[86771]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:60
	_go_fuzz_dep_.CoverTab[86772]++
											return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:61
	// _ = "end of CoverTab[86772]"
}

// makeSubSlices is a deep recursive creation/initialisation of multi-dimensional slices.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:64
// Takes the reflect.Value of the 1st dimension and a slice of the lengths of the sub dimensions
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:66
func makeSubSlices(v reflect.Value, l []int) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:66
	_go_fuzz_dep_.CoverTab[86777]++
											ty := v.Type().Elem()
											if ty.Kind() != reflect.Slice {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:68
		_go_fuzz_dep_.CoverTab[86780]++
												return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:69
		// _ = "end of CoverTab[86780]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:70
		_go_fuzz_dep_.CoverTab[86781]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:70
		// _ = "end of CoverTab[86781]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:70
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:70
	// _ = "end of CoverTab[86777]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:70
	_go_fuzz_dep_.CoverTab[86778]++
											for i := 0; i < v.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:71
		_go_fuzz_dep_.CoverTab[86782]++
												s := reflect.MakeSlice(ty, l[0], l[0])
												v.Index(i).Set(s)

												if len(l) > 1 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:75
			_go_fuzz_dep_.CoverTab[86783]++
													makeSubSlices(v.Index(i), l[1:])
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:76
			// _ = "end of CoverTab[86783]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:77
			_go_fuzz_dep_.CoverTab[86784]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:77
			// _ = "end of CoverTab[86784]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:77
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:77
		// _ = "end of CoverTab[86782]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:78
	// _ = "end of CoverTab[86778]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:78
	_go_fuzz_dep_.CoverTab[86779]++
											return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:79
	// _ = "end of CoverTab[86779]"
}

// multiDimensionalIndexPermutations returns all the permutations of the indexes of a multi-dimensional slice.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:82
// The input is a slice of integers that indicates the max size/length of each dimension
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:84
func multiDimensionalIndexPermutations(l []int) (ps [][]int) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:84
	_go_fuzz_dep_.CoverTab[86785]++
											z := make([]int, len(l), len(l))
											ps = append(ps, z)

											for i := len(l) - 1; i >= 0; i-- {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:88
		_go_fuzz_dep_.CoverTab[86787]++
												ws := make([][]int, len(ps))
												copy(ws, ps)

												for j := 1; j <= l[i]-1; j++ {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:92
			_go_fuzz_dep_.CoverTab[86788]++

													for _, p := range ws {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:94
				_go_fuzz_dep_.CoverTab[86789]++
														np := make([]int, len(p), len(p))
														copy(np, p)
														np[i] = j
														ps = append(ps, np)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:98
				// _ = "end of CoverTab[86789]"
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:99
			// _ = "end of CoverTab[86788]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:100
		// _ = "end of CoverTab[86787]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:101
	// _ = "end of CoverTab[86785]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:101
	_go_fuzz_dep_.CoverTab[86786]++
											return
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:102
	// _ = "end of CoverTab[86786]"
}

// precedingMax reads off the next conformant max value
func (dec *Decoder) precedingMax() uint32 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:106
	_go_fuzz_dep_.CoverTab[86790]++
											m := dec.conformantMax[0]
											dec.conformantMax = dec.conformantMax[1:]
											return m
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:109
	// _ = "end of CoverTab[86790]"
}

// fillFixedArray establishes if the fixed array is uni or multi dimensional and then fills it.
func (dec *Decoder) fillFixedArray(v reflect.Value, tag reflect.StructTag, def *[]deferedPtr) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:113
	_go_fuzz_dep_.CoverTab[86791]++
											l, t := parseDimensions(v)
											if t.Kind() == reflect.String {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:115
		_go_fuzz_dep_.CoverTab[86796]++
												tag = reflect.StructTag(subStringArrayTag)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:116
		// _ = "end of CoverTab[86796]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:117
		_go_fuzz_dep_.CoverTab[86797]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:117
		// _ = "end of CoverTab[86797]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:117
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:117
	// _ = "end of CoverTab[86791]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:117
	_go_fuzz_dep_.CoverTab[86792]++
											if len(l) < 1 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:118
		_go_fuzz_dep_.CoverTab[86798]++
												return errors.New("could not establish dimensions of fixed array")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:119
		// _ = "end of CoverTab[86798]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:120
		_go_fuzz_dep_.CoverTab[86799]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:120
		// _ = "end of CoverTab[86799]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:120
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:120
	// _ = "end of CoverTab[86792]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:120
	_go_fuzz_dep_.CoverTab[86793]++
											if len(l) == 1 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:121
		_go_fuzz_dep_.CoverTab[86800]++
												err := dec.fillUniDimensionalFixedArray(v, tag, def)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:123
			_go_fuzz_dep_.CoverTab[86802]++
													return fmt.Errorf("could not fill uni-dimensional fixed array: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:124
			// _ = "end of CoverTab[86802]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:125
			_go_fuzz_dep_.CoverTab[86803]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:125
			// _ = "end of CoverTab[86803]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:125
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:125
		// _ = "end of CoverTab[86800]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:125
		_go_fuzz_dep_.CoverTab[86801]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:126
		// _ = "end of CoverTab[86801]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:127
		_go_fuzz_dep_.CoverTab[86804]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:127
		// _ = "end of CoverTab[86804]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:127
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:127
	// _ = "end of CoverTab[86793]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:127
	_go_fuzz_dep_.CoverTab[86794]++

											ps := multiDimensionalIndexPermutations(l[:len(l)-1])
											for _, p := range ps {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:130
		_go_fuzz_dep_.CoverTab[86805]++

												a := v
												for _, i := range p {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:133
			_go_fuzz_dep_.CoverTab[86807]++
													a = a.Index(i)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:134
			// _ = "end of CoverTab[86807]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:135
		// _ = "end of CoverTab[86805]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:135
		_go_fuzz_dep_.CoverTab[86806]++

												err := dec.fillUniDimensionalFixedArray(a, tag, def)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:138
			_go_fuzz_dep_.CoverTab[86808]++
													return fmt.Errorf("could not fill dimension %v of multi-dimensional fixed array: %v", p, err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:139
			// _ = "end of CoverTab[86808]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:140
			_go_fuzz_dep_.CoverTab[86809]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:140
			// _ = "end of CoverTab[86809]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:140
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:140
		// _ = "end of CoverTab[86806]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:141
	// _ = "end of CoverTab[86794]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:141
	_go_fuzz_dep_.CoverTab[86795]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:142
	// _ = "end of CoverTab[86795]"
}

// readUniDimensionalFixedArray reads an array (not slice) from the byte stream.
func (dec *Decoder) fillUniDimensionalFixedArray(v reflect.Value, tag reflect.StructTag, def *[]deferedPtr) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:146
	_go_fuzz_dep_.CoverTab[86810]++
											for i := 0; i < v.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:147
		_go_fuzz_dep_.CoverTab[86812]++
												err := dec.fill(v.Index(i), tag, def)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:149
			_go_fuzz_dep_.CoverTab[86813]++
													return fmt.Errorf("could not fill index %d of fixed array: %v", i, err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:150
			// _ = "end of CoverTab[86813]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:151
			_go_fuzz_dep_.CoverTab[86814]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:151
			// _ = "end of CoverTab[86814]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:151
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:151
		// _ = "end of CoverTab[86812]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:152
	// _ = "end of CoverTab[86810]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:152
	_go_fuzz_dep_.CoverTab[86811]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:153
	// _ = "end of CoverTab[86811]"
}

// fillConformantArray establishes if the conformant array is uni or multi dimensional and then fills the slice.
func (dec *Decoder) fillConformantArray(v reflect.Value, tag reflect.StructTag, def *[]deferedPtr) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:157
	_go_fuzz_dep_.CoverTab[86815]++
											d, _ := sliceDimensions(v.Type())
											if d > 1 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:159
		_go_fuzz_dep_.CoverTab[86817]++
												err := dec.fillMultiDimensionalConformantArray(v, d, tag, def)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:161
			_go_fuzz_dep_.CoverTab[86818]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:162
			// _ = "end of CoverTab[86818]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:163
			_go_fuzz_dep_.CoverTab[86819]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:163
			// _ = "end of CoverTab[86819]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:163
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:163
		// _ = "end of CoverTab[86817]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:164
		_go_fuzz_dep_.CoverTab[86820]++
												err := dec.fillUniDimensionalConformantArray(v, tag, def)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:166
			_go_fuzz_dep_.CoverTab[86821]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:167
			// _ = "end of CoverTab[86821]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:168
			_go_fuzz_dep_.CoverTab[86822]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:168
			// _ = "end of CoverTab[86822]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:168
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:168
		// _ = "end of CoverTab[86820]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:169
	// _ = "end of CoverTab[86815]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:169
	_go_fuzz_dep_.CoverTab[86816]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:170
	// _ = "end of CoverTab[86816]"
}

// fillUniDimensionalConformantArray fills the uni-dimensional slice value.
func (dec *Decoder) fillUniDimensionalConformantArray(v reflect.Value, tag reflect.StructTag, def *[]deferedPtr) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:174
	_go_fuzz_dep_.CoverTab[86823]++
											m := dec.precedingMax()
											n := int(m)
											a := reflect.MakeSlice(v.Type(), n, n)
											for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:178
		_go_fuzz_dep_.CoverTab[86825]++
												err := dec.fill(a.Index(i), tag, def)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:180
			_go_fuzz_dep_.CoverTab[86826]++
													return fmt.Errorf("could not fill index %d of uni-dimensional conformant array: %v", i, err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:181
			// _ = "end of CoverTab[86826]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:182
			_go_fuzz_dep_.CoverTab[86827]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:182
			// _ = "end of CoverTab[86827]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:182
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:182
		// _ = "end of CoverTab[86825]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:183
	// _ = "end of CoverTab[86823]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:183
	_go_fuzz_dep_.CoverTab[86824]++
											v.Set(a)
											return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:185
	// _ = "end of CoverTab[86824]"
}

// fillMultiDimensionalConformantArray fills the multi-dimensional slice value provided from conformant array data.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:188
// The number of dimensions must be specified. This must be less than or equal to the dimensions in the slice for this
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:188
// method not to panic.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:191
func (dec *Decoder) fillMultiDimensionalConformantArray(v reflect.Value, d int, tag reflect.StructTag, def *[]deferedPtr) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:191
	_go_fuzz_dep_.CoverTab[86828]++

											l := make([]int, d, d)
											for i := range l {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:194
		_go_fuzz_dep_.CoverTab[86831]++
												l[i] = int(dec.precedingMax())
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:195
		// _ = "end of CoverTab[86831]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:196
	// _ = "end of CoverTab[86828]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:196
	_go_fuzz_dep_.CoverTab[86829]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:199
	ty := v.Type()
											v.Set(reflect.MakeSlice(ty, l[0], l[0]))

											makeSubSlices(v, l[1:])

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:205
	ps := multiDimensionalIndexPermutations(l)
	for _, p := range ps {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:206
		_go_fuzz_dep_.CoverTab[86832]++

												a := v
												for _, i := range p {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:209
			_go_fuzz_dep_.CoverTab[86834]++
													a = a.Index(i)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:210
			// _ = "end of CoverTab[86834]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:211
		// _ = "end of CoverTab[86832]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:211
		_go_fuzz_dep_.CoverTab[86833]++
												err := dec.fill(a, tag, def)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:213
			_go_fuzz_dep_.CoverTab[86835]++
													return fmt.Errorf("could not fill index %v of slice: %v", p, err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:214
			// _ = "end of CoverTab[86835]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:215
			_go_fuzz_dep_.CoverTab[86836]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:215
			// _ = "end of CoverTab[86836]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:215
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:215
		// _ = "end of CoverTab[86833]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:216
	// _ = "end of CoverTab[86829]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:216
	_go_fuzz_dep_.CoverTab[86830]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:217
	// _ = "end of CoverTab[86830]"
}

// fillVaryingArray establishes if the varying array is uni or multi dimensional and then fills the slice.
func (dec *Decoder) fillVaryingArray(v reflect.Value, tag reflect.StructTag, def *[]deferedPtr) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:221
	_go_fuzz_dep_.CoverTab[86837]++
											d, t := sliceDimensions(v.Type())
											if d > 1 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:223
		_go_fuzz_dep_.CoverTab[86839]++
												err := dec.fillMultiDimensionalVaryingArray(v, t, d, tag, def)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:225
			_go_fuzz_dep_.CoverTab[86840]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:226
			// _ = "end of CoverTab[86840]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:227
			_go_fuzz_dep_.CoverTab[86841]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:227
			// _ = "end of CoverTab[86841]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:227
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:227
		// _ = "end of CoverTab[86839]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:228
		_go_fuzz_dep_.CoverTab[86842]++
												err := dec.fillUniDimensionalVaryingArray(v, tag, def)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:230
			_go_fuzz_dep_.CoverTab[86843]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:231
			// _ = "end of CoverTab[86843]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:232
			_go_fuzz_dep_.CoverTab[86844]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:232
			// _ = "end of CoverTab[86844]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:232
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:232
		// _ = "end of CoverTab[86842]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:233
	// _ = "end of CoverTab[86837]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:233
	_go_fuzz_dep_.CoverTab[86838]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:234
	// _ = "end of CoverTab[86838]"
}

// fillUniDimensionalVaryingArray fills the uni-dimensional slice value.
func (dec *Decoder) fillUniDimensionalVaryingArray(v reflect.Value, tag reflect.StructTag, def *[]deferedPtr) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:238
	_go_fuzz_dep_.CoverTab[86845]++
											o, err := dec.readUint32()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:240
		_go_fuzz_dep_.CoverTab[86849]++
												return fmt.Errorf("could not read offset of uni-dimensional varying array: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:241
		// _ = "end of CoverTab[86849]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:242
		_go_fuzz_dep_.CoverTab[86850]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:242
		// _ = "end of CoverTab[86850]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:242
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:242
	// _ = "end of CoverTab[86845]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:242
	_go_fuzz_dep_.CoverTab[86846]++
											s, err := dec.readUint32()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:244
		_go_fuzz_dep_.CoverTab[86851]++
												return fmt.Errorf("could not establish actual count of uni-dimensional varying array: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:245
		// _ = "end of CoverTab[86851]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:246
		_go_fuzz_dep_.CoverTab[86852]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:246
		// _ = "end of CoverTab[86852]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:246
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:246
	// _ = "end of CoverTab[86846]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:246
	_go_fuzz_dep_.CoverTab[86847]++
											t := v.Type()

											n := int(s + o)
											a := reflect.MakeSlice(t, n, n)

											for i := int(o); i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:252
		_go_fuzz_dep_.CoverTab[86853]++
												err := dec.fill(a.Index(i), tag, def)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:254
			_go_fuzz_dep_.CoverTab[86854]++
													return fmt.Errorf("could not fill index %d of uni-dimensional varying array: %v", i, err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:255
			// _ = "end of CoverTab[86854]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:256
			_go_fuzz_dep_.CoverTab[86855]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:256
			// _ = "end of CoverTab[86855]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:256
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:256
		// _ = "end of CoverTab[86853]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:257
	// _ = "end of CoverTab[86847]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:257
	_go_fuzz_dep_.CoverTab[86848]++
											v.Set(a)
											return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:259
	// _ = "end of CoverTab[86848]"
}

// fillMultiDimensionalVaryingArray fills the multi-dimensional slice value provided from varying array data.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:262
// The number of dimensions must be specified. This must be less than or equal to the dimensions in the slice for this
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:262
// method not to panic.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:265
func (dec *Decoder) fillMultiDimensionalVaryingArray(v reflect.Value, t reflect.Type, d int, tag reflect.StructTag, def *[]deferedPtr) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:265
	_go_fuzz_dep_.CoverTab[86856]++

											o := make([]int, d, d)
											l := make([]int, d, d)
											for i := range l {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:269
		_go_fuzz_dep_.CoverTab[86859]++
												off, err := dec.readUint32()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:271
			_go_fuzz_dep_.CoverTab[86862]++
													return fmt.Errorf("could not read offset of dimension %d: %v", i+1, err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:272
			// _ = "end of CoverTab[86862]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:273
			_go_fuzz_dep_.CoverTab[86863]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:273
			// _ = "end of CoverTab[86863]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:273
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:273
		// _ = "end of CoverTab[86859]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:273
		_go_fuzz_dep_.CoverTab[86860]++
												o[i] = int(off)
												s, err := dec.readUint32()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:276
			_go_fuzz_dep_.CoverTab[86864]++
													return fmt.Errorf("could not read size of dimension %d: %v", i+1, err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:277
			// _ = "end of CoverTab[86864]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:278
			_go_fuzz_dep_.CoverTab[86865]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:278
			// _ = "end of CoverTab[86865]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:278
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:278
		// _ = "end of CoverTab[86860]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:278
		_go_fuzz_dep_.CoverTab[86861]++
												l[i] = int(s) + int(off)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:279
		// _ = "end of CoverTab[86861]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:280
	// _ = "end of CoverTab[86856]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:280
	_go_fuzz_dep_.CoverTab[86857]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:283
	ty := v.Type()
											v.Set(reflect.MakeSlice(ty, l[0], l[0]))

											makeSubSlices(v, l[1:])

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:289
	ps := multiDimensionalIndexPermutations(l)
	for _, p := range ps {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:290
		_go_fuzz_dep_.CoverTab[86866]++

												a := v
												var os bool	// should this permutation be skipped due to the offset of any of the dimensions?
												for i, j := range p {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:294
			_go_fuzz_dep_.CoverTab[86869]++
													if j < o[i] {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:295
				_go_fuzz_dep_.CoverTab[86871]++
														os = true
														break
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:297
				// _ = "end of CoverTab[86871]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:298
				_go_fuzz_dep_.CoverTab[86872]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:298
				// _ = "end of CoverTab[86872]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:298
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:298
			// _ = "end of CoverTab[86869]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:298
			_go_fuzz_dep_.CoverTab[86870]++
													a = a.Index(j)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:299
			// _ = "end of CoverTab[86870]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:300
		// _ = "end of CoverTab[86866]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:300
		_go_fuzz_dep_.CoverTab[86867]++
												if os {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:301
			_go_fuzz_dep_.CoverTab[86873]++

													continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:303
			// _ = "end of CoverTab[86873]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:304
			_go_fuzz_dep_.CoverTab[86874]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:304
			// _ = "end of CoverTab[86874]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:304
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:304
		// _ = "end of CoverTab[86867]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:304
		_go_fuzz_dep_.CoverTab[86868]++
												err := dec.fill(a, tag, def)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:306
			_go_fuzz_dep_.CoverTab[86875]++
													return fmt.Errorf("could not fill index %v of slice: %v", p, err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:307
			// _ = "end of CoverTab[86875]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:308
			_go_fuzz_dep_.CoverTab[86876]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:308
			// _ = "end of CoverTab[86876]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:308
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:308
		// _ = "end of CoverTab[86868]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:309
	// _ = "end of CoverTab[86857]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:309
	_go_fuzz_dep_.CoverTab[86858]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:310
	// _ = "end of CoverTab[86858]"
}

// fillConformantVaryingArray establishes if the varying array is uni or multi dimensional and then fills the slice.
func (dec *Decoder) fillConformantVaryingArray(v reflect.Value, tag reflect.StructTag, def *[]deferedPtr) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:314
	_go_fuzz_dep_.CoverTab[86877]++
											d, t := sliceDimensions(v.Type())
											if d > 1 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:316
		_go_fuzz_dep_.CoverTab[86879]++
												err := dec.fillMultiDimensionalConformantVaryingArray(v, t, d, tag, def)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:318
			_go_fuzz_dep_.CoverTab[86880]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:319
			// _ = "end of CoverTab[86880]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:320
			_go_fuzz_dep_.CoverTab[86881]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:320
			// _ = "end of CoverTab[86881]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:320
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:320
		// _ = "end of CoverTab[86879]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:321
		_go_fuzz_dep_.CoverTab[86882]++
												err := dec.fillUniDimensionalConformantVaryingArray(v, tag, def)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:323
			_go_fuzz_dep_.CoverTab[86883]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:324
			// _ = "end of CoverTab[86883]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:325
			_go_fuzz_dep_.CoverTab[86884]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:325
			// _ = "end of CoverTab[86884]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:325
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:325
		// _ = "end of CoverTab[86882]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:326
	// _ = "end of CoverTab[86877]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:326
	_go_fuzz_dep_.CoverTab[86878]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:327
	// _ = "end of CoverTab[86878]"
}

// fillUniDimensionalConformantVaryingArray fills the uni-dimensional slice value.
func (dec *Decoder) fillUniDimensionalConformantVaryingArray(v reflect.Value, tag reflect.StructTag, def *[]deferedPtr) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:331
	_go_fuzz_dep_.CoverTab[86885]++
											m := dec.precedingMax()
											o, err := dec.readUint32()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:334
		_go_fuzz_dep_.CoverTab[86890]++
												return fmt.Errorf("could not read offset of uni-dimensional conformant varying array: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:335
		// _ = "end of CoverTab[86890]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:336
		_go_fuzz_dep_.CoverTab[86891]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:336
		// _ = "end of CoverTab[86891]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:336
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:336
	// _ = "end of CoverTab[86885]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:336
	_go_fuzz_dep_.CoverTab[86886]++
											s, err := dec.readUint32()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:338
		_go_fuzz_dep_.CoverTab[86892]++
												return fmt.Errorf("could not establish actual count of uni-dimensional conformant varying array: %v", err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:339
		// _ = "end of CoverTab[86892]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:340
		_go_fuzz_dep_.CoverTab[86893]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:340
		// _ = "end of CoverTab[86893]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:340
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:340
	// _ = "end of CoverTab[86886]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:340
	_go_fuzz_dep_.CoverTab[86887]++
											if m < o+s {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:341
		_go_fuzz_dep_.CoverTab[86894]++
												return errors.New("max count is less than the offset plus actual count")
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:342
		// _ = "end of CoverTab[86894]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:343
		_go_fuzz_dep_.CoverTab[86895]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:343
		// _ = "end of CoverTab[86895]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:343
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:343
	// _ = "end of CoverTab[86887]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:343
	_go_fuzz_dep_.CoverTab[86888]++
											t := v.Type()
											n := int(s)
											a := reflect.MakeSlice(t, n, n)
											for i := int(o); i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:347
		_go_fuzz_dep_.CoverTab[86896]++
												err := dec.fill(a.Index(i), tag, def)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:349
			_go_fuzz_dep_.CoverTab[86897]++
													return fmt.Errorf("could not fill index %d of uni-dimensional conformant varying array: %v", i, err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:350
			// _ = "end of CoverTab[86897]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:351
			_go_fuzz_dep_.CoverTab[86898]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:351
			// _ = "end of CoverTab[86898]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:351
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:351
		// _ = "end of CoverTab[86896]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:352
	// _ = "end of CoverTab[86888]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:352
	_go_fuzz_dep_.CoverTab[86889]++
											v.Set(a)
											return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:354
	// _ = "end of CoverTab[86889]"
}

// fillMultiDimensionalConformantVaryingArray fills the multi-dimensional slice value provided from conformant varying array data.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:357
// The number of dimensions must be specified. This must be less than or equal to the dimensions in the slice for this
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:357
// method not to panic.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:360
func (dec *Decoder) fillMultiDimensionalConformantVaryingArray(v reflect.Value, t reflect.Type, d int, tag reflect.StructTag, def *[]deferedPtr) error {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:360
	_go_fuzz_dep_.CoverTab[86899]++

											m := make([]int, d, d)
											for i := range m {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:363
		_go_fuzz_dep_.CoverTab[86903]++
												m[i] = int(dec.precedingMax())
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:364
		// _ = "end of CoverTab[86903]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:365
	// _ = "end of CoverTab[86899]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:365
	_go_fuzz_dep_.CoverTab[86900]++
											o := make([]int, d, d)
											l := make([]int, d, d)
											for i := range l {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:368
		_go_fuzz_dep_.CoverTab[86904]++
												off, err := dec.readUint32()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:370
			_go_fuzz_dep_.CoverTab[86908]++
													return fmt.Errorf("could not read offset of dimension %d: %v", i+1, err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:371
			// _ = "end of CoverTab[86908]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:372
			_go_fuzz_dep_.CoverTab[86909]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:372
			// _ = "end of CoverTab[86909]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:372
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:372
		// _ = "end of CoverTab[86904]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:372
		_go_fuzz_dep_.CoverTab[86905]++
												o[i] = int(off)
												s, err := dec.readUint32()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:375
			_go_fuzz_dep_.CoverTab[86910]++
													return fmt.Errorf("could not read actual count of dimension %d: %v", i+1, err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:376
			// _ = "end of CoverTab[86910]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:377
			_go_fuzz_dep_.CoverTab[86911]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:377
			// _ = "end of CoverTab[86911]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:377
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:377
		// _ = "end of CoverTab[86905]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:377
		_go_fuzz_dep_.CoverTab[86906]++
												if m[i] < int(s)+int(off) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:378
			_go_fuzz_dep_.CoverTab[86912]++
													m[i] = int(s) + int(off)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:379
			// _ = "end of CoverTab[86912]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:380
			_go_fuzz_dep_.CoverTab[86913]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:380
			// _ = "end of CoverTab[86913]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:380
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:380
		// _ = "end of CoverTab[86906]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:380
		_go_fuzz_dep_.CoverTab[86907]++
												l[i] = int(s)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:381
		// _ = "end of CoverTab[86907]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:382
	// _ = "end of CoverTab[86900]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:382
	_go_fuzz_dep_.CoverTab[86901]++

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:385
	ty := v.Type()
											v.Set(reflect.MakeSlice(ty, m[0], m[0]))

											makeSubSlices(v, m[1:])

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:391
	ps := multiDimensionalIndexPermutations(m)
	for _, p := range ps {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:392
		_go_fuzz_dep_.CoverTab[86914]++

												a := v
												var os bool	// should this permutation be skipped due to the offset of any of the dimensions or max is higher than the actual count being passed
												for i, j := range p {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:396
			_go_fuzz_dep_.CoverTab[86917]++
													if j < o[i] || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:397
				_go_fuzz_dep_.CoverTab[86919]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:397
				return j >= l[i]
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:397
				// _ = "end of CoverTab[86919]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:397
			}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:397
				_go_fuzz_dep_.CoverTab[86920]++
														os = true
														break
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:399
				// _ = "end of CoverTab[86920]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:400
				_go_fuzz_dep_.CoverTab[86921]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:400
				// _ = "end of CoverTab[86921]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:400
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:400
			// _ = "end of CoverTab[86917]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:400
			_go_fuzz_dep_.CoverTab[86918]++
													a = a.Index(j)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:401
			// _ = "end of CoverTab[86918]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:402
		// _ = "end of CoverTab[86914]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:402
		_go_fuzz_dep_.CoverTab[86915]++
												if os {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:403
			_go_fuzz_dep_.CoverTab[86922]++

													continue
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:405
			// _ = "end of CoverTab[86922]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:406
			_go_fuzz_dep_.CoverTab[86923]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:406
			// _ = "end of CoverTab[86923]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:406
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:406
		// _ = "end of CoverTab[86915]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:406
		_go_fuzz_dep_.CoverTab[86916]++
												err := dec.fill(a, tag, def)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:408
			_go_fuzz_dep_.CoverTab[86924]++
													return fmt.Errorf("could not fill index %v of slice: %v", p, err)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:409
			// _ = "end of CoverTab[86924]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:410
			_go_fuzz_dep_.CoverTab[86925]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:410
			// _ = "end of CoverTab[86925]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:410
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:410
		// _ = "end of CoverTab[86916]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:411
	// _ = "end of CoverTab[86901]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:411
	_go_fuzz_dep_.CoverTab[86902]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:412
	// _ = "end of CoverTab[86902]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:413
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/arrays.go:413
var _ = _go_fuzz_dep_.CoverTab
