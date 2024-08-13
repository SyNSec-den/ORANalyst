// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/encoding/gob/type.go:5
package gob

//line /usr/local/go/src/encoding/gob/type.go:5
import (
//line /usr/local/go/src/encoding/gob/type.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/encoding/gob/type.go:5
)
//line /usr/local/go/src/encoding/gob/type.go:5
import (
//line /usr/local/go/src/encoding/gob/type.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/encoding/gob/type.go:5
)

import (
	"encoding"
	"errors"
	"fmt"
	"os"
	"reflect"
	"sync"
	"sync/atomic"
	"unicode"
	"unicode/utf8"
)

// userTypeInfo stores the information associated with a type the user has handed
//line /usr/local/go/src/encoding/gob/type.go:19
// to the package. It's computed once and stored in a map keyed by reflection
//line /usr/local/go/src/encoding/gob/type.go:19
// type.
//line /usr/local/go/src/encoding/gob/type.go:22
type userTypeInfo struct {
	user		reflect.Type	// the type the user handed us
	base		reflect.Type	// the base type after all indirections
	indir		int		// number of indirections to reach the base type
	externalEnc	int		// xGob, xBinary, or xText
	externalDec	int		// xGob, xBinary or xText
	encIndir	int8		// number of indirections to reach the receiver type; may be negative
	decIndir	int8		// number of indirections to reach the receiver type; may be negative
}

// externalEncoding bits
const (
	xGob	= 1 + iota	// GobEncoder or GobDecoder
	xBinary			// encoding.BinaryMarshaler or encoding.BinaryUnmarshaler
	xText			// encoding.TextMarshaler or encoding.TextUnmarshaler
)

var userTypeCache sync.Map	// map[reflect.Type]*userTypeInfo

// validUserType returns, and saves, the information associated with user-provided type rt.
//line /usr/local/go/src/encoding/gob/type.go:41
// If the user type is not valid, err will be non-nil. To be used when the error handler
//line /usr/local/go/src/encoding/gob/type.go:41
// is not set up.
//line /usr/local/go/src/encoding/gob/type.go:44
func validUserType(rt reflect.Type) (*userTypeInfo, error) {
//line /usr/local/go/src/encoding/gob/type.go:44
	_go_fuzz_dep_.CoverTab[85186]++
							if ui, ok := userTypeCache.Load(rt); ok {
//line /usr/local/go/src/encoding/gob/type.go:45
		_go_fuzz_dep_.CoverTab[85191]++
								return ui.(*userTypeInfo), nil
//line /usr/local/go/src/encoding/gob/type.go:46
		// _ = "end of CoverTab[85191]"
	} else {
//line /usr/local/go/src/encoding/gob/type.go:47
		_go_fuzz_dep_.CoverTab[85192]++
//line /usr/local/go/src/encoding/gob/type.go:47
		// _ = "end of CoverTab[85192]"
//line /usr/local/go/src/encoding/gob/type.go:47
	}
//line /usr/local/go/src/encoding/gob/type.go:47
	// _ = "end of CoverTab[85186]"
//line /usr/local/go/src/encoding/gob/type.go:47
	_go_fuzz_dep_.CoverTab[85187]++

//line /usr/local/go/src/encoding/gob/type.go:53
	ut := new(userTypeInfo)
							ut.base = rt
							ut.user = rt

//line /usr/local/go/src/encoding/gob/type.go:61
	slowpoke := ut.base
	for {
//line /usr/local/go/src/encoding/gob/type.go:62
		_go_fuzz_dep_.CoverTab[85193]++
								pt := ut.base
								if pt.Kind() != reflect.Pointer {
//line /usr/local/go/src/encoding/gob/type.go:64
			_go_fuzz_dep_.CoverTab[85197]++
									break
//line /usr/local/go/src/encoding/gob/type.go:65
			// _ = "end of CoverTab[85197]"
		} else {
//line /usr/local/go/src/encoding/gob/type.go:66
			_go_fuzz_dep_.CoverTab[85198]++
//line /usr/local/go/src/encoding/gob/type.go:66
			// _ = "end of CoverTab[85198]"
//line /usr/local/go/src/encoding/gob/type.go:66
		}
//line /usr/local/go/src/encoding/gob/type.go:66
		// _ = "end of CoverTab[85193]"
//line /usr/local/go/src/encoding/gob/type.go:66
		_go_fuzz_dep_.CoverTab[85194]++
								ut.base = pt.Elem()
								if ut.base == slowpoke {
//line /usr/local/go/src/encoding/gob/type.go:68
			_go_fuzz_dep_.CoverTab[85199]++

									return nil, errors.New("can't represent recursive pointer type " + ut.base.String())
//line /usr/local/go/src/encoding/gob/type.go:70
			// _ = "end of CoverTab[85199]"
		} else {
//line /usr/local/go/src/encoding/gob/type.go:71
			_go_fuzz_dep_.CoverTab[85200]++
//line /usr/local/go/src/encoding/gob/type.go:71
			// _ = "end of CoverTab[85200]"
//line /usr/local/go/src/encoding/gob/type.go:71
		}
//line /usr/local/go/src/encoding/gob/type.go:71
		// _ = "end of CoverTab[85194]"
//line /usr/local/go/src/encoding/gob/type.go:71
		_go_fuzz_dep_.CoverTab[85195]++
								if ut.indir%2 == 0 {
//line /usr/local/go/src/encoding/gob/type.go:72
			_go_fuzz_dep_.CoverTab[85201]++
									slowpoke = slowpoke.Elem()
//line /usr/local/go/src/encoding/gob/type.go:73
			// _ = "end of CoverTab[85201]"
		} else {
//line /usr/local/go/src/encoding/gob/type.go:74
			_go_fuzz_dep_.CoverTab[85202]++
//line /usr/local/go/src/encoding/gob/type.go:74
			// _ = "end of CoverTab[85202]"
//line /usr/local/go/src/encoding/gob/type.go:74
		}
//line /usr/local/go/src/encoding/gob/type.go:74
		// _ = "end of CoverTab[85195]"
//line /usr/local/go/src/encoding/gob/type.go:74
		_go_fuzz_dep_.CoverTab[85196]++
								ut.indir++
//line /usr/local/go/src/encoding/gob/type.go:75
		// _ = "end of CoverTab[85196]"
	}
//line /usr/local/go/src/encoding/gob/type.go:76
	// _ = "end of CoverTab[85187]"
//line /usr/local/go/src/encoding/gob/type.go:76
	_go_fuzz_dep_.CoverTab[85188]++

							if ok, indir := implementsInterface(ut.user, gobEncoderInterfaceType); ok {
//line /usr/local/go/src/encoding/gob/type.go:78
		_go_fuzz_dep_.CoverTab[85203]++
								ut.externalEnc, ut.encIndir = xGob, indir
//line /usr/local/go/src/encoding/gob/type.go:79
		// _ = "end of CoverTab[85203]"
	} else {
//line /usr/local/go/src/encoding/gob/type.go:80
		_go_fuzz_dep_.CoverTab[85204]++
//line /usr/local/go/src/encoding/gob/type.go:80
		if ok, indir := implementsInterface(ut.user, binaryMarshalerInterfaceType); ok {
//line /usr/local/go/src/encoding/gob/type.go:80
			_go_fuzz_dep_.CoverTab[85205]++
									ut.externalEnc, ut.encIndir = xBinary, indir
//line /usr/local/go/src/encoding/gob/type.go:81
			// _ = "end of CoverTab[85205]"
		} else {
//line /usr/local/go/src/encoding/gob/type.go:82
			_go_fuzz_dep_.CoverTab[85206]++
//line /usr/local/go/src/encoding/gob/type.go:82
			// _ = "end of CoverTab[85206]"
//line /usr/local/go/src/encoding/gob/type.go:82
		}
//line /usr/local/go/src/encoding/gob/type.go:82
		// _ = "end of CoverTab[85204]"
//line /usr/local/go/src/encoding/gob/type.go:82
	}
//line /usr/local/go/src/encoding/gob/type.go:82
	// _ = "end of CoverTab[85188]"
//line /usr/local/go/src/encoding/gob/type.go:82
	_go_fuzz_dep_.CoverTab[85189]++

//line /usr/local/go/src/encoding/gob/type.go:90
	if ok, indir := implementsInterface(ut.user, gobDecoderInterfaceType); ok {
//line /usr/local/go/src/encoding/gob/type.go:90
		_go_fuzz_dep_.CoverTab[85207]++
								ut.externalDec, ut.decIndir = xGob, indir
//line /usr/local/go/src/encoding/gob/type.go:91
		// _ = "end of CoverTab[85207]"
	} else {
//line /usr/local/go/src/encoding/gob/type.go:92
		_go_fuzz_dep_.CoverTab[85208]++
//line /usr/local/go/src/encoding/gob/type.go:92
		if ok, indir := implementsInterface(ut.user, binaryUnmarshalerInterfaceType); ok {
//line /usr/local/go/src/encoding/gob/type.go:92
			_go_fuzz_dep_.CoverTab[85209]++
									ut.externalDec, ut.decIndir = xBinary, indir
//line /usr/local/go/src/encoding/gob/type.go:93
			// _ = "end of CoverTab[85209]"
		} else {
//line /usr/local/go/src/encoding/gob/type.go:94
			_go_fuzz_dep_.CoverTab[85210]++
//line /usr/local/go/src/encoding/gob/type.go:94
			// _ = "end of CoverTab[85210]"
//line /usr/local/go/src/encoding/gob/type.go:94
		}
//line /usr/local/go/src/encoding/gob/type.go:94
		// _ = "end of CoverTab[85208]"
//line /usr/local/go/src/encoding/gob/type.go:94
	}
//line /usr/local/go/src/encoding/gob/type.go:94
	// _ = "end of CoverTab[85189]"
//line /usr/local/go/src/encoding/gob/type.go:94
	_go_fuzz_dep_.CoverTab[85190]++

//line /usr/local/go/src/encoding/gob/type.go:101
	ui, _ := userTypeCache.LoadOrStore(rt, ut)
							return ui.(*userTypeInfo), nil
//line /usr/local/go/src/encoding/gob/type.go:102
	// _ = "end of CoverTab[85190]"
}

var (
	gobEncoderInterfaceType		= reflect.TypeOf((*GobEncoder)(nil)).Elem()
	gobDecoderInterfaceType		= reflect.TypeOf((*GobDecoder)(nil)).Elem()
	binaryMarshalerInterfaceType	= reflect.TypeOf((*encoding.BinaryMarshaler)(nil)).Elem()
	binaryUnmarshalerInterfaceType	= reflect.TypeOf((*encoding.BinaryUnmarshaler)(nil)).Elem()
	textMarshalerInterfaceType	= reflect.TypeOf((*encoding.TextMarshaler)(nil)).Elem()
	textUnmarshalerInterfaceType	= reflect.TypeOf((*encoding.TextUnmarshaler)(nil)).Elem()
)

