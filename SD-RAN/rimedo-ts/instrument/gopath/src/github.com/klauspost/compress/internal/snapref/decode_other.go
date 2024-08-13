// Copyright 2016 The Snappy-Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:5
package snapref

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:5
)

// decode writes the decoding of src to dst. It assumes that the varint-encoded
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:7
// length of the decompressed bytes has already been read, and that len(dst)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:7
// equals that length.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:7
//
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:7
// It returns 0 on success or a decodeErrCodeXxx error code on failure.
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:12
func decode(dst, src []byte) int {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:12
	_go_fuzz_dep_.CoverTab[90481]++
														var d, s, offset, length int
														for s < len(src) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:14
		_go_fuzz_dep_.CoverTab[90484]++
															switch src[s] & 0x03 {
		case tagLiteral:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:16
			_go_fuzz_dep_.CoverTab[90489]++
																x := uint32(src[s] >> 2)
																switch {
			case x < 60:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:19
				_go_fuzz_dep_.CoverTab[90500]++
																	s++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:20
				// _ = "end of CoverTab[90500]"
			case x == 60:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:21
				_go_fuzz_dep_.CoverTab[90501]++
																	s += 2
																	if uint(s) > uint(len(src)) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:23
					_go_fuzz_dep_.CoverTab[90510]++
																		return decodeErrCodeCorrupt
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:24
					// _ = "end of CoverTab[90510]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:25
					_go_fuzz_dep_.CoverTab[90511]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:25
					// _ = "end of CoverTab[90511]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:25
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:25
				// _ = "end of CoverTab[90501]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:25
				_go_fuzz_dep_.CoverTab[90502]++
																	x = uint32(src[s-1])
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:26
				// _ = "end of CoverTab[90502]"
			case x == 61:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:27
				_go_fuzz_dep_.CoverTab[90503]++
																	s += 3
																	if uint(s) > uint(len(src)) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:29
					_go_fuzz_dep_.CoverTab[90512]++
																		return decodeErrCodeCorrupt
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:30
					// _ = "end of CoverTab[90512]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:31
					_go_fuzz_dep_.CoverTab[90513]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:31
					// _ = "end of CoverTab[90513]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:31
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:31
				// _ = "end of CoverTab[90503]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:31
				_go_fuzz_dep_.CoverTab[90504]++
																	x = uint32(src[s-2]) | uint32(src[s-1])<<8
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:32
				// _ = "end of CoverTab[90504]"
			case x == 62:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:33
				_go_fuzz_dep_.CoverTab[90505]++
																	s += 4
																	if uint(s) > uint(len(src)) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:35
					_go_fuzz_dep_.CoverTab[90514]++
																		return decodeErrCodeCorrupt
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:36
					// _ = "end of CoverTab[90514]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:37
					_go_fuzz_dep_.CoverTab[90515]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:37
					// _ = "end of CoverTab[90515]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:37
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:37
				// _ = "end of CoverTab[90505]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:37
				_go_fuzz_dep_.CoverTab[90506]++
																	x = uint32(src[s-3]) | uint32(src[s-2])<<8 | uint32(src[s-1])<<16
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:38
				// _ = "end of CoverTab[90506]"
			case x == 63:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:39
				_go_fuzz_dep_.CoverTab[90507]++
																	s += 5
																	if uint(s) > uint(len(src)) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:41
					_go_fuzz_dep_.CoverTab[90516]++
																		return decodeErrCodeCorrupt
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:42
					// _ = "end of CoverTab[90516]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:43
					_go_fuzz_dep_.CoverTab[90517]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:43
					// _ = "end of CoverTab[90517]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:43
				}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:43
				// _ = "end of CoverTab[90507]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:43
				_go_fuzz_dep_.CoverTab[90508]++
																	x = uint32(src[s-4]) | uint32(src[s-3])<<8 | uint32(src[s-2])<<16 | uint32(src[s-1])<<24
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:44
				// _ = "end of CoverTab[90508]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:44
			default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:44
				_go_fuzz_dep_.CoverTab[90509]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:44
				// _ = "end of CoverTab[90509]"
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:45
			// _ = "end of CoverTab[90489]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:45
			_go_fuzz_dep_.CoverTab[90490]++
																length = int(x) + 1
																if length <= 0 {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:47
				_go_fuzz_dep_.CoverTab[90518]++
																	return decodeErrCodeUnsupportedLiteralLength
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:48
				// _ = "end of CoverTab[90518]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:49
				_go_fuzz_dep_.CoverTab[90519]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:49
				// _ = "end of CoverTab[90519]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:49
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:49
			// _ = "end of CoverTab[90490]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:49
			_go_fuzz_dep_.CoverTab[90491]++
																if length > len(dst)-d || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:50
				_go_fuzz_dep_.CoverTab[90520]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:50
				return length > len(src)-s
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:50
				// _ = "end of CoverTab[90520]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:50
			}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:50
				_go_fuzz_dep_.CoverTab[90521]++
																	return decodeErrCodeCorrupt
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:51
				// _ = "end of CoverTab[90521]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:52
				_go_fuzz_dep_.CoverTab[90522]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:52
				// _ = "end of CoverTab[90522]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:52
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:52
			// _ = "end of CoverTab[90491]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:52
			_go_fuzz_dep_.CoverTab[90492]++
																copy(dst[d:], src[s:s+length])
																d += length
																s += length
																continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:56
			// _ = "end of CoverTab[90492]"

		case tagCopy1:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:58
			_go_fuzz_dep_.CoverTab[90493]++
																s += 2
																if uint(s) > uint(len(src)) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:60
				_go_fuzz_dep_.CoverTab[90523]++
																	return decodeErrCodeCorrupt
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:61
				// _ = "end of CoverTab[90523]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:62
				_go_fuzz_dep_.CoverTab[90524]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:62
				// _ = "end of CoverTab[90524]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:62
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:62
			// _ = "end of CoverTab[90493]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:62
			_go_fuzz_dep_.CoverTab[90494]++
																length = 4 + int(src[s-2])>>2&0x7
																offset = int(uint32(src[s-2])&0xe0<<3 | uint32(src[s-1]))
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:64
			// _ = "end of CoverTab[90494]"

		case tagCopy2:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:66
			_go_fuzz_dep_.CoverTab[90495]++
																s += 3
																if uint(s) > uint(len(src)) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:68
				_go_fuzz_dep_.CoverTab[90525]++
																	return decodeErrCodeCorrupt
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:69
				// _ = "end of CoverTab[90525]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:70
				_go_fuzz_dep_.CoverTab[90526]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:70
				// _ = "end of CoverTab[90526]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:70
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:70
			// _ = "end of CoverTab[90495]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:70
			_go_fuzz_dep_.CoverTab[90496]++
																length = 1 + int(src[s-3])>>2
																offset = int(uint32(src[s-2]) | uint32(src[s-1])<<8)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:72
			// _ = "end of CoverTab[90496]"

		case tagCopy4:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:74
			_go_fuzz_dep_.CoverTab[90497]++
																s += 5
																if uint(s) > uint(len(src)) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:76
				_go_fuzz_dep_.CoverTab[90527]++
																	return decodeErrCodeCorrupt
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:77
				// _ = "end of CoverTab[90527]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:78
				_go_fuzz_dep_.CoverTab[90528]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:78
				// _ = "end of CoverTab[90528]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:78
			}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:78
			// _ = "end of CoverTab[90497]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:78
			_go_fuzz_dep_.CoverTab[90498]++
																length = 1 + int(src[s-5])>>2
																offset = int(uint32(src[s-4]) | uint32(src[s-3])<<8 | uint32(src[s-2])<<16 | uint32(src[s-1])<<24)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:80
			// _ = "end of CoverTab[90498]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:80
		default:
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:80
			_go_fuzz_dep_.CoverTab[90499]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:80
			// _ = "end of CoverTab[90499]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:81
		// _ = "end of CoverTab[90484]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:81
		_go_fuzz_dep_.CoverTab[90485]++

															if offset <= 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:83
			_go_fuzz_dep_.CoverTab[90529]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:83
			return d < offset
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:83
			// _ = "end of CoverTab[90529]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:83
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:83
			_go_fuzz_dep_.CoverTab[90530]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:83
			return length > len(dst)-d
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:83
			// _ = "end of CoverTab[90530]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:83
		}() {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:83
			_go_fuzz_dep_.CoverTab[90531]++
																return decodeErrCodeCorrupt
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:84
			// _ = "end of CoverTab[90531]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:85
			_go_fuzz_dep_.CoverTab[90532]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:85
			// _ = "end of CoverTab[90532]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:85
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:85
		// _ = "end of CoverTab[90485]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:85
		_go_fuzz_dep_.CoverTab[90486]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:88
		if offset >= length {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:88
			_go_fuzz_dep_.CoverTab[90533]++
																copy(dst[d:d+length], dst[d-offset:])
																d += length
																continue
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:91
			// _ = "end of CoverTab[90533]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:92
			_go_fuzz_dep_.CoverTab[90534]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:92
			// _ = "end of CoverTab[90534]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:92
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:92
		// _ = "end of CoverTab[90486]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:92
		_go_fuzz_dep_.CoverTab[90487]++

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:101
		a := dst[d : d+length]
		b := dst[d-offset:]
		b = b[:len(a)]
		for i := range a {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:104
			_go_fuzz_dep_.CoverTab[90535]++
																a[i] = b[i]
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:105
			// _ = "end of CoverTab[90535]"
		}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:106
		// _ = "end of CoverTab[90487]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:106
		_go_fuzz_dep_.CoverTab[90488]++
															d += length
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:107
		// _ = "end of CoverTab[90488]"
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:108
	// _ = "end of CoverTab[90481]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:108
	_go_fuzz_dep_.CoverTab[90482]++
														if d != len(dst) {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:109
		_go_fuzz_dep_.CoverTab[90536]++
															return decodeErrCodeCorrupt
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:110
		// _ = "end of CoverTab[90536]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:111
		_go_fuzz_dep_.CoverTab[90537]++
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:111
		// _ = "end of CoverTab[90537]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:111
	}
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:111
	// _ = "end of CoverTab[90482]"
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:111
	_go_fuzz_dep_.CoverTab[90483]++
														return 0
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:112
	// _ = "end of CoverTab[90483]"
}

//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:113
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/klauspost/compress@v1.14.2/internal/snapref/decode_other.go:113
var _ = _go_fuzz_dep_.CoverTab
