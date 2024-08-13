// Copyright © 2014 Steve Francia <spf@spf13.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:6
// Package cast provides easy and safe casting in Go.
package cast

//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:7
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:7
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:7
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:7
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:7
)

import "time"

// ToBool casts an interface to a bool type.
func ToBool(i interface{}) bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:12
	_go_fuzz_dep_.CoverTab[118526]++
										v, _ := ToBoolE(i)
										return v
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:14
	// _ = "end of CoverTab[118526]"
}

// ToTime casts an interface to a time.Time type.
func ToTime(i interface{}) time.Time {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:18
	_go_fuzz_dep_.CoverTab[118527]++
										v, _ := ToTimeE(i)
										return v
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:20
	// _ = "end of CoverTab[118527]"
}

func ToTimeInDefaultLocation(i interface{}, location *time.Location) time.Time {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:23
	_go_fuzz_dep_.CoverTab[118528]++
										v, _ := ToTimeInDefaultLocationE(i, location)
										return v
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:25
	// _ = "end of CoverTab[118528]"
}

// ToDuration casts an interface to a time.Duration type.
func ToDuration(i interface{}) time.Duration {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:29
	_go_fuzz_dep_.CoverTab[118529]++
										v, _ := ToDurationE(i)
										return v
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:31
	// _ = "end of CoverTab[118529]"
}

// ToFloat64 casts an interface to a float64 type.
func ToFloat64(i interface{}) float64 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:35
	_go_fuzz_dep_.CoverTab[118530]++
										v, _ := ToFloat64E(i)
										return v
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:37
	// _ = "end of CoverTab[118530]"
}

// ToFloat32 casts an interface to a float32 type.
func ToFloat32(i interface{}) float32 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:41
	_go_fuzz_dep_.CoverTab[118531]++
										v, _ := ToFloat32E(i)
										return v
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:43
	// _ = "end of CoverTab[118531]"
}

// ToInt64 casts an interface to an int64 type.
func ToInt64(i interface{}) int64 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:47
	_go_fuzz_dep_.CoverTab[118532]++
										v, _ := ToInt64E(i)
										return v
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:49
	// _ = "end of CoverTab[118532]"
}

// ToInt32 casts an interface to an int32 type.
func ToInt32(i interface{}) int32 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:53
	_go_fuzz_dep_.CoverTab[118533]++
										v, _ := ToInt32E(i)
										return v
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:55
	// _ = "end of CoverTab[118533]"
}

// ToInt16 casts an interface to an int16 type.
func ToInt16(i interface{}) int16 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:59
	_go_fuzz_dep_.CoverTab[118534]++
										v, _ := ToInt16E(i)
										return v
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:61
	// _ = "end of CoverTab[118534]"
}

// ToInt8 casts an interface to an int8 type.
func ToInt8(i interface{}) int8 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:65
	_go_fuzz_dep_.CoverTab[118535]++
										v, _ := ToInt8E(i)
										return v
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:67
	// _ = "end of CoverTab[118535]"
}

// ToInt casts an interface to an int type.
func ToInt(i interface{}) int {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:71
	_go_fuzz_dep_.CoverTab[118536]++
										v, _ := ToIntE(i)
										return v
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:73
	// _ = "end of CoverTab[118536]"
}

// ToUint casts an interface to a uint type.
func ToUint(i interface{}) uint {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:77
	_go_fuzz_dep_.CoverTab[118537]++
										v, _ := ToUintE(i)
										return v
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:79
	// _ = "end of CoverTab[118537]"
}

// ToUint64 casts an interface to a uint64 type.
func ToUint64(i interface{}) uint64 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:83
	_go_fuzz_dep_.CoverTab[118538]++
										v, _ := ToUint64E(i)
										return v
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:85
	// _ = "end of CoverTab[118538]"
}

// ToUint32 casts an interface to a uint32 type.
func ToUint32(i interface{}) uint32 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:89
	_go_fuzz_dep_.CoverTab[118539]++
										v, _ := ToUint32E(i)
										return v
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:91
	// _ = "end of CoverTab[118539]"
}

