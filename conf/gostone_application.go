package conf

type Application struct {
	GoStone struct {
		Database struct {
			Url     string `yaml:"Url"`
			Timeout int    `yaml:"Timeout"`
		} `yaml:"Database"`
		Port          []int   `yaml:"Port"`
		LogLevel      string  `yaml:"LogLevel"`
		LogRotateTime int     `yaml:"LogRotateTime"`
		LogMaxAge     int     `yaml:"LogMaxAge"`
		Secret        string  `yaml:"Secret"`
		ExpiresTime   int     `yaml:"ExpiresTime"`
		BaseUrl       string  `yaml:"BaseUrl"`
		FernetPath    string  `yaml:"FernetPath"`
		TokenType     string  `yaml:"TokenType"`
		ConfPath      string  `yaml:"ConfPath"`
		LogPath       string  `yaml:"LogPath"`
		RetryConnect  int     `yaml:"RetryConnect"`
		MaxIdleConns  int     `yaml:"MaxIdleConns"`
		MaxOpenConns  int     `yaml:"MaxOpenConns"`
		SignMethod    string  `yaml:"SignMethod"`
		RefreshTime   float64 `yaml:"RefreshTime"`
		SkipAuth      bool    `yaml:"SkipAuth"`
		AdminPassword string  `yaml:"AdminPassword"`
		InitEndpoint  string  `yaml:"InitEndpoint"`
		InitRegion    string  `yaml:"InitRegion"`
	} `yaml:"GoStone"`
}