// implementsInterface reports whether the type implements the
//line /usr/local/go/src/encoding/gob/type.go:114
// gobEncoder/gobDecoder interface.
//line /usr/local/go/src/encoding/gob/type.go:114
// It also returns the number of indirections required to get to the
//line /usr/local/go/src/encoding/gob/type.go:114
// implementation.
//line /usr/local/go/src/encoding/gob/type.go:118
func implementsInterface(typ, gobEncDecType reflect.Type) (success bool, indir int8) {
//line /usr/local/go/src/encoding/gob/type.go:118
	_go_fuzz_dep_.CoverTab[85211]++
							if typ == nil {
//line /usr/local/go/src/encoding/gob/type.go:119
		_go_fuzz_dep_.CoverTab[85215]++
								return
//line /usr/local/go/src/encoding/gob/type.go:120
		// _ = "end of CoverTab[85215]"
	} else {
//line /usr/local/go/src/encoding/gob/type.go:121
		_go_fuzz_dep_.CoverTab[85216]++
//line /usr/local/go/src/encoding/gob/type.go:121
		// _ = "end of CoverTab[85216]"
//line /usr/local/go/src/encoding/gob/type.go:121
	}
//line /usr/local/go/src/encoding/gob/type.go:121
	// _ = "end of CoverTab[85211]"
//line /usr/local/go/src/encoding/gob/type.go:121
	_go_fuzz_dep_.CoverTab[85212]++
							rt := typ

//line /usr/local/go/src/encoding/gob/type.go:125
	for {
//line /usr/local/go/src/encoding/gob/type.go:125
		_go_fuzz_dep_.CoverTab[85217]++
								if rt.Implements(gobEncDecType) {
//line /usr/local/go/src/encoding/gob/type.go:126
			_go_fuzz_dep_.CoverTab[85220]++
									return true, indir
//line /usr/local/go/src/encoding/gob/type.go:127
			// _ = "end of CoverTab[85220]"
		} else {
//line /usr/local/go/src/encoding/gob/type.go:128
			_go_fuzz_dep_.CoverTab[85221]++
//line /usr/local/go/src/encoding/gob/type.go:128
			// _ = "end of CoverTab[85221]"
//line /usr/local/go/src/encoding/gob/type.go:128
		}
//line /usr/local/go/src/encoding/gob/type.go:128
		// _ = "end of CoverTab[85217]"
//line /usr/local/go/src/encoding/gob/type.go:128
		_go_fuzz_dep_.CoverTab[85218]++
								if p := rt; p.Kind() == reflect.Pointer {
//line /usr/local/go/src/encoding/gob/type.go:129
			_go_fuzz_dep_.CoverTab[85222]++
									indir++
									if indir > 100 {
//line /usr/local/go/src/encoding/gob/type.go:131
				_go_fuzz_dep_.CoverTab[85224]++
										return false, 0
//line /usr/local/go/src/encoding/gob/type.go:132
				// _ = "end of CoverTab[85224]"
			} else {
//line /usr/local/go/src/encoding/gob/type.go:133
				_go_fuzz_dep_.CoverTab[85225]++
//line /usr/local/go/src/encoding/gob/type.go:133
				// _ = "end of CoverTab[85225]"
//line /usr/local/go/src/encoding/gob/type.go:133
			}
//line /usr/local/go/src/encoding/gob/type.go:133
			// _ = "end of CoverTab[85222]"
//line /usr/local/go/src/encoding/gob/type.go:133
			_go_fuzz_dep_.CoverTab[85223]++
									rt = p.Elem()
									continue
//line /usr/local/go/src/encoding/gob/type.go:135
			// _ = "end of CoverTab[85223]"
		} else {
//line /usr/local/go/src/encoding/gob/type.go:136
			_go_fuzz_dep_.CoverTab[85226]++
//line /usr/local/go/src/encoding/gob/type.go:136
			// _ = "end of CoverTab[85226]"
//line /usr/local/go/src/encoding/gob/type.go:136
		}
//line /usr/local/go/src/encoding/gob/type.go:136
		// _ = "end of CoverTab[85218]"
//line /usr/local/go/src/encoding/gob/type.go:136
		_go_fuzz_dep_.CoverTab[85219]++
								break
//line /usr/local/go/src/encoding/gob/type.go:137
		// _ = "end of CoverTab[85219]"
	}
//line /usr/local/go/src/encoding/gob/type.go:138
	// _ = "end of CoverTab[85212]"
//line /usr/local/go/src/encoding/gob/type.go:138
	_go_fuzz_dep_.CoverTab[85213]++

							if typ.Kind() != reflect.Pointer {
//line /usr/local/go/src/encoding/gob/type.go:140
		_go_fuzz_dep_.CoverTab[85227]++

								if reflect.PointerTo(typ).Implements(gobEncDecType) {
//line /usr/local/go/src/encoding/gob/type.go:142
			_go_fuzz_dep_.CoverTab[85228]++
									return true, -1
//line /usr/local/go/src/encoding/gob/type.go:143
			// _ = "end of CoverTab[85228]"
		} else {
//line /usr/local/go/src/encoding/gob/type.go:144
			_go_fuzz_dep_.CoverTab[85229]++
//line /usr/local/go/src/encoding/gob/type.go:144
			// _ = "end of CoverTab[85229]"
//line /usr/local/go/src/encoding/gob/type.go:144
		}
//line /usr/local/go/src/encoding/gob/type.go:144
		// _ = "end of CoverTab[85227]"
	} else {
//line /usr/local/go/src/encoding/gob/type.go:145
		_go_fuzz_dep_.CoverTab[85230]++
//line /usr/local/go/src/encoding/gob/type.go:145
		// _ = "end of CoverTab[85230]"
//line /usr/local/go/src/encoding/gob/type.go:145
	}
//line /usr/local/go/src/encoding/gob/type.go:145
	// _ = "end of CoverTab[85213]"
//line /usr/local/go/src/encoding/gob/type.go:145
	_go_fuzz_dep_.CoverTab[85214]++
							return false, 0
//line /usr/local/go/src/encoding/gob/type.go:146
	// _ = "end of CoverTab[85214]"
}

// userType returns, and saves, the information associated with user-provided type rt.
//line /usr/local/go/src/encoding/gob/type.go:149
// If the user type is not valid, it calls error.
//line /usr/local/go/src/encoding/gob/type.go:151
func userType(rt reflect.Type) *userTypeInfo {
//line /usr/local/go/src/encoding/gob/type.go:151
	_go_fuzz_dep_.CoverTab[85231]++
							ut, err := validUserType(rt)
							if err != nil {
//line /usr/local/go/src/encoding/gob/type.go:153
		_go_fuzz_dep_.CoverTab[85233]++
								error_(err)
//line /usr/local/go/src/encoding/gob/type.go:154
		// _ = "end of CoverTab[85233]"
	} else {
//line /usr/local/go/src/encoding/gob/type.go:155
		_go_fuzz_dep_.CoverTab[85234]++
//line /usr/local/go/src/encoding/gob/type.go:155
		// _ = "end of CoverTab[85234]"
//line /usr/local/go/src/encoding/gob/type.go:155
	}
//line /usr/local/go/src/encoding/gob/type.go:155
	// _ = "end of CoverTab[85231]"
//line /usr/local/go/src/encoding/gob/type.go:155
	_go_fuzz_dep_.CoverTab[85232]++
							return ut
//line /usr/local/go/src/encoding/gob/type.go:156
	// _ = "end of CoverTab[85232]"
}

// A typeId represents a gob Type as an integer that can be passed on the wire.
//line /usr/local/go/src/encoding/gob/type.go:159
// Internally, typeIds are used as keys to a map to recover the underlying type info.
//line /usr/local/go/src/encoding/gob/type.go:161
type typeId int32

var nextId typeId	// incremented for each new type we build
var typeLock sync.Mutex	// set while building a type
const firstUserId = 64	// lowest id number granted to user

type gobType interface {
	id() typeId
	setId(id typeId)
	name() string
	string() string	// not public; only for debugging
	safeString(seen map[typeId]bool) string
}

var types = make(map[reflect.Type]gobType)
var idToType = make(map[typeId]gobType)
var builtinIdToType map[typeId]gobType	// set in init() after builtins are established

func setTypeId(typ gobType) {
//line /usr/local/go/src/encoding/gob/type.go:179
	_go_fuzz_dep_.CoverTab[85235]++

							if typ.id() != 0 {
//line /usr/local/go/src/encoding/gob/type.go:181
		_go_fuzz_dep_.CoverTab[85237]++
								return
//line /usr/local/go/src/encoding/gob/type.go:182
		// _ = "end of CoverTab[85237]"
	} else {
//line /usr/local/go/src/encoding/gob/type.go:183
		_go_fuzz_dep_.CoverTab[85238]++
//line /usr/local/go/src/encoding/gob/type.go:183
		// _ = "end of CoverTab[85238]"
//line /usr/local/go/src/encoding/gob/type.go:183
	}
//line /usr/local/go/src/encoding/gob/type.go:183
	// _ = "end of CoverTab[85235]"
//line /usr/local/go/src/encoding/gob/type.go:183
	_go_fuzz_dep_.CoverTab[85236]++
							nextId++
							typ.setId(nextId)
							idToType[nextId] = typ
//line /usr/local/go/src/encoding/gob/type.go:186
	// _ = "end of CoverTab[85236]"
}

func (t typeId) gobType() gobType {
//line /usr/local/go/src/encoding/gob/type.go:189
	_go_fuzz_dep_.CoverTab[85239]++
							if t == 0 {
//line /usr/local/go/src/encoding/gob/type.go:190
		_go_fuzz_dep_.CoverTab[85241]++
								return nil
//line /usr/local/go/src/encoding/gob/type.go:191
		// _ = "end of CoverTab[85241]"
	} else {
//line /usr/local/go/src/encoding/gob/type.go:192
		_go_fuzz_dep_.CoverTab[85242]++
//line /usr/local/go/src/encoding/gob/type.go:192
		// _ = "end of CoverTab[85242]"
//line /usr/local/go/src/encoding/gob/type.go:192
	}
//line /usr/local/go/src/encoding/gob/type.go:192
	// _ = "end of CoverTab[85239]"
//line /usr/local/go/src/encoding/gob/type.go:192
	_go_fuzz_dep_.CoverTab[85240]++
							return idToType[t]
//line /usr/local/go/src/encoding/gob/type.go:193
	// _ = "end of CoverTab[85240]"
}

// string returns the string representation of the type associated with the typeId.
func (t typeId) string() string {
//line /usr/local/go/src/encoding/gob/type.go:197
	_go_fuzz_dep_.CoverTab[85243]++
							if t.gobType() == nil {
//line /usr/local/go/src/encoding/gob/type.go:198
		_go_fuzz_dep_.CoverTab[85245]++
								return "<nil>"
//line /usr/local/go/src/encoding/gob/type.go:199
		// _ = "end of CoverTab[85245]"
	} else {
//line /usr/local/go/src/encoding/gob/type.go:200
		_go_fuzz_dep_.CoverTab[85246]++
//line /usr/local/go/src/encoding/gob/type.go:200
		// _ = "end of CoverTab[85246]"
//line /usr/local/go/src/encoding/gob/type.go:200
	}
//line /usr/local/go/src/encoding/gob/type.go:200
	// _ = "end of CoverTab[85243]"
//line /usr/local/go/src/encoding/gob/type.go:200
	_go_fuzz_dep_.CoverTab[85244]++
							return t.gobType().string()
//line /usr/local/go/src/encoding/gob/type.go:201
	// _ = "end of CoverTab[85244]"
}

// Name returns the name of the type associated with the typeId.
func (t typeId) name() string {
//line /usr/local/go/src/encoding/gob/type.go:205
	_go_fuzz_dep_.CoverTab[85247]++
							if t.gobType() == nil {
//line /usr/local/go/src/encoding/gob/type.go:206
		_go_fuzz_dep_.CoverTab[85249]++
								return "<nil>"
//line /usr/local/go/src/encoding/gob/type.go:207
		// _ = "end of CoverTab[85249]"
	} else {
//line /usr/local/go/src/encoding/gob/type.go:208
		_go_fuzz_dep_.CoverTab[85250]++
//line /usr/local/go/src/encoding/gob/type.go:208
		// _ = "end of CoverTab[85250]"
//line /usr/local/go/src/encoding/gob/type.go:208
	}
//line /usr/local/go/src/encoding/gob/type.go:208
	// _ = "end of CoverTab[85247]"
//line /usr/local/go/src/encoding/gob/type.go:208
	_go_fuzz_dep_.CoverTab[85248]++
							return t.gobType().name()
//line /usr/local/go/src/encoding/gob/type.go:209
	// _ = "end of CoverTab[85248]"
}

