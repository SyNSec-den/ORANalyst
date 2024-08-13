// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/encoding/json/tables.go:5
package json

//line /usr/local/go/src/encoding/json/tables.go:5
import (
//line /usr/local/go/src/encoding/json/tables.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/encoding/json/tables.go:5
)
//line /usr/local/go/src/encoding/json/tables.go:5
import (
//line /usr/local/go/src/encoding/json/tables.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/encoding/json/tables.go:5
)

import "unicode/utf8"

// safeSet holds the value true if the ASCII character with the given array
//line /usr/local/go/src/encoding/json/tables.go:9
// position can be represented inside a JSON string without any further
//line /usr/local/go/src/encoding/json/tables.go:9
// escaping.
//line /usr/local/go/src/encoding/json/tables.go:9
//
//line /usr/local/go/src/encoding/json/tables.go:9
// All values are true except for the ASCII control characters (0-31), the
//line /usr/local/go/src/encoding/json/tables.go:9
// double quote ("), and the backslash character ("\").
//line /usr/local/go/src/encoding/json/tables.go:15
var safeSet = [utf8.RuneSelf]bool{
	' ':		true,
	'!':		true,
	'"':		false,
	'#':		true,
	'$':		true,
	'%':		true,
	'&':		true,
	'\'':		true,
	'(':		true,
	')':		true,
	'*':		true,
	'+':		true,
	',':		true,
	'-':		true,
	'.':		true,
	'/':		true,
	'0':		true,
	'1':		true,
	'2':		true,
	'3':		true,
	'4':		true,
	'5':		true,
	'6':		true,
	'7':		true,
	'8':		true,
	'9':		true,
	':':		true,
	';':		true,
	'<':		true,
	'=':		true,
	'>':		true,
	'?':		true,
	'@':		true,
	'A':		true,
	'B':		true,
	'C':		true,
	'D':		true,
	'E':		true,
	'F':		true,
	'G':		true,
	'H':		true,
	'I':		true,
	'J':		true,
	'K':		true,
	'L':		true,
	'M':		true,
	'N':		true,
	'O':		true,
	'P':		true,
	'Q':		true,
	'R':		true,
	'S':		true,
	'T':		true,
	'U':		true,
	'V':		true,
	'W':		true,
	'X':		true,
	'Y':		true,
	'Z':		true,
	'[':		true,
	'\\':		false,
	']':		true,
	'^':		true,
	'_':		true,
	'`':		true,
	'a':		true,
	'b':		true,
	'c':		true,
	'd':		true,
	'e':		true,
	'f':		true,
	'g':		true,
	'h':		true,
	'i':		true,
	'j':		true,
	'k':		true,
	'l':		true,
	'm':		true,
	'n':		true,
	'o':		true,
	'p':		true,
	'q':		true,
	'r':		true,
	's':		true,
	't':		true,
	'u':		true,
	'v':		true,
	'w':		true,
	'x':		true,
	'y':		true,
	'z':		true,
	'{':		true,
	'|':		true,
	'}':		true,
	'~':		true,
	'\u007f':	true,
}

// htmlSafeSet holds the value true if the ASCII character with the given
//line /usr/local/go/src/encoding/json/tables.go:114
// array position can be safely represented inside a JSON string, embedded
//line /usr/local/go/src/encoding/json/tables.go:114
// inside of HTML <script> tags, without any additional escaping.
//line /usr/local/go/src/encoding/json/tables.go:114
//
//line /usr/local/go/src/encoding/json/tables.go:114
// All values are true except for the ASCII control characters (0-31), the
//line /usr/local/go/src/encoding/json/tables.go:114
// double quote ("), the backslash character ("\"), HTML opening and closing
//line /usr/local/go/src/encoding/json/tables.go:114
// tags ("<" and ">"), and the ampersand ("&").
//line /usr/local/go/src/encoding/json/tables.go:121
var htmlSafeSet = [utf8.RuneSelf]bool{
	' ':		true,
	'!':		true,
	'"':		false,
	'#':		true,
	'$':		true,
	'%':		true,
	'&':		false,
	'\'':		true,
	'(':		true,
	')':		true,
	'*':		true,
	'+':		true,
	',':		true,
	'-':		true,
	'.':		true,
	'/':		true,
	'0':		true,
	'1':		true,
	'2':		true,
	'3':		true,
	'4':		true,
	'5':		true,
	'6':		true,
	'7':		true,
	'8':		true,
	'9':		true,
	':':		true,
	';':		true,
	'<':		false,
	'=':		true,
	'>':		false,
	'?':		true,
	'@':		true,
	'A':		true,
	'B':		true,
	'C':		true,
	'D':		true,
	'E':		true,
	'F':		true,
	'G':		true,
	'H':		true,
	'I':		true,
	'J':		true,
	'K':		true,
	'L':		true,
	'M':		true,
	'N':		true,
	'O':		true,
	'P':		true,
	'Q':		true,
	'R':		true,
	'S':		true,
	'T':		true,
	'U':		true,
	'V':		true,
	'W':		true,
	'X':		true,
	'Y':		true,
	'Z':		true,
	'[':		true,
	'\\':		false,
	']':		true,
	'^':		true,
	'_':		true,
	'`':		true,
	'a':		true,
	'b':		true,
	'c':		true,
	'd':		true,
	'e':		true,
	'f':		true,
	'g':		true,
	'h':		true,
	'i':		true,
	'j':		true,
	'k':		true,
	'l':		true,
	'm':		true,
	'n':		true,
	'o':		true,
	'p':		true,
	'q':		true,
	'r':		true,
	's':		true,
	't':		true,
	'u':		true,
	'v':		true,
	'w':		true,
	'x':		true,
	'y':		true,
	'z':		true,
	'{':		true,
	'|':		true,
	'}':		true,
	'~':		true,
	'\u007f':	true,
}
//line /usr/local/go/src/encoding/json/tables.go:218
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/encoding/json/tables.go:218
var _ = _go_fuzz_dep_.CoverTab
