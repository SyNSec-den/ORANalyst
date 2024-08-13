// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/html/template/escape.go:5
package template

//line /usr/local/go/src/html/template/escape.go:5
import (
//line /usr/local/go/src/html/template/escape.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/html/template/escape.go:5
)
//line /usr/local/go/src/html/template/escape.go:5
import (
//line /usr/local/go/src/html/template/escape.go:5
	_atomic_ "sync/atomic"
//line /usr/local/go/src/html/template/escape.go:5
)

import (
	"bytes"
	"fmt"
	"html"
	"internal/godebug"
	"io"
	"text/template"
	"text/template/parse"
)

// escapeTemplate rewrites the named template, which must be
//line /usr/local/go/src/html/template/escape.go:17
// associated with t, to guarantee that the output of any of the named
//line /usr/local/go/src/html/template/escape.go:17
// templates is properly escaped. If no error is returned, then the named templates have
//line /usr/local/go/src/html/template/escape.go:17
// been modified. Otherwise the named templates have been rendered
//line /usr/local/go/src/html/template/escape.go:17
// unusable.
//line /usr/local/go/src/html/template/escape.go:22
func escapeTemplate(tmpl *Template, node parse.Node, name string) error {
//line /usr/local/go/src/html/template/escape.go:22
	_go_fuzz_dep_.CoverTab[30860]++
							c, _ := tmpl.esc.escapeTree(context{}, node, name, 0)
							var err error
							if c.err != nil {
//line /usr/local/go/src/html/template/escape.go:25
		_go_fuzz_dep_.CoverTab[30864]++
								err, c.err.Name = c.err, name
//line /usr/local/go/src/html/template/escape.go:26
		// _ = "end of CoverTab[30864]"
	} else {
//line /usr/local/go/src/html/template/escape.go:27
		_go_fuzz_dep_.CoverTab[30865]++
//line /usr/local/go/src/html/template/escape.go:27
		if c.state != stateText {
//line /usr/local/go/src/html/template/escape.go:27
			_go_fuzz_dep_.CoverTab[30866]++
									err = &Error{ErrEndContext, nil, name, 0, fmt.Sprintf("ends in a non-text context: %v", c)}
//line /usr/local/go/src/html/template/escape.go:28
			// _ = "end of CoverTab[30866]"
		} else {
//line /usr/local/go/src/html/template/escape.go:29
			_go_fuzz_dep_.CoverTab[30867]++
//line /usr/local/go/src/html/template/escape.go:29
			// _ = "end of CoverTab[30867]"
//line /usr/local/go/src/html/template/escape.go:29
		}
//line /usr/local/go/src/html/template/escape.go:29
		// _ = "end of CoverTab[30865]"
//line /usr/local/go/src/html/template/escape.go:29
	}
//line /usr/local/go/src/html/template/escape.go:29
	// _ = "end of CoverTab[30860]"
//line /usr/local/go/src/html/template/escape.go:29
	_go_fuzz_dep_.CoverTab[30861]++
							if err != nil {
//line /usr/local/go/src/html/template/escape.go:30
		_go_fuzz_dep_.CoverTab[30868]++

								if t := tmpl.set[name]; t != nil {
//line /usr/local/go/src/html/template/escape.go:32
			_go_fuzz_dep_.CoverTab[30870]++
									t.escapeErr = err
									t.text.Tree = nil
									t.Tree = nil
//line /usr/local/go/src/html/template/escape.go:35
			// _ = "end of CoverTab[30870]"
		} else {
//line /usr/local/go/src/html/template/escape.go:36
			_go_fuzz_dep_.CoverTab[30871]++
//line /usr/local/go/src/html/template/escape.go:36
			// _ = "end of CoverTab[30871]"
//line /usr/local/go/src/html/template/escape.go:36
		}
//line /usr/local/go/src/html/template/escape.go:36
		// _ = "end of CoverTab[30868]"
//line /usr/local/go/src/html/template/escape.go:36
		_go_fuzz_dep_.CoverTab[30869]++
								return err
//line /usr/local/go/src/html/template/escape.go:37
		// _ = "end of CoverTab[30869]"
	} else {
//line /usr/local/go/src/html/template/escape.go:38
		_go_fuzz_dep_.CoverTab[30872]++
//line /usr/local/go/src/html/template/escape.go:38
		// _ = "end of CoverTab[30872]"
//line /usr/local/go/src/html/template/escape.go:38
	}
//line /usr/local/go/src/html/template/escape.go:38
	// _ = "end of CoverTab[30861]"
//line /usr/local/go/src/html/template/escape.go:38
	_go_fuzz_dep_.CoverTab[30862]++
							tmpl.esc.commit()
							if t := tmpl.set[name]; t != nil {
//line /usr/local/go/src/html/template/escape.go:40
		_go_fuzz_dep_.CoverTab[30873]++
								t.escapeErr = escapeOK
								t.Tree = t.text.Tree
//line /usr/local/go/src/html/template/escape.go:42
		// _ = "end of CoverTab[30873]"
	} else {
//line /usr/local/go/src/html/template/escape.go:43
		_go_fuzz_dep_.CoverTab[30874]++
//line /usr/local/go/src/html/template/escape.go:43
		// _ = "end of CoverTab[30874]"
//line /usr/local/go/src/html/template/escape.go:43
	}
//line /usr/local/go/src/html/template/escape.go:43
	// _ = "end of CoverTab[30862]"
//line /usr/local/go/src/html/template/escape.go:43
	_go_fuzz_dep_.CoverTab[30863]++
							return nil
//line /usr/local/go/src/html/template/escape.go:44
	// _ = "end of CoverTab[30863]"
}

// evalArgs formats the list of arguments into a string. It is equivalent to
//line /usr/local/go/src/html/template/escape.go:47
// fmt.Sprint(args...), except that it dereferences all pointers.
//line /usr/local/go/src/html/template/escape.go:49
func evalArgs(args ...any) string {
//line /usr/local/go/src/html/template/escape.go:49
	_go_fuzz_dep_.CoverTab[30875]++

							if len(args) == 1 {
//line /usr/local/go/src/html/template/escape.go:51
		_go_fuzz_dep_.CoverTab[30878]++
								if s, ok := args[0].(string); ok {
//line /usr/local/go/src/html/template/escape.go:52
			_go_fuzz_dep_.CoverTab[30879]++
									return s
//line /usr/local/go/src/html/template/escape.go:53
			// _ = "end of CoverTab[30879]"
		} else {
//line /usr/local/go/src/html/template/escape.go:54
			_go_fuzz_dep_.CoverTab[30880]++
//line /usr/local/go/src/html/template/escape.go:54
			// _ = "end of CoverTab[30880]"
//line /usr/local/go/src/html/template/escape.go:54
		}
//line /usr/local/go/src/html/template/escape.go:54
		// _ = "end of CoverTab[30878]"
	} else {
//line /usr/local/go/src/html/template/escape.go:55
		_go_fuzz_dep_.CoverTab[30881]++
//line /usr/local/go/src/html/template/escape.go:55
		// _ = "end of CoverTab[30881]"
//line /usr/local/go/src/html/template/escape.go:55
	}
//line /usr/local/go/src/html/template/escape.go:55
	// _ = "end of CoverTab[30875]"
//line /usr/local/go/src/html/template/escape.go:55
	_go_fuzz_dep_.CoverTab[30876]++
							for i, arg := range args {
//line /usr/local/go/src/html/template/escape.go:56
		_go_fuzz_dep_.CoverTab[30882]++
								args[i] = indirectToStringerOrError(arg)
//line /usr/local/go/src/html/template/escape.go:57
		// _ = "end of CoverTab[30882]"
	}
//line /usr/local/go/src/html/template/escape.go:58
	// _ = "end of CoverTab[30876]"
//line /usr/local/go/src/html/template/escape.go:58
	_go_fuzz_dep_.CoverTab[30877]++
							return fmt.Sprint(args...)
//line /usr/local/go/src/html/template/escape.go:59
	// _ = "end of CoverTab[30877]"
}

// funcMap maps command names to functions that render their inputs safe.
var funcMap = template.FuncMap{
	"_html_template_attrescaper":		attrEscaper,
	"_html_template_commentescaper":	commentEscaper,
	"_html_template_cssescaper":		cssEscaper,
	"_html_template_cssvaluefilter":	cssValueFilter,
	"_html_template_htmlnamefilter":	htmlNameFilter,
	"_html_template_htmlescaper":		htmlEscaper,
	"_html_template_jsregexpescaper":	jsRegexpEscaper,
	"_html_template_jsstrescaper":		jsStrEscaper,
	"_html_template_jsvalescaper":		jsValEscaper,
	"_html_template_nospaceescaper":	htmlNospaceEscaper,
	"_html_template_rcdataescaper":		rcdataEscaper,
	"_html_template_srcsetescaper":		srcsetFilterAndEscaper,
	"_html_template_urlescaper":		urlEscaper,
	"_html_template_urlfilter":		urlFilter,
	"_html_template_urlnormalizer":		urlNormalizer,
	"_eval_args_":				evalArgs,
}

// escaper collects type inferences about templates and changes needed to make
//line /usr/local/go/src/html/template/escape.go:82
// templates injection safe.
//line /usr/local/go/src/html/template/escape.go:84
type escaper struct {
	// ns is the nameSpace that this escaper is associated with.
	ns	*nameSpace
	// output[templateName] is the output context for a templateName that
	// has been mangled to include its input context.
	output	map[string]context
	// derived[c.mangle(name)] maps to a template derived from the template
	// named name templateName for the start context c.
	derived	map[string]*template.Template
	// called[templateName] is a set of called mangled template names.
	called	map[string]bool
	// xxxNodeEdits are the accumulated edits to apply during commit.
	// Such edits are not applied immediately in case a template set
	// executes a given template in different escaping contexts.
	actionNodeEdits		map[*parse.ActionNode][]string
	templateNodeEdits	map[*parse.TemplateNode]string
	textNodeEdits		map[*parse.TextNode][]byte
	// rangeContext holds context about the current range loop.
	rangeContext	*rangeContext
}

// rangeContext holds information about the current range loop.
type rangeContext struct {
	outer		*rangeContext	// outer loop
	breaks		[]context	// context at each break action
	continues	[]context	// context at each continue action
}

// makeEscaper creates a blank escaper for the given set.
func makeEscaper(n *nameSpace) escaper {
//line /usr/local/go/src/html/template/escape.go:113
	_go_fuzz_dep_.CoverTab[30883]++
							return escaper{
		n,
		map[string]context{},
		map[string]*template.Template{},
		map[string]bool{},
		map[*parse.ActionNode][]string{},
		map[*parse.TemplateNode]string{},
		map[*parse.TextNode][]byte{},
		nil,
	}
//line /usr/local/go/src/html/template/escape.go:123
	// _ = "end of CoverTab[30883]"
}

// filterFailsafe is an innocuous word that is emitted in place of unsafe values
//line /usr/local/go/src/html/template/escape.go:126
// by sanitizer functions. It is not a keyword in any programming language,
//line /usr/local/go/src/html/template/escape.go:126
// contains no special characters, is not empty, and when it appears in output
//line /usr/local/go/src/html/template/escape.go:126
// it is distinct enough that a developer can find the source of the problem
//line /usr/local/go/src/html/template/escape.go:126
// via a search engine.
//line /usr/local/go/src/html/template/escape.go:131
const filterFailsafe = "ZgotmplZ"

// escape escapes a template node.
func (e *escaper) escape(c context, n parse.Node) context {
//line /usr/local/go/src/html/template/escape.go:134
	_go_fuzz_dep_.CoverTab[30884]++
							switch n := n.(type) {
	case *parse.ActionNode:
//line /usr/local/go/src/html/template/escape.go:136
		_go_fuzz_dep_.CoverTab[30886]++
								return e.escapeAction(c, n)
//line /usr/local/go/src/html/template/escape.go:137
		// _ = "end of CoverTab[30886]"
	case *parse.BreakNode:
//line /usr/local/go/src/html/template/escape.go:138
		_go_fuzz_dep_.CoverTab[30887]++
								c.n = n
								e.rangeContext.breaks = append(e.rangeContext.breaks, c)
								return context{state: stateDead}
//line /usr/local/go/src/html/template/escape.go:141
		// _ = "end of CoverTab[30887]"
	case *parse.CommentNode:
//line /usr/local/go/src/html/template/escape.go:142
		_go_fuzz_dep_.CoverTab[30888]++
								return c
//line /usr/local/go/src/html/template/escape.go:143
		// _ = "end of CoverTab[30888]"
	case *parse.ContinueNode:
//line /usr/local/go/src/html/template/escape.go:144
		_go_fuzz_dep_.CoverTab[30889]++
								c.n = n
								e.rangeContext.continues = append(e.rangeContext.breaks, c)
								return context{state: stateDead}
//line /usr/local/go/src/html/template/escape.go:147
		// _ = "end of CoverTab[30889]"
	case *parse.IfNode:
//line /usr/local/go/src/html/template/escape.go:148
		_go_fuzz_dep_.CoverTab[30890]++
								return e.escapeBranch(c, &n.BranchNode, "if")
//line /usr/local/go/src/html/template/escape.go:149
		// _ = "end of CoverTab[30890]"
	case *parse.ListNode:
//line /usr/local/go/src/html/template/escape.go:150
		_go_fuzz_dep_.CoverTab[30891]++
								return e.escapeList(c, n)
//line /usr/local/go/src/html/template/escape.go:151
		// _ = "end of CoverTab[30891]"
	case *parse.RangeNode:
//line /usr/local/go/src/html/template/escape.go:152
		_go_fuzz_dep_.CoverTab[30892]++
								return e.escapeBranch(c, &n.BranchNode, "range")
//line /usr/local/go/src/html/template/escape.go:153
		// _ = "end of CoverTab[30892]"
	case *parse.TemplateNode:
//line /usr/local/go/src/html/template/escape.go:154
		_go_fuzz_dep_.CoverTab[30893]++
								return e.escapeTemplate(c, n)
//line /usr/local/go/src/html/template/escape.go:155
		// _ = "end of CoverTab[30893]"
	case *parse.TextNode:
//line /usr/local/go/src/html/template/escape.go:156
		_go_fuzz_dep_.CoverTab[30894]++
								return e.escapeText(c, n)
//line /usr/local/go/src/html/template/escape.go:157
		// _ = "end of CoverTab[30894]"
	case *parse.WithNode:
//line /usr/local/go/src/html/template/escape.go:158
		_go_fuzz_dep_.CoverTab[30895]++
								return e.escapeBranch(c, &n.BranchNode, "with")
//line /usr/local/go/src/html/template/escape.go:159
		// _ = "end of CoverTab[30895]"
	}
//line /usr/local/go/src/html/template/escape.go:160
	// _ = "end of CoverTab[30884]"
//line /usr/local/go/src/html/template/escape.go:160
	_go_fuzz_dep_.CoverTab[30885]++
							panic("escaping " + n.String() + " is unimplemented")
//line /usr/local/go/src/html/template/escape.go:161
	// _ = "end of CoverTab[30885]"
}

var debugAllowActionJSTmpl = godebug.New("jstmpllitinterp")

