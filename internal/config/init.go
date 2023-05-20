package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

const (
	config_name = "conf"
	config_type = "yaml"
)

func Init() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName(config_name)
	viper.SetConfigType(config_type)

	if err := viper.ReadInConfig(); err != nil {
		if e, ok := err.(viper.ConfigFileNotFoundError); ok {
			return e
		} else {
			return err
		}
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Info().Str("name", e.Name).Str("string", e.String()).Msg("Config file update triggered")
	})
	log.Info().Msg("Config init success")

	return nil
}
