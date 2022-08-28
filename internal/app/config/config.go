package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

var Instance config

const (
	envConfigLocation     string = "YADCMD.CONFIG"
	defaultConfigLocation string = "configs/config.toml"
)

func init() {
	var configFile string = defaultConfigLocation
	if configLocation, exist := os.LookupEnv(envConfigLocation); exist {
		configFile = configLocation
	}
	if _, err := toml.DecodeFile(configFile, &Instance); err != nil {
		panic(err)
	}
}

type (
	config struct {
		Serve serve `toml:"serve"`
	}
	serve struct {
		Local  serveLocal  `toml:"local"`
		Remote serveRemote `toml:"remote"`
	}
	serveLocal struct {
		Unix *ServeAnyUnix `toml:"unix,omitempty"`
		TCP  *ServeAnyTCP  `toml:"tcp,omitempty"`
	}
	serveRemote struct {
		TCP *ServeAnyTCP `toml:"tcp,omitempty"`
	}

	ServeAnyUnix struct {
		Socket string `toml:"socket"`
	}
	ServeAnyTCP struct {
		IP   [4]byte `toml:"ip"`
		Port uint16  `toml:"port"`
	}
)