// escapeAction escapes an action template node.
func (e *escaper) escapeAction(c context, n *parse.ActionNode) context {
//line /usr/local/go/src/html/template/escape.go:167
	_go_fuzz_dep_.CoverTab[30896]++
							if len(n.Pipe.Decl) != 0 {
//line /usr/local/go/src/html/template/escape.go:168
		_go_fuzz_dep_.CoverTab[30901]++

								return c
//line /usr/local/go/src/html/template/escape.go:170
		// _ = "end of CoverTab[30901]"
	} else {
//line /usr/local/go/src/html/template/escape.go:171
		_go_fuzz_dep_.CoverTab[30902]++
//line /usr/local/go/src/html/template/escape.go:171
		// _ = "end of CoverTab[30902]"
//line /usr/local/go/src/html/template/escape.go:171
	}
//line /usr/local/go/src/html/template/escape.go:171
	// _ = "end of CoverTab[30896]"
//line /usr/local/go/src/html/template/escape.go:171
	_go_fuzz_dep_.CoverTab[30897]++
							c = nudge(c)

							for pos, idNode := range n.Pipe.Cmds {
//line /usr/local/go/src/html/template/escape.go:174
		_go_fuzz_dep_.CoverTab[30903]++
								node, ok := idNode.Args[0].(*parse.IdentifierNode)
								if !ok {
//line /usr/local/go/src/html/template/escape.go:176
			_go_fuzz_dep_.CoverTab[30905]++

//line /usr/local/go/src/html/template/escape.go:184
			continue
//line /usr/local/go/src/html/template/escape.go:184
			// _ = "end of CoverTab[30905]"
		} else {
//line /usr/local/go/src/html/template/escape.go:185
			_go_fuzz_dep_.CoverTab[30906]++
//line /usr/local/go/src/html/template/escape.go:185
			// _ = "end of CoverTab[30906]"
//line /usr/local/go/src/html/template/escape.go:185
		}
//line /usr/local/go/src/html/template/escape.go:185
		// _ = "end of CoverTab[30903]"
//line /usr/local/go/src/html/template/escape.go:185
		_go_fuzz_dep_.CoverTab[30904]++
								ident := node.Ident
								if _, ok := predefinedEscapers[ident]; ok {
//line /usr/local/go/src/html/template/escape.go:187
			_go_fuzz_dep_.CoverTab[30907]++
									if pos < len(n.Pipe.Cmds)-1 || func() bool {
//line /usr/local/go/src/html/template/escape.go:188
				_go_fuzz_dep_.CoverTab[30908]++
//line /usr/local/go/src/html/template/escape.go:188
				return c.state == stateAttr && func() bool {
											_go_fuzz_dep_.CoverTab[30909]++
//line /usr/local/go/src/html/template/escape.go:189
					return c.delim == delimSpaceOrTagEnd
//line /usr/local/go/src/html/template/escape.go:189
					// _ = "end of CoverTab[30909]"
//line /usr/local/go/src/html/template/escape.go:189
				}() && func() bool {
//line /usr/local/go/src/html/template/escape.go:189
					_go_fuzz_dep_.CoverTab[30910]++
//line /usr/local/go/src/html/template/escape.go:189
					return ident == "html"
//line /usr/local/go/src/html/template/escape.go:189
					// _ = "end of CoverTab[30910]"
//line /usr/local/go/src/html/template/escape.go:189
				}()
//line /usr/local/go/src/html/template/escape.go:189
				// _ = "end of CoverTab[30908]"
//line /usr/local/go/src/html/template/escape.go:189
			}() {
//line /usr/local/go/src/html/template/escape.go:189
				_go_fuzz_dep_.CoverTab[30911]++
										return context{
					state:	stateError,
					err:	errorf(ErrPredefinedEscaper, n, n.Line, "predefined escaper %q disallowed in template", ident),
				}
//line /usr/local/go/src/html/template/escape.go:193
				// _ = "end of CoverTab[30911]"
			} else {
//line /usr/local/go/src/html/template/escape.go:194
				_go_fuzz_dep_.CoverTab[30912]++
//line /usr/local/go/src/html/template/escape.go:194
				// _ = "end of CoverTab[30912]"
//line /usr/local/go/src/html/template/escape.go:194
			}
//line /usr/local/go/src/html/template/escape.go:194
			// _ = "end of CoverTab[30907]"
		} else {
//line /usr/local/go/src/html/template/escape.go:195
			_go_fuzz_dep_.CoverTab[30913]++
//line /usr/local/go/src/html/template/escape.go:195
			// _ = "end of CoverTab[30913]"
//line /usr/local/go/src/html/template/escape.go:195
		}
//line /usr/local/go/src/html/template/escape.go:195
		// _ = "end of CoverTab[30904]"
	}
//line /usr/local/go/src/html/template/escape.go:196
	// _ = "end of CoverTab[30897]"
//line /usr/local/go/src/html/template/escape.go:196
	_go_fuzz_dep_.CoverTab[30898]++
							s := make([]string, 0, 3)
							switch c.state {
	case stateError:
//line /usr/local/go/src/html/template/escape.go:199
		_go_fuzz_dep_.CoverTab[30914]++
								return c
//line /usr/local/go/src/html/template/escape.go:200
		// _ = "end of CoverTab[30914]"
	case stateURL, stateCSSDqStr, stateCSSSqStr, stateCSSDqURL, stateCSSSqURL, stateCSSURL:
//line /usr/local/go/src/html/template/escape.go:201
		_go_fuzz_dep_.CoverTab[30915]++
								switch c.urlPart {
		case urlPartNone:
//line /usr/local/go/src/html/template/escape.go:203
			_go_fuzz_dep_.CoverTab[30927]++
									s = append(s, "_html_template_urlfilter")
									fallthrough
//line /usr/local/go/src/html/template/escape.go:205
			// _ = "end of CoverTab[30927]"
		case urlPartPreQuery:
//line /usr/local/go/src/html/template/escape.go:206
			_go_fuzz_dep_.CoverTab[30928]++
									switch c.state {
			case stateCSSDqStr, stateCSSSqStr:
//line /usr/local/go/src/html/template/escape.go:208
				_go_fuzz_dep_.CoverTab[30932]++
										s = append(s, "_html_template_cssescaper")
//line /usr/local/go/src/html/template/escape.go:209
				// _ = "end of CoverTab[30932]"
			default:
//line /usr/local/go/src/html/template/escape.go:210
				_go_fuzz_dep_.CoverTab[30933]++
										s = append(s, "_html_template_urlnormalizer")
//line /usr/local/go/src/html/template/escape.go:211
				// _ = "end of CoverTab[30933]"
			}
//line /usr/local/go/src/html/template/escape.go:212
			// _ = "end of CoverTab[30928]"
		case urlPartQueryOrFrag:
//line /usr/local/go/src/html/template/escape.go:213
			_go_fuzz_dep_.CoverTab[30929]++
									s = append(s, "_html_template_urlescaper")
//line /usr/local/go/src/html/template/escape.go:214
			// _ = "end of CoverTab[30929]"
		case urlPartUnknown:
//line /usr/local/go/src/html/template/escape.go:215
			_go_fuzz_dep_.CoverTab[30930]++
									return context{
				state:	stateError,
				err:	errorf(ErrAmbigContext, n, n.Line, "%s appears in an ambiguous context within a URL", n),
			}
//line /usr/local/go/src/html/template/escape.go:219
			// _ = "end of CoverTab[30930]"
		default:
//line /usr/local/go/src/html/template/escape.go:220
			_go_fuzz_dep_.CoverTab[30931]++
									panic(c.urlPart.String())
//line /usr/local/go/src/html/template/escape.go:221
			// _ = "end of CoverTab[30931]"
		}
//line /usr/local/go/src/html/template/escape.go:222
		// _ = "end of CoverTab[30915]"
	case stateJS:
//line /usr/local/go/src/html/template/escape.go:223
		_go_fuzz_dep_.CoverTab[30916]++
								s = append(s, "_html_template_jsvalescaper")

								c.jsCtx = jsCtxDivOp
//line /usr/local/go/src/html/template/escape.go:226
		// _ = "end of CoverTab[30916]"
	case stateJSDqStr, stateJSSqStr:
//line /usr/local/go/src/html/template/escape.go:227
		_go_fuzz_dep_.CoverTab[30917]++
								s = append(s, "_html_template_jsstrescaper")
//line /usr/local/go/src/html/template/escape.go:228
		// _ = "end of CoverTab[30917]"
	case stateJSBqStr:
//line /usr/local/go/src/html/template/escape.go:229
		_go_fuzz_dep_.CoverTab[30918]++
								if debugAllowActionJSTmpl.Value() == "1" {
//line /usr/local/go/src/html/template/escape.go:230
			_go_fuzz_dep_.CoverTab[30934]++
									s = append(s, "_html_template_jsstrescaper")
//line /usr/local/go/src/html/template/escape.go:231
			// _ = "end of CoverTab[30934]"
		} else {
//line /usr/local/go/src/html/template/escape.go:232
			_go_fuzz_dep_.CoverTab[30935]++
									return context{
				state:	stateError,
				err:	errorf(errJSTmplLit, n, n.Line, "%s appears in a JS template literal", n),
			}
//line /usr/local/go/src/html/template/escape.go:236
			// _ = "end of CoverTab[30935]"
		}
//line /usr/local/go/src/html/template/escape.go:237
		// _ = "end of CoverTab[30918]"
	case stateJSRegexp:
//line /usr/local/go/src/html/template/escape.go:238
		_go_fuzz_dep_.CoverTab[30919]++
								s = append(s, "_html_template_jsregexpescaper")
//line /usr/local/go/src/html/template/escape.go:239
		// _ = "end of CoverTab[30919]"
	case stateCSS:
//line /usr/local/go/src/html/template/escape.go:240
		_go_fuzz_dep_.CoverTab[30920]++
								s = append(s, "_html_template_cssvaluefilter")
//line /usr/local/go/src/html/template/escape.go:241
		// _ = "end of CoverTab[30920]"
	case stateText:
//line /usr/local/go/src/html/template/escape.go:242
		_go_fuzz_dep_.CoverTab[30921]++
								s = append(s, "_html_template_htmlescaper")
//line /usr/local/go/src/html/template/escape.go:243
		// _ = "end of CoverTab[30921]"
	case stateRCDATA:
//line /usr/local/go/src/html/template/escape.go:244
		_go_fuzz_dep_.CoverTab[30922]++
								s = append(s, "_html_template_rcdataescaper")
//line /usr/local/go/src/html/template/escape.go:245
		// _ = "end of CoverTab[30922]"
	case stateAttr:
//line /usr/local/go/src/html/template/escape.go:246
		_go_fuzz_dep_.CoverTab[30923]++
//line /usr/local/go/src/html/template/escape.go:246
		// _ = "end of CoverTab[30923]"

	case stateAttrName, stateTag:
//line /usr/local/go/src/html/template/escape.go:248
		_go_fuzz_dep_.CoverTab[30924]++
								c.state = stateAttrName
								s = append(s, "_html_template_htmlnamefilter")
//line /usr/local/go/src/html/template/escape.go:250
		// _ = "end of CoverTab[30924]"
	case stateSrcset:
//line /usr/local/go/src/html/template/escape.go:251
		_go_fuzz_dep_.CoverTab[30925]++
								s = append(s, "_html_template_srcsetescaper")
//line /usr/local/go/src/html/template/escape.go:252
		// _ = "end of CoverTab[30925]"
	default:
//line /usr/local/go/src/html/template/escape.go:253
		_go_fuzz_dep_.CoverTab[30926]++
								if isComment(c.state) {
//line /usr/local/go/src/html/template/escape.go:254
			_go_fuzz_dep_.CoverTab[30936]++
									s = append(s, "_html_template_commentescaper")
//line /usr/local/go/src/html/template/escape.go:255
			// _ = "end of CoverTab[30936]"
		} else {
//line /usr/local/go/src/html/template/escape.go:256
			_go_fuzz_dep_.CoverTab[30937]++
									panic("unexpected state " + c.state.String())
//line /usr/local/go/src/html/template/escape.go:257
			// _ = "end of CoverTab[30937]"
		}
//line /usr/local/go/src/html/template/escape.go:258
		// _ = "end of CoverTab[30926]"
	}
//line /usr/local/go/src/html/template/escape.go:259
	// _ = "end of CoverTab[30898]"
//line /usr/local/go/src/html/template/escape.go:259
	_go_fuzz_dep_.CoverTab[30899]++
							switch c.delim {
	case delimNone:
//line /usr/local/go/src/html/template/escape.go:261
		_go_fuzz_dep_.CoverTab[30938]++
//line /usr/local/go/src/html/template/escape.go:261
		// _ = "end of CoverTab[30938]"

	case delimSpaceOrTagEnd:
//line /usr/local/go/src/html/template/escape.go:263
		_go_fuzz_dep_.CoverTab[30939]++
								s = append(s, "_html_template_nospaceescaper")
//line /usr/local/go/src/html/template/escape.go:264
		// _ = "end of CoverTab[30939]"
	default:
//line /usr/local/go/src/html/template/escape.go:265
		_go_fuzz_dep_.CoverTab[30940]++
								s = append(s, "_html_template_attrescaper")
//line /usr/local/go/src/html/template/escape.go:266
		// _ = "end of CoverTab[30940]"
	}
//line /usr/local/go/src/html/template/escape.go:267
	// _ = "end of CoverTab[30899]"
//line /usr/local/go/src/html/template/escape.go:267
	_go_fuzz_dep_.CoverTab[30900]++
							e.editActionNode(n, s)
							return c
//line /usr/local/go/src/html/template/escape.go:269
	// _ = "end of CoverTab[30900]"
}

