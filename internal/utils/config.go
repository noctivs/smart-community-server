package utils

import "github.com/spf13/viper"

var Config *viper.Viper

func InitConfig() {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./configs")

	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	Config = v
}
