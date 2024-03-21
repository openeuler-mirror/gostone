package check

import (
	"work.ctyun.cn/git/GoStack/gostone/policy/entity"
)

type AndCheck struct {
}

func (e *AndCheck) Check(values []entity.Values, context interface{}, target interface{}) bool {
	for _, v := range values {
		ch := GetValCheck(v.Operator)
		if !ch.Check(GetCheckVal(v, context, target)) {
			return false
		}
	}
	return true
}

func NewAndCheck() BaseCheck {
	return &AndCheck{}
}