// ensurePipelineContains ensures that the pipeline ends with the commands with
//line /usr/local/go/src/html/template/escape.go:272
// the identifiers in s in order. If the pipeline ends with a predefined escaper
//line /usr/local/go/src/html/template/escape.go:272
// (i.e. "html" or "urlquery"), merge it with the identifiers in s.
//line /usr/local/go/src/html/template/escape.go:275
func ensurePipelineContains(p *parse.PipeNode, s []string) {
//line /usr/local/go/src/html/template/escape.go:275
	_go_fuzz_dep_.CoverTab[30941]++
							if len(s) == 0 {
//line /usr/local/go/src/html/template/escape.go:276
		_go_fuzz_dep_.CoverTab[30946]++

								return
//line /usr/local/go/src/html/template/escape.go:278
		// _ = "end of CoverTab[30946]"
	} else {
//line /usr/local/go/src/html/template/escape.go:279
		_go_fuzz_dep_.CoverTab[30947]++
//line /usr/local/go/src/html/template/escape.go:279
		// _ = "end of CoverTab[30947]"
//line /usr/local/go/src/html/template/escape.go:279
	}
//line /usr/local/go/src/html/template/escape.go:279
	// _ = "end of CoverTab[30941]"
//line /usr/local/go/src/html/template/escape.go:279
	_go_fuzz_dep_.CoverTab[30942]++

//line /usr/local/go/src/html/template/escape.go:283
	pipelineLen := len(p.Cmds)
	if pipelineLen > 0 {
//line /usr/local/go/src/html/template/escape.go:284
		_go_fuzz_dep_.CoverTab[30948]++
								lastCmd := p.Cmds[pipelineLen-1]
								if idNode, ok := lastCmd.Args[0].(*parse.IdentifierNode); ok {
//line /usr/local/go/src/html/template/escape.go:286
			_go_fuzz_dep_.CoverTab[30949]++
									if esc := idNode.Ident; predefinedEscapers[esc] {
//line /usr/local/go/src/html/template/escape.go:287
				_go_fuzz_dep_.CoverTab[30950]++

										if len(p.Cmds) == 1 && func() bool {
//line /usr/local/go/src/html/template/escape.go:289
					_go_fuzz_dep_.CoverTab[30953]++
//line /usr/local/go/src/html/template/escape.go:289
					return len(lastCmd.Args) > 1
//line /usr/local/go/src/html/template/escape.go:289
					// _ = "end of CoverTab[30953]"
//line /usr/local/go/src/html/template/escape.go:289
				}() {
//line /usr/local/go/src/html/template/escape.go:289
					_go_fuzz_dep_.CoverTab[30954]++

//line /usr/local/go/src/html/template/escape.go:295
					lastCmd.Args[0] = parse.NewIdentifier("_eval_args_").SetTree(nil).SetPos(lastCmd.Args[0].Position())
											p.Cmds = appendCmd(p.Cmds, newIdentCmd(esc, p.Position()))
											pipelineLen++
//line /usr/local/go/src/html/template/escape.go:297
					// _ = "end of CoverTab[30954]"
				} else {
//line /usr/local/go/src/html/template/escape.go:298
					_go_fuzz_dep_.CoverTab[30955]++
//line /usr/local/go/src/html/template/escape.go:298
					// _ = "end of CoverTab[30955]"
//line /usr/local/go/src/html/template/escape.go:298
				}
//line /usr/local/go/src/html/template/escape.go:298
				// _ = "end of CoverTab[30950]"
//line /usr/local/go/src/html/template/escape.go:298
				_go_fuzz_dep_.CoverTab[30951]++

//line /usr/local/go/src/html/template/escape.go:301
				dup := false
				for i, escaper := range s {
//line /usr/local/go/src/html/template/escape.go:302
					_go_fuzz_dep_.CoverTab[30956]++
											if escFnsEq(esc, escaper) {
//line /usr/local/go/src/html/template/escape.go:303
						_go_fuzz_dep_.CoverTab[30957]++
												s[i] = idNode.Ident
												dup = true
//line /usr/local/go/src/html/template/escape.go:305
						// _ = "end of CoverTab[30957]"
					} else {
//line /usr/local/go/src/html/template/escape.go:306
						_go_fuzz_dep_.CoverTab[30958]++
//line /usr/local/go/src/html/template/escape.go:306
						// _ = "end of CoverTab[30958]"
//line /usr/local/go/src/html/template/escape.go:306
					}
//line /usr/local/go/src/html/template/escape.go:306
					// _ = "end of CoverTab[30956]"
				}
//line /usr/local/go/src/html/template/escape.go:307
				// _ = "end of CoverTab[30951]"
//line /usr/local/go/src/html/template/escape.go:307
				_go_fuzz_dep_.CoverTab[30952]++
										if dup {
//line /usr/local/go/src/html/template/escape.go:308
					_go_fuzz_dep_.CoverTab[30959]++

//line /usr/local/go/src/html/template/escape.go:311
					pipelineLen--
//line /usr/local/go/src/html/template/escape.go:311
					// _ = "end of CoverTab[30959]"
				} else {
//line /usr/local/go/src/html/template/escape.go:312
					_go_fuzz_dep_.CoverTab[30960]++
//line /usr/local/go/src/html/template/escape.go:312
					// _ = "end of CoverTab[30960]"
//line /usr/local/go/src/html/template/escape.go:312
				}
//line /usr/local/go/src/html/template/escape.go:312
				// _ = "end of CoverTab[30952]"
			} else {
//line /usr/local/go/src/html/template/escape.go:313
				_go_fuzz_dep_.CoverTab[30961]++
//line /usr/local/go/src/html/template/escape.go:313
				// _ = "end of CoverTab[30961]"
//line /usr/local/go/src/html/template/escape.go:313
			}
//line /usr/local/go/src/html/template/escape.go:313
			// _ = "end of CoverTab[30949]"
		} else {
//line /usr/local/go/src/html/template/escape.go:314
			_go_fuzz_dep_.CoverTab[30962]++
//line /usr/local/go/src/html/template/escape.go:314
			// _ = "end of CoverTab[30962]"
//line /usr/local/go/src/html/template/escape.go:314
		}
//line /usr/local/go/src/html/template/escape.go:314
		// _ = "end of CoverTab[30948]"
	} else {
//line /usr/local/go/src/html/template/escape.go:315
		_go_fuzz_dep_.CoverTab[30963]++
//line /usr/local/go/src/html/template/escape.go:315
		// _ = "end of CoverTab[30963]"
//line /usr/local/go/src/html/template/escape.go:315
	}
//line /usr/local/go/src/html/template/escape.go:315
	// _ = "end of CoverTab[30942]"
//line /usr/local/go/src/html/template/escape.go:315
	_go_fuzz_dep_.CoverTab[30943]++

							newCmds := make([]*parse.CommandNode, pipelineLen, pipelineLen+len(s))
							insertedIdents := make(map[string]bool)
							for i := 0; i < pipelineLen; i++ {
//line /usr/local/go/src/html/template/escape.go:319
		_go_fuzz_dep_.CoverTab[30964]++
								cmd := p.Cmds[i]
								newCmds[i] = cmd
								if idNode, ok := cmd.Args[0].(*parse.IdentifierNode); ok {
//line /usr/local/go/src/html/template/escape.go:322
			_go_fuzz_dep_.CoverTab[30965]++
									insertedIdents[normalizeEscFn(idNode.Ident)] = true
//line /usr/local/go/src/html/template/escape.go:323
			// _ = "end of CoverTab[30965]"
		} else {
//line /usr/local/go/src/html/template/escape.go:324
			_go_fuzz_dep_.CoverTab[30966]++
//line /usr/local/go/src/html/template/escape.go:324
			// _ = "end of CoverTab[30966]"
//line /usr/local/go/src/html/template/escape.go:324
		}
//line /usr/local/go/src/html/template/escape.go:324
		// _ = "end of CoverTab[30964]"
	}
//line /usr/local/go/src/html/template/escape.go:325
	// _ = "end of CoverTab[30943]"
//line /usr/local/go/src/html/template/escape.go:325
	_go_fuzz_dep_.CoverTab[30944]++
							for _, name := range s {
//line /usr/local/go/src/html/template/escape.go:326
		_go_fuzz_dep_.CoverTab[30967]++
								if !insertedIdents[normalizeEscFn(name)] {
//line /usr/local/go/src/html/template/escape.go:327
			_go_fuzz_dep_.CoverTab[30968]++

//line /usr/local/go/src/html/template/escape.go:332
			newCmds = appendCmd(newCmds, newIdentCmd(name, p.Position()))
//line /usr/local/go/src/html/template/escape.go:332
			// _ = "end of CoverTab[30968]"
		} else {
//line /usr/local/go/src/html/template/escape.go:333
			_go_fuzz_dep_.CoverTab[30969]++
//line /usr/local/go/src/html/template/escape.go:333
			// _ = "end of CoverTab[30969]"
//line /usr/local/go/src/html/template/escape.go:333
		}
//line /usr/local/go/src/html/template/escape.go:333
		// _ = "end of CoverTab[30967]"
	}
//line /usr/local/go/src/html/template/escape.go:334
	// _ = "end of CoverTab[30944]"
//line /usr/local/go/src/html/template/escape.go:334
	_go_fuzz_dep_.CoverTab[30945]++
							p.Cmds = newCmds
//line /usr/local/go/src/html/template/escape.go:335
	// _ = "end of CoverTab[30945]"
}

// predefinedEscapers contains template predefined escapers that are equivalent
//line /usr/local/go/src/html/template/escape.go:338
// to some contextual escapers. Keep in sync with equivEscapers.
//line /usr/local/go/src/html/template/escape.go:340
var predefinedEscapers = map[string]bool{
	"html":		true,
	"urlquery":	true,
}

// equivEscapers matches contextual escapers to equivalent predefined
//line /usr/local/go/src/html/template/escape.go:345
// template escapers.
//line /usr/local/go/src/html/template/escape.go:347
var equivEscapers = map[string]string{

//line /usr/local/go/src/html/template/escape.go:350
	"_html_template_attrescaper":	"html",
							"_html_template_htmlescaper":	"html",
							"_html_template_rcdataescaper":	"html",

//line /usr/local/go/src/html/template/escape.go:356
	"_html_template_urlescaper":	"urlquery",

//line /usr/local/go/src/html/template/escape.go:363
	"_html_template_urlnormalizer":	"urlquery",
}

// escFnsEq reports whether the two escaping functions are equivalent.
func escFnsEq(a, b string) bool {
//line /usr/local/go/src/html/template/escape.go:367
	_go_fuzz_dep_.CoverTab[30970]++
							return normalizeEscFn(a) == normalizeEscFn(b)
//line /usr/local/go/src/html/template/escape.go:368
	// _ = "end of CoverTab[30970]"
}

// normalizeEscFn(a) is equal to normalizeEscFn(b) for any pair of names of
//line /usr/local/go/src/html/template/escape.go:371
// escaper functions a and b that are equivalent.
//line /usr/local/go/src/html/template/escape.go:373
func normalizeEscFn(e string) string {
//line /usr/local/go/src/html/template/escape.go:373
	_go_fuzz_dep_.CoverTab[30971]++
							if norm := equivEscapers[e]; norm != "" {
//line /usr/local/go/src/html/template/escape.go:374
		_go_fuzz_dep_.CoverTab[30973]++
								return norm
//line /usr/local/go/src/html/template/escape.go:375
		// _ = "end of CoverTab[30973]"
	} else {
//line /usr/local/go/src/html/template/escape.go:376
		_go_fuzz_dep_.CoverTab[30974]++
//line /usr/local/go/src/html/template/escape.go:376
		// _ = "end of CoverTab[30974]"
//line /usr/local/go/src/html/template/escape.go:376
	}
//line /usr/local/go/src/html/template/escape.go:376
	// _ = "end of CoverTab[30971]"
//line /usr/local/go/src/html/template/escape.go:376
	_go_fuzz_dep_.CoverTab[30972]++
							return e
//line /usr/local/go/src/html/template/escape.go:377
	// _ = "end of CoverTab[30972]"
}

// redundantFuncs[a][b] implies that funcMap[b](funcMap[a](x)) == funcMap[a](x)
//line /usr/local/go/src/html/template/escape.go:380
// for all x.
//line /usr/local/go/src/html/template/escape.go:382
var redundantFuncs = map[string]map[string]bool{
	"_html_template_commentescaper": {
		"_html_template_attrescaper":	true,
		"_html_template_htmlescaper":	true,
	},
	"_html_template_cssescaper": {
		"_html_template_attrescaper": true,
	},
	"_html_template_jsregexpescaper": {
		"_html_template_attrescaper": true,
	},
	"_html_template_jsstrescaper": {
		"_html_template_attrescaper": true,
	},
	"_html_template_urlescaper": {
		"_html_template_urlnormalizer": true,
	},
}

// appendCmd appends the given command to the end of the command pipeline
//line /usr/local/go/src/html/template/escape.go:401
// unless it is redundant with the last command.
//line /usr/local/go/src/html/template/escape.go:403
func appendCmd(cmds []*parse.CommandNode, cmd *parse.CommandNode) []*parse.CommandNode {
//line /usr/local/go/src/html/template/escape.go:403
	_go_fuzz_dep_.CoverTab[30975]++
							if n := len(cmds); n != 0 {
//line /usr/local/go/src/html/template/escape.go:404
		_go_fuzz_dep_.CoverTab[30977]++
								last, okLast := cmds[n-1].Args[0].(*parse.IdentifierNode)
								next, okNext := cmd.Args[0].(*parse.IdentifierNode)
								if okLast && func() bool {
//line /usr/local/go/src/html/template/escape.go:407
			_go_fuzz_dep_.CoverTab[30978]++
//line /usr/local/go/src/html/template/escape.go:407
			return okNext
//line /usr/local/go/src/html/template/escape.go:407
			// _ = "end of CoverTab[30978]"
//line /usr/local/go/src/html/template/escape.go:407
		}() && func() bool {
//line /usr/local/go/src/html/template/escape.go:407
			_go_fuzz_dep_.CoverTab[30979]++
//line /usr/local/go/src/html/template/escape.go:407
			return redundantFuncs[last.Ident][next.Ident]
//line /usr/local/go/src/html/template/escape.go:407
			// _ = "end of CoverTab[30979]"
//line /usr/local/go/src/html/template/escape.go:407
		}() {
//line /usr/local/go/src/html/template/escape.go:407
			_go_fuzz_dep_.CoverTab[30980]++
									return cmds
//line /usr/local/go/src/html/template/escape.go:408
			// _ = "end of CoverTab[30980]"
		} else {
//line /usr/local/go/src/html/template/escape.go:409
			_go_fuzz_dep_.CoverTab[30981]++
//line /usr/local/go/src/html/template/escape.go:409
			// _ = "end of CoverTab[30981]"
//line /usr/local/go/src/html/template/escape.go:409
		}
//line /usr/local/go/src/html/template/escape.go:409
		// _ = "end of CoverTab[30977]"
	} else {
//line /usr/local/go/src/html/template/escape.go:410
		_go_fuzz_dep_.CoverTab[30982]++
//line /usr/local/go/src/html/template/escape.go:410
		// _ = "end of CoverTab[30982]"
//line /usr/local/go/src/html/template/escape.go:410
	}
//line /usr/local/go/src/html/template/escape.go:410
	// _ = "end of CoverTab[30975]"
//line /usr/local/go/src/html/template/escape.go:410
	_go_fuzz_dep_.CoverTab[30976]++
							return append(cmds, cmd)
//line /usr/local/go/src/html/template/escape.go:411
	// _ = "end of CoverTab[30976]"
}

// newIdentCmd produces a command containing a single identifier node.
func newIdentCmd(identifier string, pos parse.Pos) *parse.CommandNode {
//line /usr/local/go/src/html/template/escape.go:415
	_go_fuzz_dep_.CoverTab[30983]++
							return &parse.CommandNode{
		NodeType:	parse.NodeCommand,
		Args:		[]parse.Node{parse.NewIdentifier(identifier).SetTree(nil).SetPos(pos)},
	}
//line /usr/local/go/src/html/template/escape.go:419
	// _ = "end of CoverTab[30983]"
}

