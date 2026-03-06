package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type config struct {
	Secret string
	PublicKeyLocation string
	PrivateKeyLocation string
}

var _config config

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/Users/lujiaxin/GolandProjects/gin_realworld/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	if err := viper.Unmarshal(&_config); err != nil{
		panic(err)
	}

	fmt.Println(viper.Get("secret"))
}

func GetSecret() string{
	return _config.Secret
}

func GetPrivateKeyLocation() string {
	return _config.PrivateKeyLocation
}

func GetPublicKeyLocation() string {
	return _config.PublicKeyLocation
}