// CommonType holds elements of all types.
//line /usr/local/go/src/encoding/gob/type.go:212
// It is a historical artifact, kept for binary compatibility and exported
//line /usr/local/go/src/encoding/gob/type.go:212
// only for the benefit of the package's encoding of type descriptors. It is
//line /usr/local/go/src/encoding/gob/type.go:212
// not intended for direct use by clients.
//line /usr/local/go/src/encoding/gob/type.go:216
type CommonType struct {
	Name	string
	Id	typeId
}

func (t *CommonType) id() typeId {
//line /usr/local/go/src/encoding/gob/type.go:221
	_go_fuzz_dep_.CoverTab[85251]++
//line /usr/local/go/src/encoding/gob/type.go:221
	return t.Id
//line /usr/local/go/src/encoding/gob/type.go:221
	// _ = "end of CoverTab[85251]"
//line /usr/local/go/src/encoding/gob/type.go:221
}

func (t *CommonType) setId(id typeId) {
//line /usr/local/go/src/encoding/gob/type.go:223
	_go_fuzz_dep_.CoverTab[85252]++
//line /usr/local/go/src/encoding/gob/type.go:223
	t.Id = id
//line /usr/local/go/src/encoding/gob/type.go:223
	// _ = "end of CoverTab[85252]"
//line /usr/local/go/src/encoding/gob/type.go:223
}

func (t *CommonType) string() string {
//line /usr/local/go/src/encoding/gob/type.go:225
	_go_fuzz_dep_.CoverTab[85253]++
//line /usr/local/go/src/encoding/gob/type.go:225
	return t.Name
//line /usr/local/go/src/encoding/gob/type.go:225
	// _ = "end of CoverTab[85253]"
//line /usr/local/go/src/encoding/gob/type.go:225
}

func (t *CommonType) safeString(seen map[typeId]bool) string {
//line /usr/local/go/src/encoding/gob/type.go:227
	_go_fuzz_dep_.CoverTab[85254]++
							return t.Name
//line /usr/local/go/src/encoding/gob/type.go:228
	// _ = "end of CoverTab[85254]"
}

func (t *CommonType) name() string {
//line /usr/local/go/src/encoding/gob/type.go:231
	_go_fuzz_dep_.CoverTab[85255]++
//line /usr/local/go/src/encoding/gob/type.go:231
	return t.Name
//line /usr/local/go/src/encoding/gob/type.go:231
	// _ = "end of CoverTab[85255]"
//line /usr/local/go/src/encoding/gob/type.go:231
}

//line /usr/local/go/src/encoding/gob/type.go:236
var (
	// Primordial types, needed during initialization.
	// Always passed as pointers so the interface{} type
	// goes through without losing its interfaceness.
	tBool		= bootstrapType("bool", (*bool)(nil), 1)
	tInt		= bootstrapType("int", (*int)(nil), 2)
	tUint		= bootstrapType("uint", (*uint)(nil), 3)
	tFloat		= bootstrapType("float", (*float64)(nil), 4)
	tBytes		= bootstrapType("bytes", (*[]byte)(nil), 5)
	tString		= bootstrapType("string", (*string)(nil), 6)
	tComplex	= bootstrapType("complex", (*complex128)(nil), 7)
	tInterface	= bootstrapType("interface", (*any)(nil), 8)
	// Reserve some Ids for compatible expansion
	tReserved7	= bootstrapType("_reserved1", (*struct{ r7 int })(nil), 9)
	tReserved6	= bootstrapType("_reserved1", (*struct{ r6 int })(nil), 10)
	tReserved5	= bootstrapType("_reserved1", (*struct{ r5 int })(nil), 11)
	tReserved4	= bootstrapType("_reserved1", (*struct{ r4 int })(nil), 12)
	tReserved3	= bootstrapType("_reserved1", (*struct{ r3 int })(nil), 13)
	tReserved2	= bootstrapType("_reserved1", (*struct{ r2 int })(nil), 14)
	tReserved1	= bootstrapType("_reserved1", (*struct{ r1 int })(nil), 15)
)

// Predefined because it's needed by the Decoder
var tWireType = mustGetTypeInfo(reflect.TypeOf(wireType{})).id
var wireTypeUserInfo *userTypeInfo	// userTypeInfo of (*wireType)

func init() {

	checkId(16, tWireType)
	checkId(17, mustGetTypeInfo(reflect.TypeOf(arrayType{})).id)
	checkId(18, mustGetTypeInfo(reflect.TypeOf(CommonType{})).id)
	checkId(19, mustGetTypeInfo(reflect.TypeOf(sliceType{})).id)
	checkId(20, mustGetTypeInfo(reflect.TypeOf(structType{})).id)
	checkId(21, mustGetTypeInfo(reflect.TypeOf(fieldType{})).id)
	checkId(23, mustGetTypeInfo(reflect.TypeOf(mapType{})).id)

	builtinIdToType = make(map[typeId]gobType)
	for k, v := range idToType {
		builtinIdToType[k] = v
	}

//line /usr/local/go/src/encoding/gob/type.go:279
	if nextId > firstUserId {
		panic(fmt.Sprintln("nextId too large:", nextId))
	}
	nextId = firstUserId
	registerBasics()
	wireTypeUserInfo = userType(reflect.TypeOf((*wireType)(nil)))
}

// Array type
type arrayType struct {
	CommonType
	Elem	typeId
	Len	int
}

func newArrayType(name string) *arrayType {
//line /usr/local/go/src/encoding/gob/type.go:294
	_go_fuzz_dep_.CoverTab[85256]++
							a := &arrayType{CommonType{Name: name}, 0, 0}
							return a
//line /usr/local/go/src/encoding/gob/type.go:296
	// _ = "end of CoverTab[85256]"
}

func (a *arrayType) init(elem gobType, len int) {

	setTypeId(a)
	a.Elem = elem.id()
	a.Len = len
}

func (a *arrayType) safeString(seen map[typeId]bool) string {
//line /usr/local/go/src/encoding/gob/type.go:306
	_go_fuzz_dep_.CoverTab[85257]++
							if seen[a.Id] {
//line /usr/local/go/src/encoding/gob/type.go:307
		_go_fuzz_dep_.CoverTab[85259]++
								return a.Name
//line /usr/local/go/src/encoding/gob/type.go:308
		// _ = "end of CoverTab[85259]"
	} else {
//line /usr/local/go/src/encoding/gob/type.go:309
		_go_fuzz_dep_.CoverTab[85260]++
//line /usr/local/go/src/encoding/gob/type.go:309
		// _ = "end of CoverTab[85260]"
//line /usr/local/go/src/encoding/gob/type.go:309
	}
//line /usr/local/go/src/encoding/gob/type.go:309
	// _ = "end of CoverTab[85257]"
//line /usr/local/go/src/encoding/gob/type.go:309
	_go_fuzz_dep_.CoverTab[85258]++
							seen[a.Id] = true
							return fmt.Sprintf("[%d]%s", a.Len, a.Elem.gobType().safeString(seen))
//line /usr/local/go/src/encoding/gob/type.go:311
	// _ = "end of CoverTab[85258]"
}

func (a *arrayType) string() string {
//line /usr/local/go/src/encoding/gob/type.go:314
	_go_fuzz_dep_.CoverTab[85261]++
//line /usr/local/go/src/encoding/gob/type.go:314
	return a.safeString(make(map[typeId]bool))
//line /usr/local/go/src/encoding/gob/type.go:314
	// _ = "end of CoverTab[85261]"
//line /usr/local/go/src/encoding/gob/type.go:314
}

// GobEncoder type (something that implements the GobEncoder interface)
type gobEncoderType struct {
	CommonType
}

func newGobEncoderType(name string) *gobEncoderType {
//line /usr/local/go/src/encoding/gob/type.go:321
	_go_fuzz_dep_.CoverTab[85262]++
							g := &gobEncoderType{CommonType{Name: name}}
							setTypeId(g)
							return g
//line /usr/local/go/src/encoding/gob/type.go:324
	// _ = "end of CoverTab[85262]"
}

func (g *gobEncoderType) safeString(seen map[typeId]bool) string {
//line /usr/local/go/src/encoding/gob/type.go:327
	_go_fuzz_dep_.CoverTab[85263]++
							return g.Name
//line /usr/local/go/src/encoding/gob/type.go:328
	// _ = "end of CoverTab[85263]"
}

func (g *gobEncoderType) string() string {
//line /usr/local/go/src/encoding/gob/type.go:331
	_go_fuzz_dep_.CoverTab[85264]++
//line /usr/local/go/src/encoding/gob/type.go:331
	return g.Name
//line /usr/local/go/src/encoding/gob/type.go:331
	// _ = "end of CoverTab[85264]"
//line /usr/local/go/src/encoding/gob/type.go:331
}

// Map type
type mapType struct {
	CommonType
	Key	typeId
	Elem	typeId
}

func newMapType(name string) *mapType {
//line /usr/local/go/src/encoding/gob/type.go:340
	_go_fuzz_dep_.CoverTab[85265]++
							m := &mapType{CommonType{Name: name}, 0, 0}
							return m
//line /usr/local/go/src/encoding/gob/type.go:342
	// _ = "end of CoverTab[85265]"
}

func (m *mapType) init(key, elem gobType) {

	setTypeId(m)
	m.Key = key.id()
	m.Elem = elem.id()
}

func (m *mapType) safeString(seen map[typeId]bool) string {
//line /usr/local/go/src/encoding/gob/type.go:352
	_go_fuzz_dep_.CoverTab[85266]++
							if seen[m.Id] {
//line /usr/local/go/src/encoding/gob/type.go:353
		_go_fuzz_dep_.CoverTab[85268]++
								return m.Name
//line /usr/local/go/src/encoding/gob/type.go:354
		// _ = "end of CoverTab[85268]"
	} else {
//line /usr/local/go/src/encoding/gob/type.go:355
		_go_fuzz_dep_.CoverTab[85269]++
//line /usr/local/go/src/encoding/gob/type.go:355
		// _ = "end of CoverTab[85269]"
//line /usr/local/go/src/encoding/gob/type.go:355
	}
//line /usr/local/go/src/encoding/gob/type.go:355
	// _ = "end of CoverTab[85266]"
//line /usr/local/go/src/encoding/gob/type.go:355
	_go_fuzz_dep_.CoverTab[85267]++
							seen[m.Id] = true
							key := m.Key.gobType().safeString(seen)
							elem := m.Elem.gobType().safeString(seen)
							return fmt.Sprintf("map[%s]%s", key, elem)
//line /usr/local/go/src/encoding/gob/type.go:359
	// _ = "end of CoverTab[85267]"
}

func (m *mapType) string() string {
//line /usr/local/go/src/encoding/gob/type.go:362
	_go_fuzz_dep_.CoverTab[85270]++
//line /usr/local/go/src/encoding/gob/type.go:362
	return m.safeString(make(map[typeId]bool))
//line /usr/local/go/src/encoding/gob/type.go:362
	// _ = "end of CoverTab[85270]"
//line /usr/local/go/src/encoding/gob/type.go:362
}

// Slice type
type sliceType struct {
	CommonType
	Elem	typeId
}

func newSliceType(name string) *sliceType {
//line /usr/local/go/src/encoding/gob/type.go:370
	_go_fuzz_dep_.CoverTab[85271]++
							s := &sliceType{CommonType{Name: name}, 0}
							return s
//line /usr/local/go/src/encoding/gob/type.go:372
	// _ = "end of CoverTab[85271]"
}

func (s *sliceType) init(elem gobType) {

							setTypeId(s)

//line /usr/local/go/src/encoding/gob/type.go:380
	if elem.id() == 0 {
		setTypeId(elem)
	}
	s.Elem = elem.id()
}