// nudge returns the context that would result from following empty string
//line /usr/local/go/src/html/template/escape.go:422
// transitions from the input context.
//line /usr/local/go/src/html/template/escape.go:422
// For example, parsing:
//line /usr/local/go/src/html/template/escape.go:422
//
//line /usr/local/go/src/html/template/escape.go:422
//	`<a href=`
//line /usr/local/go/src/html/template/escape.go:422
//
//line /usr/local/go/src/html/template/escape.go:422
// will end in context{stateBeforeValue, attrURL}, but parsing one extra rune:
//line /usr/local/go/src/html/template/escape.go:422
//
//line /usr/local/go/src/html/template/escape.go:422
//	`<a href=x`
//line /usr/local/go/src/html/template/escape.go:422
//
//line /usr/local/go/src/html/template/escape.go:422
// will end in context{stateURL, delimSpaceOrTagEnd, ...}.
//line /usr/local/go/src/html/template/escape.go:422
// There are two transitions that happen when the 'x' is seen:
//line /usr/local/go/src/html/template/escape.go:422
// (1) Transition from a before-value state to a start-of-value state without
//line /usr/local/go/src/html/template/escape.go:422
//
//line /usr/local/go/src/html/template/escape.go:422
//	consuming any character.
//line /usr/local/go/src/html/template/escape.go:422
//
//line /usr/local/go/src/html/template/escape.go:422
// (2) Consume 'x' and transition past the first value character.
//line /usr/local/go/src/html/template/escape.go:422
// In this case, nudging produces the context after (1) happens.
//line /usr/local/go/src/html/template/escape.go:440
func nudge(c context) context {
//line /usr/local/go/src/html/template/escape.go:440
	_go_fuzz_dep_.CoverTab[30984]++
							switch c.state {
	case stateTag:
//line /usr/local/go/src/html/template/escape.go:442
		_go_fuzz_dep_.CoverTab[30986]++

								c.state = stateAttrName
//line /usr/local/go/src/html/template/escape.go:444
		// _ = "end of CoverTab[30986]"
	case stateBeforeValue:
//line /usr/local/go/src/html/template/escape.go:445
		_go_fuzz_dep_.CoverTab[30987]++

								c.state, c.delim, c.attr = attrStartStates[c.attr], delimSpaceOrTagEnd, attrNone
//line /usr/local/go/src/html/template/escape.go:447
		// _ = "end of CoverTab[30987]"
	case stateAfterName:
//line /usr/local/go/src/html/template/escape.go:448
		_go_fuzz_dep_.CoverTab[30988]++

								c.state, c.attr = stateAttrName, attrNone
//line /usr/local/go/src/html/template/escape.go:450
		// _ = "end of CoverTab[30988]"
//line /usr/local/go/src/html/template/escape.go:450
	default:
//line /usr/local/go/src/html/template/escape.go:450
		_go_fuzz_dep_.CoverTab[30989]++
//line /usr/local/go/src/html/template/escape.go:450
		// _ = "end of CoverTab[30989]"
	}
//line /usr/local/go/src/html/template/escape.go:451
	// _ = "end of CoverTab[30984]"
//line /usr/local/go/src/html/template/escape.go:451
	_go_fuzz_dep_.CoverTab[30985]++
							return c
//line /usr/local/go/src/html/template/escape.go:452
	// _ = "end of CoverTab[30985]"
}

// join joins the two contexts of a branch template node. The result is an
//line /usr/local/go/src/html/template/escape.go:455
// error context if either of the input contexts are error contexts, or if the
//line /usr/local/go/src/html/template/escape.go:455
// input contexts differ.
//line /usr/local/go/src/html/template/escape.go:458
func join(a, b context, node parse.Node, nodeName string) context {
//line /usr/local/go/src/html/template/escape.go:458
	_go_fuzz_dep_.CoverTab[30990]++
							if a.state == stateError {
//line /usr/local/go/src/html/template/escape.go:459
		_go_fuzz_dep_.CoverTab[30999]++
								return a
//line /usr/local/go/src/html/template/escape.go:460
		// _ = "end of CoverTab[30999]"
	} else {
//line /usr/local/go/src/html/template/escape.go:461
		_go_fuzz_dep_.CoverTab[31000]++
//line /usr/local/go/src/html/template/escape.go:461
		// _ = "end of CoverTab[31000]"
//line /usr/local/go/src/html/template/escape.go:461
	}
//line /usr/local/go/src/html/template/escape.go:461
	// _ = "end of CoverTab[30990]"
//line /usr/local/go/src/html/template/escape.go:461
	_go_fuzz_dep_.CoverTab[30991]++
							if b.state == stateError {
//line /usr/local/go/src/html/template/escape.go:462
		_go_fuzz_dep_.CoverTab[31001]++
								return b
//line /usr/local/go/src/html/template/escape.go:463
		// _ = "end of CoverTab[31001]"
	} else {
//line /usr/local/go/src/html/template/escape.go:464
		_go_fuzz_dep_.CoverTab[31002]++
//line /usr/local/go/src/html/template/escape.go:464
		// _ = "end of CoverTab[31002]"
//line /usr/local/go/src/html/template/escape.go:464
	}
//line /usr/local/go/src/html/template/escape.go:464
	// _ = "end of CoverTab[30991]"
//line /usr/local/go/src/html/template/escape.go:464
	_go_fuzz_dep_.CoverTab[30992]++
							if a.state == stateDead {
//line /usr/local/go/src/html/template/escape.go:465
		_go_fuzz_dep_.CoverTab[31003]++
								return b
//line /usr/local/go/src/html/template/escape.go:466
		// _ = "end of CoverTab[31003]"
	} else {
//line /usr/local/go/src/html/template/escape.go:467
		_go_fuzz_dep_.CoverTab[31004]++
//line /usr/local/go/src/html/template/escape.go:467
		// _ = "end of CoverTab[31004]"
//line /usr/local/go/src/html/template/escape.go:467
	}
//line /usr/local/go/src/html/template/escape.go:467
	// _ = "end of CoverTab[30992]"
//line /usr/local/go/src/html/template/escape.go:467
	_go_fuzz_dep_.CoverTab[30993]++
							if b.state == stateDead {
//line /usr/local/go/src/html/template/escape.go:468
		_go_fuzz_dep_.CoverTab[31005]++
								return a
//line /usr/local/go/src/html/template/escape.go:469
		// _ = "end of CoverTab[31005]"
	} else {
//line /usr/local/go/src/html/template/escape.go:470
		_go_fuzz_dep_.CoverTab[31006]++
//line /usr/local/go/src/html/template/escape.go:470
		// _ = "end of CoverTab[31006]"
//line /usr/local/go/src/html/template/escape.go:470
	}
//line /usr/local/go/src/html/template/escape.go:470
	// _ = "end of CoverTab[30993]"
//line /usr/local/go/src/html/template/escape.go:470
	_go_fuzz_dep_.CoverTab[30994]++
							if a.eq(b) {
//line /usr/local/go/src/html/template/escape.go:471
		_go_fuzz_dep_.CoverTab[31007]++
								return a
//line /usr/local/go/src/html/template/escape.go:472
		// _ = "end of CoverTab[31007]"
	} else {
//line /usr/local/go/src/html/template/escape.go:473
		_go_fuzz_dep_.CoverTab[31008]++
//line /usr/local/go/src/html/template/escape.go:473
		// _ = "end of CoverTab[31008]"
//line /usr/local/go/src/html/template/escape.go:473
	}
//line /usr/local/go/src/html/template/escape.go:473
	// _ = "end of CoverTab[30994]"
//line /usr/local/go/src/html/template/escape.go:473
	_go_fuzz_dep_.CoverTab[30995]++

							c := a
							c.urlPart = b.urlPart
							if c.eq(b) {
//line /usr/local/go/src/html/template/escape.go:477
		_go_fuzz_dep_.CoverTab[31009]++

								c.urlPart = urlPartUnknown
								return c
//line /usr/local/go/src/html/template/escape.go:480
		// _ = "end of CoverTab[31009]"
	} else {
//line /usr/local/go/src/html/template/escape.go:481
		_go_fuzz_dep_.CoverTab[31010]++
//line /usr/local/go/src/html/template/escape.go:481
		// _ = "end of CoverTab[31010]"
//line /usr/local/go/src/html/template/escape.go:481
	}
//line /usr/local/go/src/html/template/escape.go:481
	// _ = "end of CoverTab[30995]"
//line /usr/local/go/src/html/template/escape.go:481
	_go_fuzz_dep_.CoverTab[30996]++

							c = a
							c.jsCtx = b.jsCtx
							if c.eq(b) {
//line /usr/local/go/src/html/template/escape.go:485
		_go_fuzz_dep_.CoverTab[31011]++

								c.jsCtx = jsCtxUnknown
								return c
//line /usr/local/go/src/html/template/escape.go:488
		// _ = "end of CoverTab[31011]"
	} else {
//line /usr/local/go/src/html/template/escape.go:489
		_go_fuzz_dep_.CoverTab[31012]++
//line /usr/local/go/src/html/template/escape.go:489
		// _ = "end of CoverTab[31012]"
//line /usr/local/go/src/html/template/escape.go:489
	}
//line /usr/local/go/src/html/template/escape.go:489
	// _ = "end of CoverTab[30996]"
//line /usr/local/go/src/html/template/escape.go:489
	_go_fuzz_dep_.CoverTab[30997]++

//line /usr/local/go/src/html/template/escape.go:496
	if c, d := nudge(a), nudge(b); !(c.eq(a) && func() bool {
//line /usr/local/go/src/html/template/escape.go:496
		_go_fuzz_dep_.CoverTab[31013]++
//line /usr/local/go/src/html/template/escape.go:496
		return d.eq(b)
//line /usr/local/go/src/html/template/escape.go:496
		// _ = "end of CoverTab[31013]"
//line /usr/local/go/src/html/template/escape.go:496
	}()) {
//line /usr/local/go/src/html/template/escape.go:496
		_go_fuzz_dep_.CoverTab[31014]++
								if e := join(c, d, node, nodeName); e.state != stateError {
//line /usr/local/go/src/html/template/escape.go:497
			_go_fuzz_dep_.CoverTab[31015]++
									return e
//line /usr/local/go/src/html/template/escape.go:498
			// _ = "end of CoverTab[31015]"
		} else {
//line /usr/local/go/src/html/template/escape.go:499
			_go_fuzz_dep_.CoverTab[31016]++
//line /usr/local/go/src/html/template/escape.go:499
			// _ = "end of CoverTab[31016]"
//line /usr/local/go/src/html/template/escape.go:499
		}
//line /usr/local/go/src/html/template/escape.go:499
		// _ = "end of CoverTab[31014]"
	} else {
//line /usr/local/go/src/html/template/escape.go:500
		_go_fuzz_dep_.CoverTab[31017]++
//line /usr/local/go/src/html/template/escape.go:500
		// _ = "end of CoverTab[31017]"
//line /usr/local/go/src/html/template/escape.go:500
	}
//line /usr/local/go/src/html/template/escape.go:500
	// _ = "end of CoverTab[30997]"
//line /usr/local/go/src/html/template/escape.go:500
	_go_fuzz_dep_.CoverTab[30998]++

							return context{
		state:	stateError,
		err:	errorf(ErrBranchEnd, node, 0, "{{%s}} branches end in different contexts: %v, %v", nodeName, a, b),
	}
//line /usr/local/go/src/html/template/escape.go:505
	// _ = "end of CoverTab[30998]"
}

// escapeBranch escapes a branch template node: "if", "range" and "with".
func (e *escaper) escapeBranch(c context, n *parse.BranchNode, nodeName string) context {
//line /usr/local/go/src/html/template/escape.go:509
	_go_fuzz_dep_.CoverTab[31018]++
							if nodeName == "range" {
//line /usr/local/go/src/html/template/escape.go:510
		_go_fuzz_dep_.CoverTab[31021]++
								e.rangeContext = &rangeContext{outer: e.rangeContext}
//line /usr/local/go/src/html/template/escape.go:511
		// _ = "end of CoverTab[31021]"
	} else {
//line /usr/local/go/src/html/template/escape.go:512
		_go_fuzz_dep_.CoverTab[31022]++
//line /usr/local/go/src/html/template/escape.go:512
		// _ = "end of CoverTab[31022]"
//line /usr/local/go/src/html/template/escape.go:512
	}
//line /usr/local/go/src/html/template/escape.go:512
	// _ = "end of CoverTab[31018]"
//line /usr/local/go/src/html/template/escape.go:512
	_go_fuzz_dep_.CoverTab[31019]++
							c0 := e.escapeList(c, n.List)
							if nodeName == "range" {
//line /usr/local/go/src/html/template/escape.go:514
		_go_fuzz_dep_.CoverTab[31023]++
								if c0.state != stateError {
//line /usr/local/go/src/html/template/escape.go:515
			_go_fuzz_dep_.CoverTab[31027]++
									c0 = joinRange(c0, e.rangeContext)
//line /usr/local/go/src/html/template/escape.go:516
			// _ = "end of CoverTab[31027]"
		} else {
//line /usr/local/go/src/html/template/escape.go:517
			_go_fuzz_dep_.CoverTab[31028]++
//line /usr/local/go/src/html/template/escape.go:517
			// _ = "end of CoverTab[31028]"
//line /usr/local/go/src/html/template/escape.go:517
		}
//line /usr/local/go/src/html/template/escape.go:517
		// _ = "end of CoverTab[31023]"
//line /usr/local/go/src/html/template/escape.go:517
		_go_fuzz_dep_.CoverTab[31024]++
								e.rangeContext = e.rangeContext.outer
								if c0.state == stateError {
//line /usr/local/go/src/html/template/escape.go:519
			_go_fuzz_dep_.CoverTab[31029]++
									return c0
//line /usr/local/go/src/html/template/escape.go:520
			// _ = "end of CoverTab[31029]"
		} else {
//line /usr/local/go/src/html/template/escape.go:521
			_go_fuzz_dep_.CoverTab[31030]++
//line /usr/local/go/src/html/template/escape.go:521
			// _ = "end of CoverTab[31030]"
//line /usr/local/go/src/html/template/escape.go:521
		}
//line /usr/local/go/src/html/template/escape.go:521
		// _ = "end of CoverTab[31024]"
//line /usr/local/go/src/html/template/escape.go:521
		_go_fuzz_dep_.CoverTab[31025]++

//line /usr/local/go/src/html/template/escape.go:526
		e.rangeContext = &rangeContext{outer: e.rangeContext}
		c1, _ := e.escapeListConditionally(c0, n.List, nil)
		c0 = join(c0, c1, n, nodeName)
		if c0.state == stateError {
//line /usr/local/go/src/html/template/escape.go:529
			_go_fuzz_dep_.CoverTab[31031]++
									e.rangeContext = e.rangeContext.outer

//line /usr/local/go/src/html/template/escape.go:534
			c0.err.Line = n.Line
									c0.err.Description = "on range loop re-entry: " + c0.err.Description
									return c0
//line /usr/local/go/src/html/template/escape.go:536
			// _ = "end of CoverTab[31031]"
		} else {
//line /usr/local/go/src/html/template/escape.go:537
			_go_fuzz_dep_.CoverTab[31032]++
//line /usr/local/go/src/html/template/escape.go:537
			// _ = "end of CoverTab[31032]"
//line /usr/local/go/src/html/template/escape.go:537
		}
//line /usr/local/go/src/html/template/escape.go:537
		// _ = "end of CoverTab[31025]"
//line /usr/local/go/src/html/template/escape.go:537
		_go_fuzz_dep_.CoverTab[31026]++
								c0 = joinRange(c0, e.rangeContext)
								e.rangeContext = e.rangeContext.outer
								if c0.state == stateError {
//line /usr/local/go/src/html/template/escape.go:540
			_go_fuzz_dep_.CoverTab[31033]++
									return c0
//line /usr/local/go/src/html/template/escape.go:541
			// _ = "end of CoverTab[31033]"
		} else {
//line /usr/local/go/src/html/template/escape.go:542
			_go_fuzz_dep_.CoverTab[31034]++
//line /usr/local/go/src/html/template/escape.go:542
			// _ = "end of CoverTab[31034]"
//line /usr/local/go/src/html/template/escape.go:542
		}
//line /usr/local/go/src/html/template/escape.go:542
		// _ = "end of CoverTab[31026]"
	} else {
//line /usr/local/go/src/html/template/escape.go:543
		_go_fuzz_dep_.CoverTab[31035]++
//line /usr/local/go/src/html/template/escape.go:543
		// _ = "end of CoverTab[31035]"
//line /usr/local/go/src/html/template/escape.go:543
	}
//line /usr/local/go/src/html/template/escape.go:543
	// _ = "end of CoverTab[31019]"
//line /usr/local/go/src/html/template/escape.go:543
	_go_fuzz_dep_.CoverTab[31020]++
							c1 := e.escapeList(c, n.ElseList)
							return join(c0, c1, n, nodeName)
//line /usr/local/go/src/html/template/escape.go:545
	// _ = "end of CoverTab[31020]"
}

