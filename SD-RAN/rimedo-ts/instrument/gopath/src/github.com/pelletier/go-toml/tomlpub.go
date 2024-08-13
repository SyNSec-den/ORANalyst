//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:1
package toml

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:1
)

// PubTOMLValue wrapping tomlValue in order to access all properties from outside.
type PubTOMLValue = tomlValue

func (ptv *PubTOMLValue) Value() interface{} {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:6
	_go_fuzz_dep_.CoverTab[124107]++
											return ptv.value
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:7
	// _ = "end of CoverTab[124107]"
}
func (ptv *PubTOMLValue) Comment() string {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:9
	_go_fuzz_dep_.CoverTab[124108]++
											return ptv.comment
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:10
	// _ = "end of CoverTab[124108]"
}
func (ptv *PubTOMLValue) Commented() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:12
	_go_fuzz_dep_.CoverTab[124109]++
											return ptv.commented
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:13
	// _ = "end of CoverTab[124109]"
}
func (ptv *PubTOMLValue) Multiline() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:15
	_go_fuzz_dep_.CoverTab[124110]++
											return ptv.multiline
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:16
	// _ = "end of CoverTab[124110]"
}
func (ptv *PubTOMLValue) Position() Position {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:18
	_go_fuzz_dep_.CoverTab[124111]++
											return ptv.position
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:19
	// _ = "end of CoverTab[124111]"
}

func (ptv *PubTOMLValue) SetValue(v interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:22
	_go_fuzz_dep_.CoverTab[124112]++
											ptv.value = v
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:23
	// _ = "end of CoverTab[124112]"
}
func (ptv *PubTOMLValue) SetComment(s string) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:25
	_go_fuzz_dep_.CoverTab[124113]++
											ptv.comment = s
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:26
	// _ = "end of CoverTab[124113]"
}
func (ptv *PubTOMLValue) SetCommented(c bool) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:28
	_go_fuzz_dep_.CoverTab[124114]++
											ptv.commented = c
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:29
	// _ = "end of CoverTab[124114]"
}
func (ptv *PubTOMLValue) SetMultiline(m bool) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:31
	_go_fuzz_dep_.CoverTab[124115]++
											ptv.multiline = m
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:32
	// _ = "end of CoverTab[124115]"
}
func (ptv *PubTOMLValue) SetPosition(p Position) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:34
	_go_fuzz_dep_.CoverTab[124116]++
											ptv.position = p
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:35
	// _ = "end of CoverTab[124116]"
}

// PubTree wrapping Tree in order to access all properties from outside.
type PubTree = Tree

func (pt *PubTree) Values() map[string]interface{} {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:41
	_go_fuzz_dep_.CoverTab[124117]++
											return pt.values
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:42
	// _ = "end of CoverTab[124117]"
}

func (pt *PubTree) Comment() string {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:45
	_go_fuzz_dep_.CoverTab[124118]++
											return pt.comment
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:46
	// _ = "end of CoverTab[124118]"
}

func (pt *PubTree) Commented() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:49
	_go_fuzz_dep_.CoverTab[124119]++
											return pt.commented
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:50
	// _ = "end of CoverTab[124119]"
}

func (pt *PubTree) Inline() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:53
	_go_fuzz_dep_.CoverTab[124120]++
											return pt.inline
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:54
	// _ = "end of CoverTab[124120]"
}

func (pt *PubTree) SetValues(v map[string]interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:57
	_go_fuzz_dep_.CoverTab[124121]++
											pt.values = v
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:58
	// _ = "end of CoverTab[124121]"
}

func (pt *PubTree) SetComment(c string) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:61
	_go_fuzz_dep_.CoverTab[124122]++
											pt.comment = c
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:62
	// _ = "end of CoverTab[124122]"
}

func (pt *PubTree) SetCommented(c bool) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:65
	_go_fuzz_dep_.CoverTab[124123]++
											pt.commented = c
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:66
	// _ = "end of CoverTab[124123]"
}

func (pt *PubTree) SetInline(i bool) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:69
	_go_fuzz_dep_.CoverTab[124124]++
											pt.inline = i
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:70
	// _ = "end of CoverTab[124124]"
}

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:71
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomlpub.go:71
var _ = _go_fuzz_dep_.CoverTab
