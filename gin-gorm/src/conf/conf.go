package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

// LoadConf 加载配置文件
func LoadConf() {
	viper.SetConfigName("settings")
	viper.SetConfigType("json")

	// viper.AddConfigPath("./src/conf/")
	viper.AddConfigPath("./conf/")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Read in config error %s", err.Error()))
	}
}
