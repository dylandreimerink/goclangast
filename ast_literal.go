package goclangast

import "github.com/valyala/fastjson"

var LiteralMap = map[string]func() Node{
	"IntegerLiteral":   func() Node { return &IntegerLiteral{} },
	"StringLiteral":    func() Node { return &StringLiteral{} },
	"CharacterLiteral": func() Node { return &CharacterLiteral{} },
}

type IntegerLiteral struct {
	BaseNode
	Type          Type   `json:"type"`
	ValueCategory string `json:"valueCategory"`
	Value         string `json:"value"`
}

func (l *IntegerLiteral) Unmarshal(v *fastjson.Value, ctx *ParseContext) error {
	var err error
	l.Type, err = typeFromVal(v.Get("type"), ctx)
	if err != nil {
		return err
	}

	l.ValueCategory = string(v.GetStringBytes("valueCategory"))
	l.Value = string(v.GetStringBytes("value"))
	return l.BaseNode.Unmarshal(v, ctx)
}

type StringLiteral struct {
	BaseNode
	Type          Type   `json:"type"`
	ValueCategory string `json:"valueCategory"`
	Value         string `json:"value"`
}

func (l *StringLiteral) Unmarshal(v *fastjson.Value, ctx *ParseContext) error {
	var err error
	l.Type, err = typeFromVal(v.Get("type"), ctx)
	if err != nil {
		return err
	}

	l.ValueCategory = string(v.GetStringBytes("valueCategory"))
	l.Value = string(v.GetStringBytes("value"))
	return l.BaseNode.Unmarshal(v, ctx)
}

type CharacterLiteral struct {
	BaseNode
	Type          Type   `json:"type"`
	ValueCategory string `json:"valueCategory"`
	Value         int    `json:"value"`
}

func (l *CharacterLiteral) Unmarshal(v *fastjson.Value, ctx *ParseContext) error {
	var err error
	l.Type, err = typeFromVal(v.Get("type"), ctx)
	if err != nil {
		return err
	}

	l.ValueCategory = string(v.GetStringBytes("valueCategory"))
	l.Value = v.GetInt("value")
	return l.BaseNode.Unmarshal(v, ctx)
}
