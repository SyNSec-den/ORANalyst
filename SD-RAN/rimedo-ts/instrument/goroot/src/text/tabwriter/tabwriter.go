// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/text/tabwriter/tabwriter.go:5
// Package tabwriter implements a write filter (tabwriter.Writer) that
//line /usr/local/go/src/text/tabwriter/tabwriter.go:5
// translates tabbed columns in input into properly aligned text.
//line /usr/local/go/src/text/tabwriter/tabwriter.go:5
//
//line /usr/local/go/src/text/tabwriter/tabwriter.go:5
// The package is using the Elastic Tabstops algorithm described at
//line /usr/local/go/src/text/tabwriter/tabwriter.go:5
// http://nickgravgaard.com/elastictabstops/index.html.
//line /usr/local/go/src/text/tabwriter/tabwriter.go:5
//
//line /usr/local/go/src/text/tabwriter/tabwriter.go:5
// The text/tabwriter package is frozen and is not accepting new features.
//line /usr/local/go/src/text/tabwriter/tabwriter.go:12
package tabwriter

//line /usr/local/go/src/text/tabwriter/tabwriter.go:12
import (
//line /usr/local/go/src/text/tabwriter/tabwriter.go:12
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:12
)
//line /usr/local/go/src/text/tabwriter/tabwriter.go:12
import (
//line /usr/local/go/src/text/tabwriter/tabwriter.go:12
	_atomic_ "sync/atomic"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:12
)

import (
	"io"
	"unicode/utf8"
)

//line /usr/local/go/src/text/tabwriter/tabwriter.go:22
// A cell represents a segment of text terminated by tabs or line breaks.
//line /usr/local/go/src/text/tabwriter/tabwriter.go:22
// The text itself is stored in a separate buffer; cell only describes the
//line /usr/local/go/src/text/tabwriter/tabwriter.go:22
// segment's size in bytes, its width in runes, and whether it's an htab
//line /usr/local/go/src/text/tabwriter/tabwriter.go:22
// ('\t') terminated cell.
//line /usr/local/go/src/text/tabwriter/tabwriter.go:26
type cell struct {
	size	int	// cell size in bytes
	width	int	// cell width in runes
	htab	bool	// true if the cell is terminated by an htab ('\t')
}

// A Writer is a filter that inserts padding around tab-delimited
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
// columns in its input to align them in the output.
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
//
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
// The Writer treats incoming bytes as UTF-8-encoded text consisting
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
// of cells terminated by horizontal ('\t') or vertical ('\v') tabs,
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
// and newline ('\n') or formfeed ('\f') characters; both newline and
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
// formfeed act as line breaks.
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
//
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
// Tab-terminated cells in contiguous lines constitute a column. The
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
// Writer inserts padding as needed to make all cells in a column have
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
// the same width, effectively aligning the columns. It assumes that
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
// all characters have the same width, except for tabs for which a
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
// tabwidth must be specified. Column cells must be tab-terminated, not
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
// tab-separated: non-tab terminated trailing text at the end of a line
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
// forms a cell but that cell is not part of an aligned column.
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
// For instance, in this example (where | stands for a horizontal tab):
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
//
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
//	aaaa|bbb|d
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
//	aa  |b  |dd
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
//	a   |
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
//	aa  |cccc|eee
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
//
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
// the b and c are in distinct columns (the b column is not contiguous
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
// all the way). The d and e are not in a column at all (there's no
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
// terminating tab, nor would the column be contiguous).
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
//
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
// The Writer assumes that all Unicode code points have the same width;
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
// this may not be true in some fonts or if the string contains combining
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
// characters.
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
//
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
// If DiscardEmptyColumns is set, empty columns that are terminated
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
// entirely by vertical (or "soft") tabs are discarded. Columns
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
// terminated by horizontal (or "hard") tabs are not affected by
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
// this flag.
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
//
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
// If a Writer is configured to filter HTML, HTML tags and entities
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
// are passed through. The widths of tags and entities are
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
// assumed to be zero (tags) and one (entities) for formatting purposes.
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
//
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
// A segment of text may be escaped by bracketing it with Escape
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
// characters. The tabwriter passes escaped text segments through
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
// unchanged. In particular, it does not interpret any tabs or line
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
// breaks within the segment. If the StripEscape flag is set, the
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
// Escape characters are stripped from the output; otherwise they
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
// are passed through as well. For the purpose of formatting, the
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
// width of the escaped text is always computed excluding the Escape
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
// characters.
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
//
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
// The formfeed character acts like a newline but it also terminates
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
// all columns in the current line (effectively calling Flush). Tab-
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
// terminated cells in the next line start new columns. Unless found
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
// inside an HTML tag or inside an escaped text segment, formfeed
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
// characters appear as newlines in the output.
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
//
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
// The Writer must buffer input internally, because proper spacing
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
// of one line may depend on the cells in future lines. Clients must
//line /usr/local/go/src/text/tabwriter/tabwriter.go:32
// call Flush when done calling Write.
//line /usr/local/go/src/text/tabwriter/tabwriter.go:89
type Writer struct {
	// configuration
	output		io.Writer
	minwidth	int
	tabwidth	int
	padding		int
	padbytes	[8]byte
	flags		uint

	// current state
	buf	[]byte		// collected text excluding tabs or line breaks
	pos	int		// buffer position up to which cell.width of incomplete cell has been computed
	cell	cell		// current incomplete cell; cell.width is up to buf[pos] excluding ignored sections
	endChar	byte		// terminating char of escaped sequence (Escape for escapes, '>', ';' for HTML tags/entities, or 0)
	lines	[][]cell	// list of lines; each line is a list of cells
	widths	[]int		// list of column widths in runes - re-used during formatting
}

