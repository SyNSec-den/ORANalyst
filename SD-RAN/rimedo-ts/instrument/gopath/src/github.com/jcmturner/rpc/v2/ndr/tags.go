//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:1
package ndr

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:1
)

import (
	"fmt"
	"reflect"
	"strings"
)

const ndrNameSpace = "ndr"

type tags struct {
	Values	[]string
	Map	map[string]string
}

// parse the struct field tags and extract the ndr related ones.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:16
// format of tag ndr:"value,key:value1,value2"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:18
func parseTags(st reflect.StructTag) tags {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:18
	_go_fuzz_dep_.CoverTab[87295]++
											s := st.Get(ndrNameSpace)
											t := tags{
		Values:	[]string{},
		Map:	make(map[string]string),
	}
	if s != "" {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:24
		_go_fuzz_dep_.CoverTab[87297]++
												ndrTags := strings.Trim(s, `"`)
												for _, tag := range strings.Split(ndrTags, ",") {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:26
			_go_fuzz_dep_.CoverTab[87298]++
													if strings.Contains(tag, ":") {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:27
				_go_fuzz_dep_.CoverTab[87299]++
														m := strings.SplitN(tag, ":", 2)
														t.Map[m[0]] = m[1]
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:29
				// _ = "end of CoverTab[87299]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:30
				_go_fuzz_dep_.CoverTab[87300]++
														t.Values = append(t.Values, tag)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:31
				// _ = "end of CoverTab[87300]"
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:32
			// _ = "end of CoverTab[87298]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:33
		// _ = "end of CoverTab[87297]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:34
		_go_fuzz_dep_.CoverTab[87301]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:34
		// _ = "end of CoverTab[87301]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:34
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:34
	// _ = "end of CoverTab[87295]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:34
	_go_fuzz_dep_.CoverTab[87296]++
											return t
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:35
	// _ = "end of CoverTab[87296]"
}

func appendTag(t reflect.StructTag, s string) reflect.StructTag {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:38
	_go_fuzz_dep_.CoverTab[87302]++
											ts := t.Get(ndrNameSpace)
											ts = fmt.Sprintf(`%s"%s,%s"`, ndrNameSpace, ts, s)
											return reflect.StructTag(ts)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:41
	// _ = "end of CoverTab[87302]"
}

func (t *tags) StructTag() reflect.StructTag {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:44
	_go_fuzz_dep_.CoverTab[87303]++
											mv := t.Values
											for key, val := range t.Map {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:46
		_go_fuzz_dep_.CoverTab[87305]++
												mv = append(mv, key+":"+val)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:47
		// _ = "end of CoverTab[87305]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:48
	// _ = "end of CoverTab[87303]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:48
	_go_fuzz_dep_.CoverTab[87304]++
											s := ndrNameSpace + ":" + `"` + strings.Join(mv, ",") + `"`
											return reflect.StructTag(s)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:50
	// _ = "end of CoverTab[87304]"
}

func (t *tags) delete(s string) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:53
	_go_fuzz_dep_.CoverTab[87306]++
											for i, x := range t.Values {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:54
		_go_fuzz_dep_.CoverTab[87308]++
												if x == s {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:55
			_go_fuzz_dep_.CoverTab[87309]++
													t.Values = append(t.Values[:i], t.Values[i+1:]...)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:56
			// _ = "end of CoverTab[87309]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:57
			_go_fuzz_dep_.CoverTab[87310]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:57
			// _ = "end of CoverTab[87310]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:57
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:57
		// _ = "end of CoverTab[87308]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:58
	// _ = "end of CoverTab[87306]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:58
	_go_fuzz_dep_.CoverTab[87307]++
											delete(t.Map, s)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:59
	// _ = "end of CoverTab[87307]"
}

func (t *tags) HasValue(s string) bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:62
	_go_fuzz_dep_.CoverTab[87311]++
											for _, v := range t.Values {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:63
		_go_fuzz_dep_.CoverTab[87313]++
												if v == s {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:64
			_go_fuzz_dep_.CoverTab[87314]++
													return true
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:65
			// _ = "end of CoverTab[87314]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:66
			_go_fuzz_dep_.CoverTab[87315]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:66
			// _ = "end of CoverTab[87315]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:66
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:66
		// _ = "end of CoverTab[87313]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:67
	// _ = "end of CoverTab[87311]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:67
	_go_fuzz_dep_.CoverTab[87312]++
											return false
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:68
	// _ = "end of CoverTab[87312]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:69
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/rpc/v2@v2.0.3/ndr/tags.go:69
var _ = _go_fuzz_dep_.CoverTab
