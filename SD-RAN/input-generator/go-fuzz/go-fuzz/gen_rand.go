package main

import (
	"math"
)

func (m *Mutator) genRandInt8(lb, ub int) int8 {
	if lb != 0 || ub != 0 {
		return int8(m.r.Intn(ub-lb+1) + lb)
	}
	switch m.r.Intn(3) {
	case 0:
		return int8(m.r.Intn(math.MaxInt8 + 1))
	default:
		return -int8(m.r.Intn(math.MaxInt8+1)) - 1
	}
}

func (m *Mutator) genRandInt16(lb, ub int) int16 {
	if lb != 0 || ub != 0 {
		return int16(m.r.Intn(ub-lb+1) + lb)
	}
	switch m.r.Intn(3) {
	case 0:
		return int16(m.r.Intn(math.MaxInt16 + 1))
	default:
		return -int16(m.r.Intn(math.MaxInt16+1)) - 1
	}
}

func (m *Mutator) genRandInt32(lb, ub int) int32 {
	if lb != 0 || ub != 0 {
		return int32(m.r.Intn(ub-lb+1) + lb)
	}
	switch m.r.Intn(3) {
	case 0:
		return int32(m.r.Intn(math.MaxInt32 + 1))
	default:
		return -int32(m.r.Intn(math.MaxInt32+1)) - 1
	}
}

func (m *Mutator) genRandInt64(lb, ub int) int64 {
	if lb != 0 || ub != 0 {
		return m.rr.Int63n(int64(ub-lb+1)) + int64(lb)
	}
	switch m.r.Intn(3) {
	case 0:
		return m.rr.Int63()
	default:
		return -m.rr.Int63() - 1
	}
}

func (m *Mutator) genRandUint8(lb, ub int) uint8 {
	if lb != 0 || ub != 0 {
		return uint8(m.r.Intn(ub-lb+1) + lb)
	}
	return uint8(m.r.Intn(math.MaxUint8 + 1))
}

func (m *Mutator) genRandUint16(lb, ub int) uint16 {
	if lb != 0 || ub != 0 {
		return uint16(m.r.Intn(ub-lb+1) + lb)
	}
	return uint16(m.r.Intn(math.MaxUint16 + 1))
}

func (m *Mutator) genRandUint32(lb, ub int) uint32 {
	if lb != 0 || ub != 0 {
		return uint32(m.rr.Intn(ub-lb+1) + lb)
	}
	return uint32(m.rr.Intn(math.MaxUint32 + 1))
}

func (m *Mutator) genRandUint64(lb, ub int) uint64 {
	if lb != 0 || ub != 0 {
		return uint64(m.rr.Int63n(int64(ub-lb+1)) + int64(lb))
	}
	return uint64(m.rr.Uint64())
}

func (m *Mutator) genRandFloat32(lb, ub int) float32 {
	if lb != 0 || ub != 0 {
		panic("genRandFloat32 lb ub: not implemented")
	}
	return m.rr.Float32()
}

func (m *Mutator) genRandFloat64(lb, ub int) float64 {
	if lb != 0 || ub != 0 {
		panic("genRandFloat64 lb ub: not implemented")
	}
	return m.rr.Float64()
}
