package goclangast

import "github.com/valyala/fastjson"

var DeclMap = map[string]func() Node{
	"TranslationUnitDecl": func() Node { return &TranslationUnitDecl{} },
	"TypedefDecl":         func() Node { return &TypedefDecl{} },
	"EnumDecl":            func() Node { return &EnumDecl{} },
	"EnumConstantDecl":    func() Node { return &EnumConstantDecl{} },
	"RecordDecl":          func() Node { return &RecordDecl{} },
	"FieldDecl":           func() Node { return &FieldDecl{} },
	"FunctionDecl":        func() Node { return &FunctionDecl{} },
	"VarDecl":             func() Node { return &VarDecl{} },
	"ParmVarDecl":         func() Node { return &ParmVarDecl{} },
	"EmptyDecl":           func() Node { return &EmptyDecl{} },
	"StaticAssertDecl":    func() Node { return &StaticAssertDecl{} },
	"LabelDecl":           func() Node { return &LabelDecl{} },
	"IndirectFieldDecl":   func() Node { return &IndirectFieldDecl{} },
}

type Decl struct {
	ID   string `json:"id"`
	Kind string `json:"kind"`
	Name string `json:"name"`
}

func (d *Decl) Unmarshal(v *fastjson.Value, ctx *ParseContext) error {
	d.ID = string(v.GetStringBytes("id"))
	d.Kind = ctx.InternBytes(v.GetStringBytes("kind"))
	d.Name = string(v.GetStringBytes("name"))
	return nil
}

type TranslationUnitDecl struct {
	BaseNode
}

type TypedefDecl struct {
	BaseNode
	IsImplicit bool   `json:"isImplicit"`
	Name       string `json:"name"`
	Type       Type   `json:"type"`
}

func (d *TypedefDecl) Unmarshal(v *fastjson.Value, ctx *ParseContext) error {
	var err error
	d.IsImplicit = v.GetBool("isImplicit")
	d.Name = string(v.GetStringBytes("name"))
	d.Type, err = typeFromVal(v.Get("type"), ctx)
	if err != nil {
		return err
	}
	return d.BaseNode.Unmarshal(v, ctx)
}

type EnumDecl struct {
	BaseNode
}

type EnumConstantDecl struct {
	BaseNode
	Name string `json:"name"`
	Type Type   `json:"type"`
}

func (d *EnumConstantDecl) Unmarshal(v *fastjson.Value, ctx *ParseContext) error {
	var err error
	d.Name = string(v.GetStringBytes("name"))
	d.Type, err = typeFromVal(v.Get("type"), ctx)
	if err != nil {
		return err
	}
	return d.BaseNode.Unmarshal(v, ctx)
}

type RecordDecl struct {
	BaseNode
	Name               string `json:"name"`
	TagUsed            string `json:"tagUsed"`
	CompleteDefinition bool   `json:"completeDefinition"`
}

func (d *RecordDecl) Unmarshal(v *fastjson.Value, ctx *ParseContext) error {
	d.Name = string(v.GetStringBytes("name"))
	d.TagUsed = string(v.GetStringBytes("tagUsed"))
	d.CompleteDefinition = v.GetBool("completeDefinition")
	return d.BaseNode.Unmarshal(v, ctx)
}

type FieldDecl struct {
	BaseNode
	Name       string `json:"name"`
	Type       Type   `json:"type"`
	IsBitfield bool   `json:"isBitfield"`
}

func (d *FieldDecl) Unmarshal(v *fastjson.Value, ctx *ParseContext) error {
	var err error
	d.Name = string(v.GetStringBytes("name"))
	d.IsBitfield = v.GetBool("isBitfield")
	d.Type, err = typeFromVal(v.Get("type"), ctx)
	if err != nil {
		return err
	}
	return d.BaseNode.Unmarshal(v, ctx)
}

