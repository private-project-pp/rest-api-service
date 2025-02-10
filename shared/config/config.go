package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type ConfigApp struct {
	Service  ServiceConfig  `mapstructure:"service" yaml:"service"`
	Internal InternalConfig `mapstructure:"db" yaml:"db"`
}

type ServiceConfig struct {
	Name string `mapstructure:"name" yaml:"name"`
	Port string `mapstructure:"port" yaml:"port"`
}

type InternalConfig struct {
	UserRpcService RpcServiceConfig `mapstructure:"user_rpc_service" yaml:"user_rpc_service"`
}

type RpcServiceConfig struct {
	Address string `mapstructure:"address" yaml:"address"`
}

var (
	Service  ServiceConfig
	Internal InternalConfig
)

func SetupConfig() (out ConfigApp, err error) {
	viper.AddRemoteProvider("vault", "https://127.0.0.1:8443", "/secret/whatever/config.json")
	viper.SetConfigType("json")
	err = viper.ReadRemoteConfig()
	if err != nil {
		err = readLocalConfig()
	}
	if err != nil {
		return out, err
	}

	viper.Unmarshal(&out)

	Service = out.Service
	Internal = out.Internal
	return out, nil
}

func readLocalConfig() (err error) {
	path, err := os.Getwd()
	if err != nil {
		return err
	}
	viper.SetConfigName("config.json")
	viper.SetConfigType("json")
	viper.AddConfigPath(filepath.Join(path, "shared/config"))
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}
