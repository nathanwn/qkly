package workspace

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/afero"
	"github.com/spf13/viper"
)

type Config struct {
	v *viper.Viper
}

func NewConfig(fs afero.Fs) (*Config, error) {
	configFileDir := "."
	configFileName := "qkly.yaml"
	configType := "yaml"

	configFilePath := filepath.Join(configFileDir, configFileName)
	v := viper.New()

	v.SetFs(fs)
	v.SetConfigFile(configFilePath)
	v.SetConfigType(configType)
	v.SetDefault("port", "10042")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Printf(
				"Error: The config file \"%s\" does not exist\n",
				configFilePath,
			)
		}
		return nil, err
	}
	return &Config{v}, nil
}

func (conf Config) Port() string {
	return conf.v.GetString("port")
}
