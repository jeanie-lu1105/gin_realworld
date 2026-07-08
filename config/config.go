package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type config struct {
	Secret             string
	PublicKeyLocation  string
	PrivateKeyLocation string
	RedisAddr          string
	MySQLDSN           string `mapstructure:"mysqlDSN"`
}

var _config config

func init() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	currentPath, _ := os.Getwd()
	viper.AddConfigPath(currentPath)
	if exePath, err := os.Executable(); err == nil {
		viper.AddConfigPath(filepath.Dir(exePath))
	}
	os.Setenv("GO_GIN_PATH", "/Users/lujiaxin")
	if gopath := os.Getenv("GO_GIN_PATH"); gopath != "" {
		viper.AddConfigPath(filepath.Join(gopath, "GolandProjects", "gin_realworld"))
	}
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	if err := viper.Unmarshal(&_config); err != nil {
		panic(err)
	}
}

func GetSecret() string {
	return _config.Secret
}

func GetPrivateKeyLocation() string {
	currentPath, _ := os.Getwd()
	return filepath.Join(currentPath, _config.PrivateKeyLocation)
}

func GetPublicKeyLocation() string {
	currentPath, _ := os.Getwd()
	return filepath.Join(currentPath, _config.PublicKeyLocation)
}

func GetRedisAddr() string {
	return _config.RedisAddr
}

func GetMySQLDSN() string {
	return _config.MySQLDSN
}
