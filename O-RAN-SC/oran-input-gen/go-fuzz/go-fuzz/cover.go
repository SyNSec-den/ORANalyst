// Copyright 2015 go-fuzz project authors. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"os"

	. "github.com/dvyukov/go-fuzz/go-fuzz-defs"
	. "github.com/dvyukov/go-fuzz/internal/go-fuzz-types"
)

// storeBit sets the bit at position i to 1.
func storeBit(base []byte, i int) []byte {
	if i < 1 || i > 16 {
		panic("i should be between 1 and 16")
	}
	if len(base) != 2 {
		panic("base should be of length 2")
	}

	i--
	byteIndex := i / 8
	bitIndex := uint(i % 8)
	base[byteIndex] |= (1 << bitIndex) // Set the bit to 1.
	return base
}

// storeAllLowerBit sets all bits before position i to 1.
func storeAllLowerBit(base []byte, i int) []byte {
	if i < 1 || i > 16 {
		panic("i should be between 1 and 16")
	}
	if len(base) != 2 {
		panic("base should be of length 2")
	}

	for j := 0; j < i; j++ {
		byteIndex := j / 8
		bitIndex := uint(j % 8)
		base[byteIndex] |= (1 << bitIndex) // Set the bit to 1.
	}
	return base
}

// getBit checks if the bit at position i is 1.
func getBit(base []byte, i int) bool {
	if i < 1 || i > 16 {
		panic("i should be between 1 and 16")
	}
	if len(base) != 2 {
		panic("base should be of length 2")
	}

	i--
	byteIndex := i / 8
	bitIndex := uint(i % 8)
	return (base[byteIndex] & (1 << bitIndex)) != 0
}

func makeCopy(data []byte) []byte {
	return append([]byte{}, data...)
}

func makeHistCopy(data [][]byte) [][]byte {
	hist := make([][]byte, len(data))
	for i, d := range data {
		hist[i] = makeCopy(d)
	}
	return hist
}

func compareCover(hist [][]byte, cur []byte) bool {
	log.Printf("compareCover")
	if len(hist) != CoverSize || len(cur) != CoverSize {
		log.Fatalf("bad cover table size (%v, %v)", len(hist), len(cur))
	}
	res := compareCoverBody(hist, cur)
	if false {
		// This check can legitimately fail if the test process has
		// some background goroutines that continue to write to the
		// cover array (cur storage is in shared memory).
		if compareCoverDump(hist, cur) != res {
			panic("bad")
		}
	}
	return res
}

func compareCoverDump(hist [][]byte, cur []byte) bool {
	for i, v := range hist {
		if cur[i] > 0 && !getBit(v, getIndex(cur[i])) {
			return true
		}
	}
	return false
}

func updateMaxCover(hist [][]byte, cur []byte) int {
	if len(hist) != CoverSize || len(cur) != CoverSize {
		log.Fatalf("bad cover table size (%v, %v)", len(hist), len(cur))
	}
	cnt := 0
	prev := 0
	// for i, x := range cur {
	// 	x = roundUpCover(x)
	// 	v := base[i]
	// 	if v != 0 || x > 0 {
	// 		cnt++
	// 	}
	// 	if v < x {
	// 		base[i] = x
	// 	}
	// }
	for i, v := range cur {
		if isCovered(hist[i]) {
			prev++
			cnt++
		}
		if v > 0 {
			if !getBit(hist[i], getIndex(v)) {
				if !isCovered(hist[i]) {
					cnt++
				}
				// log.Printf("updateMaxCover: %v, %v", i, v)
				hist[i] = storeAllLowerBit(hist[i], getIndex(v))
			}
		}
	}
	log.Printf("updateMaxCover: %v, %v", prev, cnt)
	return cnt
}

func isCovered(idx []byte) bool {
	if len(idx) != 2 {
		panic("idx should be of length 2")
	}
	return idx[0] != 0 || idx[1] != 0
}

// Quantize the counters. Otherwise we get too inflated corpus.
func roundUpCover(x byte) byte {
	if !*flagCoverCounters && x > 0 {
		return 255
	}

	if x <= 5 {
		return x
	} else if x <= 8 {
		return 8
	} else if x <= 16 {
		return 16
	} else if x <= 32 {
		return 32
	} else if x <= 64 {
		return 64
	}
	return 255
}

func getIndex(x byte) int {
	if x <= 10 {
		return int(x)
	} else if x <= 16 {
		return 11
	} else if x <= 32 {
		return 12
	} else if x <= 64 {
		return 13
	} else if x <= 128 {
		return 14
	}
	return 15
}

func findNewCover(hist [][]byte, cover []byte) (res [][]byte, notEmpty bool) {
	// res = make([]byte, CoverSize)
	// for i, b := range base {
	// 	c := cover[i]
	// 	if c > b {
	// 		res[i] = c
	// 		notEmpty = true
	// 	}
	// }
	// return
	res = make([][]byte, CoverSize)
	for i, v := range res {
		v = make([]byte, 2)
		res[i] = v
	}

	for i, v := range cover {
		if v > 0 && !getBit(hist[i], getIndex(v)) {
			res[i] = storeAllLowerBit(res[i], getIndex(v))
			notEmpty = true
		}
	}
	return res, notEmpty
}

func worseCover(base [][]byte, cover []byte) bool {
	// for i, b := range base {
	// 	c := cover[i]
	// 	if c < b {
	// 		return true
	// 	}
	// }
	// return false
	for i, v := range cover {
		if v > 0 && !getBit(base[i], getIndex(v)) {
			return true
		}
	}
	return false
}

func dumpCover(outf string, blocks map[int][]CoverBlock, cover []byte) {
	// Exclude files that have no coverage at all.
	files := make(map[string]bool)
	for i, v := range cover {
		if v == 0 {
			continue
		}
		for _, b := range blocks[i] {
			files[b.File] = true
		}
	}

	out, err := os.Create(outf)
	if err != nil {
		log.Fatalf("failed to create coverage file: %v", err)
	}
	defer out.Close()
	const showCounters = false
	if showCounters {
		fmt.Fprintf(out, "mode: count\n")
	} else {
		fmt.Fprintf(out, "mode: set\n")
	}
	for i, v := range cover {
		for _, b := range blocks[i] {
			if !files[b.File] {
				continue
			}
			if !showCounters && v != 0 {
				v = 1
			}
			fmt.Fprintf(out, "%s:%v.%v,%v.%v %v %v\n",
				b.File, b.StartLine, b.StartCol, b.EndLine, b.EndCol, b.NumStmt, v)
		}
	}
}

func dumpSonar(outf string, sites []SonarSite) {
	out, err := os.Create(outf)
	if err != nil {
		log.Fatalf("failed to create coverage file: %v", err)
	}
	defer out.Close()
	fmt.Fprintf(out, "mode: set\n")
	for i := range sites {
		s := &sites[i]
		cnt := 0  // red color
		stmt := 1 // account in percentage calculation
		if s.takenTotal[0] == 0 && s.takenTotal[1] == 0 {
			stmt = 0 // don't account in percentage calculation
			cnt = 1  // grey color
		} else if s.takenTotal[0] > 0 && s.takenTotal[1] > 0 {
			cnt = 100 // green color
		}
		fmt.Fprintf(out, "%v %v %v\n", s.loc, stmt, cnt)
	}
}