// addLine adds a new line.
//line /usr/local/go/src/text/tabwriter/tabwriter.go:107
// flushed is a hint indicating whether the underlying writer was just flushed.
//line /usr/local/go/src/text/tabwriter/tabwriter.go:107
// If so, the previous line is not likely to be a good indicator of the new line's cells.
//line /usr/local/go/src/text/tabwriter/tabwriter.go:110
func (b *Writer) addLine(flushed bool) {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:110
	_go_fuzz_dep_.CoverTab[45085]++

//line /usr/local/go/src/text/tabwriter/tabwriter.go:114
	if n := len(b.lines) + 1; n <= cap(b.lines) {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:114
		_go_fuzz_dep_.CoverTab[45087]++
									b.lines = b.lines[:n]
									b.lines[n-1] = b.lines[n-1][:0]
//line /usr/local/go/src/text/tabwriter/tabwriter.go:116
		// _ = "end of CoverTab[45087]"
	} else {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:117
		_go_fuzz_dep_.CoverTab[45088]++
									b.lines = append(b.lines, nil)
//line /usr/local/go/src/text/tabwriter/tabwriter.go:118
		// _ = "end of CoverTab[45088]"
	}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:119
	// _ = "end of CoverTab[45085]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:119
	_go_fuzz_dep_.CoverTab[45086]++

								if !flushed {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:121
		_go_fuzz_dep_.CoverTab[45089]++

//line /usr/local/go/src/text/tabwriter/tabwriter.go:126
		if n := len(b.lines); n >= 2 {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:126
			_go_fuzz_dep_.CoverTab[45090]++
										if prev := len(b.lines[n-2]); prev > cap(b.lines[n-1]) {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:127
				_go_fuzz_dep_.CoverTab[45091]++
											b.lines[n-1] = make([]cell, 0, prev)
//line /usr/local/go/src/text/tabwriter/tabwriter.go:128
				// _ = "end of CoverTab[45091]"
			} else {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:129
				_go_fuzz_dep_.CoverTab[45092]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:129
				// _ = "end of CoverTab[45092]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:129
			}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:129
			// _ = "end of CoverTab[45090]"
		} else {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:130
			_go_fuzz_dep_.CoverTab[45093]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:130
			// _ = "end of CoverTab[45093]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:130
		}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:130
		// _ = "end of CoverTab[45089]"
	} else {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:131
		_go_fuzz_dep_.CoverTab[45094]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:131
		// _ = "end of CoverTab[45094]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:131
	}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:131
	// _ = "end of CoverTab[45086]"
}

// Reset the current state.
func (b *Writer) reset() {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:135
	_go_fuzz_dep_.CoverTab[45095]++
								b.buf = b.buf[:0]
								b.pos = 0
								b.cell = cell{}
								b.endChar = 0
								b.lines = b.lines[0:0]
								b.widths = b.widths[0:0]
								b.addLine(true)
//line /usr/local/go/src/text/tabwriter/tabwriter.go:142
	// _ = "end of CoverTab[45095]"
}

//line /usr/local/go/src/text/tabwriter/tabwriter.go:168
// Formatting can be controlled with these flags.
const (
	// Ignore html tags and treat entities (starting with '&'
	// and ending in ';') as single characters (width = 1).
	FilterHTML	uint	= 1 << iota

	// Strip Escape characters bracketing escaped text segments
	// instead of passing them through unchanged with the text.
	StripEscape

	// Force right-alignment of cell content.
	// Default is left-alignment.
	AlignRight

	// Handle empty columns as if they were not present in
	// the input in the first place.
	DiscardEmptyColumns

	// Always use tabs for indentation columns (i.e., padding of
	// leading empty cells on the left) independent of padchar.
	TabIndent

	// Print a vertical bar ('|') between columns (after formatting).
	// Discarded columns appear as zero-width columns ("||").
	Debug
)

// A Writer must be initialized with a call to Init. The first parameter (output)
//line /usr/local/go/src/text/tabwriter/tabwriter.go:195
// specifies the filter output. The remaining parameters control the formatting:
//line /usr/local/go/src/text/tabwriter/tabwriter.go:195
//
//line /usr/local/go/src/text/tabwriter/tabwriter.go:195
//	minwidth	minimal cell width including any padding
//line /usr/local/go/src/text/tabwriter/tabwriter.go:195
//	tabwidth	width of tab characters (equivalent number of spaces)
//line /usr/local/go/src/text/tabwriter/tabwriter.go:195
//	padding		padding added to a cell before computing its width
//line /usr/local/go/src/text/tabwriter/tabwriter.go:195
//	padchar		ASCII char used for padding
//line /usr/local/go/src/text/tabwriter/tabwriter.go:195
//			if padchar == '\t', the Writer will assume that the
//line /usr/local/go/src/text/tabwriter/tabwriter.go:195
//			width of a '\t' in the formatted output is tabwidth,
//line /usr/local/go/src/text/tabwriter/tabwriter.go:195
//			and cells are left-aligned independent of align_left
//line /usr/local/go/src/text/tabwriter/tabwriter.go:195
//			(for correct-looking results, tabwidth must correspond
//line /usr/local/go/src/text/tabwriter/tabwriter.go:195
//			to the tab width in the viewer displaying the result)
//line /usr/local/go/src/text/tabwriter/tabwriter.go:195
//	flags		formatting control
//line /usr/local/go/src/text/tabwriter/tabwriter.go:208
func (b *Writer) Init(output io.Writer, minwidth, tabwidth, padding int, padchar byte, flags uint) *Writer {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:208
	_go_fuzz_dep_.CoverTab[45096]++
								if minwidth < 0 || func() bool {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:209
		_go_fuzz_dep_.CoverTab[45100]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:209
		return tabwidth < 0
//line /usr/local/go/src/text/tabwriter/tabwriter.go:209
		// _ = "end of CoverTab[45100]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:209
	}() || func() bool {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:209
		_go_fuzz_dep_.CoverTab[45101]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:209
		return padding < 0
//line /usr/local/go/src/text/tabwriter/tabwriter.go:209
		// _ = "end of CoverTab[45101]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:209
	}() {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:209
		_go_fuzz_dep_.CoverTab[45102]++
									panic("negative minwidth, tabwidth, or padding")
//line /usr/local/go/src/text/tabwriter/tabwriter.go:210
		// _ = "end of CoverTab[45102]"
	} else {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:211
		_go_fuzz_dep_.CoverTab[45103]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:211
		// _ = "end of CoverTab[45103]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:211
	}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:211
	// _ = "end of CoverTab[45096]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:211
	_go_fuzz_dep_.CoverTab[45097]++
								b.output = output
								b.minwidth = minwidth
								b.tabwidth = tabwidth
								b.padding = padding
								for i := range b.padbytes {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:216
		_go_fuzz_dep_.CoverTab[45104]++
									b.padbytes[i] = padchar
//line /usr/local/go/src/text/tabwriter/tabwriter.go:217
		// _ = "end of CoverTab[45104]"
	}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:218
	// _ = "end of CoverTab[45097]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:218
	_go_fuzz_dep_.CoverTab[45098]++
								if padchar == '\t' {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:219
		_go_fuzz_dep_.CoverTab[45105]++

									flags &^= AlignRight
//line /usr/local/go/src/text/tabwriter/tabwriter.go:221
		// _ = "end of CoverTab[45105]"
	} else {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:222
		_go_fuzz_dep_.CoverTab[45106]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:222
		// _ = "end of CoverTab[45106]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:222
	}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:222
	// _ = "end of CoverTab[45098]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:222
	_go_fuzz_dep_.CoverTab[45099]++
								b.flags = flags

								b.reset()

								return b
//line /usr/local/go/src/text/tabwriter/tabwriter.go:227
	// _ = "end of CoverTab[45099]"
}

