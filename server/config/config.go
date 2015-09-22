package config

import "github.com/BurntSushi/toml"

//Config this is a config structure
type Config struct {
	//[server]
	Server struct {
		Bind string `toml:"bind"`
	} `toml:"server"`

	//[key]
	Key struct {
		PublicPath  string `toml:"public_key"`
		PrivatePath string `toml:"private_key"`
	} `toml:"key"`
}

//New read a configuration file and returns a Config object
func New(configFile string) (*Config, error) {
	config := &Config{}

	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		return nil, err
	}

	return config, nil
}
