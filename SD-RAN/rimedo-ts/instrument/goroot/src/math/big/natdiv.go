// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/math/big/natdiv.go:499
package big

//line /usr/local/go/src/math/big/natdiv.go:499
import (
//line /usr/local/go/src/math/big/natdiv.go:499
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/math/big/natdiv.go:499
)
//line /usr/local/go/src/math/big/natdiv.go:499
import (
//line /usr/local/go/src/math/big/natdiv.go:499
	_atomic_ "sync/atomic"
//line /usr/local/go/src/math/big/natdiv.go:499
)

import "math/bits"

// rem returns r such that r = u%v.
//line /usr/local/go/src/math/big/natdiv.go:503
// It uses z as the storage for r.
//line /usr/local/go/src/math/big/natdiv.go:505
func (z nat) rem(u, v nat) (r nat) {
//line /usr/local/go/src/math/big/natdiv.go:505
	_go_fuzz_dep_.CoverTab[6323]++
							if alias(z, u) {
//line /usr/local/go/src/math/big/natdiv.go:506
		_go_fuzz_dep_.CoverTab[6325]++
								z = nil
//line /usr/local/go/src/math/big/natdiv.go:507
		// _ = "end of CoverTab[6325]"
	} else {
//line /usr/local/go/src/math/big/natdiv.go:508
		_go_fuzz_dep_.CoverTab[6326]++
//line /usr/local/go/src/math/big/natdiv.go:508
		// _ = "end of CoverTab[6326]"
//line /usr/local/go/src/math/big/natdiv.go:508
	}
//line /usr/local/go/src/math/big/natdiv.go:508
	// _ = "end of CoverTab[6323]"
//line /usr/local/go/src/math/big/natdiv.go:508
	_go_fuzz_dep_.CoverTab[6324]++
							qp := getNat(0)
							q, r := qp.div(z, u, v)
							*qp = q
							putNat(qp)
							return r
//line /usr/local/go/src/math/big/natdiv.go:513
	// _ = "end of CoverTab[6324]"
}

// div returns q, r such that q = ⌊u/v⌋ and r = u%v = u - q·v.
//line /usr/local/go/src/math/big/natdiv.go:516
// It uses z and z2 as the storage for q and r.
//line /usr/local/go/src/math/big/natdiv.go:518
func (z nat) div(z2, u, v nat) (q, r nat) {
//line /usr/local/go/src/math/big/natdiv.go:518
	_go_fuzz_dep_.CoverTab[6327]++
							if len(v) == 0 {
//line /usr/local/go/src/math/big/natdiv.go:519
		_go_fuzz_dep_.CoverTab[6331]++
								panic("division by zero")
//line /usr/local/go/src/math/big/natdiv.go:520
		// _ = "end of CoverTab[6331]"
	} else {
//line /usr/local/go/src/math/big/natdiv.go:521
		_go_fuzz_dep_.CoverTab[6332]++
//line /usr/local/go/src/math/big/natdiv.go:521
		// _ = "end of CoverTab[6332]"
//line /usr/local/go/src/math/big/natdiv.go:521
	}
//line /usr/local/go/src/math/big/natdiv.go:521
	// _ = "end of CoverTab[6327]"
//line /usr/local/go/src/math/big/natdiv.go:521
	_go_fuzz_dep_.CoverTab[6328]++

							if u.cmp(v) < 0 {
//line /usr/local/go/src/math/big/natdiv.go:523
		_go_fuzz_dep_.CoverTab[6333]++
								q = z[:0]
								r = z2.set(u)
								return
//line /usr/local/go/src/math/big/natdiv.go:526
		// _ = "end of CoverTab[6333]"
	} else {
//line /usr/local/go/src/math/big/natdiv.go:527
		_go_fuzz_dep_.CoverTab[6334]++
//line /usr/local/go/src/math/big/natdiv.go:527
		// _ = "end of CoverTab[6334]"
//line /usr/local/go/src/math/big/natdiv.go:527
	}
//line /usr/local/go/src/math/big/natdiv.go:527
	// _ = "end of CoverTab[6328]"
//line /usr/local/go/src/math/big/natdiv.go:527
	_go_fuzz_dep_.CoverTab[6329]++

							if len(v) == 1 {
//line /usr/local/go/src/math/big/natdiv.go:529
		_go_fuzz_dep_.CoverTab[6335]++
		// Short division: long optimized for a single-word divisor.
								// In that case, the 2-by-1 guess is all we need at each step.
								var r2 Word
								q, r2 = z.divW(u, v[0])
								r = z2.setWord(r2)
								return
//line /usr/local/go/src/math/big/natdiv.go:535
		// _ = "end of CoverTab[6335]"
	} else {
//line /usr/local/go/src/math/big/natdiv.go:536
		_go_fuzz_dep_.CoverTab[6336]++
//line /usr/local/go/src/math/big/natdiv.go:536
		// _ = "end of CoverTab[6336]"
//line /usr/local/go/src/math/big/natdiv.go:536
	}
//line /usr/local/go/src/math/big/natdiv.go:536
	// _ = "end of CoverTab[6329]"
//line /usr/local/go/src/math/big/natdiv.go:536
	_go_fuzz_dep_.CoverTab[6330]++

							q, r = z.divLarge(z2, u, v)
							return
//line /usr/local/go/src/math/big/natdiv.go:539
	// _ = "end of CoverTab[6330]"
}

