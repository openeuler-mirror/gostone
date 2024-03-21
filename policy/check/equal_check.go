package check

import (
	log "github.com/sirupsen/logrus"
	"work.ctyun.cn/git/GoStack/gostone/policy/entity"
)

type EqualCheck struct {
}

func (e *EqualCheck) Check(values []entity.Values, context interface{}, target interface{}) bool {
	if len(values) != 1 {
		log.Error("equal check size error")
		return false
	}
	c := GetValCheck(values[0].Operator)
	return c.Check(GetCheckVal(values[0], context, target))
}

func NewEqualCheck() BaseCheck {
	return &EqualCheck{}
}
