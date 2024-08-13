// Protocol Buffers for Go with Gadgets
//
// Copyright (c) 2018, The GoGo Authors. All rights reserved.
// http://github.com/gogo/protobuf
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:29
package types

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:29
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:29
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:29
)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:29
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:29
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:29
)

func NewPopulatedStdDouble(r randyWrappers, easy bool) *float64 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:31
	_go_fuzz_dep_.CoverTab[141789]++
												v := NewPopulatedDoubleValue(r, easy)
												return &v.Value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:33
	// _ = "end of CoverTab[141789]"
}

func SizeOfStdDouble(v float64) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:36
	_go_fuzz_dep_.CoverTab[141790]++
												pv := &DoubleValue{Value: v}
												return pv.Size()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:38
	// _ = "end of CoverTab[141790]"
}

func StdDoubleMarshal(v float64) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:41
	_go_fuzz_dep_.CoverTab[141791]++
												size := SizeOfStdDouble(v)
												buf := make([]byte, size)
												_, err := StdDoubleMarshalTo(v, buf)
												return buf, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:45
	// _ = "end of CoverTab[141791]"
}

func StdDoubleMarshalTo(v float64, data []byte) (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:48
	_go_fuzz_dep_.CoverTab[141792]++
												pv := &DoubleValue{Value: v}
												return pv.MarshalTo(data)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:50
	// _ = "end of CoverTab[141792]"
}

func StdDoubleUnmarshal(v *float64, data []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:53
	_go_fuzz_dep_.CoverTab[141793]++
												pv := &DoubleValue{}
												if err := pv.Unmarshal(data); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:55
		_go_fuzz_dep_.CoverTab[141795]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:56
		// _ = "end of CoverTab[141795]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:57
		_go_fuzz_dep_.CoverTab[141796]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:57
		// _ = "end of CoverTab[141796]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:57
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:57
	// _ = "end of CoverTab[141793]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:57
	_go_fuzz_dep_.CoverTab[141794]++
												*v = pv.Value
												return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:59
	// _ = "end of CoverTab[141794]"
}
func NewPopulatedStdFloat(r randyWrappers, easy bool) *float32 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:61
	_go_fuzz_dep_.CoverTab[141797]++
												v := NewPopulatedFloatValue(r, easy)
												return &v.Value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:63
	// _ = "end of CoverTab[141797]"
}

func SizeOfStdFloat(v float32) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:66
	_go_fuzz_dep_.CoverTab[141798]++
												pv := &FloatValue{Value: v}
												return pv.Size()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:68
	// _ = "end of CoverTab[141798]"
}

func StdFloatMarshal(v float32) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:71
	_go_fuzz_dep_.CoverTab[141799]++
												size := SizeOfStdFloat(v)
												buf := make([]byte, size)
												_, err := StdFloatMarshalTo(v, buf)
												return buf, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:75
	// _ = "end of CoverTab[141799]"
}

func StdFloatMarshalTo(v float32, data []byte) (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:78
	_go_fuzz_dep_.CoverTab[141800]++
												pv := &FloatValue{Value: v}
												return pv.MarshalTo(data)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:80
	// _ = "end of CoverTab[141800]"
}

func StdFloatUnmarshal(v *float32, data []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:83
	_go_fuzz_dep_.CoverTab[141801]++
												pv := &FloatValue{}
												if err := pv.Unmarshal(data); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:85
		_go_fuzz_dep_.CoverTab[141803]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:86
		// _ = "end of CoverTab[141803]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:87
		_go_fuzz_dep_.CoverTab[141804]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:87
		// _ = "end of CoverTab[141804]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:87
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:87
	// _ = "end of CoverTab[141801]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:87
	_go_fuzz_dep_.CoverTab[141802]++
												*v = pv.Value
												return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:89
	// _ = "end of CoverTab[141802]"
}
func NewPopulatedStdInt64(r randyWrappers, easy bool) *int64 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:91
	_go_fuzz_dep_.CoverTab[141805]++
												v := NewPopulatedInt64Value(r, easy)
												return &v.Value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:93
	// _ = "end of CoverTab[141805]"
}

