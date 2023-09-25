package configs

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// define the structure of the config to be set
type AppConfig struct {
	Port       string
	DbHost     string `split_words:"true"`
	DbUser     string `split_words:"true"`
	DbPassword string `split_words:"true"`
	DbName     string `split_words:"true"`
	DbPort     string `split_words:"true"`
}

// creating a global config
var App *AppConfig

func loadAppConfig() {
	App = &AppConfig{}

	// process the config variable starting with APP_
	err := envconfig.Process("app", App)
	if err != nil {
		log.Fatal(err.Error())
	}
}
