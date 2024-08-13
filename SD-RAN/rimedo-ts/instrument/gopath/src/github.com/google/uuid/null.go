// Copyright 2021 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:5
package uuid

//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:5
)

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

var jsonNull = []byte("null")

// NullUUID represents a UUID that may be null.
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:16
// NullUUID implements the SQL driver.Scanner interface so
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:16
// it can be used as a scan destination:
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:16
//
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:16
//	var u uuid.NullUUID
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:16
//	err := db.QueryRow("SELECT name FROM foo WHERE id=?", id).Scan(&u)
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:16
//	...
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:16
//	if u.Valid {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:16
//	   // use u.UUID
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:16
//	} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:16
//	   // NULL value
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:16
//	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:29
type NullUUID struct {
	UUID	UUID
	Valid	bool	// Valid is true if UUID is not NULL
}

// Scan implements the SQL driver.Scanner interface.
func (nu *NullUUID) Scan(value interface{}) error {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:35
	_go_fuzz_dep_.CoverTab[179359]++
										if value == nil {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:36
		_go_fuzz_dep_.CoverTab[179362]++
											nu.UUID, nu.Valid = Nil, false
											return nil
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:38
		// _ = "end of CoverTab[179362]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:39
		_go_fuzz_dep_.CoverTab[179363]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:39
		// _ = "end of CoverTab[179363]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:39
	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:39
	// _ = "end of CoverTab[179359]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:39
	_go_fuzz_dep_.CoverTab[179360]++

										err := nu.UUID.Scan(value)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:42
		_go_fuzz_dep_.CoverTab[179364]++
											nu.Valid = false
											return err
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:44
		// _ = "end of CoverTab[179364]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:45
		_go_fuzz_dep_.CoverTab[179365]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:45
		// _ = "end of CoverTab[179365]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:45
	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:45
	// _ = "end of CoverTab[179360]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:45
	_go_fuzz_dep_.CoverTab[179361]++

										nu.Valid = true
										return nil
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:48
	// _ = "end of CoverTab[179361]"
}

// Value implements the driver Valuer interface.
func (nu NullUUID) Value() (driver.Value, error) {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:52
	_go_fuzz_dep_.CoverTab[179366]++
										if !nu.Valid {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:53
		_go_fuzz_dep_.CoverTab[179368]++
											return nil, nil
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:54
		// _ = "end of CoverTab[179368]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:55
		_go_fuzz_dep_.CoverTab[179369]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:55
		// _ = "end of CoverTab[179369]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:55
	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:55
	// _ = "end of CoverTab[179366]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:55
	_go_fuzz_dep_.CoverTab[179367]++

										return nu.UUID.Value()
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:57
	// _ = "end of CoverTab[179367]"
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (nu NullUUID) MarshalBinary() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:61
	_go_fuzz_dep_.CoverTab[179370]++
										if nu.Valid {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:62
		_go_fuzz_dep_.CoverTab[179372]++
											return nu.UUID[:], nil
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:63
		// _ = "end of CoverTab[179372]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:64
		_go_fuzz_dep_.CoverTab[179373]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:64
		// _ = "end of CoverTab[179373]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:64
	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:64
	// _ = "end of CoverTab[179370]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:64
	_go_fuzz_dep_.CoverTab[179371]++

										return []byte(nil), nil
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:66
	// _ = "end of CoverTab[179371]"
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (nu *NullUUID) UnmarshalBinary(data []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:70
	_go_fuzz_dep_.CoverTab[179374]++
										if len(data) != 16 {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:71
		_go_fuzz_dep_.CoverTab[179376]++
											return fmt.Errorf("invalid UUID (got %d bytes)", len(data))
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:72
		// _ = "end of CoverTab[179376]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:73
		_go_fuzz_dep_.CoverTab[179377]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:73
		// _ = "end of CoverTab[179377]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:73
	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:73
	// _ = "end of CoverTab[179374]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:73
	_go_fuzz_dep_.CoverTab[179375]++
										copy(nu.UUID[:], data)
										nu.Valid = true
										return nil
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:76
	// _ = "end of CoverTab[179375]"
}

// MarshalText implements encoding.TextMarshaler.
func (nu NullUUID) MarshalText() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:80
	_go_fuzz_dep_.CoverTab[179378]++
										if nu.Valid {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:81
		_go_fuzz_dep_.CoverTab[179380]++
											return nu.UUID.MarshalText()
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:82
		// _ = "end of CoverTab[179380]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:83
		_go_fuzz_dep_.CoverTab[179381]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:83
		// _ = "end of CoverTab[179381]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:83
	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:83
	// _ = "end of CoverTab[179378]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:83
	_go_fuzz_dep_.CoverTab[179379]++

										return jsonNull, nil
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:85
	// _ = "end of CoverTab[179379]"
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (nu *NullUUID) UnmarshalText(data []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:89
	_go_fuzz_dep_.CoverTab[179382]++
										id, err := ParseBytes(data)
										if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:91
		_go_fuzz_dep_.CoverTab[179384]++
											nu.Valid = false
											return err
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:93
		// _ = "end of CoverTab[179384]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:94
		_go_fuzz_dep_.CoverTab[179385]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:94
		// _ = "end of CoverTab[179385]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:94
	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:94
	// _ = "end of CoverTab[179382]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:94
	_go_fuzz_dep_.CoverTab[179383]++
										nu.UUID = id
										nu.Valid = true
										return nil
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:97
	// _ = "end of CoverTab[179383]"
}

// MarshalJSON implements json.Marshaler.
func (nu NullUUID) MarshalJSON() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:101
	_go_fuzz_dep_.CoverTab[179386]++
										if nu.Valid {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:102
		_go_fuzz_dep_.CoverTab[179388]++
											return json.Marshal(nu.UUID)
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:103
		// _ = "end of CoverTab[179388]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:104
		_go_fuzz_dep_.CoverTab[179389]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:104
		// _ = "end of CoverTab[179389]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:104
	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:104
	// _ = "end of CoverTab[179386]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:104
	_go_fuzz_dep_.CoverTab[179387]++

										return jsonNull, nil
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:106
	// _ = "end of CoverTab[179387]"
}

// UnmarshalJSON implements json.Unmarshaler.
func (nu *NullUUID) UnmarshalJSON(data []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:110
	_go_fuzz_dep_.CoverTab[179390]++
										if bytes.Equal(data, jsonNull) {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:111
		_go_fuzz_dep_.CoverTab[179392]++
											*nu = NullUUID{}
											return nil
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:113
		// _ = "end of CoverTab[179392]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:114
		_go_fuzz_dep_.CoverTab[179393]++
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:114
		// _ = "end of CoverTab[179393]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:114
	}
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:114
	// _ = "end of CoverTab[179390]"
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:114
	_go_fuzz_dep_.CoverTab[179391]++
										err := json.Unmarshal(data, &nu.UUID)
										nu.Valid = err == nil
										return err
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:117
	// _ = "end of CoverTab[179391]"
}

//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:118
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/google/uuid@v1.3.0/null.go:118
var _ = _go_fuzz_dep_.CoverTab