func SizeOfStdInt64(v int64) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:96
	_go_fuzz_dep_.CoverTab[141806]++
												pv := &Int64Value{Value: v}
												return pv.Size()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:98
	// _ = "end of CoverTab[141806]"
}

func StdInt64Marshal(v int64) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:101
	_go_fuzz_dep_.CoverTab[141807]++
												size := SizeOfStdInt64(v)
												buf := make([]byte, size)
												_, err := StdInt64MarshalTo(v, buf)
												return buf, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:105
	// _ = "end of CoverTab[141807]"
}

func StdInt64MarshalTo(v int64, data []byte) (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:108
	_go_fuzz_dep_.CoverTab[141808]++
												pv := &Int64Value{Value: v}
												return pv.MarshalTo(data)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:110
	// _ = "end of CoverTab[141808]"
}

func StdInt64Unmarshal(v *int64, data []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:113
	_go_fuzz_dep_.CoverTab[141809]++
												pv := &Int64Value{}
												if err := pv.Unmarshal(data); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:115
		_go_fuzz_dep_.CoverTab[141811]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:116
		// _ = "end of CoverTab[141811]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:117
		_go_fuzz_dep_.CoverTab[141812]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:117
		// _ = "end of CoverTab[141812]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:117
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:117
	// _ = "end of CoverTab[141809]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:117
	_go_fuzz_dep_.CoverTab[141810]++
												*v = pv.Value
												return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:119
	// _ = "end of CoverTab[141810]"
}
func NewPopulatedStdUInt64(r randyWrappers, easy bool) *uint64 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:121
	_go_fuzz_dep_.CoverTab[141813]++
												v := NewPopulatedUInt64Value(r, easy)
												return &v.Value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:123
	// _ = "end of CoverTab[141813]"
}

func SizeOfStdUInt64(v uint64) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:126
	_go_fuzz_dep_.CoverTab[141814]++
												pv := &UInt64Value{Value: v}
												return pv.Size()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:128
	// _ = "end of CoverTab[141814]"
}

func StdUInt64Marshal(v uint64) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:131
	_go_fuzz_dep_.CoverTab[141815]++
												size := SizeOfStdUInt64(v)
												buf := make([]byte, size)
												_, err := StdUInt64MarshalTo(v, buf)
												return buf, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:135
	// _ = "end of CoverTab[141815]"
}

func StdUInt64MarshalTo(v uint64, data []byte) (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:138
	_go_fuzz_dep_.CoverTab[141816]++
												pv := &UInt64Value{Value: v}
												return pv.MarshalTo(data)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:140
	// _ = "end of CoverTab[141816]"
}

func StdUInt64Unmarshal(v *uint64, data []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:143
	_go_fuzz_dep_.CoverTab[141817]++
												pv := &UInt64Value{}
												if err := pv.Unmarshal(data); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:145
		_go_fuzz_dep_.CoverTab[141819]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:146
		// _ = "end of CoverTab[141819]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:147
		_go_fuzz_dep_.CoverTab[141820]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:147
		// _ = "end of CoverTab[141820]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:147
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:147
	// _ = "end of CoverTab[141817]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:147
	_go_fuzz_dep_.CoverTab[141818]++
												*v = pv.Value
												return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:149
	// _ = "end of CoverTab[141818]"
}
func NewPopulatedStdInt32(r randyWrappers, easy bool) *int32 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:151
	_go_fuzz_dep_.CoverTab[141821]++
												v := NewPopulatedInt32Value(r, easy)
												return &v.Value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:153
	// _ = "end of CoverTab[141821]"
}

func SizeOfStdInt32(v int32) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:156
	_go_fuzz_dep_.CoverTab[141822]++
												pv := &Int32Value{Value: v}
												return pv.Size()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:158
	// _ = "end of CoverTab[141822]"
}

func StdInt32Marshal(v int32) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:161
	_go_fuzz_dep_.CoverTab[141823]++
												size := SizeOfStdInt32(v)
												buf := make([]byte, size)
												_, err := StdInt32MarshalTo(v, buf)
												return buf, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:165
	// _ = "end of CoverTab[141823]"
}