func (s *sliceType) safeString(seen map[typeId]bool) string {
//line /usr/local/go/src/encoding/gob/type.go:386
	_go_fuzz_dep_.CoverTab[85272]++
							if seen[s.Id] {
//line /usr/local/go/src/encoding/gob/type.go:387
		_go_fuzz_dep_.CoverTab[85274]++
								return s.Name
//line /usr/local/go/src/encoding/gob/type.go:388
		// _ = "end of CoverTab[85274]"
	} else {
//line /usr/local/go/src/encoding/gob/type.go:389
		_go_fuzz_dep_.CoverTab[85275]++
//line /usr/local/go/src/encoding/gob/type.go:389
		// _ = "end of CoverTab[85275]"
//line /usr/local/go/src/encoding/gob/type.go:389
	}
//line /usr/local/go/src/encoding/gob/type.go:389
	// _ = "end of CoverTab[85272]"
//line /usr/local/go/src/encoding/gob/type.go:389
	_go_fuzz_dep_.CoverTab[85273]++
							seen[s.Id] = true
							return fmt.Sprintf("[]%s", s.Elem.gobType().safeString(seen))
//line /usr/local/go/src/encoding/gob/type.go:391
	// _ = "end of CoverTab[85273]"
}

func (s *sliceType) string() string {
//line /usr/local/go/src/encoding/gob/type.go:394
	_go_fuzz_dep_.CoverTab[85276]++
//line /usr/local/go/src/encoding/gob/type.go:394
	return s.safeString(make(map[typeId]bool))
//line /usr/local/go/src/encoding/gob/type.go:394
	// _ = "end of CoverTab[85276]"
//line /usr/local/go/src/encoding/gob/type.go:394
}

// Struct type
type fieldType struct {
	Name	string
	Id	typeId
}

type structType struct {
	CommonType
	Field	[]*fieldType
}

func (s *structType) safeString(seen map[typeId]bool) string {
//line /usr/local/go/src/encoding/gob/type.go:407
	_go_fuzz_dep_.CoverTab[85277]++
							if s == nil {
//line /usr/local/go/src/encoding/gob/type.go:408
		_go_fuzz_dep_.CoverTab[85281]++
								return "<nil>"
//line /usr/local/go/src/encoding/gob/type.go:409
		// _ = "end of CoverTab[85281]"
	} else {
//line /usr/local/go/src/encoding/gob/type.go:410
		_go_fuzz_dep_.CoverTab[85282]++
//line /usr/local/go/src/encoding/gob/type.go:410
		// _ = "end of CoverTab[85282]"
//line /usr/local/go/src/encoding/gob/type.go:410
	}
//line /usr/local/go/src/encoding/gob/type.go:410
	// _ = "end of CoverTab[85277]"
//line /usr/local/go/src/encoding/gob/type.go:410
	_go_fuzz_dep_.CoverTab[85278]++
							if _, ok := seen[s.Id]; ok {
//line /usr/local/go/src/encoding/gob/type.go:411
		_go_fuzz_dep_.CoverTab[85283]++
								return s.Name
//line /usr/local/go/src/encoding/gob/type.go:412
		// _ = "end of CoverTab[85283]"
	} else {
//line /usr/local/go/src/encoding/gob/type.go:413
		_go_fuzz_dep_.CoverTab[85284]++
//line /usr/local/go/src/encoding/gob/type.go:413
		// _ = "end of CoverTab[85284]"
//line /usr/local/go/src/encoding/gob/type.go:413
	}
//line /usr/local/go/src/encoding/gob/type.go:413
	// _ = "end of CoverTab[85278]"
//line /usr/local/go/src/encoding/gob/type.go:413
	_go_fuzz_dep_.CoverTab[85279]++
							seen[s.Id] = true
							str := s.Name + " = struct { "
							for _, f := range s.Field {
//line /usr/local/go/src/encoding/gob/type.go:416
		_go_fuzz_dep_.CoverTab[85285]++
								str += fmt.Sprintf("%s %s; ", f.Name, f.Id.gobType().safeString(seen))
//line /usr/local/go/src/encoding/gob/type.go:417
		// _ = "end of CoverTab[85285]"
	}
//line /usr/local/go/src/encoding/gob/type.go:418
	// _ = "end of CoverTab[85279]"
//line /usr/local/go/src/encoding/gob/type.go:418
	_go_fuzz_dep_.CoverTab[85280]++
							str += "}"
							return str
//line /usr/local/go/src/encoding/gob/type.go:420
	// _ = "end of CoverTab[85280]"
}

func (s *structType) string() string {
//line /usr/local/go/src/encoding/gob/type.go:423
	_go_fuzz_dep_.CoverTab[85286]++
//line /usr/local/go/src/encoding/gob/type.go:423
	return s.safeString(make(map[typeId]bool))
//line /usr/local/go/src/encoding/gob/type.go:423
	// _ = "end of CoverTab[85286]"
//line /usr/local/go/src/encoding/gob/type.go:423
}

func newStructType(name string) *structType {
//line /usr/local/go/src/encoding/gob/type.go:425
	_go_fuzz_dep_.CoverTab[85287]++
							s := &structType{CommonType{Name: name}, nil}

//line /usr/local/go/src/encoding/gob/type.go:429
	setTypeId(s)
							return s
//line /usr/local/go/src/encoding/gob/type.go:430
	// _ = "end of CoverTab[85287]"
}