// debugging support (keep code around)
func (b *Writer) dump() {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:231
	_go_fuzz_dep_.CoverTab[45107]++
								pos := 0
								for i, line := range b.lines {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:233
		_go_fuzz_dep_.CoverTab[45109]++
									print("(", i, ") ")
									for _, c := range line {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:235
			_go_fuzz_dep_.CoverTab[45111]++
										print("[", string(b.buf[pos:pos+c.size]), "]")
										pos += c.size
//line /usr/local/go/src/text/tabwriter/tabwriter.go:237
			// _ = "end of CoverTab[45111]"
		}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:238
		// _ = "end of CoverTab[45109]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:238
		_go_fuzz_dep_.CoverTab[45110]++
									print("\n")
//line /usr/local/go/src/text/tabwriter/tabwriter.go:239
		// _ = "end of CoverTab[45110]"
	}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:240
	// _ = "end of CoverTab[45107]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:240
	_go_fuzz_dep_.CoverTab[45108]++
								print("\n")
//line /usr/local/go/src/text/tabwriter/tabwriter.go:241
	// _ = "end of CoverTab[45108]"
}

// local error wrapper so we can distinguish errors we want to return
//line /usr/local/go/src/text/tabwriter/tabwriter.go:244
// as errors from genuine panics (which we don't want to return as errors)
//line /usr/local/go/src/text/tabwriter/tabwriter.go:246
type osError struct {
	err error
}

func (b *Writer) write0(buf []byte) {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:250
	_go_fuzz_dep_.CoverTab[45112]++
								n, err := b.output.Write(buf)
								if n != len(buf) && func() bool {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:252
		_go_fuzz_dep_.CoverTab[45114]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:252
		return err == nil
//line /usr/local/go/src/text/tabwriter/tabwriter.go:252
		// _ = "end of CoverTab[45114]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:252
	}() {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:252
		_go_fuzz_dep_.CoverTab[45115]++
									err = io.ErrShortWrite
//line /usr/local/go/src/text/tabwriter/tabwriter.go:253
		// _ = "end of CoverTab[45115]"
	} else {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:254
		_go_fuzz_dep_.CoverTab[45116]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:254
		// _ = "end of CoverTab[45116]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:254
	}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:254
	// _ = "end of CoverTab[45112]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:254
	_go_fuzz_dep_.CoverTab[45113]++
								if err != nil {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:255
		_go_fuzz_dep_.CoverTab[45117]++
									panic(osError{err})
//line /usr/local/go/src/text/tabwriter/tabwriter.go:256
		// _ = "end of CoverTab[45117]"
	} else {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:257
		_go_fuzz_dep_.CoverTab[45118]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:257
		// _ = "end of CoverTab[45118]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:257
	}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:257
	// _ = "end of CoverTab[45113]"
}

func (b *Writer) writeN(src []byte, n int) {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:260
	_go_fuzz_dep_.CoverTab[45119]++
								for n > len(src) {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:261
		_go_fuzz_dep_.CoverTab[45121]++
									b.write0(src)
									n -= len(src)
//line /usr/local/go/src/text/tabwriter/tabwriter.go:263
		// _ = "end of CoverTab[45121]"
	}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:264
	// _ = "end of CoverTab[45119]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:264
	_go_fuzz_dep_.CoverTab[45120]++
								b.write0(src[0:n])
//line /usr/local/go/src/text/tabwriter/tabwriter.go:265
	// _ = "end of CoverTab[45120]"
}

var (
	newline	= []byte{'\n'}
	tabs	= []byte("\t\t\t\t\t\t\t\t")
)

