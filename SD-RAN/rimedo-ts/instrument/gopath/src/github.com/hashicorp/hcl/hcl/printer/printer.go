//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:1
// Package printer implements printing of AST nodes to HCL format.
package printer

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:2
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:2
)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:2
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:2
)

import (
	"bytes"
	"io"
	"text/tabwriter"

	"github.com/hashicorp/hcl/hcl/ast"
	"github.com/hashicorp/hcl/hcl/parser"
)

var DefaultConfig = Config{
	SpacesWidth: 2,
}

// A Config node controls the output of Fprint.
type Config struct {
	SpacesWidth int	// if set, it will use spaces instead of tabs for alignment
}

func (c *Config) Fprint(output io.Writer, node ast.Node) error {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:22
	_go_fuzz_dep_.CoverTab[122468]++
												p := &printer{
		cfg:			*c,
		comments:		make([]*ast.CommentGroup, 0),
		standaloneComments:	make([]*ast.CommentGroup, 0),
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:28
	}

	p.collectComments(node)

	if _, err := output.Write(p.unindent(p.output(node))); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:32
		_go_fuzz_dep_.CoverTab[122471]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:33
		// _ = "end of CoverTab[122471]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:34
		_go_fuzz_dep_.CoverTab[122472]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:34
		// _ = "end of CoverTab[122472]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:34
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:34
	// _ = "end of CoverTab[122468]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:34
	_go_fuzz_dep_.CoverTab[122469]++

	// flush tabwriter, if any
	var err error
	if tw, _ := output.(*tabwriter.Writer); tw != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:38
		_go_fuzz_dep_.CoverTab[122473]++
													err = tw.Flush()
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:39
		// _ = "end of CoverTab[122473]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:40
		_go_fuzz_dep_.CoverTab[122474]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:40
		// _ = "end of CoverTab[122474]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:40
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:40
	// _ = "end of CoverTab[122469]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:40
	_go_fuzz_dep_.CoverTab[122470]++

												return err
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:42
	// _ = "end of CoverTab[122470]"
}

// Fprint "pretty-prints" an HCL node to output
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:45
// It calls Config.Fprint with default settings.
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:47
func Fprint(output io.Writer, node ast.Node) error {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:47
	_go_fuzz_dep_.CoverTab[122475]++
												return DefaultConfig.Fprint(output, node)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:48
	// _ = "end of CoverTab[122475]"
}

// Format formats src HCL and returns the result.
func Format(src []byte) ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:52
	_go_fuzz_dep_.CoverTab[122476]++
												node, err := parser.Parse(src)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:54
		_go_fuzz_dep_.CoverTab[122479]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:55
		// _ = "end of CoverTab[122479]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:56
		_go_fuzz_dep_.CoverTab[122480]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:56
		// _ = "end of CoverTab[122480]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:56
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:56
	// _ = "end of CoverTab[122476]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:56
	_go_fuzz_dep_.CoverTab[122477]++

												var buf bytes.Buffer
												if err := DefaultConfig.Fprint(&buf, node); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:59
		_go_fuzz_dep_.CoverTab[122481]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:60
		// _ = "end of CoverTab[122481]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:61
		_go_fuzz_dep_.CoverTab[122482]++
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:61
		// _ = "end of CoverTab[122482]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:61
	}
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:61
	// _ = "end of CoverTab[122477]"
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:61
	_go_fuzz_dep_.CoverTab[122478]++

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:64
	buf.WriteString("\n")
												return buf.Bytes(), nil
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:65
	// _ = "end of CoverTab[122478]"
}

//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:66
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/hashicorp/hcl@v1.0.0/hcl/printer/printer.go:66
var _ = _go_fuzz_dep_.CoverTab
