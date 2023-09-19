package infrastructure

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	HOST        string `mapstructure:"HOST"`
	PORT        string `mapstructure:"PORT"`
	APIVersion  string `mapstructure:"API_VERSION"`
	MongoURI    string `mapstructure:"MONGO_URI"`
	Environment string `mapstructure:"ENVIRONMENT"`
}

// NewEnv creates a new environment
func NewEnv() Env {
	env := Env{}
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("☠️ Env config file not found: ", err)
		} else {
			log.Fatal("☠️ Env config file error: ", err)
		}
	}

	if err := viper.Unmarshal(&env); err != nil {
		log.Fatal("☠️ environment can't be loaded: ", err)
	}

	log.Printf("%+v \n", env)
	return env
}
