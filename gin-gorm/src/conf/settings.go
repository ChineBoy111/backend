package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

// ReadConf 读取配置文件
func ReadConf() {
	viper.SetConfigName("settings")
	viper.SetConfigType("yaml")

	// ./main 和 ./settings.yaml 都在项目根目录下
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Read in config error %s", err.Error()))
	}
}
