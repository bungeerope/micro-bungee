package conf

type Server struct {
	System System `mapstructure:"system" json:"system" yaml:"system"`
}

type System struct {
	MultiPoint bool   `mapstructure:"multi-point" json:"multipoint" yaml:"multi-point"`
	Env        string `mapstructure:"env" json:"env" yaml:"env"`
	Addr       int    `mapstructure:"addr" json:"addr" yaml:"addr"`
	DbType     string `mapstructure:"db-type" json:"dbType" yaml:"addr"`
}