func (b *Writer) writePadding(textw, cellw int, useTabs bool) {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:273
	_go_fuzz_dep_.CoverTab[45122]++
								if b.padbytes[0] == '\t' || func() bool {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:274
		_go_fuzz_dep_.CoverTab[45124]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:274
		return useTabs
//line /usr/local/go/src/text/tabwriter/tabwriter.go:274
		// _ = "end of CoverTab[45124]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:274
	}() {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:274
		_go_fuzz_dep_.CoverTab[45125]++

									if b.tabwidth == 0 {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:276
			_go_fuzz_dep_.CoverTab[45128]++
										return
//line /usr/local/go/src/text/tabwriter/tabwriter.go:277
			// _ = "end of CoverTab[45128]"
		} else {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:278
			_go_fuzz_dep_.CoverTab[45129]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:278
			// _ = "end of CoverTab[45129]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:278
		}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:278
		// _ = "end of CoverTab[45125]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:278
		_go_fuzz_dep_.CoverTab[45126]++

									cellw = (cellw + b.tabwidth - 1) / b.tabwidth * b.tabwidth
									n := cellw - textw
									if n < 0 {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:282
			_go_fuzz_dep_.CoverTab[45130]++
										panic("internal error")
//line /usr/local/go/src/text/tabwriter/tabwriter.go:283
			// _ = "end of CoverTab[45130]"
		} else {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:284
			_go_fuzz_dep_.CoverTab[45131]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:284
			// _ = "end of CoverTab[45131]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:284
		}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:284
		// _ = "end of CoverTab[45126]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:284
		_go_fuzz_dep_.CoverTab[45127]++
									b.writeN(tabs, (n+b.tabwidth-1)/b.tabwidth)
									return
//line /usr/local/go/src/text/tabwriter/tabwriter.go:286
		// _ = "end of CoverTab[45127]"
	} else {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:287
		_go_fuzz_dep_.CoverTab[45132]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:287
		// _ = "end of CoverTab[45132]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:287
	}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:287
	// _ = "end of CoverTab[45122]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:287
	_go_fuzz_dep_.CoverTab[45123]++

//line /usr/local/go/src/text/tabwriter/tabwriter.go:290
	b.writeN(b.padbytes[0:], cellw-textw)
//line /usr/local/go/src/text/tabwriter/tabwriter.go:290
	// _ = "end of CoverTab[45123]"
}

var vbar = []byte{'|'}

func (b *Writer) writeLines(pos0 int, line0, line1 int) (pos int) {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:295
	_go_fuzz_dep_.CoverTab[45133]++
								pos = pos0
								for i := line0; i < line1; i++ {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:297
		_go_fuzz_dep_.CoverTab[45135]++
									line := b.lines[i]

//line /usr/local/go/src/text/tabwriter/tabwriter.go:301
		useTabs := b.flags&TabIndent != 0

		for j, c := range line {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:303
			_go_fuzz_dep_.CoverTab[45137]++
										if j > 0 && func() bool {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:304
				_go_fuzz_dep_.CoverTab[45139]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:304
				return b.flags&Debug != 0
//line /usr/local/go/src/text/tabwriter/tabwriter.go:304
				// _ = "end of CoverTab[45139]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:304
			}() {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:304
				_go_fuzz_dep_.CoverTab[45140]++

											b.write0(vbar)
//line /usr/local/go/src/text/tabwriter/tabwriter.go:306
				// _ = "end of CoverTab[45140]"
			} else {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:307
				_go_fuzz_dep_.CoverTab[45141]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:307
				// _ = "end of CoverTab[45141]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:307
			}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:307
			// _ = "end of CoverTab[45137]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:307
			_go_fuzz_dep_.CoverTab[45138]++

										if c.size == 0 {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:309
				_go_fuzz_dep_.CoverTab[45142]++

											if j < len(b.widths) {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:311
					_go_fuzz_dep_.CoverTab[45143]++
												b.writePadding(c.width, b.widths[j], useTabs)
//line /usr/local/go/src/text/tabwriter/tabwriter.go:312
					// _ = "end of CoverTab[45143]"
				} else {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:313
					_go_fuzz_dep_.CoverTab[45144]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:313
					// _ = "end of CoverTab[45144]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:313
				}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:313
				// _ = "end of CoverTab[45142]"
			} else {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:314
				_go_fuzz_dep_.CoverTab[45145]++

											useTabs = false
											if b.flags&AlignRight == 0 {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:317
					_go_fuzz_dep_.CoverTab[45146]++
												b.write0(b.buf[pos : pos+c.size])
												pos += c.size
												if j < len(b.widths) {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:320
						_go_fuzz_dep_.CoverTab[45147]++
													b.writePadding(c.width, b.widths[j], false)
//line /usr/local/go/src/text/tabwriter/tabwriter.go:321
						// _ = "end of CoverTab[45147]"
					} else {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:322
						_go_fuzz_dep_.CoverTab[45148]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:322
						// _ = "end of CoverTab[45148]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:322
					}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:322
					// _ = "end of CoverTab[45146]"
				} else {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:323
					_go_fuzz_dep_.CoverTab[45149]++
												if j < len(b.widths) {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:324
						_go_fuzz_dep_.CoverTab[45151]++
													b.writePadding(c.width, b.widths[j], false)
//line /usr/local/go/src/text/tabwriter/tabwriter.go:325
						// _ = "end of CoverTab[45151]"
					} else {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:326
						_go_fuzz_dep_.CoverTab[45152]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:326
						// _ = "end of CoverTab[45152]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:326
					}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:326
					// _ = "end of CoverTab[45149]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:326
					_go_fuzz_dep_.CoverTab[45150]++
												b.write0(b.buf[pos : pos+c.size])
												pos += c.size
//line /usr/local/go/src/text/tabwriter/tabwriter.go:328
					// _ = "end of CoverTab[45150]"
				}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:329
				// _ = "end of CoverTab[45145]"
			}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:330
			// _ = "end of CoverTab[45138]"
		}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:331
		// _ = "end of CoverTab[45135]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:331
		_go_fuzz_dep_.CoverTab[45136]++

									if i+1 == len(b.lines) {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:333
			_go_fuzz_dep_.CoverTab[45153]++

//line /usr/local/go/src/text/tabwriter/tabwriter.go:336
			b.write0(b.buf[pos : pos+b.cell.size])
										pos += b.cell.size
//line /usr/local/go/src/text/tabwriter/tabwriter.go:337
			// _ = "end of CoverTab[45153]"
		} else {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:338
			_go_fuzz_dep_.CoverTab[45154]++

										b.write0(newline)
//line /usr/local/go/src/text/tabwriter/tabwriter.go:340
			// _ = "end of CoverTab[45154]"
		}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:341
		// _ = "end of CoverTab[45136]"
	}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:342
	// _ = "end of CoverTab[45133]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:342
	_go_fuzz_dep_.CoverTab[45134]++
								return
//line /usr/local/go/src/text/tabwriter/tabwriter.go:343
	// _ = "end of CoverTab[45134]"
}

