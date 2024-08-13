package schema

import (
	"fmt"
	"reflect"
)

type Vertex interface {
	Label() string
	Properties() map[string]interface{}
	GetBackendMeta() interface{}
	SetBackendMeta(interface{})
}

type Edge struct {
	Source     Vertex
	Label      string
	Target     Vertex
	Properties map[string]interface{}
}

type vertexBase struct {
	backendMeta interface{}
}

func (vb *vertexBase) GetBackendMeta() interface{} {
	if vb == nil {
		fmt.Printf("error GetBackendMeta for %v, type: %v\n",
			vb, reflect.TypeOf(vb).Elem().Name())
		return nil
	}
	return vb.backendMeta
}

func (vb *vertexBase) SetBackendMeta(v interface{}) {
	vb.backendMeta = v
}

func (_ *vertexBase) Properties() map[string]interface{} {
	return map[string]interface{}{}
}

type Package struct {
	vertexBase
	SourceURL string `json:"SourceURL"`
	Version   string
	Path      string
	PkgID     string
	Functions []*Function `json:"-"`
}

func (_ *Package) Label() string {
	return "package"
}

func (p *Package) Properties() map[string]interface{} {
	return map[string]interface{}{
		"SourceURL": p.SourceURL,
		"Version":   p.Version,
		"Path":      p.Path,
		"PkgID":     p.PkgID,
	}
}

type Function struct {
	vertexBase
	Id             string `json:"_id"`
	Key            string `json:"_key"`
	Name           string
	File           string
	Pos            string
	Package        *Package        `json:"-"`
	FirstStatement *Instruction    `json:"-"`
	Statements     []*Instruction  `json:"-"`
	Calls          []*FunctionCall `json:"-"`
}

func (_ *Function) Label() string {
	return "function"
}

func (f *Function) Properties() map[string]interface{} {
	return map[string]interface{}{
		"Name": f.Name,
		"File": f.File,
		"Pos":  f.Pos,
	}
}

type Variable struct {
	vertexBase
	Id       string `json:"_id"`
	Key      string `json:"_key"`
	Name     string
	Type     string
	IsGlobal bool
}

func (_ *Variable) Label() string {
	return "variable"
}

func (v *Variable) Properties() map[string]interface{} {
	return map[string]interface{}{
		"Name":     v.Name,
		"Type":     v.Type,
		"IsGlobal": v.IsGlobal,
	}
}

type Parameter struct {
	vertexBase
	Id       string `json:"_id"`
	Key      string `json:"_key"`
	Name     string
	Type     string
	ParamNum int // used for ordering params
}

func (_ *Parameter) Label() string {
	return "parameter"
}

func (v *Parameter) Properties() map[string]interface{} {
	return map[string]interface{}{
		"Name":     v.Name,
		"Type":     v.Type,
		"ParamNum": v.ParamNum,
	}
}

type Instruction struct {
	vertexBase
	Key        string `json:"_key"`
	Id         string `json:"_id"`
	File       string
	Offset     int
	Text       string
	SSAType    string
	ASTType    string
	Pos        string
	Misc       string
	Next       []*Instruction `json:"-"`
	References []*Variable    `json:"-"`
	Assigns    []*Variable    `json:"-"`
}

func (_ *Instruction) Label() string {
	return "instruction"
}

func (s *Instruction) Properties() map[string]interface{} {
	return map[string]interface{}{
		"File":    s.File,
		"Offset":  s.Offset,
		"Text":    s.Text,
		"ASTType": s.ASTType,
		"SSAType": s.SSAType,
		"Pos":     s.Pos,
		"Misc":    s.Misc,
	}
}

type Block struct {
	vertexBase
	Id   string `json:"_id"`
	Key  string `json:"_key"`
	Name string
	File string
	Pos  string
}

func (_ *Block) Label() string {
	return "block"
}

func (b *Block) Properties() map[string]interface{} {
	return map[string]interface{}{
		"File": b.File,
		"Pos":  b.Pos,
		"Name": b.Name,
	}
}

type FunctionCall struct {
	vertexBase
	Caller *Function   `json:"-"`
	Callee *Function   `json:"-"`
	Args   []*Variable `json:"-"`
	Return *Variable   `json:"-"`
}

func (_ *FunctionCall) Label() string {
	return "functioncall"
}
