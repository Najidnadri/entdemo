package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	allEnv := [1]string{"DATABASE_URL"}
	if err != nil {
		log.Printf("WARN: error loading .env file. Loading from OS environment..")
		for i := range allEnv {
			viper.SetDefault(allEnv[i], os.Getenv(allEnv[i]))
		}
	}
	log.Println("ENVIRONMENTS >>>")
	for i := range allEnv {
		log.Printf("%s : %s", allEnv[i], viper.GetString(allEnv[i]))
	}
}

func GetDbUrl() string {
	return viper.GetString("DATABASE_URL")
}
