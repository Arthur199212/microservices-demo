package utils

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port                string `mapstructure:"PORT"`
	CartServiceAddr     string `mapstructure:"CART_SERVICE_ADDR"`
	CheckoutServiceAddr string `mapstructure:"CHECKOUT_SERVICE_ADDR"`
	CurrencyServiceAddr string `mapstructure:"CURRENCY_SERVICE_ADDR"`
	ProductsServiceAddr string `mapstructure:"PRODUCTS_SERVICE_ADDR"`
	ShippingServiceAddr string `mapstructure:"SHIPPING_SERVICE_ADDR"`
}

func LoadConfig(path string) (Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.AutomaticEnv()

	config := Config{}
	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)
	return config, err
}
