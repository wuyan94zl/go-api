package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func InitConfig(configName string) {
	viper.SetConfigName(configName)
	viper.SetConfigType("yml")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(fmt.Errorf("Fatal error when reading %s config file:%s", configName, err))
		os.Exit(1)
	}
}
