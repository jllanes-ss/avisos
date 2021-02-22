package config

import (
	"log"

	"github.com/spf13/viper"
)

func LoadEnv() error {
	// see if the .env file exist
	viper.SetConfigFile(`.env`)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	//is debug environment
	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
	return nil
}
