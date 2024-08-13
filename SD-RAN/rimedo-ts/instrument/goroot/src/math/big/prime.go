// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/math/big/prime.go:5
package big

//line /usr/local/go/src/math/big/prime.go:5
import (
//line /usr/local/go/src/math/big/prime.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/math/big/prime.go:5
)
//line /usr/local/go/src/math/big/prime.go:5
import (
//line /usr/local/go/src/math/big/prime.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/math/big/prime.go:5
)

import "math/rand"

// ProbablyPrime reports whether x is probably prime,
//line /usr/local/go/src/math/big/prime.go:9
// applying the Miller-Rabin test with n pseudorandomly chosen bases
//line /usr/local/go/src/math/big/prime.go:9
// as well as a Baillie-PSW test.
//line /usr/local/go/src/math/big/prime.go:9
//
//line /usr/local/go/src/math/big/prime.go:9
// If x is prime, ProbablyPrime returns true.
//line /usr/local/go/src/math/big/prime.go:9
// If x is chosen randomly and not prime, ProbablyPrime probably returns false.
//line /usr/local/go/src/math/big/prime.go:9
// The probability of returning true for a randomly chosen non-prime is at most ¼ⁿ.
//line /usr/local/go/src/math/big/prime.go:9
//
//line /usr/local/go/src/math/big/prime.go:9
// ProbablyPrime is 100% accurate for inputs less than 2⁶⁴.
//line /usr/local/go/src/math/big/prime.go:9
// See Menezes et al., Handbook of Applied Cryptography, 1997, pp. 145-149,
//line /usr/local/go/src/math/big/prime.go:9
// and FIPS 186-4 Appendix F for further discussion of the error probabilities.
//line /usr/local/go/src/math/big/prime.go:9
//
//line /usr/local/go/src/math/big/prime.go:9
// ProbablyPrime is not suitable for judging primes that an adversary may
//line /usr/local/go/src/math/big/prime.go:9
// have crafted to fool the test.
//line /usr/local/go/src/math/big/prime.go:9
//
//line /usr/local/go/src/math/big/prime.go:9
// As of Go 1.8, ProbablyPrime(0) is allowed and applies only a Baillie-PSW test.
//line /usr/local/go/src/math/big/prime.go:9
// Before Go 1.8, ProbablyPrime applied only the Miller-Rabin tests, and ProbablyPrime(0) panicked.
//line /usr/local/go/src/math/big/prime.go:26
func (x *Int) ProbablyPrime(n int) bool {
//line /usr/local/go/src/math/big/prime.go:26
	_go_fuzz_dep_.CoverTab[6438]++

//line /usr/local/go/src/math/big/prime.go:36
	if n < 0 {
//line /usr/local/go/src/math/big/prime.go:36
		_go_fuzz_dep_.CoverTab[6445]++
							panic("negative n for ProbablyPrime")
//line /usr/local/go/src/math/big/prime.go:37
		// _ = "end of CoverTab[6445]"
	} else {
//line /usr/local/go/src/math/big/prime.go:38
		_go_fuzz_dep_.CoverTab[6446]++
//line /usr/local/go/src/math/big/prime.go:38
		// _ = "end of CoverTab[6446]"
//line /usr/local/go/src/math/big/prime.go:38
	}
//line /usr/local/go/src/math/big/prime.go:38
	// _ = "end of CoverTab[6438]"
//line /usr/local/go/src/math/big/prime.go:38
	_go_fuzz_dep_.CoverTab[6439]++
						if x.neg || func() bool {
//line /usr/local/go/src/math/big/prime.go:39
		_go_fuzz_dep_.CoverTab[6447]++
//line /usr/local/go/src/math/big/prime.go:39
		return len(x.abs) == 0
//line /usr/local/go/src/math/big/prime.go:39
		// _ = "end of CoverTab[6447]"
//line /usr/local/go/src/math/big/prime.go:39
	}() {
//line /usr/local/go/src/math/big/prime.go:39
		_go_fuzz_dep_.CoverTab[6448]++
							return false
//line /usr/local/go/src/math/big/prime.go:40
		// _ = "end of CoverTab[6448]"
	} else {
//line /usr/local/go/src/math/big/prime.go:41
		_go_fuzz_dep_.CoverTab[6449]++
//line /usr/local/go/src/math/big/prime.go:41
		// _ = "end of CoverTab[6449]"
//line /usr/local/go/src/math/big/prime.go:41
	}
//line /usr/local/go/src/math/big/prime.go:41
	// _ = "end of CoverTab[6439]"
//line /usr/local/go/src/math/big/prime.go:41
	_go_fuzz_dep_.CoverTab[6440]++

	// primeBitMask records the primes < 64.
	const primeBitMask uint64 = 1<<2 | 1<<3 | 1<<5 | 1<<7 |
		1<<11 | 1<<13 | 1<<17 | 1<<19 | 1<<23 | 1<<29 | 1<<31 |
		1<<37 | 1<<41 | 1<<43 | 1<<47 | 1<<53 | 1<<59 | 1<<61

	w := x.abs[0]
	if len(x.abs) == 1 && func() bool {
//line /usr/local/go/src/math/big/prime.go:49
		_go_fuzz_dep_.CoverTab[6450]++
//line /usr/local/go/src/math/big/prime.go:49
		return w < 64
//line /usr/local/go/src/math/big/prime.go:49
		// _ = "end of CoverTab[6450]"
//line /usr/local/go/src/math/big/prime.go:49
	}() {
//line /usr/local/go/src/math/big/prime.go:49
		_go_fuzz_dep_.CoverTab[6451]++
							return primeBitMask&(1<<w) != 0
//line /usr/local/go/src/math/big/prime.go:50
		// _ = "end of CoverTab[6451]"
	} else {
//line /usr/local/go/src/math/big/prime.go:51
		_go_fuzz_dep_.CoverTab[6452]++
//line /usr/local/go/src/math/big/prime.go:51
		// _ = "end of CoverTab[6452]"
//line /usr/local/go/src/math/big/prime.go:51
	}
//line /usr/local/go/src/math/big/prime.go:51
	// _ = "end of CoverTab[6440]"
//line /usr/local/go/src/math/big/prime.go:51
	_go_fuzz_dep_.CoverTab[6441]++

						if w&1 == 0 {
//line /usr/local/go/src/math/big/prime.go:53
		_go_fuzz_dep_.CoverTab[6453]++
							return false
//line /usr/local/go/src/math/big/prime.go:54
		// _ = "end of CoverTab[6453]"
	} else {
//line /usr/local/go/src/math/big/prime.go:55
		_go_fuzz_dep_.CoverTab[6454]++
//line /usr/local/go/src/math/big/prime.go:55
		// _ = "end of CoverTab[6454]"
//line /usr/local/go/src/math/big/prime.go:55
	}
//line /usr/local/go/src/math/big/prime.go:55
	// _ = "end of CoverTab[6441]"
//line /usr/local/go/src/math/big/prime.go:55
	_go_fuzz_dep_.CoverTab[6442]++

						const primesA = 3 * 5 * 7 * 11 * 13 * 17 * 19 * 23 * 37
						const primesB = 29 * 31 * 41 * 43 * 47 * 53

						var rA, rB uint32
						switch _W {
	case 32:
//line /usr/local/go/src/math/big/prime.go:62
		_go_fuzz_dep_.CoverTab[6455]++
							rA = uint32(x.abs.modW(primesA))
							rB = uint32(x.abs.modW(primesB))
//line /usr/local/go/src/math/big/prime.go:64
		// _ = "end of CoverTab[6455]"
	case 64:
//line /usr/local/go/src/math/big/prime.go:65
		_go_fuzz_dep_.CoverTab[6456]++
							r := x.abs.modW((primesA * primesB) & _M)
							rA = uint32(r % primesA)
							rB = uint32(r % primesB)
//line /usr/local/go/src/math/big/prime.go:68
		// _ = "end of CoverTab[6456]"
	default:
//line /usr/local/go/src/math/big/prime.go:69
		_go_fuzz_dep_.CoverTab[6457]++
							panic("math/big: invalid word size")
//line /usr/local/go/src/math/big/prime.go:70
		// _ = "end of CoverTab[6457]"
	}
//line /usr/local/go/src/math/big/prime.go:71
	// _ = "end of CoverTab[6442]"
//line /usr/local/go/src/math/big/prime.go:71
	_go_fuzz_dep_.CoverTab[6443]++

						if rA%3 == 0 || func() bool {
//line /usr/local/go/src/math/big/prime.go:73
		_go_fuzz_dep_.CoverTab[6458]++
//line /usr/local/go/src/math/big/prime.go:73
		return rA%5 == 0
//line /usr/local/go/src/math/big/prime.go:73
		// _ = "end of CoverTab[6458]"
//line /usr/local/go/src/math/big/prime.go:73
	}() || func() bool {
//line /usr/local/go/src/math/big/prime.go:73
		_go_fuzz_dep_.CoverTab[6459]++
//line /usr/local/go/src/math/big/prime.go:73
		return rA%7 == 0
//line /usr/local/go/src/math/big/prime.go:73
		// _ = "end of CoverTab[6459]"
//line /usr/local/go/src/math/big/prime.go:73
	}() || func() bool {
//line /usr/local/go/src/math/big/prime.go:73
		_go_fuzz_dep_.CoverTab[6460]++
//line /usr/local/go/src/math/big/prime.go:73
		return rA%11 == 0
//line /usr/local/go/src/math/big/prime.go:73
		// _ = "end of CoverTab[6460]"
//line /usr/local/go/src/math/big/prime.go:73
	}() || func() bool {
//line /usr/local/go/src/math/big/prime.go:73
		_go_fuzz_dep_.CoverTab[6461]++
//line /usr/local/go/src/math/big/prime.go:73
		return rA%13 == 0
//line /usr/local/go/src/math/big/prime.go:73
		// _ = "end of CoverTab[6461]"
//line /usr/local/go/src/math/big/prime.go:73
	}() || func() bool {
//line /usr/local/go/src/math/big/prime.go:73
		_go_fuzz_dep_.CoverTab[6462]++
//line /usr/local/go/src/math/big/prime.go:73
		return rA%17 == 0
//line /usr/local/go/src/math/big/prime.go:73
		// _ = "end of CoverTab[6462]"
//line /usr/local/go/src/math/big/prime.go:73
	}() || func() bool {
//line /usr/local/go/src/math/big/prime.go:73
		_go_fuzz_dep_.CoverTab[6463]++
//line /usr/local/go/src/math/big/prime.go:73
		return rA%19 == 0
//line /usr/local/go/src/math/big/prime.go:73
		// _ = "end of CoverTab[6463]"
//line /usr/local/go/src/math/big/prime.go:73
	}() || func() bool {
//line /usr/local/go/src/math/big/prime.go:73
		_go_fuzz_dep_.CoverTab[6464]++
//line /usr/local/go/src/math/big/prime.go:73
		return rA%23 == 0
//line /usr/local/go/src/math/big/prime.go:73
		// _ = "end of CoverTab[6464]"
//line /usr/local/go/src/math/big/prime.go:73
	}() || func() bool {
//line /usr/local/go/src/math/big/prime.go:73
		_go_fuzz_dep_.CoverTab[6465]++
//line /usr/local/go/src/math/big/prime.go:73
		return rA%37 == 0
//line /usr/local/go/src/math/big/prime.go:73
		// _ = "end of CoverTab[6465]"
//line /usr/local/go/src/math/big/prime.go:73
	}() || func() bool {
//line /usr/local/go/src/math/big/prime.go:73
		_go_fuzz_dep_.CoverTab[6466]++
//line /usr/local/go/src/math/big/prime.go:73
		return rB%29 == 0
							// _ = "end of CoverTab[6466]"
//line /usr/local/go/src/math/big/prime.go:74
	}() || func() bool {
//line /usr/local/go/src/math/big/prime.go:74
		_go_fuzz_dep_.CoverTab[6467]++
//line /usr/local/go/src/math/big/prime.go:74
		return rB%31 == 0
//line /usr/local/go/src/math/big/prime.go:74
		// _ = "end of CoverTab[6467]"
//line /usr/local/go/src/math/big/prime.go:74
	}() || func() bool {
//line /usr/local/go/src/math/big/prime.go:74
		_go_fuzz_dep_.CoverTab[6468]++
//line /usr/local/go/src/math/big/prime.go:74
		return rB%41 == 0
//line /usr/local/go/src/math/big/prime.go:74
		// _ = "end of CoverTab[6468]"
//line /usr/local/go/src/math/big/prime.go:74
	}() || func() bool {
//line /usr/local/go/src/math/big/prime.go:74
		_go_fuzz_dep_.CoverTab[6469]++
//line /usr/local/go/src/math/big/prime.go:74
		return rB%43 == 0
//line /usr/local/go/src/math/big/prime.go:74
		// _ = "end of CoverTab[6469]"
//line /usr/local/go/src/math/big/prime.go:74
	}() || func() bool {
//line /usr/local/go/src/math/big/prime.go:74
		_go_fuzz_dep_.CoverTab[6470]++
//line /usr/local/go/src/math/big/prime.go:74
		return rB%47 == 0
//line /usr/local/go/src/math/big/prime.go:74
		// _ = "end of CoverTab[6470]"
//line /usr/local/go/src/math/big/prime.go:74
	}() || func() bool {
//line /usr/local/go/src/math/big/prime.go:74
		_go_fuzz_dep_.CoverTab[6471]++
//line /usr/local/go/src/math/big/prime.go:74
		return rB%53 == 0
//line /usr/local/go/src/math/big/prime.go:74
		// _ = "end of CoverTab[6471]"
//line /usr/local/go/src/math/big/prime.go:74
	}() {
//line /usr/local/go/src/math/big/prime.go:74
		_go_fuzz_dep_.CoverTab[6472]++
							return false
//line /usr/local/go/src/math/big/prime.go:75
		// _ = "end of CoverTab[6472]"
	} else {
//line /usr/local/go/src/math/big/prime.go:76
		_go_fuzz_dep_.CoverTab[6473]++
//line /usr/local/go/src/math/big/prime.go:76
		// _ = "end of CoverTab[6473]"
//line /usr/local/go/src/math/big/prime.go:76
	}
//line /usr/local/go/src/math/big/prime.go:76
	// _ = "end of CoverTab[6443]"
//line /usr/local/go/src/math/big/prime.go:76
	_go_fuzz_dep_.CoverTab[6444]++

						return x.abs.probablyPrimeMillerRabin(n+1, true) && func() bool {
//line /usr/local/go/src/math/big/prime.go:78
		_go_fuzz_dep_.CoverTab[6474]++
//line /usr/local/go/src/math/big/prime.go:78
		return x.abs.probablyPrimeLucas()
//line /usr/local/go/src/math/big/prime.go:78
		// _ = "end of CoverTab[6474]"
//line /usr/local/go/src/math/big/prime.go:78
	}()
//line /usr/local/go/src/math/big/prime.go:78
	// _ = "end of CoverTab[6444]"
}