// divW returns q, r such that q = ⌊x/y⌋ and r = x%y = x - q·y.
//line /usr/local/go/src/math/big/natdiv.go:542
// It uses z as the storage for q.
//line /usr/local/go/src/math/big/natdiv.go:542
// Note that y is a single digit (Word), not a big number.
//line /usr/local/go/src/math/big/natdiv.go:545
func (z nat) divW(x nat, y Word) (q nat, r Word) {
//line /usr/local/go/src/math/big/natdiv.go:545
	_go_fuzz_dep_.CoverTab[6337]++
							m := len(x)
							switch {
	case y == 0:
//line /usr/local/go/src/math/big/natdiv.go:548
		_go_fuzz_dep_.CoverTab[6339]++
								panic("division by zero")
//line /usr/local/go/src/math/big/natdiv.go:549
		// _ = "end of CoverTab[6339]"
	case y == 1:
//line /usr/local/go/src/math/big/natdiv.go:550
		_go_fuzz_dep_.CoverTab[6340]++
								q = z.set(x)
								return
//line /usr/local/go/src/math/big/natdiv.go:552
		// _ = "end of CoverTab[6340]"
	case m == 0:
//line /usr/local/go/src/math/big/natdiv.go:553
		_go_fuzz_dep_.CoverTab[6341]++
								q = z[:0]
								return
//line /usr/local/go/src/math/big/natdiv.go:555
		// _ = "end of CoverTab[6341]"
//line /usr/local/go/src/math/big/natdiv.go:555
	default:
//line /usr/local/go/src/math/big/natdiv.go:555
		_go_fuzz_dep_.CoverTab[6342]++
//line /usr/local/go/src/math/big/natdiv.go:555
		// _ = "end of CoverTab[6342]"
	}
//line /usr/local/go/src/math/big/natdiv.go:556
	// _ = "end of CoverTab[6337]"
//line /usr/local/go/src/math/big/natdiv.go:556
	_go_fuzz_dep_.CoverTab[6338]++

							z = z.make(m)
							r = divWVW(z, 0, x, y)
							q = z.norm()
							return
//line /usr/local/go/src/math/big/natdiv.go:561
	// _ = "end of CoverTab[6338]"
}

// modW returns x % d.
func (x nat) modW(d Word) (r Word) {
//line /usr/local/go/src/math/big/natdiv.go:565
	_go_fuzz_dep_.CoverTab[6343]++
							// TODO(agl): we don't actually need to store the q value.
							var q nat
							q = q.make(len(x))
							return divWVW(q, 0, x, d)
//line /usr/local/go/src/math/big/natdiv.go:569
	// _ = "end of CoverTab[6343]"
}

// divWVW overwrites z with ⌊x/y⌋, returning the remainder r.
//line /usr/local/go/src/math/big/natdiv.go:572
// The caller must ensure that len(z) = len(x).
//line /usr/local/go/src/math/big/natdiv.go:574
func divWVW(z []Word, xn Word, x []Word, y Word) (r Word) {
//line /usr/local/go/src/math/big/natdiv.go:574
	_go_fuzz_dep_.CoverTab[6344]++
							r = xn
							if len(x) == 1 {
//line /usr/local/go/src/math/big/natdiv.go:576
		_go_fuzz_dep_.CoverTab[6347]++
								qq, rr := bits.Div(uint(r), uint(x[0]), uint(y))
								z[0] = Word(qq)
								return Word(rr)
//line /usr/local/go/src/math/big/natdiv.go:579
		// _ = "end of CoverTab[6347]"
	} else {
//line /usr/local/go/src/math/big/natdiv.go:580
		_go_fuzz_dep_.CoverTab[6348]++
//line /usr/local/go/src/math/big/natdiv.go:580
		// _ = "end of CoverTab[6348]"
//line /usr/local/go/src/math/big/natdiv.go:580
	}
//line /usr/local/go/src/math/big/natdiv.go:580
	// _ = "end of CoverTab[6344]"
//line /usr/local/go/src/math/big/natdiv.go:580
	_go_fuzz_dep_.CoverTab[6345]++
							rec := reciprocalWord(y)
							for i := len(z) - 1; i >= 0; i-- {
//line /usr/local/go/src/math/big/natdiv.go:582
		_go_fuzz_dep_.CoverTab[6349]++
								z[i], r = divWW(r, x[i], y, rec)
//line /usr/local/go/src/math/big/natdiv.go:583
		// _ = "end of CoverTab[6349]"
	}
//line /usr/local/go/src/math/big/natdiv.go:584
	// _ = "end of CoverTab[6345]"
//line /usr/local/go/src/math/big/natdiv.go:584
	_go_fuzz_dep_.CoverTab[6346]++
							return r
//line /usr/local/go/src/math/big/natdiv.go:585
	// _ = "end of CoverTab[6346]"
}

