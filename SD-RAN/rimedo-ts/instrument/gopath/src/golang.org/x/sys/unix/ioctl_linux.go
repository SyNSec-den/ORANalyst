// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:5
package unix

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:5
)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:5
import (
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:5
)

import "unsafe"

// IoctlRetInt performs an ioctl operation specified by req on a device
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:9
// associated with opened file descriptor fd, and returns a non-negative
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:9
// integer that is returned by the ioctl syscall.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:12
func IoctlRetInt(fd int, req uint) (int, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:12
	_go_fuzz_dep_.CoverTab[45806]++
											ret, _, err := Syscall(SYS_IOCTL, uintptr(fd), uintptr(req), 0)
											if err != 0 {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:14
		_go_fuzz_dep_.CoverTab[45808]++
												return 0, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:15
		// _ = "end of CoverTab[45808]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:16
		_go_fuzz_dep_.CoverTab[45809]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:16
		// _ = "end of CoverTab[45809]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:16
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:16
	// _ = "end of CoverTab[45806]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:16
	_go_fuzz_dep_.CoverTab[45807]++
											return int(ret), nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:17
	// _ = "end of CoverTab[45807]"
}

func IoctlGetUint32(fd int, req uint) (uint32, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:20
	_go_fuzz_dep_.CoverTab[45810]++
											var value uint32
											err := ioctlPtr(fd, req, unsafe.Pointer(&value))
											return value, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:23
	// _ = "end of CoverTab[45810]"
}

func IoctlGetRTCTime(fd int) (*RTCTime, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:26
	_go_fuzz_dep_.CoverTab[45811]++
											var value RTCTime
											err := ioctlPtr(fd, RTC_RD_TIME, unsafe.Pointer(&value))
											return &value, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:29
	// _ = "end of CoverTab[45811]"
}

func IoctlSetRTCTime(fd int, value *RTCTime) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:32
	_go_fuzz_dep_.CoverTab[45812]++
											return ioctlPtr(fd, RTC_SET_TIME, unsafe.Pointer(value))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:33
	// _ = "end of CoverTab[45812]"
}

func IoctlGetRTCWkAlrm(fd int) (*RTCWkAlrm, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:36
	_go_fuzz_dep_.CoverTab[45813]++
											var value RTCWkAlrm
											err := ioctlPtr(fd, RTC_WKALM_RD, unsafe.Pointer(&value))
											return &value, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:39
	// _ = "end of CoverTab[45813]"
}

func IoctlSetRTCWkAlrm(fd int, value *RTCWkAlrm) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:42
	_go_fuzz_dep_.CoverTab[45814]++
											return ioctlPtr(fd, RTC_WKALM_SET, unsafe.Pointer(value))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:43
	// _ = "end of CoverTab[45814]"
}

// IoctlGetEthtoolDrvinfo fetches ethtool driver information for the network
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:46
// device specified by ifname.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:48
func IoctlGetEthtoolDrvinfo(fd int, ifname string) (*EthtoolDrvinfo, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:48
	_go_fuzz_dep_.CoverTab[45815]++
											ifr, err := NewIfreq(ifname)
											if err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:50
		_go_fuzz_dep_.CoverTab[45817]++
												return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:51
		// _ = "end of CoverTab[45817]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:52
		_go_fuzz_dep_.CoverTab[45818]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:52
		// _ = "end of CoverTab[45818]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:52
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:52
	// _ = "end of CoverTab[45815]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:52
	_go_fuzz_dep_.CoverTab[45816]++

											value := EthtoolDrvinfo{Cmd: ETHTOOL_GDRVINFO}
											ifrd := ifr.withData(unsafe.Pointer(&value))

											err = ioctlIfreqData(fd, SIOCETHTOOL, &ifrd)
											return &value, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:58
	// _ = "end of CoverTab[45816]"
}

// IoctlGetWatchdogInfo fetches information about a watchdog device from the
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:61
// Linux watchdog API. For more information, see:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:61
// https://www.kernel.org/doc/html/latest/watchdog/watchdog-api.html.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:64
func IoctlGetWatchdogInfo(fd int) (*WatchdogInfo, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:64
	_go_fuzz_dep_.CoverTab[45819]++
											var value WatchdogInfo
											err := ioctlPtr(fd, WDIOC_GETSUPPORT, unsafe.Pointer(&value))
											return &value, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:67
	// _ = "end of CoverTab[45819]"
}

// IoctlWatchdogKeepalive issues a keepalive ioctl to a watchdog device. For
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:70
// more information, see:
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:70
// https://www.kernel.org/doc/html/latest/watchdog/watchdog-api.html.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:73
func IoctlWatchdogKeepalive(fd int) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:73
	_go_fuzz_dep_.CoverTab[45820]++

											return ioctl(fd, WDIOC_KEEPALIVE, 0)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:75
	// _ = "end of CoverTab[45820]"
}

