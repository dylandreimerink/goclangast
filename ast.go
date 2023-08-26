package goclangast

import (
	"bytes"
	"fmt"

	"github.com/philpearl/intern"
	"github.com/valyala/fastjson"
)

type ParseContext struct {
	strIntern *intern.Intern
}

func (p *ParseContext) InternBytes(b []byte) string {
	return p.strIntern.Deduplicate(string(b))
}

type Node interface {
	GetBaseNode() *BaseNode
	Children() []Node
	Unmarshal(v *fastjson.Value, p *ParseContext) error
}

func parse(b *bytes.Buffer) (*TranslationUnitDecl, error) {
	var p fastjson.Parser
	ctx := ParseContext{
		strIntern: intern.New(8 * 1024),
	}
	v, err := p.ParseBytes(b.Bytes())
	if err != nil {
		return nil, err
	}

	node, err := parseNode(v, &ctx)
	if err != nil {
		return nil, err
	}

	return node.(*TranslationUnitDecl), nil
}

var kindMap map[string]func() Node

func getKindMap() map[string]func() Node {
	if kindMap != nil {
		return kindMap
	}

	kindMap = make(map[string]func() Node)
	maps := []map[string]func() Node{
		AttrMap,
		CommentMap,
		DeclMap,
		ExprMap,
		LiteralMap,
		OperatorMap,
		StmtMap,
		TypeMap,
	}
	for _, m := range maps {
		for k, v := range m {
			kindMap[k] = v
		}
	}

	return kindMap
}

func parseNode(v *fastjson.Value, ctx *ParseContext) (Node, error) {
	kind := v.GetStringBytes("kind")

	if len(kind) == 0 {
		return nil, nil
	}

	nodeFn, found := getKindMap()[string(kind)]
	if !found {
		return nil, fmt.Errorf("unknown node kind: %s", string(kind))
	}

	node := nodeFn()
	err := node.Unmarshal(v, ctx)

	return node, err
}

type RawNode struct {
	Kind string `json:"kind"`
}

type Loc struct {
	Offset              int    `json:"offset"`
	File                string `json:"file"`
	Line                int    `json:"line"`
	PresumedFile        string `json:"presumedFile"`
	Col                 int    `json:"col"`
	TokLen              int    `json:"tokLen"`
	IsMacroArgExpansion bool   `json:"isMacroArgExpansion"`
	IncludedFromFile    string `json:"includedFrom"`
	SpellingLoc         *Loc   `json:"spellingLoc"`
	ExpansionLoc        *Loc   `json:"expansionLoc"`
}

func locFromVal(v *fastjson.Value, ctx *ParseContext) (*Loc, error) {
	if v == nil {
		return nil, nil
	}

	obj, err := v.Object()
	if err != nil {
		return nil, err
	}

	if obj.Len() == 0 {
		return nil, nil
	}

	var loc Loc
	loc.Offset = v.GetInt("offset")
	loc.File = ctx.InternBytes(v.GetStringBytes("file"))
	loc.Line = v.GetInt("line")
	loc.PresumedFile = ctx.InternBytes(v.GetStringBytes("presumedFile"))
	loc.Col = v.GetInt("col")
	loc.TokLen = v.GetInt("tokLen")
	loc.IsMacroArgExpansion = v.GetBool("isMacroArgExpansion")

	if includedFrom := v.Get("includedFrom"); includedFrom != nil {
		loc.IncludedFromFile = ctx.InternBytes(includedFrom.GetStringBytes("file"))
	}

	if expansionLoc := v.Get("expansionLoc"); expansionLoc != nil {
		loc.ExpansionLoc, err = locFromVal(expansionLoc, ctx)
		if err != nil {
			return nil, err
		}
	}

	if spellingLoc := v.Get("spellingLoc"); spellingLoc != nil {
		loc.SpellingLoc, err = locFromVal(spellingLoc, ctx)
		if err != nil {
			return nil, err
		}
	}

	return &loc, nil
}

type Range struct {
	Begin *Loc `json:"begin"`
	End   *Loc `json:"end"`
}

func rangeFromVal(v *fastjson.Value, ctx *ParseContext) (*Range, error) {
	if v == nil {
		return nil, nil
	}

	obj, err := v.Object()
	if err != nil {
		return nil, err
	}

	begin := obj.Get("begin")
	end := obj.Get("end")

	if begin == nil || end == nil {
		return nil, nil
	}

	beginLoc, err := locFromVal(begin, ctx)
	if err != nil {
		return nil, err
	}

	endLoc, err := locFromVal(end, ctx)
	if err != nil {
		return nil, err
	}

	if beginLoc == nil || endLoc == nil {
		return nil, nil
	}

	return &Range{
		Begin: beginLoc,
		End:   endLoc,
	}, nil
}

type BaseNode struct {
	ID    string `json:"id"`
	Kind  string `json:"kind"`
	Loc   *Loc   `json:"loc"`
	Range *Range `json:"range"`
	Inner []Node `json:"inner"`
}

func (bn *BaseNode) GetBaseNode() *BaseNode {
	return bn
}

func (bn *BaseNode) Unmarshal(v *fastjson.Value, ctx *ParseContext) error {
	var err error
	bn.ID = string(v.GetStringBytes("id"))
	bn.Kind = ctx.InternBytes(v.GetStringBytes("kind"))

	if loc := v.Get("loc"); loc != nil {
		bn.Loc, err = locFromVal(loc, ctx)
		if err != nil {
			return err
		}
	}

	if rangeVal := v.Get("range"); rangeVal != nil {
		bn.Range, err = rangeFromVal(rangeVal, ctx)
		if err != nil {
			return err
		}
	}

	inner := v.GetArray("inner")
	for _, v := range inner {
		child, err := parseNode(v, ctx)
		if err != nil {
			return err
		}

		if child == nil {
			continue
		}

		bn.Inner = append(bn.Inner, child)
	}

	return nil
}

func (bn *BaseNode) Children() []Node {
	return bn.Inner
}

func PreOrderVisit(node Node, fn func(n Node, depth int) error) {
	preOrderVisit(node, 0, fn)
}

func preOrderVisit(node Node, depth int, fn func(n Node, depth int) error) {
	if fn(node, depth) != nil {
		return
	}

	for _, child := range node.Children() {
		if child == nil {
			continue
		}

		preOrderVisit(child, depth+1, fn)
	}
}
