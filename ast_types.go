package goclangast

import "github.com/valyala/fastjson"

var TypeMap = map[string]func() Node{
	"BuiltinType":       func() Node { return &BuiltinType{} },
	"RecordType":        func() Node { return &RecordType{} },
	"PointerType":       func() Node { return &PointerType{} },
	"ConstantArrayType": func() Node { return &ConstantArrayType{} },
	"TypedefType":       func() Node { return &TypedefType{} },
	"ElaboratedType":    func() Node { return &ElaboratedType{} },
	"ParenType":         func() Node { return &ParenType{} },
	"FunctionProtoType": func() Node { return &FunctionProtoType{} },
	"QualType":          func() Node { return &QualType{} },
	"EnumType":          func() Node { return &EnumType{} },
}

type Type struct {
	DesugaredQualType string `json:"desugaredQualType"`
	QualType          string `json:"qualType"`
	TypeAliasDeclId   string `json:"typeAliasDeclId"`
}

func typeFromVal(v *fastjson.Value, ctx *ParseContext) (Type, error) {
	t := Type{}
	t.DesugaredQualType = string(v.GetStringBytes("desugaredQualType"))
	t.QualType = string(v.GetStringBytes("qualType"))
	t.TypeAliasDeclId = string(v.GetStringBytes("typeAliasDeclId"))
	return t, nil
}

type BaseType struct {
	BaseNode
	Type Type `json:"type"`
}

func (t *BaseType) Unmarshal(v *fastjson.Value, ctx *ParseContext) error {
	var err error
	t.Type, err = typeFromVal(v.Get("type"), ctx)
	if err != nil {
		return err
	}
	return t.BaseNode.Unmarshal(v, ctx)
}

type BuiltinType struct {
	BaseType
}

type RecordType struct {
	BaseType
	Decl Decl `json:"decl"`
}

func (t *RecordType) Unmarshal(v *fastjson.Value, ctx *ParseContext) error {
	err := t.Decl.Unmarshal(v.Get("decl"), ctx)
	if err != nil {
		return err
	}
	return t.BaseType.Unmarshal(v, ctx)
}

type PointerType struct {
	BaseType
}

type ConstantArrayType struct {
	BaseType
	Size int `json:"size"`
}

func (t *ConstantArrayType) Unmarshal(v *fastjson.Value, ctx *ParseContext) error {
	t.Size = v.GetInt("size")
	return t.BaseType.Unmarshal(v, ctx)
}

type TypedefType struct {
	BaseType
	Decl Decl `json:"decl"`
}

func (t *TypedefType) Unmarshal(v *fastjson.Value, ctx *ParseContext) error {
	err := t.Decl.Unmarshal(v.Get("decl"), ctx)
	if err != nil {
		return err
	}
	return t.BaseType.Unmarshal(v, ctx)
}

type ElaboratedType struct {
	BaseType
	OwnedTagDecl Decl `json:"ownedTagDecl"`
}

func (t *ElaboratedType) Unmarshal(v *fastjson.Value, ctx *ParseContext) error {
	err := t.OwnedTagDecl.Unmarshal(v.Get("ownedTagDecl"), ctx)
	if err != nil {
		return err
	}
	return t.BaseType.Unmarshal(v, ctx)
}

type ParenType struct {
	BaseType
}

type FunctionProtoType struct {
	BaseType
	CC string `json:"cc"`
}

func (t *FunctionProtoType) Unmarshal(v *fastjson.Value, ctx *ParseContext) error {
	t.CC = string(v.GetStringBytes("cc"))
	return t.BaseType.Unmarshal(v, ctx)
}

type QualType struct {
	BaseNode
}

type EnumType struct {
	BaseType
	Decl Decl `json:"decl"`
}

func (t *EnumType) Unmarshal(v *fastjson.Value, ctx *ParseContext) error {
	err := t.Decl.Unmarshal(v.Get("decl"), ctx)
	if err != nil {
		return err
	}
	return t.BaseType.Unmarshal(v, ctx)
}
