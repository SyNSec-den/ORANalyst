//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_writepub.go:1
package toml

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_writepub.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_writepub.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_writepub.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_writepub.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_writepub.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_writepub.go:1
)

// ValueStringRepresentation transforms an interface{} value into its toml string representation.
func ValueStringRepresentation(v interface{}, commented string, indent string, ord MarshalOrder, arraysOneElementPerLine bool) (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_writepub.go:4
	_go_fuzz_dep_.CoverTab[124421]++
												return tomlValueStringRepresentation(v, commented, indent, ord, arraysOneElementPerLine)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_writepub.go:5
	// _ = "end of CoverTab[124421]"
}

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_writepub.go:6
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_writepub.go:6
var _ = _go_fuzz_dep_.CoverTab