// Format the text between line0 and line1 (excluding line1); pos
//line /usr/local/go/src/text/tabwriter/tabwriter.go:346
// is the buffer position corresponding to the beginning of line0.
//line /usr/local/go/src/text/tabwriter/tabwriter.go:346
// Returns the buffer position corresponding to the beginning of
//line /usr/local/go/src/text/tabwriter/tabwriter.go:346
// line1 and an error, if any.
//line /usr/local/go/src/text/tabwriter/tabwriter.go:350
func (b *Writer) format(pos0 int, line0, line1 int) (pos int) {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:350
	_go_fuzz_dep_.CoverTab[45155]++
								pos = pos0
								column := len(b.widths)
								for this := line0; this < line1; this++ {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:353
		_go_fuzz_dep_.CoverTab[45157]++
									line := b.lines[this]

									if column >= len(line)-1 {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:356
			_go_fuzz_dep_.CoverTab[45161]++
										continue
//line /usr/local/go/src/text/tabwriter/tabwriter.go:357
			// _ = "end of CoverTab[45161]"
		} else {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:358
			_go_fuzz_dep_.CoverTab[45162]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:358
			// _ = "end of CoverTab[45162]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:358
		}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:358
		// _ = "end of CoverTab[45157]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:358
		_go_fuzz_dep_.CoverTab[45158]++

//line /usr/local/go/src/text/tabwriter/tabwriter.go:367
		pos = b.writeLines(pos, line0, this)
									line0 = this

//line /usr/local/go/src/text/tabwriter/tabwriter.go:371
		width := b.minwidth
		discardable := true
		for ; this < line1; this++ {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:373
			_go_fuzz_dep_.CoverTab[45163]++
										line = b.lines[this]
										if column >= len(line)-1 {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:375
				_go_fuzz_dep_.CoverTab[45166]++
											break
//line /usr/local/go/src/text/tabwriter/tabwriter.go:376
				// _ = "end of CoverTab[45166]"
			} else {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:377
				_go_fuzz_dep_.CoverTab[45167]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:377
				// _ = "end of CoverTab[45167]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:377
			}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:377
			// _ = "end of CoverTab[45163]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:377
			_go_fuzz_dep_.CoverTab[45164]++

										c := line[column]

										if w := c.width + b.padding; w > width {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:381
				_go_fuzz_dep_.CoverTab[45168]++
											width = w
//line /usr/local/go/src/text/tabwriter/tabwriter.go:382
				// _ = "end of CoverTab[45168]"
			} else {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:383
				_go_fuzz_dep_.CoverTab[45169]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:383
				// _ = "end of CoverTab[45169]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:383
			}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:383
			// _ = "end of CoverTab[45164]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:383
			_go_fuzz_dep_.CoverTab[45165]++

										if c.width > 0 || func() bool {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:385
				_go_fuzz_dep_.CoverTab[45170]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:385
				return c.htab
//line /usr/local/go/src/text/tabwriter/tabwriter.go:385
				// _ = "end of CoverTab[45170]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:385
			}() {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:385
				_go_fuzz_dep_.CoverTab[45171]++
											discardable = false
//line /usr/local/go/src/text/tabwriter/tabwriter.go:386
				// _ = "end of CoverTab[45171]"
			} else {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:387
				_go_fuzz_dep_.CoverTab[45172]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:387
				// _ = "end of CoverTab[45172]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:387
			}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:387
			// _ = "end of CoverTab[45165]"
		}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:388
		// _ = "end of CoverTab[45158]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:388
		_go_fuzz_dep_.CoverTab[45159]++

//line /usr/local/go/src/text/tabwriter/tabwriter.go:392
		if discardable && func() bool {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:392
			_go_fuzz_dep_.CoverTab[45173]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:392
			return b.flags&DiscardEmptyColumns != 0
//line /usr/local/go/src/text/tabwriter/tabwriter.go:392
			// _ = "end of CoverTab[45173]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:392
		}() {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:392
			_go_fuzz_dep_.CoverTab[45174]++
										width = 0
//line /usr/local/go/src/text/tabwriter/tabwriter.go:393
			// _ = "end of CoverTab[45174]"
		} else {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:394
			_go_fuzz_dep_.CoverTab[45175]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:394
			// _ = "end of CoverTab[45175]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:394
		}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:394
		// _ = "end of CoverTab[45159]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:394
		_go_fuzz_dep_.CoverTab[45160]++

//line /usr/local/go/src/text/tabwriter/tabwriter.go:398
		b.widths = append(b.widths, width)
									pos = b.format(pos, line0, this)
									b.widths = b.widths[0 : len(b.widths)-1]
									line0 = this
//line /usr/local/go/src/text/tabwriter/tabwriter.go:401
		// _ = "end of CoverTab[45160]"
	}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:402
	// _ = "end of CoverTab[45155]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:402
	_go_fuzz_dep_.CoverTab[45156]++

//line /usr/local/go/src/text/tabwriter/tabwriter.go:405
	return b.writeLines(pos, line0, line1)
//line /usr/local/go/src/text/tabwriter/tabwriter.go:405
	// _ = "end of CoverTab[45156]"
}

// Append text to current cell.
func (b *Writer) append(text []byte) {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:409
	_go_fuzz_dep_.CoverTab[45176]++
								b.buf = append(b.buf, text...)
								b.cell.size += len(text)
//line /usr/local/go/src/text/tabwriter/tabwriter.go:411
	// _ = "end of CoverTab[45176]"
}