// probablyPrimeMillerRabin reports whether n passes reps rounds of the
//line /usr/local/go/src/math/big/prime.go:81
// Miller-Rabin primality test, using pseudo-randomly chosen bases.
//line /usr/local/go/src/math/big/prime.go:81
// If force2 is true, one of the rounds is forced to use base 2.
//line /usr/local/go/src/math/big/prime.go:81
// See Handbook of Applied Cryptography, p. 139, Algorithm 4.24.
//line /usr/local/go/src/math/big/prime.go:81
// The number n is known to be non-zero.
//line /usr/local/go/src/math/big/prime.go:86
func (n nat) probablyPrimeMillerRabin(reps int, force2 bool) bool {
//line /usr/local/go/src/math/big/prime.go:86
	_go_fuzz_dep_.CoverTab[6475]++
						nm1 := nat(nil).sub(n, natOne)

						k := nm1.trailingZeroBits()
						q := nat(nil).shr(nm1, k)

						nm3 := nat(nil).sub(nm1, natTwo)
						rand := rand.New(rand.NewSource(int64(n[0])))

						var x, y, quotient nat
						nm3Len := nm3.bitLen()

NextRandom:
	for i := 0; i < reps; i++ {
//line /usr/local/go/src/math/big/prime.go:99
		_go_fuzz_dep_.CoverTab[6477]++
							if i == reps-1 && func() bool {
//line /usr/local/go/src/math/big/prime.go:100
			_go_fuzz_dep_.CoverTab[6481]++
//line /usr/local/go/src/math/big/prime.go:100
			return force2
//line /usr/local/go/src/math/big/prime.go:100
			// _ = "end of CoverTab[6481]"
//line /usr/local/go/src/math/big/prime.go:100
		}() {
//line /usr/local/go/src/math/big/prime.go:100
			_go_fuzz_dep_.CoverTab[6482]++
								x = x.set(natTwo)
//line /usr/local/go/src/math/big/prime.go:101
			// _ = "end of CoverTab[6482]"
		} else {
//line /usr/local/go/src/math/big/prime.go:102
			_go_fuzz_dep_.CoverTab[6483]++
								x = x.random(rand, nm3, nm3Len)
								x = x.add(x, natTwo)
//line /usr/local/go/src/math/big/prime.go:104
			// _ = "end of CoverTab[6483]"
		}
//line /usr/local/go/src/math/big/prime.go:105
		// _ = "end of CoverTab[6477]"
//line /usr/local/go/src/math/big/prime.go:105
		_go_fuzz_dep_.CoverTab[6478]++
							y = y.expNN(x, q, n, false)
							if y.cmp(natOne) == 0 || func() bool {
//line /usr/local/go/src/math/big/prime.go:107
			_go_fuzz_dep_.CoverTab[6484]++
//line /usr/local/go/src/math/big/prime.go:107
			return y.cmp(nm1) == 0
//line /usr/local/go/src/math/big/prime.go:107
			// _ = "end of CoverTab[6484]"
//line /usr/local/go/src/math/big/prime.go:107
		}() {
//line /usr/local/go/src/math/big/prime.go:107
			_go_fuzz_dep_.CoverTab[6485]++
								continue
//line /usr/local/go/src/math/big/prime.go:108
			// _ = "end of CoverTab[6485]"
		} else {
//line /usr/local/go/src/math/big/prime.go:109
			_go_fuzz_dep_.CoverTab[6486]++
//line /usr/local/go/src/math/big/prime.go:109
			// _ = "end of CoverTab[6486]"
//line /usr/local/go/src/math/big/prime.go:109
		}
//line /usr/local/go/src/math/big/prime.go:109
		// _ = "end of CoverTab[6478]"
//line /usr/local/go/src/math/big/prime.go:109
		_go_fuzz_dep_.CoverTab[6479]++
							for j := uint(1); j < k; j++ {
//line /usr/local/go/src/math/big/prime.go:110
			_go_fuzz_dep_.CoverTab[6487]++
								y = y.sqr(y)
								quotient, y = quotient.div(y, y, n)
								if y.cmp(nm1) == 0 {
//line /usr/local/go/src/math/big/prime.go:113
				_go_fuzz_dep_.CoverTab[6489]++
									continue NextRandom
//line /usr/local/go/src/math/big/prime.go:114
				// _ = "end of CoverTab[6489]"
			} else {
//line /usr/local/go/src/math/big/prime.go:115
				_go_fuzz_dep_.CoverTab[6490]++
//line /usr/local/go/src/math/big/prime.go:115
				// _ = "end of CoverTab[6490]"
//line /usr/local/go/src/math/big/prime.go:115
			}
//line /usr/local/go/src/math/big/prime.go:115
			// _ = "end of CoverTab[6487]"
//line /usr/local/go/src/math/big/prime.go:115
			_go_fuzz_dep_.CoverTab[6488]++
								if y.cmp(natOne) == 0 {
//line /usr/local/go/src/math/big/prime.go:116
				_go_fuzz_dep_.CoverTab[6491]++
									return false
//line /usr/local/go/src/math/big/prime.go:117
				// _ = "end of CoverTab[6491]"
			} else {
//line /usr/local/go/src/math/big/prime.go:118
				_go_fuzz_dep_.CoverTab[6492]++
//line /usr/local/go/src/math/big/prime.go:118
				// _ = "end of CoverTab[6492]"
//line /usr/local/go/src/math/big/prime.go:118
			}
//line /usr/local/go/src/math/big/prime.go:118
			// _ = "end of CoverTab[6488]"
		}
//line /usr/local/go/src/math/big/prime.go:119
		// _ = "end of CoverTab[6479]"
//line /usr/local/go/src/math/big/prime.go:119
		_go_fuzz_dep_.CoverTab[6480]++
							return false
//line /usr/local/go/src/math/big/prime.go:120
		// _ = "end of CoverTab[6480]"
	}
//line /usr/local/go/src/math/big/prime.go:121
	// _ = "end of CoverTab[6475]"
//line /usr/local/go/src/math/big/prime.go:121
	_go_fuzz_dep_.CoverTab[6476]++

						return true
//line /usr/local/go/src/math/big/prime.go:123
	// _ = "end of CoverTab[6476]"
}

