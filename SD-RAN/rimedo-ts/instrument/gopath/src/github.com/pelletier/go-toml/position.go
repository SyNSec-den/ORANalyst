// Position support for go-toml

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/position.go:3
package toml

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/position.go:3
import (
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/position.go:3
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/position.go:3
)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/position.go:3
import (
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/position.go:3
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/position.go:3
)

import (
	"fmt"
)

// Position of a document element within a TOML document.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/position.go:9
//
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/position.go:9
// Line and Col are both 1-indexed positions for the element's line number and
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/position.go:9
// column number, respectively.  Values of zero or less will cause Invalid(),
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/position.go:9
// to return true.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/position.go:14
type Position struct {
	Line	int	// line within the document
	Col	int	// column within the line
}

// String representation of the position.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/position.go:19
// Displays 1-indexed line and column numbers.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/position.go:21
func (p Position) String() string {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/position.go:21
	_go_fuzz_dep_.CoverTab[123859]++
											return fmt.Sprintf("(%d, %d)", p.Line, p.Col)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/position.go:22
	// _ = "end of CoverTab[123859]"
}

// Invalid returns whether or not the position is valid (i.e. with negative or
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/position.go:25
// null values)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/position.go:27
func (p Position) Invalid() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/position.go:27
	_go_fuzz_dep_.CoverTab[123860]++
											return p.Line <= 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/position.go:28
		_go_fuzz_dep_.CoverTab[123861]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/position.go:28
		return p.Col <= 0
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/position.go:28
		// _ = "end of CoverTab[123861]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/position.go:28
	}()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/position.go:28
	// _ = "end of CoverTab[123860]"
}

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/position.go:29
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/position.go:29
var _ = _go_fuzz_dep_.CoverTab