// div returns q, r such that q = ⌊uIn/vIn⌋ and r = uIn%vIn = uIn - q·vIn.
//line /usr/local/go/src/math/big/natdiv.go:588
// It uses z and u as the storage for q and r.
//line /usr/local/go/src/math/big/natdiv.go:588
// The caller must ensure that len(vIn) ≥ 2 (use divW otherwise)
//line /usr/local/go/src/math/big/natdiv.go:588
// and that len(uIn) ≥ len(vIn) (the answer is 0, uIn otherwise).
//line /usr/local/go/src/math/big/natdiv.go:592
func (z nat) divLarge(u, uIn, vIn nat) (q, r nat) {
//line /usr/local/go/src/math/big/natdiv.go:592
	_go_fuzz_dep_.CoverTab[6350]++
							n := len(vIn)
							m := len(uIn) - n

//line /usr/local/go/src/math/big/natdiv.go:600
	shift := nlz(vIn[n-1])
							vp := getNat(n)
							v := *vp
							shlVU(v, vIn, shift)
							u = u.make(len(uIn) + 1)
							u[len(uIn)] = shlVU(u[0:len(uIn)], uIn, shift)

//line /usr/local/go/src/math/big/natdiv.go:609
	if alias(z, u) {
//line /usr/local/go/src/math/big/natdiv.go:609
		_go_fuzz_dep_.CoverTab[6353]++
								z = nil
//line /usr/local/go/src/math/big/natdiv.go:610
		// _ = "end of CoverTab[6353]"
	} else {
//line /usr/local/go/src/math/big/natdiv.go:611
		_go_fuzz_dep_.CoverTab[6354]++
//line /usr/local/go/src/math/big/natdiv.go:611
		// _ = "end of CoverTab[6354]"
//line /usr/local/go/src/math/big/natdiv.go:611
	}
//line /usr/local/go/src/math/big/natdiv.go:611
	// _ = "end of CoverTab[6350]"
//line /usr/local/go/src/math/big/natdiv.go:611
	_go_fuzz_dep_.CoverTab[6351]++
							q = z.make(m + 1)

//line /usr/local/go/src/math/big/natdiv.go:615
	if n < divRecursiveThreshold {
//line /usr/local/go/src/math/big/natdiv.go:615
		_go_fuzz_dep_.CoverTab[6355]++
								q.divBasic(u, v)
//line /usr/local/go/src/math/big/natdiv.go:616
		// _ = "end of CoverTab[6355]"
	} else {
//line /usr/local/go/src/math/big/natdiv.go:617
		_go_fuzz_dep_.CoverTab[6356]++
								q.divRecursive(u, v)
//line /usr/local/go/src/math/big/natdiv.go:618
		// _ = "end of CoverTab[6356]"
	}
//line /usr/local/go/src/math/big/natdiv.go:619
	// _ = "end of CoverTab[6351]"
//line /usr/local/go/src/math/big/natdiv.go:619
	_go_fuzz_dep_.CoverTab[6352]++
							putNat(vp)

							q = q.norm()

//line /usr/local/go/src/math/big/natdiv.go:625
	shrVU(u, u, shift)
							r = u.norm()

							return q, r
//line /usr/local/go/src/math/big/natdiv.go:628
	// _ = "end of CoverTab[6352]"
}

