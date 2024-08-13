//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:19
package channelz

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:19
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:19
)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:19
import (
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:19
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:19
)

import "fmt"

// Identifier is an opaque identifier which uniquely identifies an entity in the
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:23
// channelz database.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:25
type Identifier struct {
	typ	RefChannelType
	id	int64
	str	string
	pid	*Identifier
}

// Type returns the entity type corresponding to id.
func (id *Identifier) Type() RefChannelType {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:33
	_go_fuzz_dep_.CoverTab[62813]++
												return id.typ
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:34
	// _ = "end of CoverTab[62813]"
}

// Int returns the integer identifier corresponding to id.
func (id *Identifier) Int() int64 {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:38
	_go_fuzz_dep_.CoverTab[62814]++
												return id.id
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:39
	// _ = "end of CoverTab[62814]"
}

// String returns a string representation of the entity corresponding to id.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:42
//
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:42
// This includes some information about the parent as well. Examples:
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:42
// Top-level channel: [Channel #channel-number]
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:42
// Nested channel:    [Channel #parent-channel-number Channel #channel-number]
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:42
// Sub channel:       [Channel #parent-channel SubChannel #subchannel-number]
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:48
func (id *Identifier) String() string {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:48
	_go_fuzz_dep_.CoverTab[62815]++
												return id.str
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:49
	// _ = "end of CoverTab[62815]"
}

// Equal returns true if other is the same as id.
func (id *Identifier) Equal(other *Identifier) bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:53
	_go_fuzz_dep_.CoverTab[62816]++
												if (id != nil) != (other != nil) {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:54
		_go_fuzz_dep_.CoverTab[62819]++
													return false
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:55
		// _ = "end of CoverTab[62819]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:56
		_go_fuzz_dep_.CoverTab[62820]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:56
		// _ = "end of CoverTab[62820]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:56
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:56
	// _ = "end of CoverTab[62816]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:56
	_go_fuzz_dep_.CoverTab[62817]++
												if id == nil && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:57
		_go_fuzz_dep_.CoverTab[62821]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:57
		return other == nil
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:57
		// _ = "end of CoverTab[62821]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:57
	}() {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:57
		_go_fuzz_dep_.CoverTab[62822]++
													return true
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:58
		// _ = "end of CoverTab[62822]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:59
		_go_fuzz_dep_.CoverTab[62823]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:59
		// _ = "end of CoverTab[62823]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:59
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:59
	// _ = "end of CoverTab[62817]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:59
	_go_fuzz_dep_.CoverTab[62818]++
												return id.typ == other.typ && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:60
		_go_fuzz_dep_.CoverTab[62824]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:60
		return id.id == other.id
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:60
		// _ = "end of CoverTab[62824]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:60
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:60
		_go_fuzz_dep_.CoverTab[62825]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:60
		return id.pid == other.pid
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:60
		// _ = "end of CoverTab[62825]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:60
	}()
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:60
	// _ = "end of CoverTab[62818]"
}

// NewIdentifierForTesting returns a new opaque identifier to be used only for
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:63
// testing purposes.
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:65
func NewIdentifierForTesting(typ RefChannelType, id int64, pid *Identifier) *Identifier {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:65
	_go_fuzz_dep_.CoverTab[62826]++
												return newIdentifer(typ, id, pid)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:66
	// _ = "end of CoverTab[62826]"
}

func newIdentifer(typ RefChannelType, id int64, pid *Identifier) *Identifier {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:69
	_go_fuzz_dep_.CoverTab[62827]++
												str := fmt.Sprintf("%s #%d", typ, id)
												if pid != nil {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:71
		_go_fuzz_dep_.CoverTab[62829]++
													str = fmt.Sprintf("%s %s", pid, str)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:72
		// _ = "end of CoverTab[62829]"
	} else {
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:73
		_go_fuzz_dep_.CoverTab[62830]++
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:73
		// _ = "end of CoverTab[62830]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:73
	}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:73
	// _ = "end of CoverTab[62827]"
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:73
	_go_fuzz_dep_.CoverTab[62828]++
												return &Identifier{typ: typ, id: id, str: str, pid: pid}
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:74
	// _ = "end of CoverTab[62828]"
}

//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:75
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/google.golang.org/grpc@v1.54.0/internal/channelz/id.go:75
var _ = _go_fuzz_dep_.CoverTab