func joinRange(c0 context, rc *rangeContext) context {
//line /usr/local/go/src/html/template/escape.go:548
	_go_fuzz_dep_.CoverTab[31036]++

//line /usr/local/go/src/html/template/escape.go:552
	for _, c := range rc.breaks {
//line /usr/local/go/src/html/template/escape.go:552
		_go_fuzz_dep_.CoverTab[31039]++
								c0 = join(c0, c, c.n, "range")
								if c0.state == stateError {
//line /usr/local/go/src/html/template/escape.go:554
			_go_fuzz_dep_.CoverTab[31040]++
									c0.err.Line = c.n.(*parse.BreakNode).Line
									c0.err.Description = "at range loop break: " + c0.err.Description
									return c0
//line /usr/local/go/src/html/template/escape.go:557
			// _ = "end of CoverTab[31040]"
		} else {
//line /usr/local/go/src/html/template/escape.go:558
			_go_fuzz_dep_.CoverTab[31041]++
//line /usr/local/go/src/html/template/escape.go:558
			// _ = "end of CoverTab[31041]"
//line /usr/local/go/src/html/template/escape.go:558
		}
//line /usr/local/go/src/html/template/escape.go:558
		// _ = "end of CoverTab[31039]"
	}
//line /usr/local/go/src/html/template/escape.go:559
	// _ = "end of CoverTab[31036]"
//line /usr/local/go/src/html/template/escape.go:559
	_go_fuzz_dep_.CoverTab[31037]++
							for _, c := range rc.continues {
//line /usr/local/go/src/html/template/escape.go:560
		_go_fuzz_dep_.CoverTab[31042]++
								c0 = join(c0, c, c.n, "range")
								if c0.state == stateError {
//line /usr/local/go/src/html/template/escape.go:562
			_go_fuzz_dep_.CoverTab[31043]++
									c0.err.Line = c.n.(*parse.ContinueNode).Line
									c0.err.Description = "at range loop continue: " + c0.err.Description
									return c0
//line /usr/local/go/src/html/template/escape.go:565
			// _ = "end of CoverTab[31043]"
		} else {
//line /usr/local/go/src/html/template/escape.go:566
			_go_fuzz_dep_.CoverTab[31044]++
//line /usr/local/go/src/html/template/escape.go:566
			// _ = "end of CoverTab[31044]"
//line /usr/local/go/src/html/template/escape.go:566
		}
//line /usr/local/go/src/html/template/escape.go:566
		// _ = "end of CoverTab[31042]"
	}
//line /usr/local/go/src/html/template/escape.go:567
	// _ = "end of CoverTab[31037]"
//line /usr/local/go/src/html/template/escape.go:567
	_go_fuzz_dep_.CoverTab[31038]++
							return c0
//line /usr/local/go/src/html/template/escape.go:568
	// _ = "end of CoverTab[31038]"
}

// escapeList escapes a list template node.
func (e *escaper) escapeList(c context, n *parse.ListNode) context {
//line /usr/local/go/src/html/template/escape.go:572
	_go_fuzz_dep_.CoverTab[31045]++
							if n == nil {
//line /usr/local/go/src/html/template/escape.go:573
		_go_fuzz_dep_.CoverTab[31048]++
								return c
//line /usr/local/go/src/html/template/escape.go:574
		// _ = "end of CoverTab[31048]"
	} else {
//line /usr/local/go/src/html/template/escape.go:575
		_go_fuzz_dep_.CoverTab[31049]++
//line /usr/local/go/src/html/template/escape.go:575
		// _ = "end of CoverTab[31049]"
//line /usr/local/go/src/html/template/escape.go:575
	}
//line /usr/local/go/src/html/template/escape.go:575
	// _ = "end of CoverTab[31045]"
//line /usr/local/go/src/html/template/escape.go:575
	_go_fuzz_dep_.CoverTab[31046]++
							for _, m := range n.Nodes {
//line /usr/local/go/src/html/template/escape.go:576
		_go_fuzz_dep_.CoverTab[31050]++
								c = e.escape(c, m)
								if c.state == stateDead {
//line /usr/local/go/src/html/template/escape.go:578
			_go_fuzz_dep_.CoverTab[31051]++
									break
//line /usr/local/go/src/html/template/escape.go:579
			// _ = "end of CoverTab[31051]"
		} else {
//line /usr/local/go/src/html/template/escape.go:580
			_go_fuzz_dep_.CoverTab[31052]++
//line /usr/local/go/src/html/template/escape.go:580
			// _ = "end of CoverTab[31052]"
//line /usr/local/go/src/html/template/escape.go:580
		}
//line /usr/local/go/src/html/template/escape.go:580
		// _ = "end of CoverTab[31050]"
	}
//line /usr/local/go/src/html/template/escape.go:581
	// _ = "end of CoverTab[31046]"
//line /usr/local/go/src/html/template/escape.go:581
	_go_fuzz_dep_.CoverTab[31047]++
							return c
//line /usr/local/go/src/html/template/escape.go:582
	// _ = "end of CoverTab[31047]"
}

// escapeListConditionally escapes a list node but only preserves edits and
//line /usr/local/go/src/html/template/escape.go:585
// inferences in e if the inferences and output context satisfy filter.
//line /usr/local/go/src/html/template/escape.go:585
// It returns the best guess at an output context, and the result of the filter
//line /usr/local/go/src/html/template/escape.go:585
// which is the same as whether e was updated.
//line /usr/local/go/src/html/template/escape.go:589
func (e *escaper) escapeListConditionally(c context, n *parse.ListNode, filter func(*escaper, context) bool) (context, bool) {
//line /usr/local/go/src/html/template/escape.go:589
	_go_fuzz_dep_.CoverTab[31053]++
							e1 := makeEscaper(e.ns)
							e1.rangeContext = e.rangeContext

							for k, v := range e.output {
//line /usr/local/go/src/html/template/escape.go:593
		_go_fuzz_dep_.CoverTab[31056]++
								e1.output[k] = v
//line /usr/local/go/src/html/template/escape.go:594
		// _ = "end of CoverTab[31056]"
	}
//line /usr/local/go/src/html/template/escape.go:595
	// _ = "end of CoverTab[31053]"
//line /usr/local/go/src/html/template/escape.go:595
	_go_fuzz_dep_.CoverTab[31054]++
							c = e1.escapeList(c, n)
							ok := filter != nil && func() bool {
//line /usr/local/go/src/html/template/escape.go:597
		_go_fuzz_dep_.CoverTab[31057]++
//line /usr/local/go/src/html/template/escape.go:597
		return filter(&e1, c)
//line /usr/local/go/src/html/template/escape.go:597
		// _ = "end of CoverTab[31057]"
//line /usr/local/go/src/html/template/escape.go:597
	}()
							if ok {
//line /usr/local/go/src/html/template/escape.go:598
		_go_fuzz_dep_.CoverTab[31058]++

								for k, v := range e1.output {
//line /usr/local/go/src/html/template/escape.go:600
			_go_fuzz_dep_.CoverTab[31064]++
									e.output[k] = v
//line /usr/local/go/src/html/template/escape.go:601
			// _ = "end of CoverTab[31064]"
		}
//line /usr/local/go/src/html/template/escape.go:602
		// _ = "end of CoverTab[31058]"
//line /usr/local/go/src/html/template/escape.go:602
		_go_fuzz_dep_.CoverTab[31059]++
								for k, v := range e1.derived {
//line /usr/local/go/src/html/template/escape.go:603
			_go_fuzz_dep_.CoverTab[31065]++
									e.derived[k] = v
//line /usr/local/go/src/html/template/escape.go:604
			// _ = "end of CoverTab[31065]"
		}
//line /usr/local/go/src/html/template/escape.go:605
		// _ = "end of CoverTab[31059]"
//line /usr/local/go/src/html/template/escape.go:605
		_go_fuzz_dep_.CoverTab[31060]++
								for k, v := range e1.called {
//line /usr/local/go/src/html/template/escape.go:606
			_go_fuzz_dep_.CoverTab[31066]++
									e.called[k] = v
//line /usr/local/go/src/html/template/escape.go:607
			// _ = "end of CoverTab[31066]"
		}
//line /usr/local/go/src/html/template/escape.go:608
		// _ = "end of CoverTab[31060]"
//line /usr/local/go/src/html/template/escape.go:608
		_go_fuzz_dep_.CoverTab[31061]++
								for k, v := range e1.actionNodeEdits {
//line /usr/local/go/src/html/template/escape.go:609
			_go_fuzz_dep_.CoverTab[31067]++
									e.editActionNode(k, v)
//line /usr/local/go/src/html/template/escape.go:610
			// _ = "end of CoverTab[31067]"
		}
//line /usr/local/go/src/html/template/escape.go:611
		// _ = "end of CoverTab[31061]"
//line /usr/local/go/src/html/template/escape.go:611
		_go_fuzz_dep_.CoverTab[31062]++
								for k, v := range e1.templateNodeEdits {
//line /usr/local/go/src/html/template/escape.go:612
			_go_fuzz_dep_.CoverTab[31068]++
									e.editTemplateNode(k, v)
//line /usr/local/go/src/html/template/escape.go:613
			// _ = "end of CoverTab[31068]"
		}
//line /usr/local/go/src/html/template/escape.go:614
		// _ = "end of CoverTab[31062]"
//line /usr/local/go/src/html/template/escape.go:614
		_go_fuzz_dep_.CoverTab[31063]++
								for k, v := range e1.textNodeEdits {
//line /usr/local/go/src/html/template/escape.go:615
			_go_fuzz_dep_.CoverTab[31069]++
									e.editTextNode(k, v)
//line /usr/local/go/src/html/template/escape.go:616
			// _ = "end of CoverTab[31069]"
		}
//line /usr/local/go/src/html/template/escape.go:617
		// _ = "end of CoverTab[31063]"
	} else {
//line /usr/local/go/src/html/template/escape.go:618
		_go_fuzz_dep_.CoverTab[31070]++
//line /usr/local/go/src/html/template/escape.go:618
		// _ = "end of CoverTab[31070]"
//line /usr/local/go/src/html/template/escape.go:618
	}
//line /usr/local/go/src/html/template/escape.go:618
	// _ = "end of CoverTab[31054]"
//line /usr/local/go/src/html/template/escape.go:618
	_go_fuzz_dep_.CoverTab[31055]++
							return c, ok
//line /usr/local/go/src/html/template/escape.go:619
	// _ = "end of CoverTab[31055]"
}

// escapeTemplate escapes a {{template}} call node.
func (e *escaper) escapeTemplate(c context, n *parse.TemplateNode) context {
//line /usr/local/go/src/html/template/escape.go:623
	_go_fuzz_dep_.CoverTab[31071]++
							c, name := e.escapeTree(c, n, n.Name, n.Line)
							if name != n.Name {
//line /usr/local/go/src/html/template/escape.go:625
		_go_fuzz_dep_.CoverTab[31073]++
								e.editTemplateNode(n, name)
//line /usr/local/go/src/html/template/escape.go:626
		// _ = "end of CoverTab[31073]"
	} else {
//line /usr/local/go/src/html/template/escape.go:627
		_go_fuzz_dep_.CoverTab[31074]++
//line /usr/local/go/src/html/template/escape.go:627
		// _ = "end of CoverTab[31074]"
//line /usr/local/go/src/html/template/escape.go:627
	}
//line /usr/local/go/src/html/template/escape.go:627
	// _ = "end of CoverTab[31071]"
//line /usr/local/go/src/html/template/escape.go:627
	_go_fuzz_dep_.CoverTab[31072]++
							return c
//line /usr/local/go/src/html/template/escape.go:628
	// _ = "end of CoverTab[31072]"
}

