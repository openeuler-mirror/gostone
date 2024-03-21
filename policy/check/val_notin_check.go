package check

import (
	log "github.com/sirupsen/logrus"
	"work.ctyun.cn/git/GoStack/gostone/policy/entity"
)

type NotInCheck struct {
}

func NewNotInCheck() ValBaseCheck {
	return &NotInCheck{}
}

func (e *NotInCheck) Check(check entity.Check) bool {
	c, _ := GetCheckMap(check)
	val, ok := c[check.Key]
	if !ok {
		log.Info("target can not find:" + check.Key)
		return false
	}
	switch v := val.(type) {
	case []string:
		for _, str := range v {
			if !e.checkOne(str, check) {
				return false
			}
		}
		return true
	case string:
		return e.checkOne(v, check)
	default:
		log.Error("can not check:", v)
		return false
	}
}

func (e *NotInCheck) checkOne(v string, check entity.Check) bool {
	return isNotIn(v, check.Value)
}

func isNotIn(target string, source []string) bool {
	for _, s := range source {
		if target == s {
			return false
		}
	}
	return true
}
