// author pengchengbai@shopee.com
// date 2021/7/18

package conf

import "time"

// Config stores tcp server properties
type Config struct {
	Address    string        `yaml:"address"`
	MaxConnect uint32        `yaml:"max-connect"`
	Timeout    time.Duration `yaml:"timeout"`
}

var tcpConfig Config

func GetTcpConfig() *Config {
	return &tcpConfig
}

func (config *Config)Init(path string) {
	if len(path) == 0 {
		config.MaxConnect = 1000
		config.Address = "127.0.0.1:6399"
		config.Timeout = 100 * time.Second
	}
}