// divBasic implements long division as described above.
//line /usr/local/go/src/math/big/natdiv.go:631
// It overwrites q with ⌊u/v⌋ and overwrites u with the remainder r.
//line /usr/local/go/src/math/big/natdiv.go:631
// q must be large enough to hold ⌊u/v⌋.
//line /usr/local/go/src/math/big/natdiv.go:634
func (q nat) divBasic(u, v nat) {
//line /usr/local/go/src/math/big/natdiv.go:634
	_go_fuzz_dep_.CoverTab[6357]++
							n := len(v)
							m := len(u) - n

							qhatvp := getNat(n + 1)
							qhatv := *qhatvp

//line /usr/local/go/src/math/big/natdiv.go:642
	vn1 := v[n-1]
							rec := reciprocalWord(vn1)

//line /usr/local/go/src/math/big/natdiv.go:646
	for j := m; j >= 0; j-- {
//line /usr/local/go/src/math/big/natdiv.go:646
		_go_fuzz_dep_.CoverTab[6359]++

//line /usr/local/go/src/math/big/natdiv.go:649
		qhat := Word(_M)
		var ujn Word
		if j+n < len(u) {
//line /usr/local/go/src/math/big/natdiv.go:651
			_go_fuzz_dep_.CoverTab[6365]++
									ujn = u[j+n]
//line /usr/local/go/src/math/big/natdiv.go:652
			// _ = "end of CoverTab[6365]"
		} else {
//line /usr/local/go/src/math/big/natdiv.go:653
			_go_fuzz_dep_.CoverTab[6366]++
//line /usr/local/go/src/math/big/natdiv.go:653
			// _ = "end of CoverTab[6366]"
//line /usr/local/go/src/math/big/natdiv.go:653
		}
//line /usr/local/go/src/math/big/natdiv.go:653
		// _ = "end of CoverTab[6359]"
//line /usr/local/go/src/math/big/natdiv.go:653
		_go_fuzz_dep_.CoverTab[6360]++

//line /usr/local/go/src/math/big/natdiv.go:658
		if ujn != vn1 {
//line /usr/local/go/src/math/big/natdiv.go:658
			_go_fuzz_dep_.CoverTab[6367]++
									var rhat Word
									qhat, rhat = divWW(ujn, u[j+n-1], vn1, rec)

//line /usr/local/go/src/math/big/natdiv.go:663
			vn2 := v[n-2]
			x1, x2 := mulWW(qhat, vn2)
			ujn2 := u[j+n-2]
			for greaterThan(x1, x2, rhat, ujn2) {
//line /usr/local/go/src/math/big/natdiv.go:666
				_go_fuzz_dep_.CoverTab[6368]++
										qhat--
										prevRhat := rhat
										rhat += vn1

//line /usr/local/go/src/math/big/natdiv.go:672
				if rhat < prevRhat {
//line /usr/local/go/src/math/big/natdiv.go:672
					_go_fuzz_dep_.CoverTab[6370]++
											break
//line /usr/local/go/src/math/big/natdiv.go:673
					// _ = "end of CoverTab[6370]"
				} else {
//line /usr/local/go/src/math/big/natdiv.go:674
					_go_fuzz_dep_.CoverTab[6371]++
//line /usr/local/go/src/math/big/natdiv.go:674
					// _ = "end of CoverTab[6371]"
//line /usr/local/go/src/math/big/natdiv.go:674
				}
//line /usr/local/go/src/math/big/natdiv.go:674
				// _ = "end of CoverTab[6368]"
//line /usr/local/go/src/math/big/natdiv.go:674
				_go_fuzz_dep_.CoverTab[6369]++

//line /usr/local/go/src/math/big/natdiv.go:677
				x1, x2 = mulWW(qhat, vn2)
//line /usr/local/go/src/math/big/natdiv.go:677
				// _ = "end of CoverTab[6369]"
			}
//line /usr/local/go/src/math/big/natdiv.go:678
			// _ = "end of CoverTab[6367]"
		} else {
//line /usr/local/go/src/math/big/natdiv.go:679
			_go_fuzz_dep_.CoverTab[6372]++
//line /usr/local/go/src/math/big/natdiv.go:679
			// _ = "end of CoverTab[6372]"
//line /usr/local/go/src/math/big/natdiv.go:679
		}
//line /usr/local/go/src/math/big/natdiv.go:679
		// _ = "end of CoverTab[6360]"
//line /usr/local/go/src/math/big/natdiv.go:679
		_go_fuzz_dep_.CoverTab[6361]++

//line /usr/local/go/src/math/big/natdiv.go:682
		qhatv[n] = mulAddVWW(qhatv[0:n], v, qhat, 0)
		qhl := len(qhatv)
		if j+qhl > len(u) && func() bool {
//line /usr/local/go/src/math/big/natdiv.go:684
			_go_fuzz_dep_.CoverTab[6373]++
//line /usr/local/go/src/math/big/natdiv.go:684
			return qhatv[n] == 0
//line /usr/local/go/src/math/big/natdiv.go:684
			// _ = "end of CoverTab[6373]"
//line /usr/local/go/src/math/big/natdiv.go:684
		}() {
//line /usr/local/go/src/math/big/natdiv.go:684
			_go_fuzz_dep_.CoverTab[6374]++
									qhl--
//line /usr/local/go/src/math/big/natdiv.go:685
			// _ = "end of CoverTab[6374]"
		} else {
//line /usr/local/go/src/math/big/natdiv.go:686
			_go_fuzz_dep_.CoverTab[6375]++
//line /usr/local/go/src/math/big/natdiv.go:686
			// _ = "end of CoverTab[6375]"
//line /usr/local/go/src/math/big/natdiv.go:686
		}
//line /usr/local/go/src/math/big/natdiv.go:686
		// _ = "end of CoverTab[6361]"
//line /usr/local/go/src/math/big/natdiv.go:686
		_go_fuzz_dep_.CoverTab[6362]++

//line /usr/local/go/src/math/big/natdiv.go:691
		c := subVV(u[j:j+qhl], u[j:], qhatv)
		if c != 0 {
//line /usr/local/go/src/math/big/natdiv.go:692
			_go_fuzz_dep_.CoverTab[6376]++
									c := addVV(u[j:j+n], u[j:], v)

//line /usr/local/go/src/math/big/natdiv.go:696
			if n < qhl {
//line /usr/local/go/src/math/big/natdiv.go:696
				_go_fuzz_dep_.CoverTab[6378]++
										u[j+n] += c
//line /usr/local/go/src/math/big/natdiv.go:697
				// _ = "end of CoverTab[6378]"
			} else {
//line /usr/local/go/src/math/big/natdiv.go:698
				_go_fuzz_dep_.CoverTab[6379]++
//line /usr/local/go/src/math/big/natdiv.go:698
				// _ = "end of CoverTab[6379]"
//line /usr/local/go/src/math/big/natdiv.go:698
			}
//line /usr/local/go/src/math/big/natdiv.go:698
			// _ = "end of CoverTab[6376]"
//line /usr/local/go/src/math/big/natdiv.go:698
			_go_fuzz_dep_.CoverTab[6377]++
									qhat--
//line /usr/local/go/src/math/big/natdiv.go:699
			// _ = "end of CoverTab[6377]"
		} else {
//line /usr/local/go/src/math/big/natdiv.go:700
			_go_fuzz_dep_.CoverTab[6380]++
//line /usr/local/go/src/math/big/natdiv.go:700
			// _ = "end of CoverTab[6380]"
//line /usr/local/go/src/math/big/natdiv.go:700
		}
//line /usr/local/go/src/math/big/natdiv.go:700
		// _ = "end of CoverTab[6362]"
//line /usr/local/go/src/math/big/natdiv.go:700
		_go_fuzz_dep_.CoverTab[6363]++

//line /usr/local/go/src/math/big/natdiv.go:704
		if j == m && func() bool {
//line /usr/local/go/src/math/big/natdiv.go:704
			_go_fuzz_dep_.CoverTab[6381]++
//line /usr/local/go/src/math/big/natdiv.go:704
			return m == len(q)
//line /usr/local/go/src/math/big/natdiv.go:704
			// _ = "end of CoverTab[6381]"
//line /usr/local/go/src/math/big/natdiv.go:704
		}() && func() bool {
//line /usr/local/go/src/math/big/natdiv.go:704
			_go_fuzz_dep_.CoverTab[6382]++
//line /usr/local/go/src/math/big/natdiv.go:704
			return qhat == 0
//line /usr/local/go/src/math/big/natdiv.go:704
			// _ = "end of CoverTab[6382]"
//line /usr/local/go/src/math/big/natdiv.go:704
		}() {
//line /usr/local/go/src/math/big/natdiv.go:704
			_go_fuzz_dep_.CoverTab[6383]++
									continue
//line /usr/local/go/src/math/big/natdiv.go:705
			// _ = "end of CoverTab[6383]"
		} else {
//line /usr/local/go/src/math/big/natdiv.go:706
			_go_fuzz_dep_.CoverTab[6384]++
//line /usr/local/go/src/math/big/natdiv.go:706
			// _ = "end of CoverTab[6384]"
//line /usr/local/go/src/math/big/natdiv.go:706
		}
//line /usr/local/go/src/math/big/natdiv.go:706
		// _ = "end of CoverTab[6363]"
//line /usr/local/go/src/math/big/natdiv.go:706
		_go_fuzz_dep_.CoverTab[6364]++
								q[j] = qhat
//line /usr/local/go/src/math/big/natdiv.go:707
		// _ = "end of CoverTab[6364]"
	}
//line /usr/local/go/src/math/big/natdiv.go:708
	// _ = "end of CoverTab[6357]"
//line /usr/local/go/src/math/big/natdiv.go:708
	_go_fuzz_dep_.CoverTab[6358]++

							putNat(qhatvp)
//line /usr/local/go/src/math/big/natdiv.go:710
	// _ = "end of CoverTab[6358]"
}