// Update the cell width.
func (b *Writer) updateWidth() {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:415
	_go_fuzz_dep_.CoverTab[45177]++
								b.cell.width += utf8.RuneCount(b.buf[b.pos:])
								b.pos = len(b.buf)
//line /usr/local/go/src/text/tabwriter/tabwriter.go:417
	// _ = "end of CoverTab[45177]"
}

// To escape a text segment, bracket it with Escape characters.
//line /usr/local/go/src/text/tabwriter/tabwriter.go:420
// For instance, the tab in this string "Ignore this tab: \xff\t\xff"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:420
// does not terminate a cell and constitutes a single character of
//line /usr/local/go/src/text/tabwriter/tabwriter.go:420
// width one for formatting purposes.
//line /usr/local/go/src/text/tabwriter/tabwriter.go:420
//
//line /usr/local/go/src/text/tabwriter/tabwriter.go:420
// The value 0xff was chosen because it cannot appear in a valid UTF-8 sequence.
//line /usr/local/go/src/text/tabwriter/tabwriter.go:426
const Escape = '\xff'

// Start escaped mode.
func (b *Writer) startEscape(ch byte) {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:429
	_go_fuzz_dep_.CoverTab[45178]++
								switch ch {
	case Escape:
//line /usr/local/go/src/text/tabwriter/tabwriter.go:431
		_go_fuzz_dep_.CoverTab[45179]++
									b.endChar = Escape
//line /usr/local/go/src/text/tabwriter/tabwriter.go:432
		// _ = "end of CoverTab[45179]"
	case '<':
//line /usr/local/go/src/text/tabwriter/tabwriter.go:433
		_go_fuzz_dep_.CoverTab[45180]++
									b.endChar = '>'
//line /usr/local/go/src/text/tabwriter/tabwriter.go:434
		// _ = "end of CoverTab[45180]"
	case '&':
//line /usr/local/go/src/text/tabwriter/tabwriter.go:435
		_go_fuzz_dep_.CoverTab[45181]++
									b.endChar = ';'
//line /usr/local/go/src/text/tabwriter/tabwriter.go:436
		// _ = "end of CoverTab[45181]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:436
	default:
//line /usr/local/go/src/text/tabwriter/tabwriter.go:436
		_go_fuzz_dep_.CoverTab[45182]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:436
		// _ = "end of CoverTab[45182]"
	}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:437
	// _ = "end of CoverTab[45178]"
}

// Terminate escaped mode. If the escaped text was an HTML tag, its width
//line /usr/local/go/src/text/tabwriter/tabwriter.go:440
// is assumed to be zero for formatting purposes; if it was an HTML entity,
//line /usr/local/go/src/text/tabwriter/tabwriter.go:440
// its width is assumed to be one. In all other cases, the width is the
//line /usr/local/go/src/text/tabwriter/tabwriter.go:440
// unicode width of the text.
//line /usr/local/go/src/text/tabwriter/tabwriter.go:444
func (b *Writer) endEscape() {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:444
	_go_fuzz_dep_.CoverTab[45183]++
								switch b.endChar {
	case Escape:
//line /usr/local/go/src/text/tabwriter/tabwriter.go:446
		_go_fuzz_dep_.CoverTab[45185]++
									b.updateWidth()
									if b.flags&StripEscape == 0 {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:448
			_go_fuzz_dep_.CoverTab[45189]++
										b.cell.width -= 2
//line /usr/local/go/src/text/tabwriter/tabwriter.go:449
			// _ = "end of CoverTab[45189]"
		} else {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:450
			_go_fuzz_dep_.CoverTab[45190]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:450
			// _ = "end of CoverTab[45190]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:450
		}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:450
		// _ = "end of CoverTab[45185]"
	case '>':
//line /usr/local/go/src/text/tabwriter/tabwriter.go:451
		_go_fuzz_dep_.CoverTab[45186]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:451
		// _ = "end of CoverTab[45186]"
	case ';':
//line /usr/local/go/src/text/tabwriter/tabwriter.go:452
		_go_fuzz_dep_.CoverTab[45187]++
									b.cell.width++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:453
		// _ = "end of CoverTab[45187]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:453
	default:
//line /usr/local/go/src/text/tabwriter/tabwriter.go:453
		_go_fuzz_dep_.CoverTab[45188]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:453
		// _ = "end of CoverTab[45188]"
	}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:454
	// _ = "end of CoverTab[45183]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:454
	_go_fuzz_dep_.CoverTab[45184]++
								b.pos = len(b.buf)
								b.endChar = 0
//line /usr/local/go/src/text/tabwriter/tabwriter.go:456
	// _ = "end of CoverTab[45184]"
}

// Terminate the current cell by adding it to the list of cells of the
//line /usr/local/go/src/text/tabwriter/tabwriter.go:459
// current line. Returns the number of cells in that line.
//line /usr/local/go/src/text/tabwriter/tabwriter.go:461
func (b *Writer) terminateCell(htab bool) int {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:461
	_go_fuzz_dep_.CoverTab[45191]++
								b.cell.htab = htab
								line := &b.lines[len(b.lines)-1]
								*line = append(*line, b.cell)
								b.cell = cell{}
								return len(*line)
//line /usr/local/go/src/text/tabwriter/tabwriter.go:466
	// _ = "end of CoverTab[45191]"
}

