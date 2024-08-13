// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Bluetooth sockets and messages

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/bluetooth_linux.go:7
package unix

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/bluetooth_linux.go:7
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/bluetooth_linux.go:7
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/bluetooth_linux.go:7
)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/bluetooth_linux.go:7
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/bluetooth_linux.go:7
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/bluetooth_linux.go:7
)

// Bluetooth Protocols
const (
	BTPROTO_L2CAP	= 0
	BTPROTO_HCI	= 1
	BTPROTO_SCO	= 2
	BTPROTO_RFCOMM	= 3
	BTPROTO_BNEP	= 4
	BTPROTO_CMTP	= 5
	BTPROTO_HIDP	= 6
	BTPROTO_AVDTP	= 7
)

const (
	HCI_CHANNEL_RAW		= 0
	HCI_CHANNEL_USER	= 1
	HCI_CHANNEL_MONITOR	= 2
	HCI_CHANNEL_CONTROL	= 3
	HCI_CHANNEL_LOGGING	= 4
)

// Socketoption Level
const (
	SOL_BLUETOOTH	= 0x112
	SOL_HCI		= 0x0
	SOL_L2CAP	= 0x6
	SOL_RFCOMM	= 0x12
	SOL_SCO		= 0x11
)

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/bluetooth_linux.go:36
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/bluetooth_linux.go:36
var _ = _go_fuzz_dep_.CoverTab
