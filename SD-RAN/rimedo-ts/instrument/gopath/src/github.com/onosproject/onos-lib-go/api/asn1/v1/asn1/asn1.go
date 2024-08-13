// Copyright 2021-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:15
package asn1

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:15
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:15
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:15
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:15
)

import (
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"math"
)

// UpdateValue - replace the bytes value with values from a new []byte
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:22
// the size stays the same
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:24
func (m *BitString) UpdateValue(newBytes []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:24
	_go_fuzz_dep_.CoverTab[170998]++
														if m == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:25
		_go_fuzz_dep_.CoverTab[171001]++
															return m.Value, errors.NewInvalid("null")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:26
		// _ = "end of CoverTab[171001]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:27
		_go_fuzz_dep_.CoverTab[171002]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:27
		// _ = "end of CoverTab[171002]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:27
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:27
	// _ = "end of CoverTab[170998]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:27
	_go_fuzz_dep_.CoverTab[170999]++
														expectedLen := int(math.Ceil(float64(m.Len) / 8.0))
														if len(newBytes) != expectedLen {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:29
		_go_fuzz_dep_.CoverTab[171003]++
															return m.Value, errors.NewInvalid("too many bytes %d. Expecting %d", len(newBytes), expectedLen)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:30
		// _ = "end of CoverTab[171003]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:31
		_go_fuzz_dep_.CoverTab[171004]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:31
		// _ = "end of CoverTab[171004]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:31
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:31
	// _ = "end of CoverTab[170999]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:31
	_go_fuzz_dep_.CoverTab[171000]++
														m.Value = newBytes
														return m.Value, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:33
	// _ = "end of CoverTab[171000]"
}

// TruncateValue - truncates value of trailing bits in the BitString the size stays the same
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:36
// Assuming that BitString has a non-empty length
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:38
func (m *BitString) TruncateValue() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:38
	_go_fuzz_dep_.CoverTab[171005]++
														if m == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:39
		_go_fuzz_dep_.CoverTab[171010]++
															return m.Value, errors.NewInvalid("null")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:40
		// _ = "end of CoverTab[171010]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:41
		_go_fuzz_dep_.CoverTab[171011]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:41
		// _ = "end of CoverTab[171011]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:41
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:41
	// _ = "end of CoverTab[171005]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:41
	_go_fuzz_dep_.CoverTab[171006]++
														if m.Len == 0 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:42
		_go_fuzz_dep_.CoverTab[171012]++
															return nil, errors.NewInvalid("Length should not be 0")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:43
		// _ = "end of CoverTab[171012]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:44
		_go_fuzz_dep_.CoverTab[171013]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:44
		// _ = "end of CoverTab[171013]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:44
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:44
	// _ = "end of CoverTab[171006]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:44
	_go_fuzz_dep_.CoverTab[171007]++

														expectedBytesLen := int(math.Ceil(float64(m.Len) / 8.0))
														if len(m.Value) != expectedBytesLen {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:47
		_go_fuzz_dep_.CoverTab[171014]++
															return m.Value, errors.NewInvalid("too many bytes %d. Expecting %d", len(m.Value), expectedBytesLen)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:48
		// _ = "end of CoverTab[171014]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:49
		_go_fuzz_dep_.CoverTab[171015]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:49
		// _ = "end of CoverTab[171015]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:49
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:49
	// _ = "end of CoverTab[171007]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:49
	_go_fuzz_dep_.CoverTab[171008]++

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:52
	truncBytes := make([]byte, expectedBytesLen)
	for i := 0; i < expectedBytesLen; i++ {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:53
		_go_fuzz_dep_.CoverTab[171016]++
															truncBytes[i] = m.Value[i]
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:54
		// _ = "end of CoverTab[171016]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:55
	// _ = "end of CoverTab[171008]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:55
	_go_fuzz_dep_.CoverTab[171009]++

														bitsFull := expectedBytesLen * 8
														trailingBits := uint32(bitsFull) - m.Len

														mask := ^((1 << trailingBits) - 1)
														truncBytes[len(truncBytes)-1] = truncBytes[len(truncBytes)-1] & byte(mask)

														m.Value = truncBytes
														return m.Value, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:64
	// _ = "end of CoverTab[171009]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:65
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/api/asn1/v1/asn1/asn1.go:65
var _ = _go_fuzz_dep_.CoverTab
