package config

import "strings"

func ConfigTem(svr string) string {
	t := `
package config

import (
	"github.com/spf13/viper"
	"sample/source/log"
	"sample/source/tool"
)

var config *viper.Viper

func GetConfig() *viper.Viper {
	if config != nil {
		return config
	}
	config = viper.New()
	config.AddConfigPath("./config")
	env := tool.GetTool().GetEnv()
	switch env {
	case "pro":
		config.SetConfigName("config")
	case "test":
		config.SetConfigName("configsample")
	case "dev":
		config.SetConfigName("configDev")
	}
	err := config.ReadInConfig()
	if err != nil {
		log.GetLogger().Fatal(err)
	}

	return config
}

`

	t = strings.ReplaceAll(t, "sample", svr)
	return t
}
