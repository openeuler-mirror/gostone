package check

import (
	log "github.com/sirupsen/logrus"
	conf2 "work.ctyun.cn/git/GoStack/gostone/conf"
	"work.ctyun.cn/git/GoStack/gostone/policy/entity"
)

type RuleCheck struct {
}

func NewRuleCheck() ValBaseCheck {
	return &RuleCheck{}
}

func (e *RuleCheck) Check(check entity.Check) bool {
	if len(check.Value) != 1 {
		log.Error("equal check size error")
		return false
	}
	rule, ok := conf2.GetRule(check.Value[0])
	if !ok {
		return true
	}
	ruleCheck := GetCheck(rule.Operator)
	return ruleCheck.Check(rule.Values, check.Context, check.Target)
}