// probablyPrimeLucas reports whether n passes the "almost extra strong" Lucas probable prime test,
//line /usr/local/go/src/math/big/prime.go:126
// using Baillie-OEIS parameter selection. This corresponds to "AESLPSP" on Jacobsen's tables (link below).
//line /usr/local/go/src/math/big/prime.go:126
// The combination of this test and a Miller-Rabin/Fermat test with base 2 gives a Baillie-PSW test.
//line /usr/local/go/src/math/big/prime.go:126
//
//line /usr/local/go/src/math/big/prime.go:126
// References:
//line /usr/local/go/src/math/big/prime.go:126
//
//line /usr/local/go/src/math/big/prime.go:126
// Baillie and Wagstaff, "Lucas Pseudoprimes", Mathematics of Computation 35(152),
//line /usr/local/go/src/math/big/prime.go:126
// October 1980, pp. 1391-1417, especially page 1401.
//line /usr/local/go/src/math/big/prime.go:126
// https://www.ams.org/journals/mcom/1980-35-152/S0025-5718-1980-0583518-6/S0025-5718-1980-0583518-6.pdf
//line /usr/local/go/src/math/big/prime.go:126
//
//line /usr/local/go/src/math/big/prime.go:126
// Grantham, "Frobenius Pseudoprimes", Mathematics of Computation 70(234),
//line /usr/local/go/src/math/big/prime.go:126
// March 2000, pp. 873-891.
//line /usr/local/go/src/math/big/prime.go:126
// https://www.ams.org/journals/mcom/2001-70-234/S0025-5718-00-01197-2/S0025-5718-00-01197-2.pdf
//line /usr/local/go/src/math/big/prime.go:126
//
//line /usr/local/go/src/math/big/prime.go:126
// Baillie, "Extra strong Lucas pseudoprimes", OEIS A217719, https://oeis.org/A217719.
//line /usr/local/go/src/math/big/prime.go:126
//
//line /usr/local/go/src/math/big/prime.go:126
// Jacobsen, "Pseudoprime Statistics, Tables, and Data", http://ntheory.org/pseudoprimes.html.
//line /usr/local/go/src/math/big/prime.go:126
//
//line /usr/local/go/src/math/big/prime.go:126
// Nicely, "The Baillie-PSW Primality Test", https://web.archive.org/web/20191121062007/http://www.trnicely.net/misc/bpsw.html.
//line /usr/local/go/src/math/big/prime.go:126
// (Note that Nicely's definition of the "extra strong" test gives the wrong Jacobi condition,
//line /usr/local/go/src/math/big/prime.go:126
// as pointed out by Jacobsen.)
//line /usr/local/go/src/math/big/prime.go:126
//
//line /usr/local/go/src/math/big/prime.go:126
// Crandall and Pomerance, Prime Numbers: A Computational Perspective, 2nd ed.
//line /usr/local/go/src/math/big/prime.go:126
// Springer, 2005.
//line /usr/local/go/src/math/big/prime.go:150
func (n nat) probablyPrimeLucas() bool {
//line /usr/local/go/src/math/big/prime.go:150
	_go_fuzz_dep_.CoverTab[6493]++

						if len(n) == 0 || func() bool {
//line /usr/local/go/src/math/big/prime.go:152
		_go_fuzz_dep_.CoverTab[6500]++
//line /usr/local/go/src/math/big/prime.go:152
		return n.cmp(natOne) == 0
//line /usr/local/go/src/math/big/prime.go:152
		// _ = "end of CoverTab[6500]"
//line /usr/local/go/src/math/big/prime.go:152
	}() {
//line /usr/local/go/src/math/big/prime.go:152
		_go_fuzz_dep_.CoverTab[6501]++
							return false
//line /usr/local/go/src/math/big/prime.go:153
		// _ = "end of CoverTab[6501]"
	} else {
//line /usr/local/go/src/math/big/prime.go:154
		_go_fuzz_dep_.CoverTab[6502]++
//line /usr/local/go/src/math/big/prime.go:154
		// _ = "end of CoverTab[6502]"
//line /usr/local/go/src/math/big/prime.go:154
	}
//line /usr/local/go/src/math/big/prime.go:154
	// _ = "end of CoverTab[6493]"
//line /usr/local/go/src/math/big/prime.go:154
	_go_fuzz_dep_.CoverTab[6494]++

//line /usr/local/go/src/math/big/prime.go:157
	if n[0]&1 == 0 {
//line /usr/local/go/src/math/big/prime.go:157
		_go_fuzz_dep_.CoverTab[6503]++
							return n.cmp(natTwo) == 0
//line /usr/local/go/src/math/big/prime.go:158
		// _ = "end of CoverTab[6503]"
	} else {
//line /usr/local/go/src/math/big/prime.go:159
		_go_fuzz_dep_.CoverTab[6504]++
//line /usr/local/go/src/math/big/prime.go:159
		// _ = "end of CoverTab[6504]"
//line /usr/local/go/src/math/big/prime.go:159
	}
//line /usr/local/go/src/math/big/prime.go:159
	// _ = "end of CoverTab[6494]"
//line /usr/local/go/src/math/big/prime.go:159
	_go_fuzz_dep_.CoverTab[6495]++

//line /usr/local/go/src/math/big/prime.go:168
	p := Word(3)
	d := nat{1}
	t1 := nat(nil)
	intD := &Int{abs: d}
	intN := &Int{abs: n}
	for ; ; p++ {
//line /usr/local/go/src/math/big/prime.go:173
		_go_fuzz_dep_.CoverTab[6505]++
							if p > 10000 {
//line /usr/local/go/src/math/big/prime.go:174
			_go_fuzz_dep_.CoverTab[6509]++

//line /usr/local/go/src/math/big/prime.go:177
			panic("math/big: internal error: cannot find (D/n) = -1 for " + intN.String())
//line /usr/local/go/src/math/big/prime.go:177
			// _ = "end of CoverTab[6509]"
		} else {
//line /usr/local/go/src/math/big/prime.go:178
			_go_fuzz_dep_.CoverTab[6510]++
//line /usr/local/go/src/math/big/prime.go:178
			// _ = "end of CoverTab[6510]"
//line /usr/local/go/src/math/big/prime.go:178
		}
//line /usr/local/go/src/math/big/prime.go:178
		// _ = "end of CoverTab[6505]"
//line /usr/local/go/src/math/big/prime.go:178
		_go_fuzz_dep_.CoverTab[6506]++
							d[0] = p*p - 4
							j := Jacobi(intD, intN)
							if j == -1 {
//line /usr/local/go/src/math/big/prime.go:181
			_go_fuzz_dep_.CoverTab[6511]++
								break
//line /usr/local/go/src/math/big/prime.go:182
			// _ = "end of CoverTab[6511]"
		} else {
//line /usr/local/go/src/math/big/prime.go:183
			_go_fuzz_dep_.CoverTab[6512]++
//line /usr/local/go/src/math/big/prime.go:183
			// _ = "end of CoverTab[6512]"
//line /usr/local/go/src/math/big/prime.go:183
		}
//line /usr/local/go/src/math/big/prime.go:183
		// _ = "end of CoverTab[6506]"
//line /usr/local/go/src/math/big/prime.go:183
		_go_fuzz_dep_.CoverTab[6507]++
							if j == 0 {
//line /usr/local/go/src/math/big/prime.go:184
			_go_fuzz_dep_.CoverTab[6513]++

//line /usr/local/go/src/math/big/prime.go:190
			return len(n) == 1 && func() bool {
//line /usr/local/go/src/math/big/prime.go:190
				_go_fuzz_dep_.CoverTab[6514]++
//line /usr/local/go/src/math/big/prime.go:190
				return n[0] == p+2
//line /usr/local/go/src/math/big/prime.go:190
				// _ = "end of CoverTab[6514]"
//line /usr/local/go/src/math/big/prime.go:190
			}()
//line /usr/local/go/src/math/big/prime.go:190
			// _ = "end of CoverTab[6513]"
		} else {
//line /usr/local/go/src/math/big/prime.go:191
			_go_fuzz_dep_.CoverTab[6515]++
//line /usr/local/go/src/math/big/prime.go:191
			// _ = "end of CoverTab[6515]"
//line /usr/local/go/src/math/big/prime.go:191
		}
//line /usr/local/go/src/math/big/prime.go:191
		// _ = "end of CoverTab[6507]"
//line /usr/local/go/src/math/big/prime.go:191
		_go_fuzz_dep_.CoverTab[6508]++
							if p == 40 {
//line /usr/local/go/src/math/big/prime.go:192
			_go_fuzz_dep_.CoverTab[6516]++

//line /usr/local/go/src/math/big/prime.go:196
			t1 = t1.sqrt(n)
			t1 = t1.sqr(t1)
			if t1.cmp(n) == 0 {
//line /usr/local/go/src/math/big/prime.go:198
				_go_fuzz_dep_.CoverTab[6517]++
									return false
//line /usr/local/go/src/math/big/prime.go:199
				// _ = "end of CoverTab[6517]"
			} else {
//line /usr/local/go/src/math/big/prime.go:200
				_go_fuzz_dep_.CoverTab[6518]++
//line /usr/local/go/src/math/big/prime.go:200
				// _ = "end of CoverTab[6518]"
//line /usr/local/go/src/math/big/prime.go:200
			}
//line /usr/local/go/src/math/big/prime.go:200
			// _ = "end of CoverTab[6516]"
		} else {
//line /usr/local/go/src/math/big/prime.go:201
			_go_fuzz_dep_.CoverTab[6519]++
//line /usr/local/go/src/math/big/prime.go:201
			// _ = "end of CoverTab[6519]"
//line /usr/local/go/src/math/big/prime.go:201
		}
//line /usr/local/go/src/math/big/prime.go:201
		// _ = "end of CoverTab[6508]"
	}
//line /usr/local/go/src/math/big/prime.go:202
	// _ = "end of CoverTab[6495]"
//line /usr/local/go/src/math/big/prime.go:202
	_go_fuzz_dep_.CoverTab[6496]++

//line /usr/local/go/src/math/big/prime.go:216
	s := nat(nil).add(n, natOne)
						r := int(s.trailingZeroBits())
						s = s.shr(s, uint(r))
						nm2 := nat(nil).sub(n, natTwo)

//line /usr/local/go/src/math/big/prime.go:249
	natP := nat(nil).setWord(p)
	vk := nat(nil).setWord(2)
	vk1 := nat(nil).setWord(p)
	t2 := nat(nil)
	for i := int(s.bitLen()); i >= 0; i-- {
//line /usr/local/go/src/math/big/prime.go:253
		_go_fuzz_dep_.CoverTab[6520]++
							if s.bit(uint(i)) != 0 {
//line /usr/local/go/src/math/big/prime.go:254
			_go_fuzz_dep_.CoverTab[6521]++

//line /usr/local/go/src/math/big/prime.go:257
			t1 = t1.mul(vk, vk1)
								t1 = t1.add(t1, n)
								t1 = t1.sub(t1, natP)
								t2, vk = t2.div(vk, t1, n)

								t1 = t1.sqr(vk1)
								t1 = t1.add(t1, nm2)
								t2, vk1 = t2.div(vk1, t1, n)
//line /usr/local/go/src/math/big/prime.go:264
			// _ = "end of CoverTab[6521]"
		} else {
//line /usr/local/go/src/math/big/prime.go:265
			_go_fuzz_dep_.CoverTab[6522]++

//line /usr/local/go/src/math/big/prime.go:268
			t1 = t1.mul(vk, vk1)
								t1 = t1.add(t1, n)
								t1 = t1.sub(t1, natP)
								t2, vk1 = t2.div(vk1, t1, n)

								t1 = t1.sqr(vk)
								t1 = t1.add(t1, nm2)
								t2, vk = t2.div(vk, t1, n)
//line /usr/local/go/src/math/big/prime.go:275
			// _ = "end of CoverTab[6522]"
		}
//line /usr/local/go/src/math/big/prime.go:276
		// _ = "end of CoverTab[6520]"
	}
//line /usr/local/go/src/math/big/prime.go:277
	// _ = "end of CoverTab[6496]"
//line /usr/local/go/src/math/big/prime.go:277
	_go_fuzz_dep_.CoverTab[6497]++

//line /usr/local/go/src/math/big/prime.go:280
	if vk.cmp(natTwo) == 0 || func() bool {
//line /usr/local/go/src/math/big/prime.go:280
		_go_fuzz_dep_.CoverTab[6523]++
//line /usr/local/go/src/math/big/prime.go:280
		return vk.cmp(nm2) == 0
//line /usr/local/go/src/math/big/prime.go:280
		// _ = "end of CoverTab[6523]"
//line /usr/local/go/src/math/big/prime.go:280
	}() {
//line /usr/local/go/src/math/big/prime.go:280
		_go_fuzz_dep_.CoverTab[6524]++

//line /usr/local/go/src/math/big/prime.go:288
		t1 := t1.mul(vk, natP)
		t2 := t2.shl(vk1, 1)
		if t1.cmp(t2) < 0 {
//line /usr/local/go/src/math/big/prime.go:290
			_go_fuzz_dep_.CoverTab[6526]++
								t1, t2 = t2, t1
//line /usr/local/go/src/math/big/prime.go:291
			// _ = "end of CoverTab[6526]"
		} else {
//line /usr/local/go/src/math/big/prime.go:292
			_go_fuzz_dep_.CoverTab[6527]++
//line /usr/local/go/src/math/big/prime.go:292
			// _ = "end of CoverTab[6527]"
//line /usr/local/go/src/math/big/prime.go:292
		}
//line /usr/local/go/src/math/big/prime.go:292
		// _ = "end of CoverTab[6524]"
//line /usr/local/go/src/math/big/prime.go:292
		_go_fuzz_dep_.CoverTab[6525]++
							t1 = t1.sub(t1, t2)
							t3 := vk1
							vk1 = nil
							_ = vk1
							t2, t3 = t2.div(t3, t1, n)
							if len(t3) == 0 {
//line /usr/local/go/src/math/big/prime.go:298
			_go_fuzz_dep_.CoverTab[6528]++
								return true
//line /usr/local/go/src/math/big/prime.go:299
			// _ = "end of CoverTab[6528]"
		} else {
//line /usr/local/go/src/math/big/prime.go:300
			_go_fuzz_dep_.CoverTab[6529]++
//line /usr/local/go/src/math/big/prime.go:300
			// _ = "end of CoverTab[6529]"
//line /usr/local/go/src/math/big/prime.go:300
		}
//line /usr/local/go/src/math/big/prime.go:300
		// _ = "end of CoverTab[6525]"
	} else {
//line /usr/local/go/src/math/big/prime.go:301
		_go_fuzz_dep_.CoverTab[6530]++
//line /usr/local/go/src/math/big/prime.go:301
		// _ = "end of CoverTab[6530]"
//line /usr/local/go/src/math/big/prime.go:301
	}
//line /usr/local/go/src/math/big/prime.go:301
	// _ = "end of CoverTab[6497]"
//line /usr/local/go/src/math/big/prime.go:301
	_go_fuzz_dep_.CoverTab[6498]++

//line /usr/local/go/src/math/big/prime.go:304
	for t := 0; t < r-1; t++ {
//line /usr/local/go/src/math/big/prime.go:304
		_go_fuzz_dep_.CoverTab[6531]++
							if len(vk) == 0 {
//line /usr/local/go/src/math/big/prime.go:305
			_go_fuzz_dep_.CoverTab[6534]++
								return true
//line /usr/local/go/src/math/big/prime.go:306
			// _ = "end of CoverTab[6534]"
		} else {
//line /usr/local/go/src/math/big/prime.go:307
			_go_fuzz_dep_.CoverTab[6535]++
//line /usr/local/go/src/math/big/prime.go:307
			// _ = "end of CoverTab[6535]"
//line /usr/local/go/src/math/big/prime.go:307
		}
//line /usr/local/go/src/math/big/prime.go:307
		// _ = "end of CoverTab[6531]"
//line /usr/local/go/src/math/big/prime.go:307
		_go_fuzz_dep_.CoverTab[6532]++

//line /usr/local/go/src/math/big/prime.go:310
		if len(vk) == 1 && func() bool {
//line /usr/local/go/src/math/big/prime.go:310
			_go_fuzz_dep_.CoverTab[6536]++
//line /usr/local/go/src/math/big/prime.go:310
			return vk[0] == 2
//line /usr/local/go/src/math/big/prime.go:310
			// _ = "end of CoverTab[6536]"
//line /usr/local/go/src/math/big/prime.go:310
		}() {
//line /usr/local/go/src/math/big/prime.go:310
			_go_fuzz_dep_.CoverTab[6537]++
								return false
//line /usr/local/go/src/math/big/prime.go:311
			// _ = "end of CoverTab[6537]"
		} else {
//line /usr/local/go/src/math/big/prime.go:312
			_go_fuzz_dep_.CoverTab[6538]++
//line /usr/local/go/src/math/big/prime.go:312
			// _ = "end of CoverTab[6538]"
//line /usr/local/go/src/math/big/prime.go:312
		}
//line /usr/local/go/src/math/big/prime.go:312
		// _ = "end of CoverTab[6532]"
//line /usr/local/go/src/math/big/prime.go:312
		_go_fuzz_dep_.CoverTab[6533]++

//line /usr/local/go/src/math/big/prime.go:315
		t1 = t1.sqr(vk)
							t1 = t1.sub(t1, natTwo)
							t2, vk = t2.div(vk, t1, n)
//line /usr/local/go/src/math/big/prime.go:317
		// _ = "end of CoverTab[6533]"
	}
//line /usr/local/go/src/math/big/prime.go:318
	// _ = "end of CoverTab[6498]"
//line /usr/local/go/src/math/big/prime.go:318
	_go_fuzz_dep_.CoverTab[6499]++
						return false
//line /usr/local/go/src/math/big/prime.go:319
	// _ = "end of CoverTab[6499]"
}

//line /usr/local/go/src/math/big/prime.go:320
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/math/big/prime.go:320
var _ = _go_fuzz_dep_.CoverTab
