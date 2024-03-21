package service

import (
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

func getQuery(values url.Values, target interface{}) {
	v := reflect.ValueOf(target)
	t := reflect.TypeOf(target)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}
	for i := 0; i < t.NumField(); i++ {
		key := t.Field(i).Tag.Get("json")
		if key == "" {
			key = strings.ToLower(t.Field(i).Name)
		}
		if v.Field(i).Kind() == reflect.Int {
			in, _ := strconv.ParseInt(values.Get(key), 10, 64)
			v.Field(i).SetInt(in)
		} else {
			v.Field(i).SetString(values.Get(key))
		}

	}
}

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
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

func Struct2Map4Search(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		key := t.Field(i).Tag.Get("json")
		if key == "" {
			key = strings.ToLower(t.Field(i).Name)
		}
		if key == "enabled" {
			data[key] = strings.ToLower(v.Field(i).String())
		} else {
			data[key] = v.Field(i).Interface()
		}

	}
	return data
}
func getEnabledBool(enabled *bool) bool {
	if enabled == nil || *enabled {
		return true
	}
	return false
}

func getEnabled(enabled *bool) int {
	if enabled == nil || *enabled {
		return 1
	}
	return 0
}

func transEnabled(end int) bool {
	if end == 1 {
		return true
	}
	return false
}
func getEnabledString(enabled string) int {
	if enabled == "" || enabled == "true" {
		return 1
	}
	return 0
}
