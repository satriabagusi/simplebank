/*
Author: Satria Bagus(satria.bagus18@gmail.com)
config.go (c) 2023
Desc: description
Created:  2023-06-29T16:43:56.619Z
Modified: !date!
*/

package config

import (
	"github.com/satriabagusi/simplebank/utility"
)

type Config struct {
	ServerAddress string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Load() {
	c.ServerAddress = utility.GetEnv("SERVER_ADDRESS")
}
