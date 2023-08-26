package goclangast

import "github.com/valyala/fastjson"

var OperatorMap = map[string]func() Node{
	"BinaryOperator":            func() Node { return &BinaryOperator{} },
	"UnaryOperator":             func() Node { return &UnaryOperator{} },
	"ConditionalOperator":       func() Node { return &ConditionalOperator{} },
	"CompoundAssignOperator":    func() Node { return &CompoundAssignOperator{} },
	"BinaryConditionalOperator": func() Node { return &BinaryConditionalOperator{} },
}

type Operator struct {
	BaseNode
	Type          Type   `json:"type"`
	ValueCategory string `json:"valueCategory"`
}

func (o *Operator) Unmarshal(v *fastjson.Value, ctx *ParseContext) error {
	var err error
	o.Type, err = typeFromVal(v.Get("type"), ctx)
	if err != nil {
		return err
	}

	o.ValueCategory = string(v.GetStringBytes("valueCategory"))
	return o.BaseNode.Unmarshal(v, ctx)
}

type BinaryOperator struct {
	Operator
	Opcode string `json:"opcode"`
}

func (o *BinaryOperator) Unmarshal(v *fastjson.Value, ctx *ParseContext) error {
	o.Opcode = string(v.GetStringBytes("opcode"))
	return o.Operator.Unmarshal(v, ctx)
}

type UnaryOperator struct {
	Operator
	IsPostfix bool   `json:"isPostfix"`
	Opcode    string `json:"opcode"`
}

func (o *UnaryOperator) Unmarshal(v *fastjson.Value, ctx *ParseContext) error {
	o.IsPostfix = v.GetBool("isPostfix")
	o.Opcode = string(v.GetStringBytes("opcode"))
	return o.Operator.Unmarshal(v, ctx)
}

type ConditionalOperator struct {
	Operator
}

type CompoundAssignOperator struct {
	Operator
	Opcode            string `json:"opcode"`
	ComputeLHSType    string `json:"computeLHSType"`
	ComputeResultType string `json:"computeResultType"`
}

func (o *CompoundAssignOperator) Unmarshal(v *fastjson.Value, ctx *ParseContext) error {
	o.Opcode = string(v.GetStringBytes("opcode"))
	o.ComputeLHSType = string(v.GetStringBytes("computeLHSType"))
	o.ComputeResultType = string(v.GetStringBytes("computeResultType"))
	return o.Operator.Unmarshal(v, ctx)
}

type BinaryConditionalOperator struct {
	Operator
}
