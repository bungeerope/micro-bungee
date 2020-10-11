package conf

type Server struct {
	System System `mapstructure:"system" json:"system" yaml:"system"`
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}

type System struct {
	MultiPoint bool   `mapstructure:"multi-point" json:"multipoint" yaml:"multi-point"`
	Env        string `mapstructure:"env" json:"env" yaml:"env"`
	Port       int    `mapstructure:"port" json:"port" yaml:"port"`
	DbType     string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`
}

type Mysql struct {
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	Url          string `mapstructure:"url" json:"url" yaml:"url"`
	Dbname       string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	LogMode      bool   `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
}