type ReferencedDecl struct {
	ID   string `json:"id"`
	Kind string `json:"kind"`
	Name string `json:"name"`
	Type Type   `json:"type"`
}

func (d *ReferencedDecl) Unmarshal(v *fastjson.Value, ctx *ParseContext) error {
	var err error
	d.ID = string(v.GetStringBytes("id"))
	d.Name = string(v.GetStringBytes("name"))
	d.Type, err = typeFromVal(v.Get("type"), ctx)
	return err
}

type IndirectFieldDecl struct {
	BaseNode
	Name       string `json:"name"`
	IsImplicit bool   `json:"isImplicit"`
}

func (d *IndirectFieldDecl) Unmarshal(v *fastjson.Value, ctx *ParseContext) error {
	d.Name = string(v.GetStringBytes("name"))
	d.IsImplicit = v.GetBool("isImplicit")
	return d.BaseNode.Unmarshal(v, ctx)
}

type FunctionDecl struct {
	BaseNode
	IsUsed       bool   `json:"isUsed"`
	Name         string `json:"name"`
	MangledName  string `json:"mangledName"`
	Type         Type   `json:"type"`
	StorageClass string `json:"storageClass"`
	Inline       bool   `json:"inline"`
}

func (d *FunctionDecl) Unmarshal(v *fastjson.Value, ctx *ParseContext) error {
	var err error
	d.Name = string(v.GetStringBytes("name"))
	d.IsUsed = v.GetBool("isUsed")
	d.MangledName = string(v.GetStringBytes("mangledName"))
	d.Type, err = typeFromVal(v.Get("type"), ctx)
	if err != nil {
		return err
	}
	d.StorageClass = string(v.GetStringBytes("storageClass"))
	d.Inline = v.GetBool("inline")
	return d.BaseNode.Unmarshal(v, ctx)
}

type VarDecl struct {
	BaseNode
	IsUsed       bool   `json:"isUsed"`
	Name         string `json:"name"`
	MangledName  string `json:"mangledName"`
	Type         Type   `json:"type"`
	StorageClass string `json:"storageClass"`
	Init         string `json:"init"`
}

func (d *VarDecl) Unmarshal(v *fastjson.Value, ctx *ParseContext) error {
	var err error
	d.Name = string(v.GetStringBytes("name"))
	d.IsUsed = v.GetBool("isUsed")
	d.MangledName = string(v.GetStringBytes("mangledName"))
	d.Type, err = typeFromVal(v.Get("type"), ctx)
	if err != nil {
		return err
	}
	d.StorageClass = string(v.GetStringBytes("storageClass"))
	d.Init = string(v.GetStringBytes("init"))
	return d.BaseNode.Unmarshal(v, ctx)
}

type ParmVarDecl struct {
	BaseNode
	Type        Type   `json:"type"`
	IsUsed      bool   `json:"isUsed"`
	Name        string `json:"name"`
	MangledName string `json:"mangledName"`
}

func (d *ParmVarDecl) Unmarshal(v *fastjson.Value, ctx *ParseContext) error {
	var err error
	d.Name = string(v.GetStringBytes("name"))
	d.IsUsed = v.GetBool("isUsed")
	d.MangledName = string(v.GetStringBytes("mangledName"))
	d.Type, err = typeFromVal(v.Get("type"), ctx)
	if err != nil {
		return err
	}
	return d.BaseNode.Unmarshal(v, ctx)
}

type EmptyDecl struct {
	BaseNode
}

type StaticAssertDecl struct {
	BaseNode
}

type LabelDecl struct {
	BaseNode
	IsUsed bool   `json:"isUsed"`
	Name   string `json:"name"`
}

func (d *LabelDecl) Unmarshal(v *fastjson.Value, ctx *ParseContext) error {
	d.Name = string(v.GetStringBytes("name"))
	d.IsUsed = v.GetBool("isUsed")
	return d.BaseNode.Unmarshal(v, ctx)
}
