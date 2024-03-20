package check

import (
	"work.ctyun.cn/git/GoStack/gostone/policy/entity"
)

var checkMap = make(map[entity.Operator]BaseCheck)

type BaseCheck interface {
	Check(values []entity.Values, context interface{}, target interface{}) bool
}

func init() {
	checkMap[entity.And] = NewAndCheck()
	checkMap[entity.Or] = NewOrCheck()
	checkMap[entity.Equal] = NewEqualCheck()
}
func GetCheck(operator entity.Operator) BaseCheck {
	return checkMap[operator]
}

func GetCheckVal(value entity.Values, context interface{}, target interface{}) entity.Check {
	return entity.Check{
		Key:     value.Key,
		Value:   value.Values,
		Context: context,
		Target:  target,
	}
}
