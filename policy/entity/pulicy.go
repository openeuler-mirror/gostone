package entity

type Keystone struct {
	Keystone struct {
		Rule   []Rule   `yaml:"Rule"`
		Policy []Policy `yaml:"Policy"`
	} `yaml:"Keystone"`
}

type Rule struct {
	Name     string   `yaml:"Name"`
	Operator Operator `yaml:"Operator"`
	Values   []Values `yaml:"Values"`
}

type Values struct {
	Key      string   `yaml:"Key"`
	Operator Operator `yaml:"Operator"`
	Values   []string `yaml:"Values"`
}

type Policy struct {
	Name     string   `yaml:"Name"`
	Operator Operator `yaml:"Operator"`
	Values   []Values `yaml:"Values"`
}

type Check struct {
	Key     string
	Value   []string
	Context interface{}
	Target  interface{}
}