func (b *Writer) handlePanic(err *error, op string) {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:469
	_go_fuzz_dep_.CoverTab[45192]++
								if e := recover(); e != nil {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:470
		_go_fuzz_dep_.CoverTab[45193]++
									if op == "Flush" {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:471
			_go_fuzz_dep_.CoverTab[45196]++

										b.reset()
//line /usr/local/go/src/text/tabwriter/tabwriter.go:473
			// _ = "end of CoverTab[45196]"
		} else {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:474
			_go_fuzz_dep_.CoverTab[45197]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:474
			// _ = "end of CoverTab[45197]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:474
		}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:474
		// _ = "end of CoverTab[45193]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:474
		_go_fuzz_dep_.CoverTab[45194]++
									if nerr, ok := e.(osError); ok {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:475
			_go_fuzz_dep_.CoverTab[45198]++
										*err = nerr.err
										return
//line /usr/local/go/src/text/tabwriter/tabwriter.go:477
			// _ = "end of CoverTab[45198]"
		} else {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:478
			_go_fuzz_dep_.CoverTab[45199]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:478
			// _ = "end of CoverTab[45199]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:478
		}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:478
		// _ = "end of CoverTab[45194]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:478
		_go_fuzz_dep_.CoverTab[45195]++
									panic("tabwriter: panic during " + op)
//line /usr/local/go/src/text/tabwriter/tabwriter.go:479
		// _ = "end of CoverTab[45195]"
	} else {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:480
		_go_fuzz_dep_.CoverTab[45200]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:480
		// _ = "end of CoverTab[45200]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:480
	}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:480
	// _ = "end of CoverTab[45192]"
}

// Flush should be called after the last call to Write to ensure
//line /usr/local/go/src/text/tabwriter/tabwriter.go:483
// that any data buffered in the Writer is written to output. Any
//line /usr/local/go/src/text/tabwriter/tabwriter.go:483
// incomplete escape sequence at the end is considered
//line /usr/local/go/src/text/tabwriter/tabwriter.go:483
// complete for formatting purposes.
//line /usr/local/go/src/text/tabwriter/tabwriter.go:487
func (b *Writer) Flush() error {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:487
	_go_fuzz_dep_.CoverTab[45201]++
								return b.flush()
//line /usr/local/go/src/text/tabwriter/tabwriter.go:488
	// _ = "end of CoverTab[45201]"
}

// flush is the internal version of Flush, with a named return value which we
//line /usr/local/go/src/text/tabwriter/tabwriter.go:491
// don't want to expose.
//line /usr/local/go/src/text/tabwriter/tabwriter.go:493
func (b *Writer) flush() (err error) {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:493
	_go_fuzz_dep_.CoverTab[45202]++
								defer b.handlePanic(&err, "Flush")
								b.flushNoDefers()
								return nil
//line /usr/local/go/src/text/tabwriter/tabwriter.go:496
	// _ = "end of CoverTab[45202]"
}

// flushNoDefers is like flush, but without a deferred handlePanic call. This
//line /usr/local/go/src/text/tabwriter/tabwriter.go:499
// can be called from other methods which already have their own deferred
//line /usr/local/go/src/text/tabwriter/tabwriter.go:499
// handlePanic calls, such as Write, and avoid the extra defer work.
//line /usr/local/go/src/text/tabwriter/tabwriter.go:502
func (b *Writer) flushNoDefers() {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:502
	_go_fuzz_dep_.CoverTab[45203]++

								if b.cell.size > 0 {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:504
		_go_fuzz_dep_.CoverTab[45205]++
									if b.endChar != 0 {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:505
			_go_fuzz_dep_.CoverTab[45207]++

										b.endEscape()
//line /usr/local/go/src/text/tabwriter/tabwriter.go:507
			// _ = "end of CoverTab[45207]"
		} else {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:508
			_go_fuzz_dep_.CoverTab[45208]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:508
			// _ = "end of CoverTab[45208]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:508
		}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:508
		// _ = "end of CoverTab[45205]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:508
		_go_fuzz_dep_.CoverTab[45206]++
									b.terminateCell(false)
//line /usr/local/go/src/text/tabwriter/tabwriter.go:509
		// _ = "end of CoverTab[45206]"
	} else {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:510
		_go_fuzz_dep_.CoverTab[45209]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:510
		// _ = "end of CoverTab[45209]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:510
	}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:510
	// _ = "end of CoverTab[45203]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:510
	_go_fuzz_dep_.CoverTab[45204]++

//line /usr/local/go/src/text/tabwriter/tabwriter.go:513
	b.format(0, 0, len(b.lines))
								b.reset()
//line /usr/local/go/src/text/tabwriter/tabwriter.go:514
	// _ = "end of CoverTab[45204]"
}

var hbar = []byte("---\n")