// escapeTree escapes the named template starting in the given context as
//line /usr/local/go/src/html/template/escape.go:631
// necessary and returns its output context.
//line /usr/local/go/src/html/template/escape.go:633
func (e *escaper) escapeTree(c context, node parse.Node, name string, line int) (context, string) {
//line /usr/local/go/src/html/template/escape.go:633
	_go_fuzz_dep_.CoverTab[31075]++

//line /usr/local/go/src/html/template/escape.go:636
	dname := c.mangle(name)
	e.called[dname] = true
	if out, ok := e.output[dname]; ok {
//line /usr/local/go/src/html/template/escape.go:638
		_go_fuzz_dep_.CoverTab[31079]++

								return out, dname
//line /usr/local/go/src/html/template/escape.go:640
		// _ = "end of CoverTab[31079]"
	} else {
//line /usr/local/go/src/html/template/escape.go:641
		_go_fuzz_dep_.CoverTab[31080]++
//line /usr/local/go/src/html/template/escape.go:641
		// _ = "end of CoverTab[31080]"
//line /usr/local/go/src/html/template/escape.go:641
	}
//line /usr/local/go/src/html/template/escape.go:641
	// _ = "end of CoverTab[31075]"
//line /usr/local/go/src/html/template/escape.go:641
	_go_fuzz_dep_.CoverTab[31076]++
							t := e.template(name)
							if t == nil {
//line /usr/local/go/src/html/template/escape.go:643
		_go_fuzz_dep_.CoverTab[31081]++

//line /usr/local/go/src/html/template/escape.go:646
		if e.ns.set[name] != nil {
//line /usr/local/go/src/html/template/escape.go:646
			_go_fuzz_dep_.CoverTab[31083]++
									return context{
				state:	stateError,
				err:	errorf(ErrNoSuchTemplate, node, line, "%q is an incomplete or empty template", name),
			}, dname
//line /usr/local/go/src/html/template/escape.go:650
			// _ = "end of CoverTab[31083]"
		} else {
//line /usr/local/go/src/html/template/escape.go:651
			_go_fuzz_dep_.CoverTab[31084]++
//line /usr/local/go/src/html/template/escape.go:651
			// _ = "end of CoverTab[31084]"
//line /usr/local/go/src/html/template/escape.go:651
		}
//line /usr/local/go/src/html/template/escape.go:651
		// _ = "end of CoverTab[31081]"
//line /usr/local/go/src/html/template/escape.go:651
		_go_fuzz_dep_.CoverTab[31082]++
								return context{
			state:	stateError,
			err:	errorf(ErrNoSuchTemplate, node, line, "no such template %q", name),
		}, dname
//line /usr/local/go/src/html/template/escape.go:655
		// _ = "end of CoverTab[31082]"
	} else {
//line /usr/local/go/src/html/template/escape.go:656
		_go_fuzz_dep_.CoverTab[31085]++
//line /usr/local/go/src/html/template/escape.go:656
		// _ = "end of CoverTab[31085]"
//line /usr/local/go/src/html/template/escape.go:656
	}
//line /usr/local/go/src/html/template/escape.go:656
	// _ = "end of CoverTab[31076]"
//line /usr/local/go/src/html/template/escape.go:656
	_go_fuzz_dep_.CoverTab[31077]++
							if dname != name {
//line /usr/local/go/src/html/template/escape.go:657
		_go_fuzz_dep_.CoverTab[31086]++

//line /usr/local/go/src/html/template/escape.go:660
		dt := e.template(dname)
		if dt == nil {
//line /usr/local/go/src/html/template/escape.go:661
			_go_fuzz_dep_.CoverTab[31088]++
									dt = template.New(dname)
									dt.Tree = &parse.Tree{Name: dname, Root: t.Root.CopyList()}
									e.derived[dname] = dt
//line /usr/local/go/src/html/template/escape.go:664
			// _ = "end of CoverTab[31088]"
		} else {
//line /usr/local/go/src/html/template/escape.go:665
			_go_fuzz_dep_.CoverTab[31089]++
//line /usr/local/go/src/html/template/escape.go:665
			// _ = "end of CoverTab[31089]"
//line /usr/local/go/src/html/template/escape.go:665
		}
//line /usr/local/go/src/html/template/escape.go:665
		// _ = "end of CoverTab[31086]"
//line /usr/local/go/src/html/template/escape.go:665
		_go_fuzz_dep_.CoverTab[31087]++
								t = dt
//line /usr/local/go/src/html/template/escape.go:666
		// _ = "end of CoverTab[31087]"
	} else {
//line /usr/local/go/src/html/template/escape.go:667
		_go_fuzz_dep_.CoverTab[31090]++
//line /usr/local/go/src/html/template/escape.go:667
		// _ = "end of CoverTab[31090]"
//line /usr/local/go/src/html/template/escape.go:667
	}
//line /usr/local/go/src/html/template/escape.go:667
	// _ = "end of CoverTab[31077]"
//line /usr/local/go/src/html/template/escape.go:667
	_go_fuzz_dep_.CoverTab[31078]++
							return e.computeOutCtx(c, t), dname
//line /usr/local/go/src/html/template/escape.go:668
	// _ = "end of CoverTab[31078]"
}

// computeOutCtx takes a template and its start context and computes the output
//line /usr/local/go/src/html/template/escape.go:671
// context while storing any inferences in e.
//line /usr/local/go/src/html/template/escape.go:673
func (e *escaper) computeOutCtx(c context, t *template.Template) context {
//line /usr/local/go/src/html/template/escape.go:673
	_go_fuzz_dep_.CoverTab[31091]++

							c1, ok := e.escapeTemplateBody(c, t)
							if !ok {
//line /usr/local/go/src/html/template/escape.go:676
		_go_fuzz_dep_.CoverTab[31094]++

								if c2, ok2 := e.escapeTemplateBody(c1, t); ok2 {
//line /usr/local/go/src/html/template/escape.go:678
			_go_fuzz_dep_.CoverTab[31095]++
									c1, ok = c2, true
//line /usr/local/go/src/html/template/escape.go:679
			// _ = "end of CoverTab[31095]"
		} else {
//line /usr/local/go/src/html/template/escape.go:680
			_go_fuzz_dep_.CoverTab[31096]++
//line /usr/local/go/src/html/template/escape.go:680
			// _ = "end of CoverTab[31096]"
//line /usr/local/go/src/html/template/escape.go:680
		}
//line /usr/local/go/src/html/template/escape.go:680
		// _ = "end of CoverTab[31094]"

	} else {
//line /usr/local/go/src/html/template/escape.go:682
		_go_fuzz_dep_.CoverTab[31097]++
//line /usr/local/go/src/html/template/escape.go:682
		// _ = "end of CoverTab[31097]"
//line /usr/local/go/src/html/template/escape.go:682
	}
//line /usr/local/go/src/html/template/escape.go:682
	// _ = "end of CoverTab[31091]"
//line /usr/local/go/src/html/template/escape.go:682
	_go_fuzz_dep_.CoverTab[31092]++
							if !ok && func() bool {
//line /usr/local/go/src/html/template/escape.go:683
		_go_fuzz_dep_.CoverTab[31098]++
//line /usr/local/go/src/html/template/escape.go:683
		return c1.state != stateError
//line /usr/local/go/src/html/template/escape.go:683
		// _ = "end of CoverTab[31098]"
//line /usr/local/go/src/html/template/escape.go:683
	}() {
//line /usr/local/go/src/html/template/escape.go:683
		_go_fuzz_dep_.CoverTab[31099]++
								return context{
			state:	stateError,
			err:	errorf(ErrOutputContext, t.Tree.Root, 0, "cannot compute output context for template %s", t.Name()),
		}
//line /usr/local/go/src/html/template/escape.go:687
		// _ = "end of CoverTab[31099]"
	} else {
//line /usr/local/go/src/html/template/escape.go:688
		_go_fuzz_dep_.CoverTab[31100]++
//line /usr/local/go/src/html/template/escape.go:688
		// _ = "end of CoverTab[31100]"
//line /usr/local/go/src/html/template/escape.go:688
	}
//line /usr/local/go/src/html/template/escape.go:688
	// _ = "end of CoverTab[31092]"
//line /usr/local/go/src/html/template/escape.go:688
	_go_fuzz_dep_.CoverTab[31093]++
							return c1
//line /usr/local/go/src/html/template/escape.go:689
	// _ = "end of CoverTab[31093]"
}

// escapeTemplateBody escapes the given template assuming the given output
//line /usr/local/go/src/html/template/escape.go:692
// context, and returns the best guess at the output context and whether the
//line /usr/local/go/src/html/template/escape.go:692
// assumption was correct.
//line /usr/local/go/src/html/template/escape.go:695
func (e *escaper) escapeTemplateBody(c context, t *template.Template) (context, bool) {
//line /usr/local/go/src/html/template/escape.go:695
	_go_fuzz_dep_.CoverTab[31101]++
							filter := func(e1 *escaper, c1 context) bool {
//line /usr/local/go/src/html/template/escape.go:696
		_go_fuzz_dep_.CoverTab[31103]++
								if c1.state == stateError {
//line /usr/local/go/src/html/template/escape.go:697
			_go_fuzz_dep_.CoverTab[31106]++

									return false
//line /usr/local/go/src/html/template/escape.go:699
			// _ = "end of CoverTab[31106]"
		} else {
//line /usr/local/go/src/html/template/escape.go:700
			_go_fuzz_dep_.CoverTab[31107]++
//line /usr/local/go/src/html/template/escape.go:700
			// _ = "end of CoverTab[31107]"
//line /usr/local/go/src/html/template/escape.go:700
		}
//line /usr/local/go/src/html/template/escape.go:700
		// _ = "end of CoverTab[31103]"
//line /usr/local/go/src/html/template/escape.go:700
		_go_fuzz_dep_.CoverTab[31104]++
								if !e1.called[t.Name()] {
//line /usr/local/go/src/html/template/escape.go:701
			_go_fuzz_dep_.CoverTab[31108]++

//line /usr/local/go/src/html/template/escape.go:704
			return true
//line /usr/local/go/src/html/template/escape.go:704
			// _ = "end of CoverTab[31108]"
		} else {
//line /usr/local/go/src/html/template/escape.go:705
			_go_fuzz_dep_.CoverTab[31109]++
//line /usr/local/go/src/html/template/escape.go:705
			// _ = "end of CoverTab[31109]"
//line /usr/local/go/src/html/template/escape.go:705
		}
//line /usr/local/go/src/html/template/escape.go:705
		// _ = "end of CoverTab[31104]"
//line /usr/local/go/src/html/template/escape.go:705
		_go_fuzz_dep_.CoverTab[31105]++

								return c.eq(c1)
//line /usr/local/go/src/html/template/escape.go:707
		// _ = "end of CoverTab[31105]"
	}
//line /usr/local/go/src/html/template/escape.go:708
	// _ = "end of CoverTab[31101]"
//line /usr/local/go/src/html/template/escape.go:708
	_go_fuzz_dep_.CoverTab[31102]++

//line /usr/local/go/src/html/template/escape.go:713
	e.output[t.Name()] = c
							return e.escapeListConditionally(c, t.Tree.Root, filter)
//line /usr/local/go/src/html/template/escape.go:714
	// _ = "end of CoverTab[31102]"
}

// delimEnds maps each delim to a string of characters that terminate it.
var delimEnds = [...]string{
							delimDoubleQuote:	`"`,
							delimSingleQuote:	"'",

//line /usr/local/go/src/html/template/escape.go:728
	delimSpaceOrTagEnd:	" \t\n\f\r>",
}

var doctypeBytes = []byte("<!DOCTYPE")