// greaterThan reports whether the two digit numbers x1 x2 > y1 y2.
//line /usr/local/go/src/math/big/natdiv.go:713
// TODO(rsc): In contradiction to most of this file, x1 is the high
//line /usr/local/go/src/math/big/natdiv.go:713
// digit and x2 is the low digit. This should be fixed.
//line /usr/local/go/src/math/big/natdiv.go:716
func greaterThan(x1, x2, y1, y2 Word) bool {
//line /usr/local/go/src/math/big/natdiv.go:716
	_go_fuzz_dep_.CoverTab[6385]++
							return x1 > y1 || func() bool {
//line /usr/local/go/src/math/big/natdiv.go:717
		_go_fuzz_dep_.CoverTab[6386]++
//line /usr/local/go/src/math/big/natdiv.go:717
		return x1 == y1 && func() bool {
//line /usr/local/go/src/math/big/natdiv.go:717
			_go_fuzz_dep_.CoverTab[6387]++
//line /usr/local/go/src/math/big/natdiv.go:717
			return x2 > y2
//line /usr/local/go/src/math/big/natdiv.go:717
			// _ = "end of CoverTab[6387]"
//line /usr/local/go/src/math/big/natdiv.go:717
		}()
//line /usr/local/go/src/math/big/natdiv.go:717
		// _ = "end of CoverTab[6386]"
//line /usr/local/go/src/math/big/natdiv.go:717
	}()
//line /usr/local/go/src/math/big/natdiv.go:717
	// _ = "end of CoverTab[6385]"
}

// divRecursiveThreshold is the number of divisor digits
//line /usr/local/go/src/math/big/natdiv.go:720
// at which point divRecursive is faster than divBasic.
//line /usr/local/go/src/math/big/natdiv.go:722
const divRecursiveThreshold = 100

// divRecursive implements recursive division as described above.
//line /usr/local/go/src/math/big/natdiv.go:724
// It overwrites z with ⌊u/v⌋ and overwrites u with the remainder r.
//line /usr/local/go/src/math/big/natdiv.go:724
// z must be large enough to hold ⌊u/v⌋.
//line /usr/local/go/src/math/big/natdiv.go:724
// This function is just for allocating and freeing temporaries
//line /usr/local/go/src/math/big/natdiv.go:724
// around divRecursiveStep, the real implementation.
//line /usr/local/go/src/math/big/natdiv.go:729
func (z nat) divRecursive(u, v nat) {
//line /usr/local/go/src/math/big/natdiv.go:729
	_go_fuzz_dep_.CoverTab[6388]++

//line /usr/local/go/src/math/big/natdiv.go:733
	recDepth := 2 * bits.Len(uint(len(v)))
							tmp := getNat(3 * len(v))
							temps := make([]*nat, recDepth)

							z.clear()
							z.divRecursiveStep(u, v, 0, tmp, temps)

//line /usr/local/go/src/math/big/natdiv.go:741
	for _, n := range temps {
//line /usr/local/go/src/math/big/natdiv.go:741
		_go_fuzz_dep_.CoverTab[6390]++
								if n != nil {
//line /usr/local/go/src/math/big/natdiv.go:742
			_go_fuzz_dep_.CoverTab[6391]++
									putNat(n)
//line /usr/local/go/src/math/big/natdiv.go:743
			// _ = "end of CoverTab[6391]"
		} else {
//line /usr/local/go/src/math/big/natdiv.go:744
			_go_fuzz_dep_.CoverTab[6392]++
//line /usr/local/go/src/math/big/natdiv.go:744
			// _ = "end of CoverTab[6392]"
//line /usr/local/go/src/math/big/natdiv.go:744
		}
//line /usr/local/go/src/math/big/natdiv.go:744
		// _ = "end of CoverTab[6390]"
	}
//line /usr/local/go/src/math/big/natdiv.go:745
	// _ = "end of CoverTab[6388]"
//line /usr/local/go/src/math/big/natdiv.go:745
	_go_fuzz_dep_.CoverTab[6389]++
							putNat(tmp)
//line /usr/local/go/src/math/big/natdiv.go:746
	// _ = "end of CoverTab[6389]"
}