// IoctlFileCloneRange performs an FICLONERANGE ioctl operation to clone the
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:78
// range of data conveyed in value to the file associated with the file
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:78
// descriptor destFd. See the ioctl_ficlonerange(2) man page for details.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:81
func IoctlFileCloneRange(destFd int, value *FileCloneRange) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:81
	_go_fuzz_dep_.CoverTab[45821]++
											return ioctlPtr(destFd, FICLONERANGE, unsafe.Pointer(value))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:82
	// _ = "end of CoverTab[45821]"
}

// IoctlFileClone performs an FICLONE ioctl operation to clone the entire file
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:85
// associated with the file description srcFd to the file associated with the
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:85
// file descriptor destFd. See the ioctl_ficlone(2) man page for details.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:88
func IoctlFileClone(destFd, srcFd int) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:88
	_go_fuzz_dep_.CoverTab[45822]++
											return ioctl(destFd, FICLONE, uintptr(srcFd))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:89
	// _ = "end of CoverTab[45822]"
}

type FileDedupeRange struct {
	Src_offset	uint64
	Src_length	uint64
	Reserved1	uint16
	Reserved2	uint32
	Info		[]FileDedupeRangeInfo
}

type FileDedupeRangeInfo struct {
	Dest_fd		int64
	Dest_offset	uint64
	Bytes_deduped	uint64
	Status		int32
	Reserved	uint32
}

// IoctlFileDedupeRange performs an FIDEDUPERANGE ioctl operation to share the
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:108
// range of data conveyed in value from the file associated with the file
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:108
// descriptor srcFd to the value.Info destinations. See the
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:108
// ioctl_fideduperange(2) man page for details.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:112
func IoctlFileDedupeRange(srcFd int, value *FileDedupeRange) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:112
	_go_fuzz_dep_.CoverTab[45823]++
											buf := make([]byte, SizeofRawFileDedupeRange+
		len(value.Info)*SizeofRawFileDedupeRangeInfo)
	rawrange := (*RawFileDedupeRange)(unsafe.Pointer(&buf[0]))
	rawrange.Src_offset = value.Src_offset
	rawrange.Src_length = value.Src_length
	rawrange.Dest_count = uint16(len(value.Info))
	rawrange.Reserved1 = value.Reserved1
	rawrange.Reserved2 = value.Reserved2

	for i := range value.Info {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:122
		_go_fuzz_dep_.CoverTab[45826]++
												rawinfo := (*RawFileDedupeRangeInfo)(unsafe.Pointer(
			uintptr(unsafe.Pointer(&buf[0])) + uintptr(SizeofRawFileDedupeRange) +
				uintptr(i*SizeofRawFileDedupeRangeInfo)))
												rawinfo.Dest_fd = value.Info[i].Dest_fd
												rawinfo.Dest_offset = value.Info[i].Dest_offset
												rawinfo.Bytes_deduped = value.Info[i].Bytes_deduped
												rawinfo.Status = value.Info[i].Status
												rawinfo.Reserved = value.Info[i].Reserved
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:130
		// _ = "end of CoverTab[45826]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:131
	// _ = "end of CoverTab[45823]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:131
	_go_fuzz_dep_.CoverTab[45824]++

											err := ioctlPtr(srcFd, FIDEDUPERANGE, unsafe.Pointer(&buf[0]))

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:136
	for i := range value.Info {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:136
		_go_fuzz_dep_.CoverTab[45827]++
												rawinfo := (*RawFileDedupeRangeInfo)(unsafe.Pointer(
			uintptr(unsafe.Pointer(&buf[0])) + uintptr(SizeofRawFileDedupeRange) +
				uintptr(i*SizeofRawFileDedupeRangeInfo)))
												value.Info[i].Dest_fd = rawinfo.Dest_fd
												value.Info[i].Dest_offset = rawinfo.Dest_offset
												value.Info[i].Bytes_deduped = rawinfo.Bytes_deduped
												value.Info[i].Status = rawinfo.Status
												value.Info[i].Reserved = rawinfo.Reserved
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:144
		// _ = "end of CoverTab[45827]"
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:145
	// _ = "end of CoverTab[45824]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:145
	_go_fuzz_dep_.CoverTab[45825]++

											return err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:147
	// _ = "end of CoverTab[45825]"
}

func IoctlHIDGetDesc(fd int, value *HIDRawReportDescriptor) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:150
	_go_fuzz_dep_.CoverTab[45828]++
											return ioctlPtr(fd, HIDIOCGRDESC, unsafe.Pointer(value))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:151
	// _ = "end of CoverTab[45828]"
}

func IoctlHIDGetRawInfo(fd int) (*HIDRawDevInfo, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:154
	_go_fuzz_dep_.CoverTab[45829]++
											var value HIDRawDevInfo
											err := ioctlPtr(fd, HIDIOCGRAWINFO, unsafe.Pointer(&value))
											return &value, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:157
	// _ = "end of CoverTab[45829]"
}

func IoctlHIDGetRawName(fd int) (string, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:160
	_go_fuzz_dep_.CoverTab[45830]++
											var value [_HIDIOCGRAWNAME_LEN]byte
											err := ioctlPtr(fd, _HIDIOCGRAWNAME, unsafe.Pointer(&value[0]))
											return ByteSliceToString(value[:]), err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:163
	// _ = "end of CoverTab[45830]"
}