func StdInt32MarshalTo(v int32, data []byte) (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:168
	_go_fuzz_dep_.CoverTab[141824]++
												pv := &Int32Value{Value: v}
												return pv.MarshalTo(data)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:170
	// _ = "end of CoverTab[141824]"
}

func StdInt32Unmarshal(v *int32, data []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:173
	_go_fuzz_dep_.CoverTab[141825]++
												pv := &Int32Value{}
												if err := pv.Unmarshal(data); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:175
		_go_fuzz_dep_.CoverTab[141827]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:176
		// _ = "end of CoverTab[141827]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:177
		_go_fuzz_dep_.CoverTab[141828]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:177
		// _ = "end of CoverTab[141828]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:177
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:177
	// _ = "end of CoverTab[141825]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:177
	_go_fuzz_dep_.CoverTab[141826]++
												*v = pv.Value
												return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:179
	// _ = "end of CoverTab[141826]"
}
func NewPopulatedStdUInt32(r randyWrappers, easy bool) *uint32 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:181
	_go_fuzz_dep_.CoverTab[141829]++
												v := NewPopulatedUInt32Value(r, easy)
												return &v.Value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:183
	// _ = "end of CoverTab[141829]"
}

func SizeOfStdUInt32(v uint32) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:186
	_go_fuzz_dep_.CoverTab[141830]++
												pv := &UInt32Value{Value: v}
												return pv.Size()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:188
	// _ = "end of CoverTab[141830]"
}

func StdUInt32Marshal(v uint32) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:191
	_go_fuzz_dep_.CoverTab[141831]++
												size := SizeOfStdUInt32(v)
												buf := make([]byte, size)
												_, err := StdUInt32MarshalTo(v, buf)
												return buf, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:195
	// _ = "end of CoverTab[141831]"
}

func StdUInt32MarshalTo(v uint32, data []byte) (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:198
	_go_fuzz_dep_.CoverTab[141832]++
												pv := &UInt32Value{Value: v}
												return pv.MarshalTo(data)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:200
	// _ = "end of CoverTab[141832]"
}

func StdUInt32Unmarshal(v *uint32, data []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:203
	_go_fuzz_dep_.CoverTab[141833]++
												pv := &UInt32Value{}
												if err := pv.Unmarshal(data); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:205
		_go_fuzz_dep_.CoverTab[141835]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:206
		// _ = "end of CoverTab[141835]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:207
		_go_fuzz_dep_.CoverTab[141836]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:207
		// _ = "end of CoverTab[141836]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:207
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:207
	// _ = "end of CoverTab[141833]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:207
	_go_fuzz_dep_.CoverTab[141834]++
												*v = pv.Value
												return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:209
	// _ = "end of CoverTab[141834]"
}
func NewPopulatedStdBool(r randyWrappers, easy bool) *bool {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:211
	_go_fuzz_dep_.CoverTab[141837]++
												v := NewPopulatedBoolValue(r, easy)
												return &v.Value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:213
	// _ = "end of CoverTab[141837]"
}

func SizeOfStdBool(v bool) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:216
	_go_fuzz_dep_.CoverTab[141838]++
												pv := &BoolValue{Value: v}
												return pv.Size()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:218
	// _ = "end of CoverTab[141838]"
}

func StdBoolMarshal(v bool) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:221
	_go_fuzz_dep_.CoverTab[141839]++
												size := SizeOfStdBool(v)
												buf := make([]byte, size)
												_, err := StdBoolMarshalTo(v, buf)
												return buf, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:225
	// _ = "end of CoverTab[141839]"
}

func StdBoolMarshalTo(v bool, data []byte) (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:228
	_go_fuzz_dep_.CoverTab[141840]++
												pv := &BoolValue{Value: v}
												return pv.MarshalTo(data)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:230
	// _ = "end of CoverTab[141840]"
}

