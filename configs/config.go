package configs

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort  string
	AppDebug string

	MysqlHost     string
	MysqlPort     string
	MysqlDb       string
	MysqlUser     string
	MysqlPassword string

	PostgresDb       string
	PostgresUser     string
	PostgresPassword string
}

var AppConfig Config

func LoadConfig() error {
	viper.SetConfigName("env")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	AppConfig = Config{
		AppPort:  viper.GetString("app.port"),
		AppDebug: viper.GetString("app.debug"),

		MysqlHost:     viper.GetString("mysql.mysql_host"),
		MysqlPort:     viper.GetString("mysql.mysql_port"),
		MysqlDb:       viper.GetString("mysql.mysql_db"),
		MysqlUser:     viper.GetString("mysql.mysql_user"),
		MysqlPassword: viper.GetString("mysql.mysql_password"),
	}

	return nil
}