// ToUint16 casts an interface to a uint16 type.
func ToUint16(i interface{}) uint16 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:95
	_go_fuzz_dep_.CoverTab[118540]++
										v, _ := ToUint16E(i)
										return v
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:97
	// _ = "end of CoverTab[118540]"
}

// ToUint8 casts an interface to a uint8 type.
func ToUint8(i interface{}) uint8 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:101
	_go_fuzz_dep_.CoverTab[118541]++
										v, _ := ToUint8E(i)
										return v
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:103
	// _ = "end of CoverTab[118541]"
}

// ToString casts an interface to a string type.
func ToString(i interface{}) string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:107
	_go_fuzz_dep_.CoverTab[118542]++
										v, _ := ToStringE(i)
										return v
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:109
	// _ = "end of CoverTab[118542]"
}

// ToStringMapString casts an interface to a map[string]string type.
func ToStringMapString(i interface{}) map[string]string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:113
	_go_fuzz_dep_.CoverTab[118543]++
										v, _ := ToStringMapStringE(i)
										return v
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:115
	// _ = "end of CoverTab[118543]"
}

// ToStringMapStringSlice casts an interface to a map[string][]string type.
func ToStringMapStringSlice(i interface{}) map[string][]string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:119
	_go_fuzz_dep_.CoverTab[118544]++
										v, _ := ToStringMapStringSliceE(i)
										return v
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:121
	// _ = "end of CoverTab[118544]"
}

// ToStringMapBool casts an interface to a map[string]bool type.
func ToStringMapBool(i interface{}) map[string]bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:125
	_go_fuzz_dep_.CoverTab[118545]++
										v, _ := ToStringMapBoolE(i)
										return v
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:127
	// _ = "end of CoverTab[118545]"
}

// ToStringMapInt casts an interface to a map[string]int type.
func ToStringMapInt(i interface{}) map[string]int {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:131
	_go_fuzz_dep_.CoverTab[118546]++
										v, _ := ToStringMapIntE(i)
										return v
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:133
	// _ = "end of CoverTab[118546]"
}

// ToStringMapInt64 casts an interface to a map[string]int64 type.
func ToStringMapInt64(i interface{}) map[string]int64 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:137
	_go_fuzz_dep_.CoverTab[118547]++
										v, _ := ToStringMapInt64E(i)
										return v
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:139
	// _ = "end of CoverTab[118547]"
}

// ToStringMap casts an interface to a map[string]interface{} type.
func ToStringMap(i interface{}) map[string]interface{} {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:143
	_go_fuzz_dep_.CoverTab[118548]++
										v, _ := ToStringMapE(i)
										return v
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:145
	// _ = "end of CoverTab[118548]"
}

// ToSlice casts an interface to a []interface{} type.
func ToSlice(i interface{}) []interface{} {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:149
	_go_fuzz_dep_.CoverTab[118549]++
										v, _ := ToSliceE(i)
										return v
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:151
	// _ = "end of CoverTab[118549]"
}

// ToBoolSlice casts an interface to a []bool type.
func ToBoolSlice(i interface{}) []bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:155
	_go_fuzz_dep_.CoverTab[118550]++
										v, _ := ToBoolSliceE(i)
										return v
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:157
	// _ = "end of CoverTab[118550]"
}

// ToStringSlice casts an interface to a []string type.
func ToStringSlice(i interface{}) []string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:161
	_go_fuzz_dep_.CoverTab[118551]++
										v, _ := ToStringSliceE(i)
										return v
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:163
	// _ = "end of CoverTab[118551]"
}

// ToIntSlice casts an interface to a []int type.
func ToIntSlice(i interface{}) []int {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:167
	_go_fuzz_dep_.CoverTab[118552]++
										v, _ := ToIntSliceE(i)
										return v
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:169
	// _ = "end of CoverTab[118552]"
}

// ToDurationSlice casts an interface to a []time.Duration type.
func ToDurationSlice(i interface{}) []time.Duration {
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:173
	_go_fuzz_dep_.CoverTab[118553]++
										v, _ := ToDurationSliceE(i)
										return v
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:175
	// _ = "end of CoverTab[118553]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:176
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/cast@v1.4.1/cast.go:176
var _ = _go_fuzz_dep_.CoverTab
