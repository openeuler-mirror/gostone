package connect

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"work.ctyun.cn/git/GoStack/gostone/conf"
)

var AppConf conf.Application

func InitConf() {
	args := os.Args
	yamlPath := "./etc/application.yaml"
	if len(args) != 1 {
		yamlPath = args[1]
	}
	filename, _ := filepath.Abs(yamlPath)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &AppConf)
}
