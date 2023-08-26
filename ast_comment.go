package goclangast

import "github.com/valyala/fastjson"

var CommentMap = map[string]func() Node{
	"FullComment":              func() Node { return &FullComment{} },
	"ParagraphComment":         func() Node { return &ParagraphComment{} },
	"TextComment":              func() Node { return &TextComment{} },
	"InlineCommandComment":     func() Node { return &InlineCommandComment{} },
	"BlockCommandComment":      func() Node { return &BlockCommandComment{} },
	"VerbatimBlockComment":     func() Node { return &VerbatimBlockComment{} },
	"VerbatimBlockLineComment": func() Node { return &VerbatimBlockLineComment{} },
	"VerbatimLineComment":      func() Node { return &VerbatimLineComment{} },
}

type FullComment struct {
	BaseNode
}

type ParagraphComment struct {
	BaseNode
}

type InlineCommandComment struct {
	BaseNode
	Name       string   `json:"name"`
	RenderKind string   `json:"renderKind"`
	Args       []string `json:"args"`
}

func (c *InlineCommandComment) Unmarshal(v *fastjson.Value, ctx *ParseContext) error {
	c.Name = string(v.GetStringBytes("name"))
	c.RenderKind = string(v.GetStringBytes("renderKind"))
	c.Args = make([]string, 0)
	for _, arg := range v.GetArray("args") {
		c.Args = append(c.Args, string(arg.GetStringBytes()))
	}
	return c.BaseNode.Unmarshal(v, ctx)
}

type BlockCommandComment struct {
	BaseNode
	Name string `json:"name"`
}

type VerbatimBlockComment struct {
	BaseNode
	Name      string `json:"name"`
	CloseName string `json:"closeName"`
}

func (c *VerbatimBlockComment) Unmarshal(v *fastjson.Value, ctx *ParseContext) error {
	c.Name = string(v.GetStringBytes("name"))
	c.CloseName = string(v.GetStringBytes("closeName"))
	return c.BaseNode.Unmarshal(v, ctx)
}

type TextComment struct {
	BaseNode
	Text string `json:"text"`
}

func (c *TextComment) Unmarshal(v *fastjson.Value, ctx *ParseContext) error {
	c.Text = string(v.GetStringBytes("text"))
	return c.BaseNode.Unmarshal(v, ctx)
}

type VerbatimBlockLineComment struct {
	BaseNode
	Text string `json:"text"`
}

func (c *VerbatimBlockLineComment) Unmarshal(v *fastjson.Value, ctx *ParseContext) error {
	c.Text = string(v.GetStringBytes("text"))
	return c.BaseNode.Unmarshal(v, ctx)
}

type VerbatimLineComment struct {
	BaseNode
	Text string `json:"text"`
}

func (c *VerbatimLineComment) Unmarshal(v *fastjson.Value, ctx *ParseContext) error {
	c.Text = string(v.GetStringBytes("text"))
	return c.BaseNode.Unmarshal(v, ctx)
}
