package config

import "github.com/spf13/viper"

type Config struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE"`
	DBHost   string `mapstructure:"DB_HOST"`
}

func SetupConfig(path string) (config Config, err error) {

	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