// newTypeObject allocates a gobType for the reflection type rt.
//line /usr/local/go/src/encoding/gob/type.go:433
// Unless ut represents a GobEncoder, rt should be the base type
//line /usr/local/go/src/encoding/gob/type.go:433
// of ut.
//line /usr/local/go/src/encoding/gob/type.go:433
// This is only called from the encoding side. The decoding side
//line /usr/local/go/src/encoding/gob/type.go:433
// works through typeIds and userTypeInfos alone.
//line /usr/local/go/src/encoding/gob/type.go:438
func newTypeObject(name string, ut *userTypeInfo, rt reflect.Type) (gobType, error) {
//line /usr/local/go/src/encoding/gob/type.go:438
	_go_fuzz_dep_.CoverTab[85288]++

							if ut.externalEnc != 0 {
//line /usr/local/go/src/encoding/gob/type.go:440
		_go_fuzz_dep_.CoverTab[85291]++
								return newGobEncoderType(name), nil
//line /usr/local/go/src/encoding/gob/type.go:441
		// _ = "end of CoverTab[85291]"
	} else {
//line /usr/local/go/src/encoding/gob/type.go:442
		_go_fuzz_dep_.CoverTab[85292]++
//line /usr/local/go/src/encoding/gob/type.go:442
		// _ = "end of CoverTab[85292]"
//line /usr/local/go/src/encoding/gob/type.go:442
	}
//line /usr/local/go/src/encoding/gob/type.go:442
	// _ = "end of CoverTab[85288]"
//line /usr/local/go/src/encoding/gob/type.go:442
	_go_fuzz_dep_.CoverTab[85289]++
							var err error
							var type0, type1 gobType
							defer func() {
//line /usr/local/go/src/encoding/gob/type.go:445
		_go_fuzz_dep_.CoverTab[85293]++
								if err != nil {
//line /usr/local/go/src/encoding/gob/type.go:446
			_go_fuzz_dep_.CoverTab[85294]++
									delete(types, rt)
//line /usr/local/go/src/encoding/gob/type.go:447
			// _ = "end of CoverTab[85294]"
		} else {
//line /usr/local/go/src/encoding/gob/type.go:448
			_go_fuzz_dep_.CoverTab[85295]++
//line /usr/local/go/src/encoding/gob/type.go:448
			// _ = "end of CoverTab[85295]"
//line /usr/local/go/src/encoding/gob/type.go:448
		}
//line /usr/local/go/src/encoding/gob/type.go:448
		// _ = "end of CoverTab[85293]"
	}()
//line /usr/local/go/src/encoding/gob/type.go:449
	// _ = "end of CoverTab[85289]"
//line /usr/local/go/src/encoding/gob/type.go:449
	_go_fuzz_dep_.CoverTab[85290]++

//line /usr/local/go/src/encoding/gob/type.go:452
	switch t := rt; t.Kind() {

	case reflect.Bool:
//line /usr/local/go/src/encoding/gob/type.go:454
		_go_fuzz_dep_.CoverTab[85296]++
								return tBool.gobType(), nil
//line /usr/local/go/src/encoding/gob/type.go:455
		// _ = "end of CoverTab[85296]"

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /usr/local/go/src/encoding/gob/type.go:457
		_go_fuzz_dep_.CoverTab[85297]++
								return tInt.gobType(), nil
//line /usr/local/go/src/encoding/gob/type.go:458
		// _ = "end of CoverTab[85297]"

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
//line /usr/local/go/src/encoding/gob/type.go:460
		_go_fuzz_dep_.CoverTab[85298]++
								return tUint.gobType(), nil
//line /usr/local/go/src/encoding/gob/type.go:461
		// _ = "end of CoverTab[85298]"

	case reflect.Float32, reflect.Float64:
//line /usr/local/go/src/encoding/gob/type.go:463
		_go_fuzz_dep_.CoverTab[85299]++
								return tFloat.gobType(), nil
//line /usr/local/go/src/encoding/gob/type.go:464
		// _ = "end of CoverTab[85299]"

	case reflect.Complex64, reflect.Complex128:
//line /usr/local/go/src/encoding/gob/type.go:466
		_go_fuzz_dep_.CoverTab[85300]++
								return tComplex.gobType(), nil
//line /usr/local/go/src/encoding/gob/type.go:467
		// _ = "end of CoverTab[85300]"

	case reflect.String:
//line /usr/local/go/src/encoding/gob/type.go:469
		_go_fuzz_dep_.CoverTab[85301]++
								return tString.gobType(), nil
//line /usr/local/go/src/encoding/gob/type.go:470
		// _ = "end of CoverTab[85301]"

	case reflect.Interface:
//line /usr/local/go/src/encoding/gob/type.go:472
		_go_fuzz_dep_.CoverTab[85302]++
								return tInterface.gobType(), nil
//line /usr/local/go/src/encoding/gob/type.go:473
		// _ = "end of CoverTab[85302]"

	case reflect.Array:
//line /usr/local/go/src/encoding/gob/type.go:475
		_go_fuzz_dep_.CoverTab[85303]++
								at := newArrayType(name)
								types[rt] = at
								type0, err = getBaseType("", t.Elem())
								if err != nil {
//line /usr/local/go/src/encoding/gob/type.go:479
			_go_fuzz_dep_.CoverTab[85314]++
									return nil, err
//line /usr/local/go/src/encoding/gob/type.go:480
			// _ = "end of CoverTab[85314]"
		} else {
//line /usr/local/go/src/encoding/gob/type.go:481
			_go_fuzz_dep_.CoverTab[85315]++
//line /usr/local/go/src/encoding/gob/type.go:481
			// _ = "end of CoverTab[85315]"
//line /usr/local/go/src/encoding/gob/type.go:481
		}
//line /usr/local/go/src/encoding/gob/type.go:481
		// _ = "end of CoverTab[85303]"
//line /usr/local/go/src/encoding/gob/type.go:481
		_go_fuzz_dep_.CoverTab[85304]++

//line /usr/local/go/src/encoding/gob/type.go:490
		at.init(type0, t.Len())
								return at, nil
//line /usr/local/go/src/encoding/gob/type.go:491
		// _ = "end of CoverTab[85304]"

	case reflect.Map:
//line /usr/local/go/src/encoding/gob/type.go:493
		_go_fuzz_dep_.CoverTab[85305]++
								mt := newMapType(name)
								types[rt] = mt
								type0, err = getBaseType("", t.Key())
								if err != nil {
//line /usr/local/go/src/encoding/gob/type.go:497
			_go_fuzz_dep_.CoverTab[85316]++
									return nil, err
//line /usr/local/go/src/encoding/gob/type.go:498
			// _ = "end of CoverTab[85316]"
		} else {
//line /usr/local/go/src/encoding/gob/type.go:499
			_go_fuzz_dep_.CoverTab[85317]++
//line /usr/local/go/src/encoding/gob/type.go:499
			// _ = "end of CoverTab[85317]"
//line /usr/local/go/src/encoding/gob/type.go:499
		}
//line /usr/local/go/src/encoding/gob/type.go:499
		// _ = "end of CoverTab[85305]"
//line /usr/local/go/src/encoding/gob/type.go:499
		_go_fuzz_dep_.CoverTab[85306]++
								type1, err = getBaseType("", t.Elem())
								if err != nil {
//line /usr/local/go/src/encoding/gob/type.go:501
			_go_fuzz_dep_.CoverTab[85318]++
									return nil, err
//line /usr/local/go/src/encoding/gob/type.go:502
			// _ = "end of CoverTab[85318]"
		} else {
//line /usr/local/go/src/encoding/gob/type.go:503
			_go_fuzz_dep_.CoverTab[85319]++
//line /usr/local/go/src/encoding/gob/type.go:503
			// _ = "end of CoverTab[85319]"
//line /usr/local/go/src/encoding/gob/type.go:503
		}
//line /usr/local/go/src/encoding/gob/type.go:503
		// _ = "end of CoverTab[85306]"
//line /usr/local/go/src/encoding/gob/type.go:503
		_go_fuzz_dep_.CoverTab[85307]++
								mt.init(type0, type1)
								return mt, nil
//line /usr/local/go/src/encoding/gob/type.go:505
		// _ = "end of CoverTab[85307]"

	case reflect.Slice:
//line /usr/local/go/src/encoding/gob/type.go:507
		_go_fuzz_dep_.CoverTab[85308]++

								if t.Elem().Kind() == reflect.Uint8 {
//line /usr/local/go/src/encoding/gob/type.go:509
			_go_fuzz_dep_.CoverTab[85320]++
									return tBytes.gobType(), nil
//line /usr/local/go/src/encoding/gob/type.go:510
			// _ = "end of CoverTab[85320]"
		} else {
//line /usr/local/go/src/encoding/gob/type.go:511
			_go_fuzz_dep_.CoverTab[85321]++
//line /usr/local/go/src/encoding/gob/type.go:511
			// _ = "end of CoverTab[85321]"
//line /usr/local/go/src/encoding/gob/type.go:511
		}
//line /usr/local/go/src/encoding/gob/type.go:511
		// _ = "end of CoverTab[85308]"
//line /usr/local/go/src/encoding/gob/type.go:511
		_go_fuzz_dep_.CoverTab[85309]++
								st := newSliceType(name)
								types[rt] = st
								type0, err = getBaseType(t.Elem().Name(), t.Elem())
								if err != nil {
//line /usr/local/go/src/encoding/gob/type.go:515
			_go_fuzz_dep_.CoverTab[85322]++
									return nil, err
//line /usr/local/go/src/encoding/gob/type.go:516
			// _ = "end of CoverTab[85322]"
		} else {
//line /usr/local/go/src/encoding/gob/type.go:517
			_go_fuzz_dep_.CoverTab[85323]++
//line /usr/local/go/src/encoding/gob/type.go:517
			// _ = "end of CoverTab[85323]"
//line /usr/local/go/src/encoding/gob/type.go:517
		}
//line /usr/local/go/src/encoding/gob/type.go:517
		// _ = "end of CoverTab[85309]"
//line /usr/local/go/src/encoding/gob/type.go:517
		_go_fuzz_dep_.CoverTab[85310]++
								st.init(type0)
								return st, nil
//line /usr/local/go/src/encoding/gob/type.go:519
		// _ = "end of CoverTab[85310]"

	case reflect.Struct:
//line /usr/local/go/src/encoding/gob/type.go:521
		_go_fuzz_dep_.CoverTab[85311]++
								st := newStructType(name)
								types[rt] = st
								idToType[st.id()] = st
								for i := 0; i < t.NumField(); i++ {
//line /usr/local/go/src/encoding/gob/type.go:525
			_go_fuzz_dep_.CoverTab[85324]++
									f := t.Field(i)
									if !isSent(&f) {
//line /usr/local/go/src/encoding/gob/type.go:527
				_go_fuzz_dep_.CoverTab[85329]++
										continue
//line /usr/local/go/src/encoding/gob/type.go:528
				// _ = "end of CoverTab[85329]"
			} else {
//line /usr/local/go/src/encoding/gob/type.go:529
				_go_fuzz_dep_.CoverTab[85330]++
//line /usr/local/go/src/encoding/gob/type.go:529
				// _ = "end of CoverTab[85330]"
//line /usr/local/go/src/encoding/gob/type.go:529
			}
//line /usr/local/go/src/encoding/gob/type.go:529
			// _ = "end of CoverTab[85324]"
//line /usr/local/go/src/encoding/gob/type.go:529
			_go_fuzz_dep_.CoverTab[85325]++
									typ := userType(f.Type).base
									tname := typ.Name()
									if tname == "" {
//line /usr/local/go/src/encoding/gob/type.go:532
				_go_fuzz_dep_.CoverTab[85331]++
										t := userType(f.Type).base
										tname = t.String()
//line /usr/local/go/src/encoding/gob/type.go:534
				// _ = "end of CoverTab[85331]"
			} else {
//line /usr/local/go/src/encoding/gob/type.go:535
				_go_fuzz_dep_.CoverTab[85332]++
//line /usr/local/go/src/encoding/gob/type.go:535
				// _ = "end of CoverTab[85332]"
//line /usr/local/go/src/encoding/gob/type.go:535
			}
//line /usr/local/go/src/encoding/gob/type.go:535
			// _ = "end of CoverTab[85325]"
//line /usr/local/go/src/encoding/gob/type.go:535
			_go_fuzz_dep_.CoverTab[85326]++
									gt, err := getBaseType(tname, f.Type)
									if err != nil {
//line /usr/local/go/src/encoding/gob/type.go:537
				_go_fuzz_dep_.CoverTab[85333]++
										return nil, err
//line /usr/local/go/src/encoding/gob/type.go:538
				// _ = "end of CoverTab[85333]"
			} else {
//line /usr/local/go/src/encoding/gob/type.go:539
				_go_fuzz_dep_.CoverTab[85334]++
//line /usr/local/go/src/encoding/gob/type.go:539
				// _ = "end of CoverTab[85334]"
//line /usr/local/go/src/encoding/gob/type.go:539
			}
//line /usr/local/go/src/encoding/gob/type.go:539
			// _ = "end of CoverTab[85326]"
//line /usr/local/go/src/encoding/gob/type.go:539
			_go_fuzz_dep_.CoverTab[85327]++

//line /usr/local/go/src/encoding/gob/type.go:544
			if gt.id() == 0 {
//line /usr/local/go/src/encoding/gob/type.go:544
				_go_fuzz_dep_.CoverTab[85335]++
										setTypeId(gt)
//line /usr/local/go/src/encoding/gob/type.go:545
				// _ = "end of CoverTab[85335]"
			} else {
//line /usr/local/go/src/encoding/gob/type.go:546
				_go_fuzz_dep_.CoverTab[85336]++
//line /usr/local/go/src/encoding/gob/type.go:546
				// _ = "end of CoverTab[85336]"
//line /usr/local/go/src/encoding/gob/type.go:546
			}
//line /usr/local/go/src/encoding/gob/type.go:546
			// _ = "end of CoverTab[85327]"
//line /usr/local/go/src/encoding/gob/type.go:546
			_go_fuzz_dep_.CoverTab[85328]++
									st.Field = append(st.Field, &fieldType{f.Name, gt.id()})
//line /usr/local/go/src/encoding/gob/type.go:547
			// _ = "end of CoverTab[85328]"
		}
//line /usr/local/go/src/encoding/gob/type.go:548
		// _ = "end of CoverTab[85311]"
//line /usr/local/go/src/encoding/gob/type.go:548
		_go_fuzz_dep_.CoverTab[85312]++
								return st, nil
//line /usr/local/go/src/encoding/gob/type.go:549
		// _ = "end of CoverTab[85312]"

	default:
//line /usr/local/go/src/encoding/gob/type.go:551
		_go_fuzz_dep_.CoverTab[85313]++
								return nil, errors.New("gob NewTypeObject can't handle type: " + rt.String())
//line /usr/local/go/src/encoding/gob/type.go:552
		// _ = "end of CoverTab[85313]"
	}
//line /usr/local/go/src/encoding/gob/type.go:553
	// _ = "end of CoverTab[85290]"
}

// isExported reports whether this is an exported - upper case - name.
func isExported(name string) bool {
//line /usr/local/go/src/encoding/gob/type.go:557
	_go_fuzz_dep_.CoverTab[85337]++
							rune, _ := utf8.DecodeRuneInString(name)
							return unicode.IsUpper(rune)
//line /usr/local/go/src/encoding/gob/type.go:559
	// _ = "end of CoverTab[85337]"
}

// isSent reports whether this struct field is to be transmitted.
//line /usr/local/go/src/encoding/gob/type.go:562
// It will be transmitted only if it is exported and not a chan or func field
//line /usr/local/go/src/encoding/gob/type.go:562
// or pointer to chan or func.
//line /usr/local/go/src/encoding/gob/type.go:565
func isSent(field *reflect.StructField) bool {
//line /usr/local/go/src/encoding/gob/type.go:565
	_go_fuzz_dep_.CoverTab[85338]++
							if !isExported(field.Name) {
//line /usr/local/go/src/encoding/gob/type.go:566
		_go_fuzz_dep_.CoverTab[85342]++
								return false
//line /usr/local/go/src/encoding/gob/type.go:567
		// _ = "end of CoverTab[85342]"
	} else {
//line /usr/local/go/src/encoding/gob/type.go:568
		_go_fuzz_dep_.CoverTab[85343]++
//line /usr/local/go/src/encoding/gob/type.go:568
		// _ = "end of CoverTab[85343]"
//line /usr/local/go/src/encoding/gob/type.go:568
	}
//line /usr/local/go/src/encoding/gob/type.go:568
	// _ = "end of CoverTab[85338]"
//line /usr/local/go/src/encoding/gob/type.go:568
	_go_fuzz_dep_.CoverTab[85339]++

//line /usr/local/go/src/encoding/gob/type.go:571
	typ := field.Type
	for typ.Kind() == reflect.Pointer {
//line /usr/local/go/src/encoding/gob/type.go:572
		_go_fuzz_dep_.CoverTab[85344]++
								typ = typ.Elem()
//line /usr/local/go/src/encoding/gob/type.go:573
		// _ = "end of CoverTab[85344]"
	}
//line /usr/local/go/src/encoding/gob/type.go:574
	// _ = "end of CoverTab[85339]"
//line /usr/local/go/src/encoding/gob/type.go:574
	_go_fuzz_dep_.CoverTab[85340]++
							if typ.Kind() == reflect.Chan || func() bool {
//line /usr/local/go/src/encoding/gob/type.go:575
		_go_fuzz_dep_.CoverTab[85345]++
//line /usr/local/go/src/encoding/gob/type.go:575
		return typ.Kind() == reflect.Func
//line /usr/local/go/src/encoding/gob/type.go:575
		// _ = "end of CoverTab[85345]"
//line /usr/local/go/src/encoding/gob/type.go:575
	}() {
//line /usr/local/go/src/encoding/gob/type.go:575
		_go_fuzz_dep_.CoverTab[85346]++
								return false
//line /usr/local/go/src/encoding/gob/type.go:576
		// _ = "end of CoverTab[85346]"
	} else {
//line /usr/local/go/src/encoding/gob/type.go:577
		_go_fuzz_dep_.CoverTab[85347]++
//line /usr/local/go/src/encoding/gob/type.go:577
		// _ = "end of CoverTab[85347]"
//line /usr/local/go/src/encoding/gob/type.go:577
	}
//line /usr/local/go/src/encoding/gob/type.go:577
	// _ = "end of CoverTab[85340]"
//line /usr/local/go/src/encoding/gob/type.go:577
	_go_fuzz_dep_.CoverTab[85341]++
							return true
//line /usr/local/go/src/encoding/gob/type.go:578
	// _ = "end of CoverTab[85341]"
}

