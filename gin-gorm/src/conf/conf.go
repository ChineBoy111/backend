package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigName("settings")
	viper.SetConfigType("json")
	// viper.AddConfigPath("./src/conf/")
	viper.AddConfigPath("./conf/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Read in config error %s", err.Error()))
	}
	// fmt.Println(viper.GetString("server.port"))
}
