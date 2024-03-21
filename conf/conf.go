package conf

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"work.ctyun.cn/git/GoStack/gostone/policy/entity"
)

var (
	ruleMap   = make(map[string]entity.Rule)
	policyMap = make(map[string]entity.Policy)
	confPath  string
)

func InitPath(path string) {
	confPath = path
	initConf()
}

func initConf() {
	filename, _ := filepath.Abs(confPath + "/policy.yaml")
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Error(err)
	}
	var conf entity.Keystone
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		panic(err)
	}
	for _, rule := range conf.Keystone.Rule {
		values := make([]entity.Values, 0)
		for _, val := range rule.Values {
			val.Key = Capitalize(val.Key)
			values = append(values, val)
		}
		rule.Values = values
		ruleMap[rule.Name] = rule
	}
	for _, p := range conf.Keystone.Policy {
		values := make([]entity.Values, 0)
		for _, val := range p.Values {
			val.Key = Capitalize(val.Key)
			values = append(values, val)
		}
		p.Values = values
		policyMap[p.Name] = p
	}
}

func GetRule(name string) (entity.Rule, bool) {
	r, ok := ruleMap[name]
	return r, ok
}

func GetPolicy(name string) (entity.Policy, bool) {
	p, ok := policyMap[name]
	return p, ok
}

// Capitalize 字符首字母大写
func Capitalize(str string) string {
	var upperStr string
	vv := []rune(str)
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 {
				vv[i] -= 32
				upperStr += string(vv[i])
			} else {
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}