// getBaseType returns the Gob type describing the given reflect.Type's base type.
//line /usr/local/go/src/encoding/gob/type.go:581
// typeLock must be held.
//line /usr/local/go/src/encoding/gob/type.go:583
func getBaseType(name string, rt reflect.Type) (gobType, error) {
//line /usr/local/go/src/encoding/gob/type.go:583
	_go_fuzz_dep_.CoverTab[85348]++
							ut := userType(rt)
							return getType(name, ut, ut.base)
//line /usr/local/go/src/encoding/gob/type.go:585
	// _ = "end of CoverTab[85348]"
}

// getType returns the Gob type describing the given reflect.Type.
//line /usr/local/go/src/encoding/gob/type.go:588
// Should be called only when handling GobEncoders/Decoders,
//line /usr/local/go/src/encoding/gob/type.go:588
// which may be pointers. All other types are handled through the
//line /usr/local/go/src/encoding/gob/type.go:588
// base type, never a pointer.
//line /usr/local/go/src/encoding/gob/type.go:588
// typeLock must be held.
//line /usr/local/go/src/encoding/gob/type.go:593
func getType(name string, ut *userTypeInfo, rt reflect.Type) (gobType, error) {
//line /usr/local/go/src/encoding/gob/type.go:593
	_go_fuzz_dep_.CoverTab[85349]++
							typ, present := types[rt]
							if present {
//line /usr/local/go/src/encoding/gob/type.go:595
		_go_fuzz_dep_.CoverTab[85352]++
								return typ, nil
//line /usr/local/go/src/encoding/gob/type.go:596
		// _ = "end of CoverTab[85352]"
	} else {
//line /usr/local/go/src/encoding/gob/type.go:597
		_go_fuzz_dep_.CoverTab[85353]++
//line /usr/local/go/src/encoding/gob/type.go:597
		// _ = "end of CoverTab[85353]"
//line /usr/local/go/src/encoding/gob/type.go:597
	}
//line /usr/local/go/src/encoding/gob/type.go:597
	// _ = "end of CoverTab[85349]"
//line /usr/local/go/src/encoding/gob/type.go:597
	_go_fuzz_dep_.CoverTab[85350]++
							typ, err := newTypeObject(name, ut, rt)
							if err == nil {
//line /usr/local/go/src/encoding/gob/type.go:599
		_go_fuzz_dep_.CoverTab[85354]++
								types[rt] = typ
//line /usr/local/go/src/encoding/gob/type.go:600
		// _ = "end of CoverTab[85354]"
	} else {
//line /usr/local/go/src/encoding/gob/type.go:601
		_go_fuzz_dep_.CoverTab[85355]++
//line /usr/local/go/src/encoding/gob/type.go:601
		// _ = "end of CoverTab[85355]"
//line /usr/local/go/src/encoding/gob/type.go:601
	}
//line /usr/local/go/src/encoding/gob/type.go:601
	// _ = "end of CoverTab[85350]"
//line /usr/local/go/src/encoding/gob/type.go:601
	_go_fuzz_dep_.CoverTab[85351]++
							return typ, err
//line /usr/local/go/src/encoding/gob/type.go:602
	// _ = "end of CoverTab[85351]"
}

func checkId(want, got typeId) {
//line /usr/local/go/src/encoding/gob/type.go:605
	_go_fuzz_dep_.CoverTab[85356]++
							if want != got {
//line /usr/local/go/src/encoding/gob/type.go:606
		_go_fuzz_dep_.CoverTab[85357]++
								fmt.Fprintf(os.Stderr, "checkId: %d should be %d\n", int(got), int(want))
								panic("bootstrap type wrong id: " + got.name() + " " + got.string() + " not " + want.string())
//line /usr/local/go/src/encoding/gob/type.go:608
		// _ = "end of CoverTab[85357]"
	} else {
//line /usr/local/go/src/encoding/gob/type.go:609
		_go_fuzz_dep_.CoverTab[85358]++
//line /usr/local/go/src/encoding/gob/type.go:609
		// _ = "end of CoverTab[85358]"
//line /usr/local/go/src/encoding/gob/type.go:609
	}
//line /usr/local/go/src/encoding/gob/type.go:609
	// _ = "end of CoverTab[85356]"
}

// used for building the basic types; called only from init().  the incoming
//line /usr/local/go/src/encoding/gob/type.go:612
// interface always refers to a pointer.
//line /usr/local/go/src/encoding/gob/type.go:614
func bootstrapType(name string, e any, expect typeId) typeId {
//line /usr/local/go/src/encoding/gob/type.go:614
	_go_fuzz_dep_.CoverTab[85359]++
							rt := reflect.TypeOf(e).Elem()
							_, present := types[rt]
							if present {
//line /usr/local/go/src/encoding/gob/type.go:617
		_go_fuzz_dep_.CoverTab[85361]++
								panic("bootstrap type already present: " + name + ", " + rt.String())
//line /usr/local/go/src/encoding/gob/type.go:618
		// _ = "end of CoverTab[85361]"
	} else {
//line /usr/local/go/src/encoding/gob/type.go:619
		_go_fuzz_dep_.CoverTab[85362]++
//line /usr/local/go/src/encoding/gob/type.go:619
		// _ = "end of CoverTab[85362]"
//line /usr/local/go/src/encoding/gob/type.go:619
	}
//line /usr/local/go/src/encoding/gob/type.go:619
	// _ = "end of CoverTab[85359]"
//line /usr/local/go/src/encoding/gob/type.go:619
	_go_fuzz_dep_.CoverTab[85360]++
							typ := &CommonType{Name: name}
							types[rt] = typ
							setTypeId(typ)
							checkId(expect, nextId)
							userType(rt)
							return nextId
//line /usr/local/go/src/encoding/gob/type.go:625
	// _ = "end of CoverTab[85360]"
}

// Representation of the information we send and receive about this type.
//line /usr/local/go/src/encoding/gob/type.go:628
// Each value we send is preceded by its type definition: an encoded int.
//line /usr/local/go/src/encoding/gob/type.go:628
// However, the very first time we send the value, we first send the pair
//line /usr/local/go/src/encoding/gob/type.go:628
// (-id, wireType).
//line /usr/local/go/src/encoding/gob/type.go:628
// For bootstrapping purposes, we assume that the recipient knows how
//line /usr/local/go/src/encoding/gob/type.go:628
// to decode a wireType; it is exactly the wireType struct here, interpreted
//line /usr/local/go/src/encoding/gob/type.go:628
// using the gob rules for sending a structure, except that we assume the
//line /usr/local/go/src/encoding/gob/type.go:628
// ids for wireType and structType etc. are known. The relevant pieces
//line /usr/local/go/src/encoding/gob/type.go:628
// are built in encode.go's init() function.
//line /usr/local/go/src/encoding/gob/type.go:628
// To maintain binary compatibility, if you extend this type, always put
//line /usr/local/go/src/encoding/gob/type.go:628
// the new fields last.
//line /usr/local/go/src/encoding/gob/type.go:639
type wireType struct {
	ArrayT			*arrayType
	SliceT			*sliceType
	StructT			*structType
	MapT			*mapType
	GobEncoderT		*gobEncoderType
	BinaryMarshalerT	*gobEncoderType
	TextMarshalerT		*gobEncoderType
}

func (w *wireType) string() string {
//line /usr/local/go/src/encoding/gob/type.go:649
	_go_fuzz_dep_.CoverTab[85363]++
							const unknown = "unknown type"
							if w == nil {
//line /usr/local/go/src/encoding/gob/type.go:651
		_go_fuzz_dep_.CoverTab[85366]++
								return unknown
//line /usr/local/go/src/encoding/gob/type.go:652
		// _ = "end of CoverTab[85366]"
	} else {
//line /usr/local/go/src/encoding/gob/type.go:653
		_go_fuzz_dep_.CoverTab[85367]++
//line /usr/local/go/src/encoding/gob/type.go:653
		// _ = "end of CoverTab[85367]"
//line /usr/local/go/src/encoding/gob/type.go:653
	}
//line /usr/local/go/src/encoding/gob/type.go:653
	// _ = "end of CoverTab[85363]"
//line /usr/local/go/src/encoding/gob/type.go:653
	_go_fuzz_dep_.CoverTab[85364]++
							switch {
	case w.ArrayT != nil:
//line /usr/local/go/src/encoding/gob/type.go:655
		_go_fuzz_dep_.CoverTab[85368]++
								return w.ArrayT.Name
//line /usr/local/go/src/encoding/gob/type.go:656
		// _ = "end of CoverTab[85368]"
	case w.SliceT != nil:
//line /usr/local/go/src/encoding/gob/type.go:657
		_go_fuzz_dep_.CoverTab[85369]++
								return w.SliceT.Name
//line /usr/local/go/src/encoding/gob/type.go:658
		// _ = "end of CoverTab[85369]"
	case w.StructT != nil:
//line /usr/local/go/src/encoding/gob/type.go:659
		_go_fuzz_dep_.CoverTab[85370]++
								return w.StructT.Name
//line /usr/local/go/src/encoding/gob/type.go:660
		// _ = "end of CoverTab[85370]"
	case w.MapT != nil:
//line /usr/local/go/src/encoding/gob/type.go:661
		_go_fuzz_dep_.CoverTab[85371]++
								return w.MapT.Name
//line /usr/local/go/src/encoding/gob/type.go:662
		// _ = "end of CoverTab[85371]"
	case w.GobEncoderT != nil:
//line /usr/local/go/src/encoding/gob/type.go:663
		_go_fuzz_dep_.CoverTab[85372]++
								return w.GobEncoderT.Name
//line /usr/local/go/src/encoding/gob/type.go:664
		// _ = "end of CoverTab[85372]"
	case w.BinaryMarshalerT != nil:
//line /usr/local/go/src/encoding/gob/type.go:665
		_go_fuzz_dep_.CoverTab[85373]++
								return w.BinaryMarshalerT.Name
//line /usr/local/go/src/encoding/gob/type.go:666
		// _ = "end of CoverTab[85373]"
	case w.TextMarshalerT != nil:
//line /usr/local/go/src/encoding/gob/type.go:667
		_go_fuzz_dep_.CoverTab[85374]++
								return w.TextMarshalerT.Name
//line /usr/local/go/src/encoding/gob/type.go:668
		// _ = "end of CoverTab[85374]"
//line /usr/local/go/src/encoding/gob/type.go:668
	default:
//line /usr/local/go/src/encoding/gob/type.go:668
		_go_fuzz_dep_.CoverTab[85375]++
//line /usr/local/go/src/encoding/gob/type.go:668
		// _ = "end of CoverTab[85375]"
	}
//line /usr/local/go/src/encoding/gob/type.go:669
	// _ = "end of CoverTab[85364]"
//line /usr/local/go/src/encoding/gob/type.go:669
	_go_fuzz_dep_.CoverTab[85365]++
							return unknown
//line /usr/local/go/src/encoding/gob/type.go:670
	// _ = "end of CoverTab[85365]"
}

type typeInfo struct {
	id	typeId
	encInit	sync.Mutex	// protects creation of encoder
	encoder	atomic.Pointer[encEngine]
	wire	*wireType
}

// typeInfoMap is an atomic pointer to map[reflect.Type]*typeInfo.
//line /usr/local/go/src/encoding/gob/type.go:680
// It's updated copy-on-write. Readers just do an atomic load
//line /usr/local/go/src/encoding/gob/type.go:680
// to get the current version of the map. Writers make a full copy of
//line /usr/local/go/src/encoding/gob/type.go:680
// the map and atomically update the pointer to point to the new map.
//line /usr/local/go/src/encoding/gob/type.go:680
// Under heavy read contention, this is significantly faster than a map
//line /usr/local/go/src/encoding/gob/type.go:680
// protected by a mutex.
//line /usr/local/go/src/encoding/gob/type.go:686
var typeInfoMap atomic.Value

