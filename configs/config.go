package configs

import (
	"sync"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// Config is a struct that will receive configuration options via environment
// variables.
type Config struct {
	CatAPI struct {
		APIKey string `mapstructure:"API_KEY"`
		URL    string `mapstructure:"URL"`
	} `mapstructure:"CAT_API"`
	Discord struct {
		Prefix string `mapstructure:"PREFIX"`
		Token  string `mapstructure:"TOKEN"`
	}
	NovelAI struct {
		Key string `mapstructure:"KEY"`
	} `mapstructure:"NOVELAI"`
	Saucenao struct {
		APIKey string `mapstructure:"API_KEY"`
	} `mapstructure:"SAUCENAO"`
	Server struct {
		Env      string `mapstructure:"ENV"`
		LogLevel string `mapstructure:"LOG_LEVEL"`
	}
}

var (
	conf Config
	once sync.Once
)

// Get are responsible to load env and get data an return the struct
func Get() *Config {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal().Err(err).Msg("Failed reading config file")
	}

	once.Do(func() {
		log.Info().Msg("Service configuration initialized.")
		err = viper.Unmarshal(&conf)
		if err != nil {
			log.Fatal().Err(err).Msg("")
		}
	})

	return &conf
}
