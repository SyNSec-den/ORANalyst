// Copyright 2015 go-fuzz project authors. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.

package main

func compareCoverBody(hist [][]byte, cur []byte) bool {
	return compareCoverDump(hist, cur)
}