// divRecursiveStep is the actual implementation of recursive division.
//line /usr/local/go/src/math/big/natdiv.go:749
// It adds ⌊u/v⌋ to z and overwrites u with the remainder r.
//line /usr/local/go/src/math/big/natdiv.go:749
// z must be large enough to hold ⌊u/v⌋.
//line /usr/local/go/src/math/big/natdiv.go:749
// It uses temps[depth] (allocating if needed) as a temporary live across
//line /usr/local/go/src/math/big/natdiv.go:749
// the recursive call. It also uses tmp, but not live across the recursion.
//line /usr/local/go/src/math/big/natdiv.go:754
func (z nat) divRecursiveStep(u, v nat, depth int, tmp *nat, temps []*nat) {
//line /usr/local/go/src/math/big/natdiv.go:754
	_go_fuzz_dep_.CoverTab[6393]++

//line /usr/local/go/src/math/big/natdiv.go:758
	u = u.norm()
	v = v.norm()
	if len(u) == 0 {
//line /usr/local/go/src/math/big/natdiv.go:760
		_go_fuzz_dep_.CoverTab[6403]++
								z.clear()
								return
//line /usr/local/go/src/math/big/natdiv.go:762
		// _ = "end of CoverTab[6403]"
	} else {
//line /usr/local/go/src/math/big/natdiv.go:763
		_go_fuzz_dep_.CoverTab[6404]++
//line /usr/local/go/src/math/big/natdiv.go:763
		// _ = "end of CoverTab[6404]"
//line /usr/local/go/src/math/big/natdiv.go:763
	}
//line /usr/local/go/src/math/big/natdiv.go:763
	// _ = "end of CoverTab[6393]"
//line /usr/local/go/src/math/big/natdiv.go:763
	_go_fuzz_dep_.CoverTab[6394]++

//line /usr/local/go/src/math/big/natdiv.go:766
	n := len(v)
	if n < divRecursiveThreshold {
//line /usr/local/go/src/math/big/natdiv.go:767
		_go_fuzz_dep_.CoverTab[6405]++
								z.divBasic(u, v)
								return
//line /usr/local/go/src/math/big/natdiv.go:769
		// _ = "end of CoverTab[6405]"
	} else {
//line /usr/local/go/src/math/big/natdiv.go:770
		_go_fuzz_dep_.CoverTab[6406]++
//line /usr/local/go/src/math/big/natdiv.go:770
		// _ = "end of CoverTab[6406]"
//line /usr/local/go/src/math/big/natdiv.go:770
	}
//line /usr/local/go/src/math/big/natdiv.go:770
	// _ = "end of CoverTab[6394]"
//line /usr/local/go/src/math/big/natdiv.go:770
	_go_fuzz_dep_.CoverTab[6395]++

//line /usr/local/go/src/math/big/natdiv.go:773
	m := len(u) - n
	if m < 0 {
//line /usr/local/go/src/math/big/natdiv.go:774
		_go_fuzz_dep_.CoverTab[6407]++
								return
//line /usr/local/go/src/math/big/natdiv.go:775
		// _ = "end of CoverTab[6407]"
	} else {
//line /usr/local/go/src/math/big/natdiv.go:776
		_go_fuzz_dep_.CoverTab[6408]++
//line /usr/local/go/src/math/big/natdiv.go:776
		// _ = "end of CoverTab[6408]"
//line /usr/local/go/src/math/big/natdiv.go:776
	}
//line /usr/local/go/src/math/big/natdiv.go:776
	// _ = "end of CoverTab[6395]"
//line /usr/local/go/src/math/big/natdiv.go:776
	_go_fuzz_dep_.CoverTab[6396]++

//line /usr/local/go/src/math/big/natdiv.go:784
	B := n / 2

//line /usr/local/go/src/math/big/natdiv.go:787
	if temps[depth] == nil {
//line /usr/local/go/src/math/big/natdiv.go:787
		_go_fuzz_dep_.CoverTab[6409]++
								temps[depth] = getNat(n)
//line /usr/local/go/src/math/big/natdiv.go:788
		// _ = "end of CoverTab[6409]"
	} else {
//line /usr/local/go/src/math/big/natdiv.go:789
		_go_fuzz_dep_.CoverTab[6410]++
								*temps[depth] = temps[depth].make(B + 1)
//line /usr/local/go/src/math/big/natdiv.go:790
		// _ = "end of CoverTab[6410]"
	}
//line /usr/local/go/src/math/big/natdiv.go:791
	// _ = "end of CoverTab[6396]"
//line /usr/local/go/src/math/big/natdiv.go:791
	_go_fuzz_dep_.CoverTab[6397]++

//line /usr/local/go/src/math/big/natdiv.go:799
	j := m
	for j > B {
//line /usr/local/go/src/math/big/natdiv.go:800
		_go_fuzz_dep_.CoverTab[6411]++

//line /usr/local/go/src/math/big/natdiv.go:812
		s := (B - 1)

//line /usr/local/go/src/math/big/natdiv.go:815
		uu := u[j-B:]

//line /usr/local/go/src/math/big/natdiv.go:818
		qhat := *temps[depth]
								qhat.clear()
								qhat.divRecursiveStep(uu[s:B+n], v[s:], depth+1, tmp, temps)
								qhat = qhat.norm()

//line /usr/local/go/src/math/big/natdiv.go:835
		qhatv := tmp.make(3 * n)
		qhatv.clear()
		qhatv = qhatv.mul(qhat, v[:s])
		for i := 0; i < 2; i++ {
//line /usr/local/go/src/math/big/natdiv.go:838
			_go_fuzz_dep_.CoverTab[6415]++
									e := qhatv.cmp(uu.norm())
									if e <= 0 {
//line /usr/local/go/src/math/big/natdiv.go:840
				_go_fuzz_dep_.CoverTab[6418]++
										break
//line /usr/local/go/src/math/big/natdiv.go:841
				// _ = "end of CoverTab[6418]"
			} else {
//line /usr/local/go/src/math/big/natdiv.go:842
				_go_fuzz_dep_.CoverTab[6419]++
//line /usr/local/go/src/math/big/natdiv.go:842
				// _ = "end of CoverTab[6419]"
//line /usr/local/go/src/math/big/natdiv.go:842
			}
//line /usr/local/go/src/math/big/natdiv.go:842
			// _ = "end of CoverTab[6415]"
//line /usr/local/go/src/math/big/natdiv.go:842
			_go_fuzz_dep_.CoverTab[6416]++
									subVW(qhat, qhat, 1)
									c := subVV(qhatv[:s], qhatv[:s], v[:s])
									if len(qhatv) > s {
//line /usr/local/go/src/math/big/natdiv.go:845
				_go_fuzz_dep_.CoverTab[6420]++
										subVW(qhatv[s:], qhatv[s:], c)
//line /usr/local/go/src/math/big/natdiv.go:846
				// _ = "end of CoverTab[6420]"
			} else {
//line /usr/local/go/src/math/big/natdiv.go:847
				_go_fuzz_dep_.CoverTab[6421]++
//line /usr/local/go/src/math/big/natdiv.go:847
				// _ = "end of CoverTab[6421]"
//line /usr/local/go/src/math/big/natdiv.go:847
			}
//line /usr/local/go/src/math/big/natdiv.go:847
			// _ = "end of CoverTab[6416]"
//line /usr/local/go/src/math/big/natdiv.go:847
			_go_fuzz_dep_.CoverTab[6417]++
									addAt(uu[s:], v[s:], 0)
//line /usr/local/go/src/math/big/natdiv.go:848
			// _ = "end of CoverTab[6417]"
		}
//line /usr/local/go/src/math/big/natdiv.go:849
		// _ = "end of CoverTab[6411]"
//line /usr/local/go/src/math/big/natdiv.go:849
		_go_fuzz_dep_.CoverTab[6412]++
								if qhatv.cmp(uu.norm()) > 0 {
//line /usr/local/go/src/math/big/natdiv.go:850
			_go_fuzz_dep_.CoverTab[6422]++
									panic("impossible")
//line /usr/local/go/src/math/big/natdiv.go:851
			// _ = "end of CoverTab[6422]"
		} else {
//line /usr/local/go/src/math/big/natdiv.go:852
			_go_fuzz_dep_.CoverTab[6423]++
//line /usr/local/go/src/math/big/natdiv.go:852
			// _ = "end of CoverTab[6423]"
//line /usr/local/go/src/math/big/natdiv.go:852
		}
//line /usr/local/go/src/math/big/natdiv.go:852
		// _ = "end of CoverTab[6412]"
//line /usr/local/go/src/math/big/natdiv.go:852
		_go_fuzz_dep_.CoverTab[6413]++
								c := subVV(uu[:len(qhatv)], uu[:len(qhatv)], qhatv)
								if c > 0 {
//line /usr/local/go/src/math/big/natdiv.go:854
			_go_fuzz_dep_.CoverTab[6424]++
									subVW(uu[len(qhatv):], uu[len(qhatv):], c)
//line /usr/local/go/src/math/big/natdiv.go:855
			// _ = "end of CoverTab[6424]"
		} else {
//line /usr/local/go/src/math/big/natdiv.go:856
			_go_fuzz_dep_.CoverTab[6425]++
//line /usr/local/go/src/math/big/natdiv.go:856
			// _ = "end of CoverTab[6425]"
//line /usr/local/go/src/math/big/natdiv.go:856
		}
//line /usr/local/go/src/math/big/natdiv.go:856
		// _ = "end of CoverTab[6413]"
//line /usr/local/go/src/math/big/natdiv.go:856
		_go_fuzz_dep_.CoverTab[6414]++
								addAt(z, qhat, j-B)
								j -= B
//line /usr/local/go/src/math/big/natdiv.go:858
		// _ = "end of CoverTab[6414]"
	}
//line /usr/local/go/src/math/big/natdiv.go:859
	// _ = "end of CoverTab[6397]"
//line /usr/local/go/src/math/big/natdiv.go:859
	_go_fuzz_dep_.CoverTab[6398]++

//line /usr/local/go/src/math/big/natdiv.go:865
	s := B - 1
	qhat := *temps[depth]
	qhat.clear()
	qhat.divRecursiveStep(u[s:].norm(), v[s:], depth+1, tmp, temps)
	qhat = qhat.norm()
	qhatv := tmp.make(3 * n)
	qhatv.clear()
	qhatv = qhatv.mul(qhat, v[:s])

	for i := 0; i < 2; i++ {
//line /usr/local/go/src/math/big/natdiv.go:874
		_go_fuzz_dep_.CoverTab[6426]++
								if e := qhatv.cmp(u.norm()); e > 0 {
//line /usr/local/go/src/math/big/natdiv.go:875
			_go_fuzz_dep_.CoverTab[6427]++
									subVW(qhat, qhat, 1)
									c := subVV(qhatv[:s], qhatv[:s], v[:s])
									if len(qhatv) > s {
//line /usr/local/go/src/math/big/natdiv.go:878
				_go_fuzz_dep_.CoverTab[6429]++
										subVW(qhatv[s:], qhatv[s:], c)
//line /usr/local/go/src/math/big/natdiv.go:879
				// _ = "end of CoverTab[6429]"
			} else {
//line /usr/local/go/src/math/big/natdiv.go:880
				_go_fuzz_dep_.CoverTab[6430]++
//line /usr/local/go/src/math/big/natdiv.go:880
				// _ = "end of CoverTab[6430]"
//line /usr/local/go/src/math/big/natdiv.go:880
			}
//line /usr/local/go/src/math/big/natdiv.go:880
			// _ = "end of CoverTab[6427]"
//line /usr/local/go/src/math/big/natdiv.go:880
			_go_fuzz_dep_.CoverTab[6428]++
									addAt(u[s:], v[s:], 0)
//line /usr/local/go/src/math/big/natdiv.go:881
			// _ = "end of CoverTab[6428]"
		} else {
//line /usr/local/go/src/math/big/natdiv.go:882
			_go_fuzz_dep_.CoverTab[6431]++
//line /usr/local/go/src/math/big/natdiv.go:882
			// _ = "end of CoverTab[6431]"
//line /usr/local/go/src/math/big/natdiv.go:882
		}
//line /usr/local/go/src/math/big/natdiv.go:882
		// _ = "end of CoverTab[6426]"
	}
//line /usr/local/go/src/math/big/natdiv.go:883
	// _ = "end of CoverTab[6398]"
//line /usr/local/go/src/math/big/natdiv.go:883
	_go_fuzz_dep_.CoverTab[6399]++
							if qhatv.cmp(u.norm()) > 0 {
//line /usr/local/go/src/math/big/natdiv.go:884
		_go_fuzz_dep_.CoverTab[6432]++
								panic("impossible")
//line /usr/local/go/src/math/big/natdiv.go:885
		// _ = "end of CoverTab[6432]"
	} else {
//line /usr/local/go/src/math/big/natdiv.go:886
		_go_fuzz_dep_.CoverTab[6433]++
//line /usr/local/go/src/math/big/natdiv.go:886
		// _ = "end of CoverTab[6433]"
//line /usr/local/go/src/math/big/natdiv.go:886
	}
//line /usr/local/go/src/math/big/natdiv.go:886
	// _ = "end of CoverTab[6399]"
//line /usr/local/go/src/math/big/natdiv.go:886
	_go_fuzz_dep_.CoverTab[6400]++
							c := subVV(u[0:len(qhatv)], u[0:len(qhatv)], qhatv)
							if c > 0 {
//line /usr/local/go/src/math/big/natdiv.go:888
		_go_fuzz_dep_.CoverTab[6434]++
								c = subVW(u[len(qhatv):], u[len(qhatv):], c)
//line /usr/local/go/src/math/big/natdiv.go:889
		// _ = "end of CoverTab[6434]"
	} else {
//line /usr/local/go/src/math/big/natdiv.go:890
		_go_fuzz_dep_.CoverTab[6435]++
//line /usr/local/go/src/math/big/natdiv.go:890
		// _ = "end of CoverTab[6435]"
//line /usr/local/go/src/math/big/natdiv.go:890
	}
//line /usr/local/go/src/math/big/natdiv.go:890
	// _ = "end of CoverTab[6400]"
//line /usr/local/go/src/math/big/natdiv.go:890
	_go_fuzz_dep_.CoverTab[6401]++
							if c > 0 {
//line /usr/local/go/src/math/big/natdiv.go:891
		_go_fuzz_dep_.CoverTab[6436]++
								panic("impossible")
//line /usr/local/go/src/math/big/natdiv.go:892
		// _ = "end of CoverTab[6436]"
	} else {
//line /usr/local/go/src/math/big/natdiv.go:893
		_go_fuzz_dep_.CoverTab[6437]++
//line /usr/local/go/src/math/big/natdiv.go:893
		// _ = "end of CoverTab[6437]"
//line /usr/local/go/src/math/big/natdiv.go:893
	}
//line /usr/local/go/src/math/big/natdiv.go:893
	// _ = "end of CoverTab[6401]"
//line /usr/local/go/src/math/big/natdiv.go:893
	_go_fuzz_dep_.CoverTab[6402]++

//line /usr/local/go/src/math/big/natdiv.go:896
	addAt(z, qhat.norm(), 0)
//line /usr/local/go/src/math/big/natdiv.go:896
	// _ = "end of CoverTab[6402]"
}

//line /usr/local/go/src/math/big/natdiv.go:897
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/math/big/natdiv.go:897
var _ = _go_fuzz_dep_.CoverTab