func lookupTypeInfo(rt reflect.Type) *typeInfo {
//line /usr/local/go/src/encoding/gob/type.go:688
	_go_fuzz_dep_.CoverTab[85376]++
							m, _ := typeInfoMap.Load().(map[reflect.Type]*typeInfo)
							return m[rt]
//line /usr/local/go/src/encoding/gob/type.go:690
	// _ = "end of CoverTab[85376]"
}

func getTypeInfo(ut *userTypeInfo) (*typeInfo, error) {
//line /usr/local/go/src/encoding/gob/type.go:693
	_go_fuzz_dep_.CoverTab[85377]++
							rt := ut.base
							if ut.externalEnc != 0 {
//line /usr/local/go/src/encoding/gob/type.go:695
		_go_fuzz_dep_.CoverTab[85380]++

								rt = ut.user
//line /usr/local/go/src/encoding/gob/type.go:697
		// _ = "end of CoverTab[85380]"
	} else {
//line /usr/local/go/src/encoding/gob/type.go:698
		_go_fuzz_dep_.CoverTab[85381]++
//line /usr/local/go/src/encoding/gob/type.go:698
		// _ = "end of CoverTab[85381]"
//line /usr/local/go/src/encoding/gob/type.go:698
	}
//line /usr/local/go/src/encoding/gob/type.go:698
	// _ = "end of CoverTab[85377]"
//line /usr/local/go/src/encoding/gob/type.go:698
	_go_fuzz_dep_.CoverTab[85378]++
							if info := lookupTypeInfo(rt); info != nil {
//line /usr/local/go/src/encoding/gob/type.go:699
		_go_fuzz_dep_.CoverTab[85382]++
								return info, nil
//line /usr/local/go/src/encoding/gob/type.go:700
		// _ = "end of CoverTab[85382]"
	} else {
//line /usr/local/go/src/encoding/gob/type.go:701
		_go_fuzz_dep_.CoverTab[85383]++
//line /usr/local/go/src/encoding/gob/type.go:701
		// _ = "end of CoverTab[85383]"
//line /usr/local/go/src/encoding/gob/type.go:701
	}
//line /usr/local/go/src/encoding/gob/type.go:701
	// _ = "end of CoverTab[85378]"
//line /usr/local/go/src/encoding/gob/type.go:701
	_go_fuzz_dep_.CoverTab[85379]++
							return buildTypeInfo(ut, rt)
//line /usr/local/go/src/encoding/gob/type.go:702
	// _ = "end of CoverTab[85379]"
}

// buildTypeInfo constructs the type information for the type
//line /usr/local/go/src/encoding/gob/type.go:705
// and stores it in the type info map.
//line /usr/local/go/src/encoding/gob/type.go:707
func buildTypeInfo(ut *userTypeInfo, rt reflect.Type) (*typeInfo, error) {
//line /usr/local/go/src/encoding/gob/type.go:707
	_go_fuzz_dep_.CoverTab[85384]++
							typeLock.Lock()
							defer typeLock.Unlock()

							if info := lookupTypeInfo(rt); info != nil {
//line /usr/local/go/src/encoding/gob/type.go:711
		_go_fuzz_dep_.CoverTab[85389]++
								return info, nil
//line /usr/local/go/src/encoding/gob/type.go:712
		// _ = "end of CoverTab[85389]"
	} else {
//line /usr/local/go/src/encoding/gob/type.go:713
		_go_fuzz_dep_.CoverTab[85390]++
//line /usr/local/go/src/encoding/gob/type.go:713
		// _ = "end of CoverTab[85390]"
//line /usr/local/go/src/encoding/gob/type.go:713
	}
//line /usr/local/go/src/encoding/gob/type.go:713
	// _ = "end of CoverTab[85384]"
//line /usr/local/go/src/encoding/gob/type.go:713
	_go_fuzz_dep_.CoverTab[85385]++

							gt, err := getBaseType(rt.Name(), rt)
							if err != nil {
//line /usr/local/go/src/encoding/gob/type.go:716
		_go_fuzz_dep_.CoverTab[85391]++
								return nil, err
//line /usr/local/go/src/encoding/gob/type.go:717
		// _ = "end of CoverTab[85391]"
	} else {
//line /usr/local/go/src/encoding/gob/type.go:718
		_go_fuzz_dep_.CoverTab[85392]++
//line /usr/local/go/src/encoding/gob/type.go:718
		// _ = "end of CoverTab[85392]"
//line /usr/local/go/src/encoding/gob/type.go:718
	}
//line /usr/local/go/src/encoding/gob/type.go:718
	// _ = "end of CoverTab[85385]"
//line /usr/local/go/src/encoding/gob/type.go:718
	_go_fuzz_dep_.CoverTab[85386]++
							info := &typeInfo{id: gt.id()}

							if ut.externalEnc != 0 {
//line /usr/local/go/src/encoding/gob/type.go:721
		_go_fuzz_dep_.CoverTab[85393]++
								userType, err := getType(rt.Name(), ut, rt)
								if err != nil {
//line /usr/local/go/src/encoding/gob/type.go:723
			_go_fuzz_dep_.CoverTab[85396]++
									return nil, err
//line /usr/local/go/src/encoding/gob/type.go:724
			// _ = "end of CoverTab[85396]"
		} else {
//line /usr/local/go/src/encoding/gob/type.go:725
			_go_fuzz_dep_.CoverTab[85397]++
//line /usr/local/go/src/encoding/gob/type.go:725
			// _ = "end of CoverTab[85397]"
//line /usr/local/go/src/encoding/gob/type.go:725
		}
//line /usr/local/go/src/encoding/gob/type.go:725
		// _ = "end of CoverTab[85393]"
//line /usr/local/go/src/encoding/gob/type.go:725
		_go_fuzz_dep_.CoverTab[85394]++
								gt := userType.id().gobType().(*gobEncoderType)
								switch ut.externalEnc {
		case xGob:
//line /usr/local/go/src/encoding/gob/type.go:728
			_go_fuzz_dep_.CoverTab[85398]++
									info.wire = &wireType{GobEncoderT: gt}
//line /usr/local/go/src/encoding/gob/type.go:729
			// _ = "end of CoverTab[85398]"
		case xBinary:
//line /usr/local/go/src/encoding/gob/type.go:730
			_go_fuzz_dep_.CoverTab[85399]++
									info.wire = &wireType{BinaryMarshalerT: gt}
//line /usr/local/go/src/encoding/gob/type.go:731
			// _ = "end of CoverTab[85399]"
		case xText:
//line /usr/local/go/src/encoding/gob/type.go:732
			_go_fuzz_dep_.CoverTab[85400]++
									info.wire = &wireType{TextMarshalerT: gt}
//line /usr/local/go/src/encoding/gob/type.go:733
			// _ = "end of CoverTab[85400]"
//line /usr/local/go/src/encoding/gob/type.go:733
		default:
//line /usr/local/go/src/encoding/gob/type.go:733
			_go_fuzz_dep_.CoverTab[85401]++
//line /usr/local/go/src/encoding/gob/type.go:733
			// _ = "end of CoverTab[85401]"
		}
//line /usr/local/go/src/encoding/gob/type.go:734
		// _ = "end of CoverTab[85394]"
//line /usr/local/go/src/encoding/gob/type.go:734
		_go_fuzz_dep_.CoverTab[85395]++
								rt = ut.user
//line /usr/local/go/src/encoding/gob/type.go:735
		// _ = "end of CoverTab[85395]"
	} else {
//line /usr/local/go/src/encoding/gob/type.go:736
		_go_fuzz_dep_.CoverTab[85402]++
								t := info.id.gobType()
								switch typ := rt; typ.Kind() {
		case reflect.Array:
//line /usr/local/go/src/encoding/gob/type.go:739
			_go_fuzz_dep_.CoverTab[85403]++
									info.wire = &wireType{ArrayT: t.(*arrayType)}
//line /usr/local/go/src/encoding/gob/type.go:740
			// _ = "end of CoverTab[85403]"
		case reflect.Map:
//line /usr/local/go/src/encoding/gob/type.go:741
			_go_fuzz_dep_.CoverTab[85404]++
									info.wire = &wireType{MapT: t.(*mapType)}
//line /usr/local/go/src/encoding/gob/type.go:742
			// _ = "end of CoverTab[85404]"
		case reflect.Slice:
//line /usr/local/go/src/encoding/gob/type.go:743
			_go_fuzz_dep_.CoverTab[85405]++

									if typ.Elem().Kind() != reflect.Uint8 {
//line /usr/local/go/src/encoding/gob/type.go:745
				_go_fuzz_dep_.CoverTab[85408]++
										info.wire = &wireType{SliceT: t.(*sliceType)}
//line /usr/local/go/src/encoding/gob/type.go:746
				// _ = "end of CoverTab[85408]"
			} else {
//line /usr/local/go/src/encoding/gob/type.go:747
				_go_fuzz_dep_.CoverTab[85409]++
//line /usr/local/go/src/encoding/gob/type.go:747
				// _ = "end of CoverTab[85409]"
//line /usr/local/go/src/encoding/gob/type.go:747
			}
//line /usr/local/go/src/encoding/gob/type.go:747
			// _ = "end of CoverTab[85405]"
		case reflect.Struct:
//line /usr/local/go/src/encoding/gob/type.go:748
			_go_fuzz_dep_.CoverTab[85406]++
									info.wire = &wireType{StructT: t.(*structType)}
//line /usr/local/go/src/encoding/gob/type.go:749
			// _ = "end of CoverTab[85406]"
//line /usr/local/go/src/encoding/gob/type.go:749
		default:
//line /usr/local/go/src/encoding/gob/type.go:749
			_go_fuzz_dep_.CoverTab[85407]++
//line /usr/local/go/src/encoding/gob/type.go:749
			// _ = "end of CoverTab[85407]"
		}
//line /usr/local/go/src/encoding/gob/type.go:750
		// _ = "end of CoverTab[85402]"
	}
//line /usr/local/go/src/encoding/gob/type.go:751
	// _ = "end of CoverTab[85386]"
//line /usr/local/go/src/encoding/gob/type.go:751
	_go_fuzz_dep_.CoverTab[85387]++

//line /usr/local/go/src/encoding/gob/type.go:754
	newm := make(map[reflect.Type]*typeInfo)
	m, _ := typeInfoMap.Load().(map[reflect.Type]*typeInfo)
	for k, v := range m {
//line /usr/local/go/src/encoding/gob/type.go:756
		_go_fuzz_dep_.CoverTab[85410]++
								newm[k] = v
//line /usr/local/go/src/encoding/gob/type.go:757
		// _ = "end of CoverTab[85410]"
	}
//line /usr/local/go/src/encoding/gob/type.go:758
	// _ = "end of CoverTab[85387]"
//line /usr/local/go/src/encoding/gob/type.go:758
	_go_fuzz_dep_.CoverTab[85388]++
							newm[rt] = info
							typeInfoMap.Store(newm)
							return info, nil
//line /usr/local/go/src/encoding/gob/type.go:761
	// _ = "end of CoverTab[85388]"
}

// Called only when a panic is acceptable and unexpected.
func mustGetTypeInfo(rt reflect.Type) *typeInfo {
//line /usr/local/go/src/encoding/gob/type.go:765
	_go_fuzz_dep_.CoverTab[85411]++
							t, err := getTypeInfo(userType(rt))
							if err != nil {
//line /usr/local/go/src/encoding/gob/type.go:767
		_go_fuzz_dep_.CoverTab[85413]++
								panic("getTypeInfo: " + err.Error())
//line /usr/local/go/src/encoding/gob/type.go:768
		// _ = "end of CoverTab[85413]"
	} else {
//line /usr/local/go/src/encoding/gob/type.go:769
		_go_fuzz_dep_.CoverTab[85414]++
//line /usr/local/go/src/encoding/gob/type.go:769
		// _ = "end of CoverTab[85414]"
//line /usr/local/go/src/encoding/gob/type.go:769
	}
//line /usr/local/go/src/encoding/gob/type.go:769
	// _ = "end of CoverTab[85411]"
//line /usr/local/go/src/encoding/gob/type.go:769
	_go_fuzz_dep_.CoverTab[85412]++
							return t
//line /usr/local/go/src/encoding/gob/type.go:770
	// _ = "end of CoverTab[85412]"
}

