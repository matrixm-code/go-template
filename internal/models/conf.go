package models

type DbConfig struct {
	Addr  string `yaml:"addr"`
	Max   int    `yaml:"max"`
	Idle  int    `yaml:"idle"`
	Debug bool   `yaml:"debug"`
}

type LogConfig struct {
	Level string `json:"level" yaml:"level"`
	Development bool `json:"development" yaml:"development"`
	Encoding string `json:"encoding" yaml:"encoding"`
	OutputPaths []string `json:"outputPaths" yaml:"outputPaths"`
	ErrorOutputPaths []string `json:"errorOutputPaths" yaml:"errorOutputPaths"`
}