package check

import (
	log "github.com/sirupsen/logrus"
	"regexp"
	"work.ctyun.cn/git/GoStack/gostone/policy/entity"
)

const valReg = "^%(.+)"

var valRegx = regexp.MustCompile(valReg)

type ValEqualCheck struct {
}

func NewValEqualCheck() ValBaseCheck {
	return &ValEqualCheck{}
}

func (e *ValEqualCheck) Check(check entity.Check) bool {
	c, t := GetCheckMap(check)
	if len(check.Value) != 1 {
		log.Error("equal check size error")
		return false
	}
	val, ok := c[check.Key]
	if !ok {
		log.Info("target can not find:" + check.Key)
		return false
	}
	switch v := val.(type) {
	case []string:
		for _, str := range v {
			if e.checkOne(str, check, t) {
				return true
			}
		}
		return false
	case string:
		return e.checkOne(v, check, t)
	default:
		log.Error("can not check:", v)
		return false
	}

}

func (e *ValEqualCheck) checkOne(v string, check entity.Check, t map[string]interface{}) bool {
	if valRegx.Match([]byte(check.Value[0])) {
		//当为匹配target中字段时
		targetKey := valRegx.FindStringSubmatch(check.Value[0])[1]
		return v == t[targetKey]
	}
	return v == check.Value[0]
}