// GobEncoder is the interface describing data that provides its own
//line /usr/local/go/src/encoding/gob/type.go:773
// representation for encoding values for transmission to a GobDecoder.
//line /usr/local/go/src/encoding/gob/type.go:773
// A type that implements GobEncoder and GobDecoder has complete
//line /usr/local/go/src/encoding/gob/type.go:773
// control over the representation of its data and may therefore
//line /usr/local/go/src/encoding/gob/type.go:773
// contain things such as private fields, channels, and functions,
//line /usr/local/go/src/encoding/gob/type.go:773
// which are not usually transmissible in gob streams.
//line /usr/local/go/src/encoding/gob/type.go:773
//
//line /usr/local/go/src/encoding/gob/type.go:773
// Note: Since gobs can be stored permanently, it is good design
//line /usr/local/go/src/encoding/gob/type.go:773
// to guarantee the encoding used by a GobEncoder is stable as the
//line /usr/local/go/src/encoding/gob/type.go:773
// software evolves. For instance, it might make sense for GobEncode
//line /usr/local/go/src/encoding/gob/type.go:773
// to include a version number in the encoding.
//line /usr/local/go/src/encoding/gob/type.go:784
type GobEncoder interface {
	// GobEncode returns a byte slice representing the encoding of the
	// receiver for transmission to a GobDecoder, usually of the same
	// concrete type.
	GobEncode() ([]byte, error)
}

// GobDecoder is the interface describing data that provides its own
//line /usr/local/go/src/encoding/gob/type.go:791
// routine for decoding transmitted values sent by a GobEncoder.
//line /usr/local/go/src/encoding/gob/type.go:793
type GobDecoder interface {
	// GobDecode overwrites the receiver, which must be a pointer,
	// with the value represented by the byte slice, which was written
	// by GobEncode, usually for the same concrete type.
	GobDecode([]byte) error
}

var (
	nameToConcreteType	sync.Map	// map[string]reflect.Type
	concreteTypeToName	sync.Map	// map[reflect.Type]string
)

// RegisterName is like Register but uses the provided name rather than the
//line /usr/local/go/src/encoding/gob/type.go:805
// type's default.
//line /usr/local/go/src/encoding/gob/type.go:807
func RegisterName(name string, value any) {
//line /usr/local/go/src/encoding/gob/type.go:807
	_go_fuzz_dep_.CoverTab[85415]++
							if name == "" {
//line /usr/local/go/src/encoding/gob/type.go:808
		_go_fuzz_dep_.CoverTab[85418]++

								panic("attempt to register empty name")
//line /usr/local/go/src/encoding/gob/type.go:810
		// _ = "end of CoverTab[85418]"
	} else {
//line /usr/local/go/src/encoding/gob/type.go:811
		_go_fuzz_dep_.CoverTab[85419]++
//line /usr/local/go/src/encoding/gob/type.go:811
		// _ = "end of CoverTab[85419]"
//line /usr/local/go/src/encoding/gob/type.go:811
	}
//line /usr/local/go/src/encoding/gob/type.go:811
	// _ = "end of CoverTab[85415]"
//line /usr/local/go/src/encoding/gob/type.go:811
	_go_fuzz_dep_.CoverTab[85416]++

							ut := userType(reflect.TypeOf(value))

//line /usr/local/go/src/encoding/gob/type.go:819
	if t, dup := nameToConcreteType.LoadOrStore(name, reflect.TypeOf(value)); dup && func() bool {
//line /usr/local/go/src/encoding/gob/type.go:819
		_go_fuzz_dep_.CoverTab[85420]++
//line /usr/local/go/src/encoding/gob/type.go:819
		return t != ut.user
//line /usr/local/go/src/encoding/gob/type.go:819
		// _ = "end of CoverTab[85420]"
//line /usr/local/go/src/encoding/gob/type.go:819
	}() {
//line /usr/local/go/src/encoding/gob/type.go:819
		_go_fuzz_dep_.CoverTab[85421]++
								panic(fmt.Sprintf("gob: registering duplicate types for %q: %s != %s", name, t, ut.user))
//line /usr/local/go/src/encoding/gob/type.go:820
		// _ = "end of CoverTab[85421]"
	} else {
//line /usr/local/go/src/encoding/gob/type.go:821
		_go_fuzz_dep_.CoverTab[85422]++
//line /usr/local/go/src/encoding/gob/type.go:821
		// _ = "end of CoverTab[85422]"
//line /usr/local/go/src/encoding/gob/type.go:821
	}
//line /usr/local/go/src/encoding/gob/type.go:821
	// _ = "end of CoverTab[85416]"
//line /usr/local/go/src/encoding/gob/type.go:821
	_go_fuzz_dep_.CoverTab[85417]++

//line /usr/local/go/src/encoding/gob/type.go:824
	if n, dup := concreteTypeToName.LoadOrStore(ut.base, name); dup && func() bool {
//line /usr/local/go/src/encoding/gob/type.go:824
		_go_fuzz_dep_.CoverTab[85423]++
//line /usr/local/go/src/encoding/gob/type.go:824
		return n != name
//line /usr/local/go/src/encoding/gob/type.go:824
		// _ = "end of CoverTab[85423]"
//line /usr/local/go/src/encoding/gob/type.go:824
	}() {
//line /usr/local/go/src/encoding/gob/type.go:824
		_go_fuzz_dep_.CoverTab[85424]++
								nameToConcreteType.Delete(name)
								panic(fmt.Sprintf("gob: registering duplicate names for %s: %q != %q", ut.user, n, name))
//line /usr/local/go/src/encoding/gob/type.go:826
		// _ = "end of CoverTab[85424]"
	} else {
//line /usr/local/go/src/encoding/gob/type.go:827
		_go_fuzz_dep_.CoverTab[85425]++
//line /usr/local/go/src/encoding/gob/type.go:827
		// _ = "end of CoverTab[85425]"
//line /usr/local/go/src/encoding/gob/type.go:827
	}
//line /usr/local/go/src/encoding/gob/type.go:827
	// _ = "end of CoverTab[85417]"
}

// Register records a type, identified by a value for that type, under its
//line /usr/local/go/src/encoding/gob/type.go:830
// internal type name. That name will identify the concrete type of a value
//line /usr/local/go/src/encoding/gob/type.go:830
// sent or received as an interface variable. Only types that will be
//line /usr/local/go/src/encoding/gob/type.go:830
// transferred as implementations of interface values need to be registered.
//line /usr/local/go/src/encoding/gob/type.go:830
// Expecting to be used only during initialization, it panics if the mapping
//line /usr/local/go/src/encoding/gob/type.go:830
// between types and names is not a bijection.
//line /usr/local/go/src/encoding/gob/type.go:836
func Register(value any) {
//line /usr/local/go/src/encoding/gob/type.go:836
	_go_fuzz_dep_.CoverTab[85426]++

							rt := reflect.TypeOf(value)
							name := rt.String()

//line /usr/local/go/src/encoding/gob/type.go:843
	star := ""
	if rt.Name() == "" {
//line /usr/local/go/src/encoding/gob/type.go:844
		_go_fuzz_dep_.CoverTab[85429]++
								if pt := rt; pt.Kind() == reflect.Pointer {
//line /usr/local/go/src/encoding/gob/type.go:845
			_go_fuzz_dep_.CoverTab[85430]++
									star = "*"

//line /usr/local/go/src/encoding/gob/type.go:863
			rt = pt
//line /usr/local/go/src/encoding/gob/type.go:863
			// _ = "end of CoverTab[85430]"
		} else {
//line /usr/local/go/src/encoding/gob/type.go:864
			_go_fuzz_dep_.CoverTab[85431]++
//line /usr/local/go/src/encoding/gob/type.go:864
			// _ = "end of CoverTab[85431]"
//line /usr/local/go/src/encoding/gob/type.go:864
		}
//line /usr/local/go/src/encoding/gob/type.go:864
		// _ = "end of CoverTab[85429]"
	} else {
//line /usr/local/go/src/encoding/gob/type.go:865
		_go_fuzz_dep_.CoverTab[85432]++
//line /usr/local/go/src/encoding/gob/type.go:865
		// _ = "end of CoverTab[85432]"
//line /usr/local/go/src/encoding/gob/type.go:865
	}
//line /usr/local/go/src/encoding/gob/type.go:865
	// _ = "end of CoverTab[85426]"
//line /usr/local/go/src/encoding/gob/type.go:865
	_go_fuzz_dep_.CoverTab[85427]++
							if rt.Name() != "" {
//line /usr/local/go/src/encoding/gob/type.go:866
		_go_fuzz_dep_.CoverTab[85433]++
								if rt.PkgPath() == "" {
//line /usr/local/go/src/encoding/gob/type.go:867
			_go_fuzz_dep_.CoverTab[85434]++
									name = star + rt.Name()
//line /usr/local/go/src/encoding/gob/type.go:868
			// _ = "end of CoverTab[85434]"
		} else {
//line /usr/local/go/src/encoding/gob/type.go:869
			_go_fuzz_dep_.CoverTab[85435]++
									name = star + rt.PkgPath() + "." + rt.Name()
//line /usr/local/go/src/encoding/gob/type.go:870
			// _ = "end of CoverTab[85435]"
		}
//line /usr/local/go/src/encoding/gob/type.go:871
		// _ = "end of CoverTab[85433]"
	} else {
//line /usr/local/go/src/encoding/gob/type.go:872
		_go_fuzz_dep_.CoverTab[85436]++
//line /usr/local/go/src/encoding/gob/type.go:872
		// _ = "end of CoverTab[85436]"
//line /usr/local/go/src/encoding/gob/type.go:872
	}
//line /usr/local/go/src/encoding/gob/type.go:872
	// _ = "end of CoverTab[85427]"
//line /usr/local/go/src/encoding/gob/type.go:872
	_go_fuzz_dep_.CoverTab[85428]++

							RegisterName(name, value)
//line /usr/local/go/src/encoding/gob/type.go:874
	// _ = "end of CoverTab[85428]"
}

func registerBasics() {
//line /usr/local/go/src/encoding/gob/type.go:877
	_go_fuzz_dep_.CoverTab[85437]++
							Register(int(0))
							Register(int8(0))
							Register(int16(0))
							Register(int32(0))
							Register(int64(0))
							Register(uint(0))
							Register(uint8(0))
							Register(uint16(0))
							Register(uint32(0))
							Register(uint64(0))
							Register(float32(0))
							Register(float64(0))
							Register(complex64(0i))
							Register(complex128(0i))
							Register(uintptr(0))
							Register(false)
							Register("")
							Register([]byte(nil))
							Register([]int(nil))
							Register([]int8(nil))
							Register([]int16(nil))
							Register([]int32(nil))
							Register([]int64(nil))
							Register([]uint(nil))
							Register([]uint8(nil))
							Register([]uint16(nil))
							Register([]uint32(nil))
							Register([]uint64(nil))
							Register([]float32(nil))
							Register([]float64(nil))
							Register([]complex64(nil))
							Register([]complex128(nil))
							Register([]uintptr(nil))
							Register([]bool(nil))
							Register([]string(nil))
//line /usr/local/go/src/encoding/gob/type.go:912
	// _ = "end of CoverTab[85437]"
}

//line /usr/local/go/src/encoding/gob/type.go:913
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/encoding/gob/type.go:913
var _ = _go_fuzz_dep_.CoverTab
