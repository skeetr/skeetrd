package harvesterd

import (
	"harvesterd/logger"
)

import "code.google.com/p/gcfg"

type Config struct {
	Section logger.LoggerConfig
}

var configInstance *Config = new(Config)

func GetConfig() *Config {
	return configInstance
}

func (self *Config) Load(ini string) {
	err := gcfg.ReadStringInto(self, ini)
	if err != nil {
		logger.Critical("error: cannot parse config", err)
	}
}

func (self *Config) LoadFile(filename string) {
	err := gcfg.ReadFileInto(self, filename)
	if err != nil {
		logger.Critical("erro:", err)
	}
}
