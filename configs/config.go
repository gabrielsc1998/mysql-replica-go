package configs

import (
	"github.com/spf13/viper"
)

type Conf struct {
	DatabaseMasterHost string `mapstructure:"MYSQL_MASTER_HOST"`
	DatabaseMasterPort string `mapstructure:"MYSQL_MASTER_PORT"`
	DatabaseMasterUser string `mapstructure:"MYSQL_MASTER_USER"`
	DatabaseMasterPass string `mapstructure:"MYSQL_MASTER_PASSWORD"`
	DatabaseMasterName string `mapstructure:"MYSQL_MASTER_DATABASE"`
	DatabaseSlaveHost  string `mapstructure:"MYSQL_SLAVE_HOST"`
	DatabaseSlavePort  string `mapstructure:"MYSQL_SLAVE_PORT"`
	DatabaseSlaveUser  string `mapstructure:"MYSQL_SLAVE_USER"`
	DatabaseSlavePass  string `mapstructure:"MYSQL_SLAVE_PASSWORD"`
	DatabaseSlaveName  string `mapstructure:"MYSQL_SLAVE_DATABASE"`
}

func LoadConfig(path string) (*Conf, error) {
	var cfg *Conf
	viper.SetConfigType("env")
	viper.AddConfigPath("../../../../..")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg, err
}
