package model

import (
	"github.com/Knetic/govaluate"
)

type Condition struct {
	expression string
}

func NewCondition(expression string) (*Condition, error) {
	if err := validate(expression); err != nil {
		return nil, err
	}
	return &Condition{
		expression: expression,
	}, nil
}

func (c Condition) Expression() string {
	return c.expression
}
func (c Condition) IsEmpty() bool {
	return c.expression == ""
}
func (c Condition) Eval(sub, obj, env []Attribute) bool {
	if c.IsEmpty() {
		return true
	}

	goeExpr, err := govaluate.NewEvaluableExpression(c.expression)
	if err != nil {
		return false
	}

	parameters := make(map[string]interface{}, 8)
	for _, attr := range sub {
		parameters[SubVarNamePrefix+attr.Name()] = attr.Value()
	}
	for _, attr := range obj {
		parameters[ObjVarNamePrefix+attr.Name()] = attr.Value()
	}
	for _, attr := range env {
		parameters[EnvVarNamePrefix+attr.Name()] = attr.Value()
	}

	result, err := goeExpr.Evaluate(parameters)
	if err != nil {
		return false
	}
	boolResult, ok := result.(bool)
	if !ok {
		return false
	}
	return boolResult
}
