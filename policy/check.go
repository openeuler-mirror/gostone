package policy

import (
	"work.ctyun.cn/git/GoStack/gostone/conf"
	"work.ctyun.cn/git/GoStack/gostone/policy/check"
)

func Check(name string, context interface{}, target interface{}) bool {
	policy, ok := conf.GetPolicy(name)
	if !ok {
		return true
	}
	ck := check.GetCheck(policy.Operator)
	return ck.Check(policy.Values, context, target)
}
