package check

import (
	"work.ctyun.cn/git/GoStack/gostone/policy/entity"
)

type OrCheck struct {
}

func (e *OrCheck) Check(values []entity.Values, context interface{}, target interface{}) bool {
	for _, v := range values {
		ch := GetValCheck(v.Operator)
		if ch.Check(GetCheckVal(v, context, target)) {
			return true
		}
	}
	return false
}

func NewOrCheck() BaseCheck {
	return &OrCheck{}
}
