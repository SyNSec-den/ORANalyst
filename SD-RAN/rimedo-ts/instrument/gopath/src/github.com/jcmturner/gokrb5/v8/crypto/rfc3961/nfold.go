//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:1
package rfc3961

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:1
)

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:16
// Nfold expands the key to ensure it is not smaller than one cipher block.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:16
// Defined in RFC 3961.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:16
//
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:16
// m input bytes that will be "stretched" to the least common multiple of n bits and the bit length of m.
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:20
func Nfold(m []byte, n int) []byte {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:20
	_go_fuzz_dep_.CoverTab[85614]++
													k := len(m) * 8

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:24
	lcm := lcm(n, k)
	relicate := lcm / k
	var sumBytes []byte

	for i := 0; i < relicate; i++ {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:28
		_go_fuzz_dep_.CoverTab[85617]++
														rotation := 13 * i
														sumBytes = append(sumBytes, rotateRight(m, rotation)...)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:30
		// _ = "end of CoverTab[85617]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:31
	// _ = "end of CoverTab[85614]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:31
	_go_fuzz_dep_.CoverTab[85615]++

													nfold := make([]byte, n/8)
													sum := make([]byte, n/8)
													for i := 0; i < lcm/n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:35
		_go_fuzz_dep_.CoverTab[85618]++
														for j := 0; j < n/8; j++ {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:36
			_go_fuzz_dep_.CoverTab[85620]++
															sum[j] = sumBytes[j+(i*len(sum))]
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:37
			// _ = "end of CoverTab[85620]"
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:38
		// _ = "end of CoverTab[85618]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:38
		_go_fuzz_dep_.CoverTab[85619]++
														nfold = onesComplementAddition(nfold, sum)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:39
		// _ = "end of CoverTab[85619]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:40
	// _ = "end of CoverTab[85615]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:40
	_go_fuzz_dep_.CoverTab[85616]++
													return nfold
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:41
	// _ = "end of CoverTab[85616]"
}

func onesComplementAddition(n1, n2 []byte) []byte {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:44
	_go_fuzz_dep_.CoverTab[85621]++
													numBits := len(n1) * 8
													out := make([]byte, numBits/8)
													carry := 0
													for i := numBits - 1; i > -1; i-- {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:48
		_go_fuzz_dep_.CoverTab[85624]++
														n1b := getBit(&n1, i)
														n2b := getBit(&n2, i)
														s := n1b + n2b + carry

														if s == 0 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:53
			_go_fuzz_dep_.CoverTab[85625]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:53
			return s == 1
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:53
			// _ = "end of CoverTab[85625]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:53
		}() {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:53
			_go_fuzz_dep_.CoverTab[85626]++
															setBit(&out, i, s)
															carry = 0
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:55
			// _ = "end of CoverTab[85626]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:56
			_go_fuzz_dep_.CoverTab[85627]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:56
			if s == 2 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:56
				_go_fuzz_dep_.CoverTab[85628]++
																carry = 1
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:57
				// _ = "end of CoverTab[85628]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:58
				_go_fuzz_dep_.CoverTab[85629]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:58
				if s == 3 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:58
					_go_fuzz_dep_.CoverTab[85630]++
																	setBit(&out, i, 1)
																	carry = 1
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:60
					// _ = "end of CoverTab[85630]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:61
					_go_fuzz_dep_.CoverTab[85631]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:61
					// _ = "end of CoverTab[85631]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:61
				}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:61
				// _ = "end of CoverTab[85629]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:61
			}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:61
			// _ = "end of CoverTab[85627]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:61
		}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:61
		// _ = "end of CoverTab[85624]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:62
	// _ = "end of CoverTab[85621]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:62
	_go_fuzz_dep_.CoverTab[85622]++
													if carry == 1 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:63
		_go_fuzz_dep_.CoverTab[85632]++
														carryArray := make([]byte, len(n1))
														carryArray[len(carryArray)-1] = 1
														out = onesComplementAddition(out, carryArray)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:66
		// _ = "end of CoverTab[85632]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:67
		_go_fuzz_dep_.CoverTab[85633]++
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:67
		// _ = "end of CoverTab[85633]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:67
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:67
	// _ = "end of CoverTab[85622]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:67
	_go_fuzz_dep_.CoverTab[85623]++
													return out
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:68
	// _ = "end of CoverTab[85623]"
}

func rotateRight(b []byte, step int) []byte {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:71
	_go_fuzz_dep_.CoverTab[85634]++
													out := make([]byte, len(b))
													bitLen := len(b) * 8
													for i := 0; i < bitLen; i++ {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:74
		_go_fuzz_dep_.CoverTab[85636]++
														v := getBit(&b, i)
														setBit(&out, (i+step)%bitLen, v)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:76
		// _ = "end of CoverTab[85636]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:77
	// _ = "end of CoverTab[85634]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:77
	_go_fuzz_dep_.CoverTab[85635]++
													return out
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:78
	// _ = "end of CoverTab[85635]"
}

func lcm(x, y int) int {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:81
	_go_fuzz_dep_.CoverTab[85637]++
													return (x * y) / gcd(x, y)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:82
	// _ = "end of CoverTab[85637]"
}

func gcd(x, y int) int {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:85
	_go_fuzz_dep_.CoverTab[85638]++
													for y != 0 {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:86
		_go_fuzz_dep_.CoverTab[85640]++
														x, y = y, x%y
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:87
		// _ = "end of CoverTab[85640]"
	}
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:88
	// _ = "end of CoverTab[85638]"
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:88
	_go_fuzz_dep_.CoverTab[85639]++
													return x
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:89
	// _ = "end of CoverTab[85639]"
}

func getBit(b *[]byte, p int) int {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:92
	_go_fuzz_dep_.CoverTab[85641]++
													pByte := p / 8
													pBit := uint(p % 8)
													vByte := (*b)[pByte]
													vInt := int(vByte >> (8 - (pBit + 1)) & 0x0001)
													return vInt
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:97
	// _ = "end of CoverTab[85641]"
}

func setBit(b *[]byte, p, v int) {
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:100
	_go_fuzz_dep_.CoverTab[85642]++
													pByte := p / 8
													pBit := uint(p % 8)
													oldByte := (*b)[pByte]
													var newByte byte
													newByte = byte(v<<(8-(pBit+1))) | oldByte
													(*b)[pByte] = newByte
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:106
	// _ = "end of CoverTab[85642]"
}

//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:107
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/jcmturner/gokrb5/v8@v8.4.2/crypto/rfc3961/nfold.go:107
var _ = _go_fuzz_dep_.CoverTab
