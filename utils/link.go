package utils

import (
	"database/sql"
	"reflect"
	"strings"
	"work.ctyun.cn/git/GoStack/gostone/conf"
)

type LinkPath = string

const (
	UserPath         LinkPath = "/v3/users/"
	DomainPath       LinkPath = "/v3/domains/"
	ProjectPath      LinkPath = "/v3/projects/"
	EndpointPath     LinkPath = "/v3/endpoints/"
	ServicePath      LinkPath = "/v3/services/"
	RegionPath       LinkPath = "/v3/regions/"
	RolePath         LinkPath = "/v3/roles/"
	UserProjectsPath LinkPath = "/v3/users/%s/projects"
)

func SetSingleLink(source interface{}, path LinkPath, key string, ignoreString []string) (result map[string]interface{}) {
	target := StructToMap(source)
	if ignoreString != nil {
		for _, field := range ignoreString {
			if _, ok := target[field]; ok {
				delete(target, field)
			}
		}
	}
	target["links"] = map[string]interface{}{
		"self": conf.Url + string(path) + target["id"].(string),
	}
	enable, ok := target["enabled"]
	if ok {
		switch enable.(type) {
		case int:
			if enable.(int) == 1 {
				target["enabled"] = true
			} else {
				target["enabled"] = false
			}
		case int64:
			if enable.(int64) == 1 {
				target["enabled"] = true
			} else {
				target["enabled"] = false
			}
		}
	}
	isDomain, ok := target["is_domain"]
	if ok {
		switch isDomain.(type) {
		case int:
			if isDomain.(int) == 1 {
				target["is_domain"] = true
			} else {
				target["is_domain"] = false
			}
		case int64:
			if isDomain.(int64) == 1 {
				target["is_domain"] = true
			} else {
				target["is_domain"] = false
			}
		}
	}
	parentId, ok := target["parent_id"]
	if ok {
		pd := parentId.(sql.NullString)
		target["parent_id"] = pd.String
	}
	result = make(map[string]interface{})
	result[key] = target
	return
}

func SetArrayLink(source interface{}, path LinkPath, name string, ignoreString []string) (result map[string]interface{}) {
	v := reflect.ValueOf(source)
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		return
	}
	list := make([]map[string]interface{}, 0)
	for i := 0; i < v.Len(); i++ {
		in := v.Index(i).Interface()
		var target map[string]interface{}
		if v.Index(i).Kind() == reflect.Map {
			target = in.(map[string]interface{})
		} else {
			target = StructToMap(in)
		}
		enable, ok := target["enabled"]
		if ok {
			switch enable.(type) {
			case int:
				if enable.(int) == 1 {
					target["enabled"] = true
				} else {
					target["enabled"] = false
				}
			case int64:
				if enable.(int64) == 1 {
					target["enabled"] = true
				} else {
					target["enabled"] = false
				}
			}
		}
		isDomain, ok := target["is_domain"]
		if ok {
			switch isDomain.(type) {
			case int:
				if isDomain.(int) == 1 {
					target["is_domain"] = true
				} else {
					target["is_domain"] = false
				}
			case int64:
				if isDomain.(int64) == 1 {
					target["is_domain"] = true
				} else {
					target["is_domain"] = false
				}
			}
		}
		if ignoreString != nil {
			for _, field := range ignoreString {
				if _, ok := target[field]; ok {
					delete(target, field)
				}
			}
		}
		parentId, ok := target["parent_id"]
		if ok {
			pd := parentId.(sql.NullString)
			target["parent_id"] = pd.String
		}
		target["links"] = map[string]interface{}{
			"self": conf.Url + string(path) + target["id"].(string),
		}
		list = append(list, target)
	}
	result = make(map[string]interface{})
	result[name] = list
	result["links"] = map[string]interface{}{
		"next":     nil,
		"previous": nil,
		"self":     string(path),
	}
	return
}

func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Map {
		return obj.(map[string]interface{})
	}
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		key := t.Field(i).Tag.Get("json")
		if key == "" {
			key = strings.ToLower(t.Field(i).Name)
		}
		data[key] = v.Field(i).Interface()
	}
	return data
}