// escapeText escapes a text template node.
func (e *escaper) escapeText(c context, n *parse.TextNode) context {
//line /usr/local/go/src/html/template/escape.go:734
	_go_fuzz_dep_.CoverTab[31110]++
							s, written, i, b := n.Text, 0, 0, new(bytes.Buffer)
							for i != len(s) {
//line /usr/local/go/src/html/template/escape.go:736
		_go_fuzz_dep_.CoverTab[31113]++
								c1, nread := contextAfterText(c, s[i:])
								i1 := i + nread
								if c.state == stateText || func() bool {
//line /usr/local/go/src/html/template/escape.go:739
			_go_fuzz_dep_.CoverTab[31117]++
//line /usr/local/go/src/html/template/escape.go:739
			return c.state == stateRCDATA
//line /usr/local/go/src/html/template/escape.go:739
			// _ = "end of CoverTab[31117]"
//line /usr/local/go/src/html/template/escape.go:739
		}() {
//line /usr/local/go/src/html/template/escape.go:739
			_go_fuzz_dep_.CoverTab[31118]++
									end := i1
									if c1.state != c.state {
//line /usr/local/go/src/html/template/escape.go:741
				_go_fuzz_dep_.CoverTab[31120]++
										for j := end - 1; j >= i; j-- {
//line /usr/local/go/src/html/template/escape.go:742
					_go_fuzz_dep_.CoverTab[31121]++
											if s[j] == '<' {
//line /usr/local/go/src/html/template/escape.go:743
						_go_fuzz_dep_.CoverTab[31122]++
												end = j
												break
//line /usr/local/go/src/html/template/escape.go:745
						// _ = "end of CoverTab[31122]"
					} else {
//line /usr/local/go/src/html/template/escape.go:746
						_go_fuzz_dep_.CoverTab[31123]++
//line /usr/local/go/src/html/template/escape.go:746
						// _ = "end of CoverTab[31123]"
//line /usr/local/go/src/html/template/escape.go:746
					}
//line /usr/local/go/src/html/template/escape.go:746
					// _ = "end of CoverTab[31121]"
				}
//line /usr/local/go/src/html/template/escape.go:747
				// _ = "end of CoverTab[31120]"
			} else {
//line /usr/local/go/src/html/template/escape.go:748
				_go_fuzz_dep_.CoverTab[31124]++
//line /usr/local/go/src/html/template/escape.go:748
				// _ = "end of CoverTab[31124]"
//line /usr/local/go/src/html/template/escape.go:748
			}
//line /usr/local/go/src/html/template/escape.go:748
			// _ = "end of CoverTab[31118]"
//line /usr/local/go/src/html/template/escape.go:748
			_go_fuzz_dep_.CoverTab[31119]++
									for j := i; j < end; j++ {
//line /usr/local/go/src/html/template/escape.go:749
				_go_fuzz_dep_.CoverTab[31125]++
										if s[j] == '<' && func() bool {
//line /usr/local/go/src/html/template/escape.go:750
					_go_fuzz_dep_.CoverTab[31126]++
//line /usr/local/go/src/html/template/escape.go:750
					return !bytes.HasPrefix(bytes.ToUpper(s[j:]), doctypeBytes)
//line /usr/local/go/src/html/template/escape.go:750
					// _ = "end of CoverTab[31126]"
//line /usr/local/go/src/html/template/escape.go:750
				}() {
//line /usr/local/go/src/html/template/escape.go:750
					_go_fuzz_dep_.CoverTab[31127]++
											b.Write(s[written:j])
											b.WriteString("&lt;")
											written = j + 1
//line /usr/local/go/src/html/template/escape.go:753
					// _ = "end of CoverTab[31127]"
				} else {
//line /usr/local/go/src/html/template/escape.go:754
					_go_fuzz_dep_.CoverTab[31128]++
//line /usr/local/go/src/html/template/escape.go:754
					// _ = "end of CoverTab[31128]"
//line /usr/local/go/src/html/template/escape.go:754
				}
//line /usr/local/go/src/html/template/escape.go:754
				// _ = "end of CoverTab[31125]"
			}
//line /usr/local/go/src/html/template/escape.go:755
			// _ = "end of CoverTab[31119]"
		} else {
//line /usr/local/go/src/html/template/escape.go:756
			_go_fuzz_dep_.CoverTab[31129]++
//line /usr/local/go/src/html/template/escape.go:756
			if isComment(c.state) && func() bool {
//line /usr/local/go/src/html/template/escape.go:756
				_go_fuzz_dep_.CoverTab[31130]++
//line /usr/local/go/src/html/template/escape.go:756
				return c.delim == delimNone
//line /usr/local/go/src/html/template/escape.go:756
				// _ = "end of CoverTab[31130]"
//line /usr/local/go/src/html/template/escape.go:756
			}() {
//line /usr/local/go/src/html/template/escape.go:756
				_go_fuzz_dep_.CoverTab[31131]++
										switch c.state {
				case stateJSBlockCmt:
//line /usr/local/go/src/html/template/escape.go:758
					_go_fuzz_dep_.CoverTab[31133]++

//line /usr/local/go/src/html/template/escape.go:766
					if bytes.ContainsAny(s[written:i1], "\n\r\u2028\u2029") {
//line /usr/local/go/src/html/template/escape.go:766
						_go_fuzz_dep_.CoverTab[31136]++
												b.WriteByte('\n')
//line /usr/local/go/src/html/template/escape.go:767
						// _ = "end of CoverTab[31136]"
					} else {
//line /usr/local/go/src/html/template/escape.go:768
						_go_fuzz_dep_.CoverTab[31137]++
												b.WriteByte(' ')
//line /usr/local/go/src/html/template/escape.go:769
						// _ = "end of CoverTab[31137]"
					}
//line /usr/local/go/src/html/template/escape.go:770
					// _ = "end of CoverTab[31133]"
				case stateCSSBlockCmt:
//line /usr/local/go/src/html/template/escape.go:771
					_go_fuzz_dep_.CoverTab[31134]++
											b.WriteByte(' ')
//line /usr/local/go/src/html/template/escape.go:772
					// _ = "end of CoverTab[31134]"
//line /usr/local/go/src/html/template/escape.go:772
				default:
//line /usr/local/go/src/html/template/escape.go:772
					_go_fuzz_dep_.CoverTab[31135]++
//line /usr/local/go/src/html/template/escape.go:772
					// _ = "end of CoverTab[31135]"
				}
//line /usr/local/go/src/html/template/escape.go:773
				// _ = "end of CoverTab[31131]"
//line /usr/local/go/src/html/template/escape.go:773
				_go_fuzz_dep_.CoverTab[31132]++
										written = i1
//line /usr/local/go/src/html/template/escape.go:774
				// _ = "end of CoverTab[31132]"
			} else {
//line /usr/local/go/src/html/template/escape.go:775
				_go_fuzz_dep_.CoverTab[31138]++
//line /usr/local/go/src/html/template/escape.go:775
				// _ = "end of CoverTab[31138]"
//line /usr/local/go/src/html/template/escape.go:775
			}
//line /usr/local/go/src/html/template/escape.go:775
			// _ = "end of CoverTab[31129]"
//line /usr/local/go/src/html/template/escape.go:775
		}
//line /usr/local/go/src/html/template/escape.go:775
		// _ = "end of CoverTab[31113]"
//line /usr/local/go/src/html/template/escape.go:775
		_go_fuzz_dep_.CoverTab[31114]++
								if c.state != c1.state && func() bool {
//line /usr/local/go/src/html/template/escape.go:776
			_go_fuzz_dep_.CoverTab[31139]++
//line /usr/local/go/src/html/template/escape.go:776
			return isComment(c1.state)
//line /usr/local/go/src/html/template/escape.go:776
			// _ = "end of CoverTab[31139]"
//line /usr/local/go/src/html/template/escape.go:776
		}() && func() bool {
//line /usr/local/go/src/html/template/escape.go:776
			_go_fuzz_dep_.CoverTab[31140]++
//line /usr/local/go/src/html/template/escape.go:776
			return c1.delim == delimNone
//line /usr/local/go/src/html/template/escape.go:776
			// _ = "end of CoverTab[31140]"
//line /usr/local/go/src/html/template/escape.go:776
		}() {
//line /usr/local/go/src/html/template/escape.go:776
			_go_fuzz_dep_.CoverTab[31141]++

									cs := i1 - 2
									if c1.state == stateHTMLCmt {
//line /usr/local/go/src/html/template/escape.go:779
				_go_fuzz_dep_.CoverTab[31143]++

										cs -= 2
//line /usr/local/go/src/html/template/escape.go:781
				// _ = "end of CoverTab[31143]"
			} else {
//line /usr/local/go/src/html/template/escape.go:782
				_go_fuzz_dep_.CoverTab[31144]++
//line /usr/local/go/src/html/template/escape.go:782
				// _ = "end of CoverTab[31144]"
//line /usr/local/go/src/html/template/escape.go:782
			}
//line /usr/local/go/src/html/template/escape.go:782
			// _ = "end of CoverTab[31141]"
//line /usr/local/go/src/html/template/escape.go:782
			_go_fuzz_dep_.CoverTab[31142]++
									b.Write(s[written:cs])
									written = i1
//line /usr/local/go/src/html/template/escape.go:784
			// _ = "end of CoverTab[31142]"
		} else {
//line /usr/local/go/src/html/template/escape.go:785
			_go_fuzz_dep_.CoverTab[31145]++
//line /usr/local/go/src/html/template/escape.go:785
			// _ = "end of CoverTab[31145]"
//line /usr/local/go/src/html/template/escape.go:785
		}
//line /usr/local/go/src/html/template/escape.go:785
		// _ = "end of CoverTab[31114]"
//line /usr/local/go/src/html/template/escape.go:785
		_go_fuzz_dep_.CoverTab[31115]++
								if i == i1 && func() bool {
//line /usr/local/go/src/html/template/escape.go:786
			_go_fuzz_dep_.CoverTab[31146]++
//line /usr/local/go/src/html/template/escape.go:786
			return c.state == c1.state
//line /usr/local/go/src/html/template/escape.go:786
			// _ = "end of CoverTab[31146]"
//line /usr/local/go/src/html/template/escape.go:786
		}() {
//line /usr/local/go/src/html/template/escape.go:786
			_go_fuzz_dep_.CoverTab[31147]++
									panic(fmt.Sprintf("infinite loop from %v to %v on %q..%q", c, c1, s[:i], s[i:]))
//line /usr/local/go/src/html/template/escape.go:787
			// _ = "end of CoverTab[31147]"
		} else {
//line /usr/local/go/src/html/template/escape.go:788
			_go_fuzz_dep_.CoverTab[31148]++
//line /usr/local/go/src/html/template/escape.go:788
			// _ = "end of CoverTab[31148]"
//line /usr/local/go/src/html/template/escape.go:788
		}
//line /usr/local/go/src/html/template/escape.go:788
		// _ = "end of CoverTab[31115]"
//line /usr/local/go/src/html/template/escape.go:788
		_go_fuzz_dep_.CoverTab[31116]++
								c, i = c1, i1
//line /usr/local/go/src/html/template/escape.go:789
		// _ = "end of CoverTab[31116]"
	}
//line /usr/local/go/src/html/template/escape.go:790
	// _ = "end of CoverTab[31110]"
//line /usr/local/go/src/html/template/escape.go:790
	_go_fuzz_dep_.CoverTab[31111]++

							if written != 0 && func() bool {
//line /usr/local/go/src/html/template/escape.go:792
		_go_fuzz_dep_.CoverTab[31149]++
//line /usr/local/go/src/html/template/escape.go:792
		return c.state != stateError
//line /usr/local/go/src/html/template/escape.go:792
		// _ = "end of CoverTab[31149]"
//line /usr/local/go/src/html/template/escape.go:792
	}() {
//line /usr/local/go/src/html/template/escape.go:792
		_go_fuzz_dep_.CoverTab[31150]++
								if !isComment(c.state) || func() bool {
//line /usr/local/go/src/html/template/escape.go:793
			_go_fuzz_dep_.CoverTab[31152]++
//line /usr/local/go/src/html/template/escape.go:793
			return c.delim != delimNone
//line /usr/local/go/src/html/template/escape.go:793
			// _ = "end of CoverTab[31152]"
//line /usr/local/go/src/html/template/escape.go:793
		}() {
//line /usr/local/go/src/html/template/escape.go:793
			_go_fuzz_dep_.CoverTab[31153]++
									b.Write(n.Text[written:])
//line /usr/local/go/src/html/template/escape.go:794
			// _ = "end of CoverTab[31153]"
		} else {
//line /usr/local/go/src/html/template/escape.go:795
			_go_fuzz_dep_.CoverTab[31154]++
//line /usr/local/go/src/html/template/escape.go:795
			// _ = "end of CoverTab[31154]"
//line /usr/local/go/src/html/template/escape.go:795
		}
//line /usr/local/go/src/html/template/escape.go:795
		// _ = "end of CoverTab[31150]"
//line /usr/local/go/src/html/template/escape.go:795
		_go_fuzz_dep_.CoverTab[31151]++
								e.editTextNode(n, b.Bytes())
//line /usr/local/go/src/html/template/escape.go:796
		// _ = "end of CoverTab[31151]"
	} else {
//line /usr/local/go/src/html/template/escape.go:797
		_go_fuzz_dep_.CoverTab[31155]++
//line /usr/local/go/src/html/template/escape.go:797
		// _ = "end of CoverTab[31155]"
//line /usr/local/go/src/html/template/escape.go:797
	}
//line /usr/local/go/src/html/template/escape.go:797
	// _ = "end of CoverTab[31111]"
//line /usr/local/go/src/html/template/escape.go:797
	_go_fuzz_dep_.CoverTab[31112]++
							return c
//line /usr/local/go/src/html/template/escape.go:798
	// _ = "end of CoverTab[31112]"
}

// contextAfterText starts in context c, consumes some tokens from the front of
//line /usr/local/go/src/html/template/escape.go:801
// s, then returns the context after those tokens and the unprocessed suffix.
//line /usr/local/go/src/html/template/escape.go:803
func contextAfterText(c context, s []byte) (context, int) {
//line /usr/local/go/src/html/template/escape.go:803
	_go_fuzz_dep_.CoverTab[31156]++
							if c.delim == delimNone {
//line /usr/local/go/src/html/template/escape.go:804
		_go_fuzz_dep_.CoverTab[31163]++
								c1, i := tSpecialTagEnd(c, s)
								if i == 0 {
//line /usr/local/go/src/html/template/escape.go:806
			_go_fuzz_dep_.CoverTab[31165]++

//line /usr/local/go/src/html/template/escape.go:809
			return c1, 0
//line /usr/local/go/src/html/template/escape.go:809
			// _ = "end of CoverTab[31165]"
		} else {
//line /usr/local/go/src/html/template/escape.go:810
			_go_fuzz_dep_.CoverTab[31166]++
//line /usr/local/go/src/html/template/escape.go:810
			// _ = "end of CoverTab[31166]"
//line /usr/local/go/src/html/template/escape.go:810
		}
//line /usr/local/go/src/html/template/escape.go:810
		// _ = "end of CoverTab[31163]"
//line /usr/local/go/src/html/template/escape.go:810
		_go_fuzz_dep_.CoverTab[31164]++

								return transitionFunc[c.state](c, s[:i])
//line /usr/local/go/src/html/template/escape.go:812
		// _ = "end of CoverTab[31164]"
	} else {
//line /usr/local/go/src/html/template/escape.go:813
		_go_fuzz_dep_.CoverTab[31167]++
//line /usr/local/go/src/html/template/escape.go:813
		// _ = "end of CoverTab[31167]"
//line /usr/local/go/src/html/template/escape.go:813
	}
//line /usr/local/go/src/html/template/escape.go:813
	// _ = "end of CoverTab[31156]"
//line /usr/local/go/src/html/template/escape.go:813
	_go_fuzz_dep_.CoverTab[31157]++

//line /usr/local/go/src/html/template/escape.go:817
	i := bytes.IndexAny(s, delimEnds[c.delim])
	if i == -1 {
//line /usr/local/go/src/html/template/escape.go:818
		_go_fuzz_dep_.CoverTab[31168]++
								i = len(s)
//line /usr/local/go/src/html/template/escape.go:819
		// _ = "end of CoverTab[31168]"
	} else {
//line /usr/local/go/src/html/template/escape.go:820
		_go_fuzz_dep_.CoverTab[31169]++
//line /usr/local/go/src/html/template/escape.go:820
		// _ = "end of CoverTab[31169]"
//line /usr/local/go/src/html/template/escape.go:820
	}
//line /usr/local/go/src/html/template/escape.go:820
	// _ = "end of CoverTab[31157]"
//line /usr/local/go/src/html/template/escape.go:820
	_go_fuzz_dep_.CoverTab[31158]++
							if c.delim == delimSpaceOrTagEnd {
//line /usr/local/go/src/html/template/escape.go:821
		_go_fuzz_dep_.CoverTab[31170]++

//line /usr/local/go/src/html/template/escape.go:829
		if j := bytes.IndexAny(s[:i], "\"'<=`"); j >= 0 {
//line /usr/local/go/src/html/template/escape.go:829
			_go_fuzz_dep_.CoverTab[31171]++
									return context{
				state:	stateError,
				err:	errorf(ErrBadHTML, nil, 0, "%q in unquoted attr: %q", s[j:j+1], s[:i]),
			}, len(s)
//line /usr/local/go/src/html/template/escape.go:833
			// _ = "end of CoverTab[31171]"
		} else {
//line /usr/local/go/src/html/template/escape.go:834
			_go_fuzz_dep_.CoverTab[31172]++
//line /usr/local/go/src/html/template/escape.go:834
			// _ = "end of CoverTab[31172]"
//line /usr/local/go/src/html/template/escape.go:834
		}
//line /usr/local/go/src/html/template/escape.go:834
		// _ = "end of CoverTab[31170]"
	} else {
//line /usr/local/go/src/html/template/escape.go:835
		_go_fuzz_dep_.CoverTab[31173]++
//line /usr/local/go/src/html/template/escape.go:835
		// _ = "end of CoverTab[31173]"
//line /usr/local/go/src/html/template/escape.go:835
	}
//line /usr/local/go/src/html/template/escape.go:835
	// _ = "end of CoverTab[31158]"
//line /usr/local/go/src/html/template/escape.go:835
	_go_fuzz_dep_.CoverTab[31159]++
							if i == len(s) {
//line /usr/local/go/src/html/template/escape.go:836
		_go_fuzz_dep_.CoverTab[31174]++

//line /usr/local/go/src/html/template/escape.go:841
		for u := []byte(html.UnescapeString(string(s))); len(u) != 0; {
//line /usr/local/go/src/html/template/escape.go:841
			_go_fuzz_dep_.CoverTab[31176]++
									c1, i1 := transitionFunc[c.state](c, u)
									c, u = c1, u[i1:]
//line /usr/local/go/src/html/template/escape.go:843
			// _ = "end of CoverTab[31176]"
		}
//line /usr/local/go/src/html/template/escape.go:844
		// _ = "end of CoverTab[31174]"
//line /usr/local/go/src/html/template/escape.go:844
		_go_fuzz_dep_.CoverTab[31175]++
								return c, len(s)
//line /usr/local/go/src/html/template/escape.go:845
		// _ = "end of CoverTab[31175]"
	} else {
//line /usr/local/go/src/html/template/escape.go:846
		_go_fuzz_dep_.CoverTab[31177]++
//line /usr/local/go/src/html/template/escape.go:846
		// _ = "end of CoverTab[31177]"
//line /usr/local/go/src/html/template/escape.go:846
	}
//line /usr/local/go/src/html/template/escape.go:846
	// _ = "end of CoverTab[31159]"
//line /usr/local/go/src/html/template/escape.go:846
	_go_fuzz_dep_.CoverTab[31160]++

							element := c.element

//line /usr/local/go/src/html/template/escape.go:851
	if c.state == stateAttr && func() bool {
//line /usr/local/go/src/html/template/escape.go:851
		_go_fuzz_dep_.CoverTab[31178]++
//line /usr/local/go/src/html/template/escape.go:851
		return c.element == elementScript
//line /usr/local/go/src/html/template/escape.go:851
		// _ = "end of CoverTab[31178]"
//line /usr/local/go/src/html/template/escape.go:851
	}() && func() bool {
//line /usr/local/go/src/html/template/escape.go:851
		_go_fuzz_dep_.CoverTab[31179]++
//line /usr/local/go/src/html/template/escape.go:851
		return c.attr == attrScriptType
//line /usr/local/go/src/html/template/escape.go:851
		// _ = "end of CoverTab[31179]"
//line /usr/local/go/src/html/template/escape.go:851
	}() && func() bool {
//line /usr/local/go/src/html/template/escape.go:851
		_go_fuzz_dep_.CoverTab[31180]++
//line /usr/local/go/src/html/template/escape.go:851
		return !isJSType(string(s[:i]))
//line /usr/local/go/src/html/template/escape.go:851
		// _ = "end of CoverTab[31180]"
//line /usr/local/go/src/html/template/escape.go:851
	}() {
//line /usr/local/go/src/html/template/escape.go:851
		_go_fuzz_dep_.CoverTab[31181]++
								element = elementNone
//line /usr/local/go/src/html/template/escape.go:852
		// _ = "end of CoverTab[31181]"
	} else {
//line /usr/local/go/src/html/template/escape.go:853
		_go_fuzz_dep_.CoverTab[31182]++
//line /usr/local/go/src/html/template/escape.go:853
		// _ = "end of CoverTab[31182]"
//line /usr/local/go/src/html/template/escape.go:853
	}
//line /usr/local/go/src/html/template/escape.go:853
	// _ = "end of CoverTab[31160]"
//line /usr/local/go/src/html/template/escape.go:853
	_go_fuzz_dep_.CoverTab[31161]++

							if c.delim != delimSpaceOrTagEnd {
//line /usr/local/go/src/html/template/escape.go:855
		_go_fuzz_dep_.CoverTab[31183]++

								i++
//line /usr/local/go/src/html/template/escape.go:857
		// _ = "end of CoverTab[31183]"
	} else {
//line /usr/local/go/src/html/template/escape.go:858
		_go_fuzz_dep_.CoverTab[31184]++
//line /usr/local/go/src/html/template/escape.go:858
		// _ = "end of CoverTab[31184]"
//line /usr/local/go/src/html/template/escape.go:858
	}
//line /usr/local/go/src/html/template/escape.go:858
	// _ = "end of CoverTab[31161]"
//line /usr/local/go/src/html/template/escape.go:858
	_go_fuzz_dep_.CoverTab[31162]++

//line /usr/local/go/src/html/template/escape.go:861
	return context{state: stateTag, element: element}, i
//line /usr/local/go/src/html/template/escape.go:861
	// _ = "end of CoverTab[31162]"
}

