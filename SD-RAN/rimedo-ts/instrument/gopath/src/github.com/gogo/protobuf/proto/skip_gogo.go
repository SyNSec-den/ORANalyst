// Protocol Buffers for Go with Gadgets
//
// Copyright (c) 2013, The GoGo Authors. All rights reserved.
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

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:29
package proto

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:29
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:29
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:29
)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:29
import (
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:29
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:29
)

import (
	"fmt"
	"io"
)

func Skip(data []byte) (n int, err error) {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:36
	_go_fuzz_dep_.CoverTab[109113]++
												l := len(data)
												index := 0
												for index < l {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:39
		_go_fuzz_dep_.CoverTab[109115]++
													var wire uint64
													for shift := uint(0); ; shift += 7 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:41
			_go_fuzz_dep_.CoverTab[109117]++
														if index >= l {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:42
				_go_fuzz_dep_.CoverTab[109119]++
															return 0, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:43
				// _ = "end of CoverTab[109119]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:44
				_go_fuzz_dep_.CoverTab[109120]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:44
				// _ = "end of CoverTab[109120]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:44
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:44
			// _ = "end of CoverTab[109117]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:44
			_go_fuzz_dep_.CoverTab[109118]++
														b := data[index]
														index++
														wire |= (uint64(b) & 0x7F) << shift
														if b < 0x80 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:48
				_go_fuzz_dep_.CoverTab[109121]++
															break
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:49
				// _ = "end of CoverTab[109121]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:50
				_go_fuzz_dep_.CoverTab[109122]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:50
				// _ = "end of CoverTab[109122]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:50
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:50
			// _ = "end of CoverTab[109118]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:51
		// _ = "end of CoverTab[109115]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:51
		_go_fuzz_dep_.CoverTab[109116]++
													wireType := int(wire & 0x7)
													switch wireType {
		case 0:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:54
			_go_fuzz_dep_.CoverTab[109123]++
														for {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:55
				_go_fuzz_dep_.CoverTab[109133]++
															if index >= l {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:56
					_go_fuzz_dep_.CoverTab[109135]++
																return 0, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:57
					// _ = "end of CoverTab[109135]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:58
					_go_fuzz_dep_.CoverTab[109136]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:58
					// _ = "end of CoverTab[109136]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:58
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:58
				// _ = "end of CoverTab[109133]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:58
				_go_fuzz_dep_.CoverTab[109134]++
															index++
															if data[index-1] < 0x80 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:60
					_go_fuzz_dep_.CoverTab[109137]++
																break
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:61
					// _ = "end of CoverTab[109137]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:62
					_go_fuzz_dep_.CoverTab[109138]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:62
					// _ = "end of CoverTab[109138]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:62
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:62
				// _ = "end of CoverTab[109134]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:63
			// _ = "end of CoverTab[109123]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:63
			_go_fuzz_dep_.CoverTab[109124]++
														return index, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:64
			// _ = "end of CoverTab[109124]"
		case 1:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:65
			_go_fuzz_dep_.CoverTab[109125]++
														index += 8
														return index, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:67
			// _ = "end of CoverTab[109125]"
		case 2:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:68
			_go_fuzz_dep_.CoverTab[109126]++
														var length int
														for shift := uint(0); ; shift += 7 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:70
				_go_fuzz_dep_.CoverTab[109139]++
															if index >= l {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:71
					_go_fuzz_dep_.CoverTab[109141]++
																return 0, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:72
					// _ = "end of CoverTab[109141]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:73
					_go_fuzz_dep_.CoverTab[109142]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:73
					// _ = "end of CoverTab[109142]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:73
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:73
				// _ = "end of CoverTab[109139]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:73
				_go_fuzz_dep_.CoverTab[109140]++
															b := data[index]
															index++
															length |= (int(b) & 0x7F) << shift
															if b < 0x80 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:77
					_go_fuzz_dep_.CoverTab[109143]++
																break
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:78
					// _ = "end of CoverTab[109143]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:79
					_go_fuzz_dep_.CoverTab[109144]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:79
					// _ = "end of CoverTab[109144]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:79
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:79
				// _ = "end of CoverTab[109140]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:80
			// _ = "end of CoverTab[109126]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:80
			_go_fuzz_dep_.CoverTab[109127]++
														index += length
														return index, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:82
			// _ = "end of CoverTab[109127]"
		case 3:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:83
			_go_fuzz_dep_.CoverTab[109128]++
														for {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:84
				_go_fuzz_dep_.CoverTab[109145]++
															var innerWire uint64
															var start int = index
															for shift := uint(0); ; shift += 7 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:87
					_go_fuzz_dep_.CoverTab[109149]++
																if index >= l {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:88
						_go_fuzz_dep_.CoverTab[109151]++
																	return 0, io.ErrUnexpectedEOF
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:89
						// _ = "end of CoverTab[109151]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:90
						_go_fuzz_dep_.CoverTab[109152]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:90
						// _ = "end of CoverTab[109152]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:90
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:90
					// _ = "end of CoverTab[109149]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:90
					_go_fuzz_dep_.CoverTab[109150]++
																b := data[index]
																index++
																innerWire |= (uint64(b) & 0x7F) << shift
																if b < 0x80 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:94
						_go_fuzz_dep_.CoverTab[109153]++
																	break
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:95
						// _ = "end of CoverTab[109153]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:96
						_go_fuzz_dep_.CoverTab[109154]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:96
						// _ = "end of CoverTab[109154]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:96
					}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:96
					// _ = "end of CoverTab[109150]"
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:97
				// _ = "end of CoverTab[109145]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:97
				_go_fuzz_dep_.CoverTab[109146]++
															innerWireType := int(innerWire & 0x7)
															if innerWireType == 4 {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:99
					_go_fuzz_dep_.CoverTab[109155]++
																break
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:100
					// _ = "end of CoverTab[109155]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:101
					_go_fuzz_dep_.CoverTab[109156]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:101
					// _ = "end of CoverTab[109156]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:101
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:101
				// _ = "end of CoverTab[109146]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:101
				_go_fuzz_dep_.CoverTab[109147]++
															next, err := Skip(data[start:])
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:103
					_go_fuzz_dep_.CoverTab[109157]++
																return 0, err
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:104
					// _ = "end of CoverTab[109157]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:105
					_go_fuzz_dep_.CoverTab[109158]++
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:105
					// _ = "end of CoverTab[109158]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:105
				}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:105
				// _ = "end of CoverTab[109147]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:105
				_go_fuzz_dep_.CoverTab[109148]++
															index = start + next
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:106
				// _ = "end of CoverTab[109148]"
			}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:107
			// _ = "end of CoverTab[109128]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:107
			_go_fuzz_dep_.CoverTab[109129]++
														return index, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:108
			// _ = "end of CoverTab[109129]"
		case 4:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:109
			_go_fuzz_dep_.CoverTab[109130]++
														return index, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:110
			// _ = "end of CoverTab[109130]"
		case 5:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:111
			_go_fuzz_dep_.CoverTab[109131]++
														index += 4
														return index, nil
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:113
			// _ = "end of CoverTab[109131]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:114
			_go_fuzz_dep_.CoverTab[109132]++
														return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:115
			// _ = "end of CoverTab[109132]"
		}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:116
		// _ = "end of CoverTab[109116]"
	}
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:117
	// _ = "end of CoverTab[109113]"
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:117
	_go_fuzz_dep_.CoverTab[109114]++
												panic("unreachable")
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:118
	// _ = "end of CoverTab[109114]"
}

//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:119
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto/skip_gogo.go:119
var _ = _go_fuzz_dep_.CoverTab