func StdBoolUnmarshal(v *bool, data []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:233
	_go_fuzz_dep_.CoverTab[141841]++
												pv := &BoolValue{}
												if err := pv.Unmarshal(data); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:235
		_go_fuzz_dep_.CoverTab[141843]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:236
		// _ = "end of CoverTab[141843]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:237
		_go_fuzz_dep_.CoverTab[141844]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:237
		// _ = "end of CoverTab[141844]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:237
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:237
	// _ = "end of CoverTab[141841]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:237
	_go_fuzz_dep_.CoverTab[141842]++
												*v = pv.Value
												return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:239
	// _ = "end of CoverTab[141842]"
}
func NewPopulatedStdString(r randyWrappers, easy bool) *string {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:241
	_go_fuzz_dep_.CoverTab[141845]++
												v := NewPopulatedStringValue(r, easy)
												return &v.Value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:243
	// _ = "end of CoverTab[141845]"
}

func SizeOfStdString(v string) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:246
	_go_fuzz_dep_.CoverTab[141846]++
												pv := &StringValue{Value: v}
												return pv.Size()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:248
	// _ = "end of CoverTab[141846]"
}

func StdStringMarshal(v string) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:251
	_go_fuzz_dep_.CoverTab[141847]++
												size := SizeOfStdString(v)
												buf := make([]byte, size)
												_, err := StdStringMarshalTo(v, buf)
												return buf, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:255
	// _ = "end of CoverTab[141847]"
}

func StdStringMarshalTo(v string, data []byte) (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:258
	_go_fuzz_dep_.CoverTab[141848]++
												pv := &StringValue{Value: v}
												return pv.MarshalTo(data)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:260
	// _ = "end of CoverTab[141848]"
}

func StdStringUnmarshal(v *string, data []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:263
	_go_fuzz_dep_.CoverTab[141849]++
												pv := &StringValue{}
												if err := pv.Unmarshal(data); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:265
		_go_fuzz_dep_.CoverTab[141851]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:266
		// _ = "end of CoverTab[141851]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:267
		_go_fuzz_dep_.CoverTab[141852]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:267
		// _ = "end of CoverTab[141852]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:267
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:267
	// _ = "end of CoverTab[141849]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:267
	_go_fuzz_dep_.CoverTab[141850]++
												*v = pv.Value
												return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:269
	// _ = "end of CoverTab[141850]"
}
func NewPopulatedStdBytes(r randyWrappers, easy bool) *[]byte {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:271
	_go_fuzz_dep_.CoverTab[141853]++
												v := NewPopulatedBytesValue(r, easy)
												return &v.Value
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:273
	// _ = "end of CoverTab[141853]"
}

func SizeOfStdBytes(v []byte) int {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:276
	_go_fuzz_dep_.CoverTab[141854]++
												pv := &BytesValue{Value: v}
												return pv.Size()
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:278
	// _ = "end of CoverTab[141854]"
}

func StdBytesMarshal(v []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:281
	_go_fuzz_dep_.CoverTab[141855]++
												size := SizeOfStdBytes(v)
												buf := make([]byte, size)
												_, err := StdBytesMarshalTo(v, buf)
												return buf, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:285
	// _ = "end of CoverTab[141855]"
}

func StdBytesMarshalTo(v []byte, data []byte) (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:288
	_go_fuzz_dep_.CoverTab[141856]++
												pv := &BytesValue{Value: v}
												return pv.MarshalTo(data)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:290
	// _ = "end of CoverTab[141856]"
}

func StdBytesUnmarshal(v *[]byte, data []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:293
	_go_fuzz_dep_.CoverTab[141857]++
												pv := &BytesValue{}
												if err := pv.Unmarshal(data); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:295
		_go_fuzz_dep_.CoverTab[141859]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:296
		// _ = "end of CoverTab[141859]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:297
		_go_fuzz_dep_.CoverTab[141860]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:297
		// _ = "end of CoverTab[141860]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:297
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:297
	// _ = "end of CoverTab[141857]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:297
	_go_fuzz_dep_.CoverTab[141858]++
												*v = pv.Value
												return nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:299
	// _ = "end of CoverTab[141858]"
}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:300
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/types/wrappers_gogo.go:300
var _ = _go_fuzz_dep_.CoverTab