// editActionNode records a change to an action pipeline for later commit.
func (e *escaper) editActionNode(n *parse.ActionNode, cmds []string) {
//line /usr/local/go/src/html/template/escape.go:865
	_go_fuzz_dep_.CoverTab[31185]++
							if _, ok := e.actionNodeEdits[n]; ok {
//line /usr/local/go/src/html/template/escape.go:866
		_go_fuzz_dep_.CoverTab[31187]++
								panic(fmt.Sprintf("node %s shared between templates", n))
//line /usr/local/go/src/html/template/escape.go:867
		// _ = "end of CoverTab[31187]"
	} else {
//line /usr/local/go/src/html/template/escape.go:868
		_go_fuzz_dep_.CoverTab[31188]++
//line /usr/local/go/src/html/template/escape.go:868
		// _ = "end of CoverTab[31188]"
//line /usr/local/go/src/html/template/escape.go:868
	}
//line /usr/local/go/src/html/template/escape.go:868
	// _ = "end of CoverTab[31185]"
//line /usr/local/go/src/html/template/escape.go:868
	_go_fuzz_dep_.CoverTab[31186]++
							e.actionNodeEdits[n] = cmds
//line /usr/local/go/src/html/template/escape.go:869
	// _ = "end of CoverTab[31186]"
}

// editTemplateNode records a change to a {{template}} callee for later commit.
func (e *escaper) editTemplateNode(n *parse.TemplateNode, callee string) {
//line /usr/local/go/src/html/template/escape.go:873
	_go_fuzz_dep_.CoverTab[31189]++
							if _, ok := e.templateNodeEdits[n]; ok {
//line /usr/local/go/src/html/template/escape.go:874
		_go_fuzz_dep_.CoverTab[31191]++
								panic(fmt.Sprintf("node %s shared between templates", n))
//line /usr/local/go/src/html/template/escape.go:875
		// _ = "end of CoverTab[31191]"
	} else {
//line /usr/local/go/src/html/template/escape.go:876
		_go_fuzz_dep_.CoverTab[31192]++
//line /usr/local/go/src/html/template/escape.go:876
		// _ = "end of CoverTab[31192]"
//line /usr/local/go/src/html/template/escape.go:876
	}
//line /usr/local/go/src/html/template/escape.go:876
	// _ = "end of CoverTab[31189]"
//line /usr/local/go/src/html/template/escape.go:876
	_go_fuzz_dep_.CoverTab[31190]++
							e.templateNodeEdits[n] = callee
//line /usr/local/go/src/html/template/escape.go:877
	// _ = "end of CoverTab[31190]"
}

// editTextNode records a change to a text node for later commit.
func (e *escaper) editTextNode(n *parse.TextNode, text []byte) {
//line /usr/local/go/src/html/template/escape.go:881
	_go_fuzz_dep_.CoverTab[31193]++
							if _, ok := e.textNodeEdits[n]; ok {
//line /usr/local/go/src/html/template/escape.go:882
		_go_fuzz_dep_.CoverTab[31195]++
								panic(fmt.Sprintf("node %s shared between templates", n))
//line /usr/local/go/src/html/template/escape.go:883
		// _ = "end of CoverTab[31195]"
	} else {
//line /usr/local/go/src/html/template/escape.go:884
		_go_fuzz_dep_.CoverTab[31196]++
//line /usr/local/go/src/html/template/escape.go:884
		// _ = "end of CoverTab[31196]"
//line /usr/local/go/src/html/template/escape.go:884
	}
//line /usr/local/go/src/html/template/escape.go:884
	// _ = "end of CoverTab[31193]"
//line /usr/local/go/src/html/template/escape.go:884
	_go_fuzz_dep_.CoverTab[31194]++
							e.textNodeEdits[n] = text
//line /usr/local/go/src/html/template/escape.go:885
	// _ = "end of CoverTab[31194]"
}

// commit applies changes to actions and template calls needed to contextually
//line /usr/local/go/src/html/template/escape.go:888
// autoescape content and adds any derived templates to the set.
//line /usr/local/go/src/html/template/escape.go:890
func (e *escaper) commit() {
//line /usr/local/go/src/html/template/escape.go:890
	_go_fuzz_dep_.CoverTab[31197]++
							for name := range e.output {
//line /usr/local/go/src/html/template/escape.go:891
		_go_fuzz_dep_.CoverTab[31203]++
								e.template(name).Funcs(funcMap)
//line /usr/local/go/src/html/template/escape.go:892
		// _ = "end of CoverTab[31203]"
	}
//line /usr/local/go/src/html/template/escape.go:893
	// _ = "end of CoverTab[31197]"
//line /usr/local/go/src/html/template/escape.go:893
	_go_fuzz_dep_.CoverTab[31198]++

//line /usr/local/go/src/html/template/escape.go:896
	tmpl := e.arbitraryTemplate()
	for _, t := range e.derived {
//line /usr/local/go/src/html/template/escape.go:897
		_go_fuzz_dep_.CoverTab[31204]++
								if _, err := tmpl.text.AddParseTree(t.Name(), t.Tree); err != nil {
//line /usr/local/go/src/html/template/escape.go:898
			_go_fuzz_dep_.CoverTab[31205]++
									panic("error adding derived template")
//line /usr/local/go/src/html/template/escape.go:899
			// _ = "end of CoverTab[31205]"
		} else {
//line /usr/local/go/src/html/template/escape.go:900
			_go_fuzz_dep_.CoverTab[31206]++
//line /usr/local/go/src/html/template/escape.go:900
			// _ = "end of CoverTab[31206]"
//line /usr/local/go/src/html/template/escape.go:900
		}
//line /usr/local/go/src/html/template/escape.go:900
		// _ = "end of CoverTab[31204]"
	}
//line /usr/local/go/src/html/template/escape.go:901
	// _ = "end of CoverTab[31198]"
//line /usr/local/go/src/html/template/escape.go:901
	_go_fuzz_dep_.CoverTab[31199]++
							for n, s := range e.actionNodeEdits {
//line /usr/local/go/src/html/template/escape.go:902
		_go_fuzz_dep_.CoverTab[31207]++
								ensurePipelineContains(n.Pipe, s)
//line /usr/local/go/src/html/template/escape.go:903
		// _ = "end of CoverTab[31207]"
	}
//line /usr/local/go/src/html/template/escape.go:904
	// _ = "end of CoverTab[31199]"
//line /usr/local/go/src/html/template/escape.go:904
	_go_fuzz_dep_.CoverTab[31200]++
							for n, name := range e.templateNodeEdits {
//line /usr/local/go/src/html/template/escape.go:905
		_go_fuzz_dep_.CoverTab[31208]++
								n.Name = name
//line /usr/local/go/src/html/template/escape.go:906
		// _ = "end of CoverTab[31208]"
	}
//line /usr/local/go/src/html/template/escape.go:907
	// _ = "end of CoverTab[31200]"
//line /usr/local/go/src/html/template/escape.go:907
	_go_fuzz_dep_.CoverTab[31201]++
							for n, s := range e.textNodeEdits {
//line /usr/local/go/src/html/template/escape.go:908
		_go_fuzz_dep_.CoverTab[31209]++
								n.Text = s
//line /usr/local/go/src/html/template/escape.go:909
		// _ = "end of CoverTab[31209]"
	}
//line /usr/local/go/src/html/template/escape.go:910
	// _ = "end of CoverTab[31201]"
//line /usr/local/go/src/html/template/escape.go:910
	_go_fuzz_dep_.CoverTab[31202]++

//line /usr/local/go/src/html/template/escape.go:913
	e.called = make(map[string]bool)
							e.actionNodeEdits = make(map[*parse.ActionNode][]string)
							e.templateNodeEdits = make(map[*parse.TemplateNode]string)
							e.textNodeEdits = make(map[*parse.TextNode][]byte)
//line /usr/local/go/src/html/template/escape.go:916
	// _ = "end of CoverTab[31202]"
}

// template returns the named template given a mangled template name.
func (e *escaper) template(name string) *template.Template {
//line /usr/local/go/src/html/template/escape.go:920
	_go_fuzz_dep_.CoverTab[31210]++

//line /usr/local/go/src/html/template/escape.go:923
	t := e.arbitraryTemplate().text.Lookup(name)
	if t == nil {
//line /usr/local/go/src/html/template/escape.go:924
		_go_fuzz_dep_.CoverTab[31212]++
								t = e.derived[name]
//line /usr/local/go/src/html/template/escape.go:925
		// _ = "end of CoverTab[31212]"
	} else {
//line /usr/local/go/src/html/template/escape.go:926
		_go_fuzz_dep_.CoverTab[31213]++
//line /usr/local/go/src/html/template/escape.go:926
		// _ = "end of CoverTab[31213]"
//line /usr/local/go/src/html/template/escape.go:926
	}
//line /usr/local/go/src/html/template/escape.go:926
	// _ = "end of CoverTab[31210]"
//line /usr/local/go/src/html/template/escape.go:926
	_go_fuzz_dep_.CoverTab[31211]++
							return t
//line /usr/local/go/src/html/template/escape.go:927
	// _ = "end of CoverTab[31211]"
}

// arbitraryTemplate returns an arbitrary template from the name space
//line /usr/local/go/src/html/template/escape.go:930
// associated with e and panics if no templates are found.
//line /usr/local/go/src/html/template/escape.go:932
func (e *escaper) arbitraryTemplate() *Template {
//line /usr/local/go/src/html/template/escape.go:932
	_go_fuzz_dep_.CoverTab[31214]++
							for _, t := range e.ns.set {
//line /usr/local/go/src/html/template/escape.go:933
		_go_fuzz_dep_.CoverTab[31216]++
								return t
//line /usr/local/go/src/html/template/escape.go:934
		// _ = "end of CoverTab[31216]"
	}
//line /usr/local/go/src/html/template/escape.go:935
	// _ = "end of CoverTab[31214]"
//line /usr/local/go/src/html/template/escape.go:935
	_go_fuzz_dep_.CoverTab[31215]++
							panic("no templates in name space")
//line /usr/local/go/src/html/template/escape.go:936
	// _ = "end of CoverTab[31215]"
}

//line /usr/local/go/src/html/template/escape.go:942
// HTMLEscape writes to w the escaped HTML equivalent of the plain text data b.
func HTMLEscape(w io.Writer, b []byte) {
//line /usr/local/go/src/html/template/escape.go:943
	_go_fuzz_dep_.CoverTab[31217]++
							template.HTMLEscape(w, b)
//line /usr/local/go/src/html/template/escape.go:944
	// _ = "end of CoverTab[31217]"
}

// HTMLEscapeString returns the escaped HTML equivalent of the plain text data s.
func HTMLEscapeString(s string) string {
//line /usr/local/go/src/html/template/escape.go:948
	_go_fuzz_dep_.CoverTab[31218]++
							return template.HTMLEscapeString(s)
//line /usr/local/go/src/html/template/escape.go:949
	// _ = "end of CoverTab[31218]"
}

// HTMLEscaper returns the escaped HTML equivalent of the textual
//line /usr/local/go/src/html/template/escape.go:952
// representation of its arguments.
//line /usr/local/go/src/html/template/escape.go:954
func HTMLEscaper(args ...any) string {
//line /usr/local/go/src/html/template/escape.go:954
	_go_fuzz_dep_.CoverTab[31219]++
							return template.HTMLEscaper(args...)
//line /usr/local/go/src/html/template/escape.go:955
	// _ = "end of CoverTab[31219]"
}

// JSEscape writes to w the escaped JavaScript equivalent of the plain text data b.
func JSEscape(w io.Writer, b []byte) {
//line /usr/local/go/src/html/template/escape.go:959
	_go_fuzz_dep_.CoverTab[31220]++
							template.JSEscape(w, b)
//line /usr/local/go/src/html/template/escape.go:960
	// _ = "end of CoverTab[31220]"
}

// JSEscapeString returns the escaped JavaScript equivalent of the plain text data s.
func JSEscapeString(s string) string {
//line /usr/local/go/src/html/template/escape.go:964
	_go_fuzz_dep_.CoverTab[31221]++
							return template.JSEscapeString(s)
//line /usr/local/go/src/html/template/escape.go:965
	// _ = "end of CoverTab[31221]"
}

// JSEscaper returns the escaped JavaScript equivalent of the textual
//line /usr/local/go/src/html/template/escape.go:968
// representation of its arguments.
//line /usr/local/go/src/html/template/escape.go:970
func JSEscaper(args ...any) string {
//line /usr/local/go/src/html/template/escape.go:970
	_go_fuzz_dep_.CoverTab[31222]++
							return template.JSEscaper(args...)
//line /usr/local/go/src/html/template/escape.go:971
	// _ = "end of CoverTab[31222]"
}

// URLQueryEscaper returns the escaped value of the textual representation of
//line /usr/local/go/src/html/template/escape.go:974
// its arguments in a form suitable for embedding in a URL query.
//line /usr/local/go/src/html/template/escape.go:976
func URLQueryEscaper(args ...any) string {
//line /usr/local/go/src/html/template/escape.go:976
	_go_fuzz_dep_.CoverTab[31223]++
							return template.URLQueryEscaper(args...)
//line /usr/local/go/src/html/template/escape.go:977
	// _ = "end of CoverTab[31223]"
}

//line /usr/local/go/src/html/template/escape.go:978
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/html/template/escape.go:978
var _ = _go_fuzz_dep_.CoverTab
