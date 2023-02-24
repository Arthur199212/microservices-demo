package utils

import "github.com/spf13/viper"

type Config struct {
	Port                 string `mapstructure:"PORT"`
	AllowTestCardNumbers bool   `mapstructure:"ALLOW_TEST_CARD_NUMBERS"`
}

func LoadConfig(path string) (Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.AutomaticEnv()

	config := Config{}
	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	err := viper.Unmarshal(&config)
	return config, err
}
