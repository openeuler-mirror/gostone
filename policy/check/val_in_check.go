package check

import (
	log "github.com/sirupsen/logrus"
	"work.ctyun.cn/git/GoStack/gostone/policy/entity"
)

type InCheck struct {
}

func NewInCheck() ValBaseCheck {
	return &InCheck{}
}

func (e *InCheck) Check(check entity.Check) bool {
	c, _ := GetCheckMap(check)
	val, ok := c[check.Key]
	if !ok {
		log.Info("target can not find:" + check.Key)
		return false
	}
	switch v := val.(type) {
	case []string:
		for _, str := range v {
			if e.checkOne(str, check) {
				return true
			}
		}
		return false
	case string:
		return e.checkOne(v, check)
	default:
		log.Error("can not check:", v)
		return false
	}
}

func (e *InCheck) checkOne(v string, check entity.Check) bool {
	return isIn(v, check.Value)
}

func isIn(target string, source []string) bool {
	for _, s := range source {
		if target == s {
			return true
		}
	}
	return false
}