func IoctlHIDGetRawPhys(fd int) (string, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:166
	_go_fuzz_dep_.CoverTab[45831]++
											var value [_HIDIOCGRAWPHYS_LEN]byte
											err := ioctlPtr(fd, _HIDIOCGRAWPHYS, unsafe.Pointer(&value[0]))
											return ByteSliceToString(value[:]), err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:169
	// _ = "end of CoverTab[45831]"
}

func IoctlHIDGetRawUniq(fd int) (string, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:172
	_go_fuzz_dep_.CoverTab[45832]++
											var value [_HIDIOCGRAWUNIQ_LEN]byte
											err := ioctlPtr(fd, _HIDIOCGRAWUNIQ, unsafe.Pointer(&value[0]))
											return ByteSliceToString(value[:]), err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:175
	// _ = "end of CoverTab[45832]"
}

// IoctlIfreq performs an ioctl using an Ifreq structure for input and/or
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:178
// output. See the netdevice(7) man page for details.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:180
func IoctlIfreq(fd int, req uint, value *Ifreq) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:180
	_go_fuzz_dep_.CoverTab[45833]++

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:183
	return ioctlPtr(fd, req, unsafe.Pointer(&value.raw))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:183
	// _ = "end of CoverTab[45833]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:188
// ioctlIfreqData performs an ioctl using an ifreqData structure for input
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:188
// and/or output. See the netdevice(7) man page for details.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:190
func ioctlIfreqData(fd int, req uint, value *ifreqData) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:190
	_go_fuzz_dep_.CoverTab[45834]++

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:193
	return ioctlPtr(fd, req, unsafe.Pointer(value))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:193
	// _ = "end of CoverTab[45834]"
}

// IoctlKCMClone attaches a new file descriptor to a multiplexor by cloning an
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:196
// existing KCM socket, returning a structure containing the file descriptor of
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:196
// the new socket.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:199
func IoctlKCMClone(fd int) (*KCMClone, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:199
	_go_fuzz_dep_.CoverTab[45835]++
											var info KCMClone
											if err := ioctlPtr(fd, SIOCKCMCLONE, unsafe.Pointer(&info)); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:201
		_go_fuzz_dep_.CoverTab[45837]++
												return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:202
		// _ = "end of CoverTab[45837]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:203
		_go_fuzz_dep_.CoverTab[45838]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:203
		// _ = "end of CoverTab[45838]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:203
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:203
	// _ = "end of CoverTab[45835]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:203
	_go_fuzz_dep_.CoverTab[45836]++

											return &info, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:205
	// _ = "end of CoverTab[45836]"
}

// IoctlKCMAttach attaches a TCP socket and associated BPF program file
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:208
// descriptor to a multiplexor.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:210
func IoctlKCMAttach(fd int, info KCMAttach) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:210
	_go_fuzz_dep_.CoverTab[45839]++
											return ioctlPtr(fd, SIOCKCMATTACH, unsafe.Pointer(&info))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:211
	// _ = "end of CoverTab[45839]"
}

// IoctlKCMUnattach unattaches a TCP socket file descriptor from a multiplexor.
func IoctlKCMUnattach(fd int, info KCMUnattach) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:215
	_go_fuzz_dep_.CoverTab[45840]++
											return ioctlPtr(fd, SIOCKCMUNATTACH, unsafe.Pointer(&info))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:216
	// _ = "end of CoverTab[45840]"
}

// IoctlLoopGetStatus64 gets the status of the loop device associated with the
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:219
// file descriptor fd using the LOOP_GET_STATUS64 operation.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:221
func IoctlLoopGetStatus64(fd int) (*LoopInfo64, error) {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:221
	_go_fuzz_dep_.CoverTab[45841]++
											var value LoopInfo64
											if err := ioctlPtr(fd, LOOP_GET_STATUS64, unsafe.Pointer(&value)); err != nil {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:223
		_go_fuzz_dep_.CoverTab[45843]++
												return nil, err
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:224
		// _ = "end of CoverTab[45843]"
	} else {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:225
		_go_fuzz_dep_.CoverTab[45844]++
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:225
		// _ = "end of CoverTab[45844]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:225
	}
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:225
	// _ = "end of CoverTab[45841]"
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:225
	_go_fuzz_dep_.CoverTab[45842]++
											return &value, nil
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:226
	// _ = "end of CoverTab[45842]"
}

// IoctlLoopSetStatus64 sets the status of the loop device associated with the
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:229
// file descriptor fd using the LOOP_SET_STATUS64 operation.
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:231
func IoctlLoopSetStatus64(fd int, value *LoopInfo64) error {
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:231
	_go_fuzz_dep_.CoverTab[45845]++
											return ioctlPtr(fd, LOOP_SET_STATUS64, unsafe.Pointer(value))
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:232
	// _ = "end of CoverTab[45845]"
}

//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:233
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/golang.org/x/sys@v0.8.0/unix/ioctl_linux.go:233
var _ = _go_fuzz_dep_.CoverTab