// Write writes buf to the writer b.
//line /usr/local/go/src/text/tabwriter/tabwriter.go:519
// The only errors returned are ones encountered
//line /usr/local/go/src/text/tabwriter/tabwriter.go:519
// while writing to the underlying output stream.
//line /usr/local/go/src/text/tabwriter/tabwriter.go:522
func (b *Writer) Write(buf []byte) (n int, err error) {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:522
	_go_fuzz_dep_.CoverTab[45210]++
								defer b.handlePanic(&err, "Write")

//line /usr/local/go/src/text/tabwriter/tabwriter.go:526
	n = 0
	for i, ch := range buf {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:527
		_go_fuzz_dep_.CoverTab[45212]++
									if b.endChar == 0 {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:528
			_go_fuzz_dep_.CoverTab[45213]++

										switch ch {
			case '\t', '\v', '\n', '\f':
//line /usr/local/go/src/text/tabwriter/tabwriter.go:531
				_go_fuzz_dep_.CoverTab[45214]++

											b.append(buf[n:i])
											b.updateWidth()
											n = i + 1
											ncells := b.terminateCell(ch == '\t')
											if ch == '\n' || func() bool {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:537
					_go_fuzz_dep_.CoverTab[45219]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:537
					return ch == '\f'
//line /usr/local/go/src/text/tabwriter/tabwriter.go:537
					// _ = "end of CoverTab[45219]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:537
				}() {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:537
					_go_fuzz_dep_.CoverTab[45220]++

												b.addLine(ch == '\f')
												if ch == '\f' || func() bool {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:540
						_go_fuzz_dep_.CoverTab[45221]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:540
						return ncells == 1
//line /usr/local/go/src/text/tabwriter/tabwriter.go:540
						// _ = "end of CoverTab[45221]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:540
					}() {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:540
						_go_fuzz_dep_.CoverTab[45222]++

//line /usr/local/go/src/text/tabwriter/tabwriter.go:546
						b.flushNoDefers()
						if ch == '\f' && func() bool {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:547
							_go_fuzz_dep_.CoverTab[45223]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:547
							return b.flags&Debug != 0
//line /usr/local/go/src/text/tabwriter/tabwriter.go:547
							// _ = "end of CoverTab[45223]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:547
						}() {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:547
							_go_fuzz_dep_.CoverTab[45224]++

														b.write0(hbar)
//line /usr/local/go/src/text/tabwriter/tabwriter.go:549
							// _ = "end of CoverTab[45224]"
						} else {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:550
							_go_fuzz_dep_.CoverTab[45225]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:550
							// _ = "end of CoverTab[45225]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:550
						}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:550
						// _ = "end of CoverTab[45222]"
					} else {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:551
						_go_fuzz_dep_.CoverTab[45226]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:551
						// _ = "end of CoverTab[45226]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:551
					}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:551
					// _ = "end of CoverTab[45220]"
				} else {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:552
					_go_fuzz_dep_.CoverTab[45227]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:552
					// _ = "end of CoverTab[45227]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:552
				}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:552
				// _ = "end of CoverTab[45214]"

			case Escape:
//line /usr/local/go/src/text/tabwriter/tabwriter.go:554
				_go_fuzz_dep_.CoverTab[45215]++

											b.append(buf[n:i])
											b.updateWidth()
											n = i
											if b.flags&StripEscape != 0 {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:559
					_go_fuzz_dep_.CoverTab[45228]++
												n++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:560
					// _ = "end of CoverTab[45228]"
				} else {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:561
					_go_fuzz_dep_.CoverTab[45229]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:561
					// _ = "end of CoverTab[45229]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:561
				}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:561
				// _ = "end of CoverTab[45215]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:561
				_go_fuzz_dep_.CoverTab[45216]++
											b.startEscape(Escape)
//line /usr/local/go/src/text/tabwriter/tabwriter.go:562
				// _ = "end of CoverTab[45216]"

			case '<', '&':
//line /usr/local/go/src/text/tabwriter/tabwriter.go:564
				_go_fuzz_dep_.CoverTab[45217]++

											if b.flags&FilterHTML != 0 {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:566
					_go_fuzz_dep_.CoverTab[45230]++

												b.append(buf[n:i])
												b.updateWidth()
												n = i
												b.startEscape(ch)
//line /usr/local/go/src/text/tabwriter/tabwriter.go:571
					// _ = "end of CoverTab[45230]"
				} else {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:572
					_go_fuzz_dep_.CoverTab[45231]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:572
					// _ = "end of CoverTab[45231]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:572
				}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:572
				// _ = "end of CoverTab[45217]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:572
			default:
//line /usr/local/go/src/text/tabwriter/tabwriter.go:572
				_go_fuzz_dep_.CoverTab[45218]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:572
				// _ = "end of CoverTab[45218]"
			}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:573
			// _ = "end of CoverTab[45213]"

		} else {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:575
			_go_fuzz_dep_.CoverTab[45232]++

										if ch == b.endChar {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:577
				_go_fuzz_dep_.CoverTab[45233]++

											j := i + 1
											if ch == Escape && func() bool {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:580
					_go_fuzz_dep_.CoverTab[45235]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:580
					return b.flags&StripEscape != 0
//line /usr/local/go/src/text/tabwriter/tabwriter.go:580
					// _ = "end of CoverTab[45235]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:580
				}() {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:580
					_go_fuzz_dep_.CoverTab[45236]++
												j = i
//line /usr/local/go/src/text/tabwriter/tabwriter.go:581
					// _ = "end of CoverTab[45236]"
				} else {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:582
					_go_fuzz_dep_.CoverTab[45237]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:582
					// _ = "end of CoverTab[45237]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:582
				}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:582
				// _ = "end of CoverTab[45233]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:582
				_go_fuzz_dep_.CoverTab[45234]++
											b.append(buf[n:j])
											n = i + 1
											b.endEscape()
//line /usr/local/go/src/text/tabwriter/tabwriter.go:585
				// _ = "end of CoverTab[45234]"
			} else {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:586
				_go_fuzz_dep_.CoverTab[45238]++
//line /usr/local/go/src/text/tabwriter/tabwriter.go:586
				// _ = "end of CoverTab[45238]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:586
			}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:586
			// _ = "end of CoverTab[45232]"
		}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:587
		// _ = "end of CoverTab[45212]"
	}
//line /usr/local/go/src/text/tabwriter/tabwriter.go:588
	// _ = "end of CoverTab[45210]"
//line /usr/local/go/src/text/tabwriter/tabwriter.go:588
	_go_fuzz_dep_.CoverTab[45211]++

//line /usr/local/go/src/text/tabwriter/tabwriter.go:591
	b.append(buf[n:])
								n = len(buf)
								return
//line /usr/local/go/src/text/tabwriter/tabwriter.go:593
	// _ = "end of CoverTab[45211]"
}

// NewWriter allocates and initializes a new tabwriter.Writer.
//line /usr/local/go/src/text/tabwriter/tabwriter.go:596
// The parameters are the same as for the Init function.
//line /usr/local/go/src/text/tabwriter/tabwriter.go:598
func NewWriter(output io.Writer, minwidth, tabwidth, padding int, padchar byte, flags uint) *Writer {
//line /usr/local/go/src/text/tabwriter/tabwriter.go:598
	_go_fuzz_dep_.CoverTab[45239]++
								return new(Writer).Init(output, minwidth, tabwidth, padding, padchar, flags)
//line /usr/local/go/src/text/tabwriter/tabwriter.go:599
	// _ = "end of CoverTab[45239]"
}

//line /usr/local/go/src/text/tabwriter/tabwriter.go:600
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/text/tabwriter/tabwriter.go:600
var _ = _go_fuzz_dep_.CoverTab
