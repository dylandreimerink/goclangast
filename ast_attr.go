package goclangast

import "github.com/valyala/fastjson"

var AttrMap = map[string]func() Node{
	"AlignedAttr":              func() Node { return &AlignedAttr{} },
	"AlwaysInlineAttr":         func() Node { return &AlwaysInlineAttr{} },
	"SectionAttr":              func() Node { return &SectionAttr{} },
	"UsedAttr":                 func() Node { return &UsedAttr{} },
	"BuiltinAttr":              func() Node { return &BuiltinAttr{} },
	"NoThrowAttr":              func() Node { return &NoThrowAttr{} },
	"ConstAttr":                func() Node { return &ConstAttr{} },
	"UnusedAttr":               func() Node { return &UnusedAttr{} },
	"FormatAttr":               func() Node { return &FormatAttr{} },
	"FallThroughAttr":          func() Node { return &FallThroughAttr{} },
	"NoBuiltinAttr":            func() Node { return &NoBuiltinAttr{} },
	"PackedAttr":               func() Node { return &PackedAttr{} },
	"LoopHintAttr":             func() Node { return &LoopHintAttr{} },
	"WarnUnusedResultAttr":     func() Node { return &WarnUnusedResultAttr{} },
	"GNUInlineAttr":            func() Node { return &GNUInlineAttr{} },
	"NoInstrumentFunctionAttr": func() Node { return &NoInstrumentFunctionAttr{} },
	"AsmLabelAttr":             func() Node { return &AsmLabelAttr{} },
	"RestrictAttr":             func() Node { return &RestrictAttr{} },
	"AllocSizeAttr":            func() Node { return &AllocSizeAttr{} },
	"PureAttr":                 func() Node { return &PureAttr{} },
	"NonNullAttr":              func() Node { return &NonNullAttr{} },
	"AssumeAlignedAttr":        func() Node { return &AssumeAlignedAttr{} },
	"TransparentUnionAttr":     func() Node { return &TransparentUnionAttr{} },
	"WeakAttr":                 func() Node { return &WeakAttr{} },
	"NoInlineAttr":             func() Node { return &NoInlineAttr{} },
	"DeprecatedAttr":           func() Node { return &DeprecatedAttr{} },
}

type AlwaysInlineAttr struct {
	BaseNode
}

type SectionAttr struct {
	BaseNode
}

type UsedAttr struct {
	BaseNode
}

type UnusedAttr struct {
	BaseNode
}

type FormatAttr struct {
	BaseNode
}

type FallThroughAttr struct {
	BaseNode
}

type NoBuiltinAttr struct {
	BaseNode
}

type PackedAttr struct {
	BaseNode
}

type WarnUnusedResultAttr struct {
	BaseNode
}

type GNUInlineAttr struct {
	BaseNode
}

type NoInstrumentFunctionAttr struct {
	BaseNode
}

type AsmLabelAttr struct {
	BaseNode
}

type RestrictAttr struct {
	BaseNode
}

type AllocSizeAttr struct {
	BaseNode
}

type PureAttr struct {
	BaseNode
}

type NonNullAttr struct {
	BaseNode
}

type AssumeAlignedAttr struct {
	BaseNode
}

type TransparentUnionAttr struct {
	BaseNode
}

type WeakAttr struct {
	BaseNode
}

type NoInlineAttr struct {
	BaseNode
}

type AlignedAttr struct {
	BaseNode
}

type AttrImplicit struct {
	BaseNode
	Implicit bool `json:"implicit"`
}

func (a *AttrImplicit) Unmarshal(v *fastjson.Value, ctx *ParseContext) error {
	a.Implicit = v.GetBool("implicit")
	return a.BaseNode.Unmarshal(v, ctx)
}

type BuiltinAttr struct {
	AttrImplicit
}

type NoThrowAttr struct {
	AttrImplicit
}

type ConstAttr struct {
	AttrImplicit
}

type LoopHintAttr struct {
	AttrImplicit
}

type DeprecatedAttr struct {
	BaseNode
}
