package entity

type Operator string

const (
	Equal    Operator = "equal"
	In       Operator = "in"
	NotEqual Operator = "not"
	NotIn    Operator = "notin"
	And      Operator = "and"
	Or       Operator = "or"
	RuleOp   Operator = "rule"
)
