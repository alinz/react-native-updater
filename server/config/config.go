package config

import "github.com/BurntSushi/toml"

//Config this is a config structure
type Config struct {
	//[server]
	Server struct {
		Bind string `toml:"bind"`
	} `toml:"server"`

	//[rsa]
	RSA struct {
		Passphrase  string `toml:"passphrase"`
		PublicPath  string `toml:"public_key"`
		PrivatePath string `toml:"private_key"`
	} `toml:"rsa"`

	//[aes]
	AES struct {
		SecureKey string `toml:"secure_key"`
	} `toml:"aes"`

	//[db]
	DB struct {
		Database string   `toml:"database"`
		Hosts    []string `toml:"hosts"`
		Username string   `toml:"username"`
		Password string   `toml:"password"`
	} `toml:"db"`
}

//New read a configuration file and returns a Config object
func New(configFile string) (*Config, error) {
	config := &Config{}

	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		return nil, err
	}

	return config, nil
}
