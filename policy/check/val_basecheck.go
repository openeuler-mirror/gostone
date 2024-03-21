package check

import (
	"reflect"
	"work.ctyun.cn/git/GoStack/gostone/policy/entity"
)

var valCheckMap = make(map[entity.Operator]ValBaseCheck)

func init() {
	valCheckMap[entity.Equal] = NewValEqualCheck()
	valCheckMap[entity.RuleOp] = NewRuleCheck()
	valCheckMap[entity.In] = NewInCheck()
	valCheckMap[entity.NotIn] = NewNotInCheck()
	valCheckMap[entity.NotEqual] = NewNotCheck()
}

func GetValCheck(operator entity.Operator) ValBaseCheck {
	return valCheckMap[operator]
}

type ValBaseCheck interface {
	Check(check entity.Check) bool
}

func GetCheckMap(check entity.Check) (map[string]interface{}, map[string]interface{}) {
	return StructToMap(check.Context), StructToMap(check.Target)
}

func StructToMap(obj interface{}) map[string]interface{} {
	if obj == nil {
		return map[string]interface{}{}
	}
	contextMap, ok := obj.(map[string]interface{})
	if ok {
		return contextMap
	}
	m := make(map[string]interface{})
	elem := reflect.ValueOf(obj)
	relType := reflect.TypeOf(obj)
	if elem.Kind() == reflect.Ptr {
		elem = elem.Elem()
		relType = relType.Elem()
	}
	for i := 0; i < relType.NumField(); i++ {
		m[relType.Field(i).Name] = elem.Field(i).Interface()
	}
	return m
}
