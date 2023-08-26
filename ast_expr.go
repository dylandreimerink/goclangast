package goclangast

import "github.com/valyala/fastjson"

var ExprMap = map[string]func() Node{
	"ConstantExpr":             func() Node { return &ConstantExpr{} },
	"DeclRefExpr":              func() Node { return &DeclRefExpr{} },
	"ImplicitCastExpr":         func() Node { return &ImplicitCastExpr{} },
	"ParenExpr":                func() Node { return &ParenExpr{} },
	"CStyleCastExpr":           func() Node { return &CStyleCastExpr{} },
	"CallExpr":                 func() Node { return &CallExpr{} },
	"MemberExpr":               func() Node { return &MemberExpr{} },
	"ArraySubscriptExpr":       func() Node { return &ArraySubscriptExpr{} },
	"UnaryExprOrTypeTraitExpr": func() Node { return &UnaryExprOrTypeTraitExpr{} },
	"StmtExpr":                 func() Node { return &StmtExpr{} },
	"InitListExpr":             func() Node { return &InitListExpr{} },
	"ImplicitValueInitExpr":    func() Node { return &ImplicitValueInitExpr{} },
	"OffsetOfExpr":             func() Node { return &OffsetOfExpr{} },
	"CompoundLiteralExpr":      func() Node { return &CompoundLiteralExpr{} },
	"AddrLabelExpr":            func() Node { return &AddrLabelExpr{} },
	"GenericSelectionExpr":     func() Node { return &GenericSelectionExpr{} },
	"PredefinedExpr":           func() Node { return &PredefinedExpr{} },
	"ChooseExpr":               func() Node { return &ChooseExpr{} },
	"TypeTraitExpr":            func() Node { return &TypeTraitExpr{} },
	"RecoveryExpr":             func() Node { return &RecoveryExpr{} },
	"OpaqueValueExpr":          func() Node { return &OpaqueValueExpr{} },
}

type Expr struct {
	BaseNode
	Type          Type   `json:"type"`
	ValueCategory string `json:"valueCategory"`
}

func (e *Expr) Unmarshal(v *fastjson.Value, ctx *ParseContext) error {
	var err error
	e.ValueCategory = ctx.InternBytes(v.GetStringBytes("valueCategory"))
	e.Type, err = typeFromVal(v.Get("type"), ctx)
	if err != nil {
		return err
	}
	return e.BaseNode.Unmarshal(v, ctx)
}

type ConstantExpr struct {
	Expr
	Value string `json:"value"`
}

func (d *ConstantExpr) Unmarshal(v *fastjson.Value, ctx *ParseContext) error {
	d.Value = string(v.GetStringBytes("value"))
	return d.Expr.Unmarshal(v, ctx)
}

type DeclRefExpr struct {
	Expr
	ReferencedDecl  ReferencedDecl `json:"referencedDecl"`
	NonOdrUseReason string         `json:"nonOdrUseReason"`
}

func (d *DeclRefExpr) Unmarshal(v *fastjson.Value, ctx *ParseContext) error {
	err := d.ReferencedDecl.Unmarshal(v.Get("referencedDecl"), ctx)
	if err != nil {
		return err
	}

	d.NonOdrUseReason = string(v.GetStringBytes("value"))
	return d.Expr.Unmarshal(v, ctx)
}

type ImplicitCastExpr struct {
	Expr
	CastKind string `json:"castKind"`
}

func (d *ImplicitCastExpr) Unmarshal(v *fastjson.Value, ctx *ParseContext) error {
	d.CastKind = string(v.GetStringBytes("castKind"))
	return d.Expr.Unmarshal(v, ctx)
}

type ParenExpr struct {
	Expr
}

type CStyleCastExpr struct {
	Expr
	CastKind string `json:"castKind"`
}

func (d *CStyleCastExpr) Unmarshal(v *fastjson.Value, ctx *ParseContext) error {
	d.CastKind = string(v.GetStringBytes("castKind"))
	return d.Expr.Unmarshal(v, ctx)
}

type CallExpr struct {
	Expr
}

type MemberExpr struct {
	Expr
	IsArrow              bool   `json:"isArrow"`
	ReferencedMemberDecl string `json:"referencedMemberDecl"`
}

func (d *MemberExpr) Unmarshal(v *fastjson.Value, ctx *ParseContext) error {
	d.IsArrow = v.GetBool("isArrow")
	d.ReferencedMemberDecl = string(v.GetStringBytes("referencedMemberDecl"))
	return d.Expr.Unmarshal(v, ctx)
}

type ArraySubscriptExpr struct {
	Expr
}

type UnaryExprOrTypeTraitExpr struct {
	Expr
	Name    string `json:"name"`
	ArgType Type   `json:"argType"`
}

func (d *UnaryExprOrTypeTraitExpr) Unmarshal(v *fastjson.Value, ctx *ParseContext) error {
	var err error
	d.ArgType, err = typeFromVal(v.Get("argType"), ctx)
	if err != nil {
		return err
	}
	d.Name = string(v.GetStringBytes("name"))
	return d.Expr.Unmarshal(v, ctx)
}

type StmtExpr struct {
	Expr
}

type InitListExpr struct {
	Expr
}

type ImplicitValueInitExpr struct {
	Expr
}

type OffsetOfExpr struct {
	Expr
}

type CompoundLiteralExpr struct {
	Expr
}

type OpaqueValueExpr struct {
	Expr
}

type ChooseExpr struct {
	Expr
}

type TypeTraitExpr struct {
	BaseNode
}

type RecoveryExpr struct {
	BaseNode
}

type AddrLabelExpr struct {
	Expr
	LabelDeclId string `json:"labelDeclId"`
	Name        string `json:"name"`
}

func (d *AddrLabelExpr) Unmarshal(v *fastjson.Value, ctx *ParseContext) error {
	d.LabelDeclId = string(v.GetStringBytes("labelDeclId"))
	d.Name = string(v.GetStringBytes("name"))
	return d.Expr.Unmarshal(v, ctx)
}

type GenericSelectionExpr struct {
	Expr
}

type PredefinedExpr struct {
	BaseNode
}
