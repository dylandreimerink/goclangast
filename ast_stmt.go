package goclangast

import "github.com/valyala/fastjson"

var StmtMap = map[string]func() Node{
	"CompoundStmt":   func() Node { return &CompoundStmt{} },
	"GCCAsmStmt":     func() Node { return &GCCAsmStmt{} },
	"ReturnStmt":     func() Node { return &ReturnStmt{} },
	"DeclStmt":       func() Node { return &DeclStmt{} },
	"IfStmt":         func() Node { return &IfStmt{} },
	"SwitchStmt":     func() Node { return &SwitchStmt{} },
	"CaseStmt":       func() Node { return &CaseStmt{} },
	"AttributedStmt": func() Node { return &AttributedStmt{} },
	"NullStmt":       func() Node { return &NullStmt{} },
	"BreakStmt":      func() Node { return &BreakStmt{} },
	"GotoStmt":       func() Node { return &GotoStmt{} },
	"DefaultStmt":    func() Node { return &DefaultStmt{} },
	"DoStmt":         func() Node { return &DoStmt{} },
	"ForStmt":        func() Node { return &ForStmt{} },
	"LabelStmt":      func() Node { return &LabelStmt{} },
	"ContinueStmt":   func() Node { return &ContinueStmt{} },
	"WhileStmt":      func() Node { return &WhileStmt{} },
}

type CompoundStmt struct {
	BaseNode
}

type GCCAsmStmt struct {
	BaseNode
}

type ReturnStmt struct {
	BaseNode
}

type DeclStmt struct {
	BaseNode
}

type IfStmt struct {
	BaseNode
}

type SwitchStmt struct {
	BaseNode
}

type CaseStmt struct {
	BaseNode
}

type AttributedStmt struct {
	BaseNode
}

type NullStmt struct {
	BaseNode
}

type BreakStmt struct {
	BaseNode
}

type GotoStmt struct {
	BaseNode
	TargetLabelDeclId string `json:"targetLabelDeclId"`
}

func (s *GotoStmt) Unmarshal(v *fastjson.Value, ctx *ParseContext) error {
	s.TargetLabelDeclId = string(v.GetStringBytes("targetLabelDeclId"))
	return s.BaseNode.Unmarshal(v, ctx)
}

type DefaultStmt struct {
	BaseNode
}

type DoStmt struct {
	BaseNode
}

type ForStmt struct {
	BaseNode
}

type LabelStmt struct {
	BaseNode
	Name   string `json:"name"`
	DeclId string `json:"declId"`
}

func (s *LabelStmt) Unmarshal(v *fastjson.Value, ctx *ParseContext) error {
	s.Name = string(v.GetStringBytes("name"))
	s.DeclId = string(v.GetStringBytes("declId"))
	return s.BaseNode.Unmarshal(v, ctx)
}

type ContinueStmt struct {
	BaseNode
}

type WhileStmt struct {
	BaseNode
}
