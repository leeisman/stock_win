package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"stock_win/pkg/db"
	"stock_win/pkg/scraper/config"
)

type Config struct {
	fx.Out
	Scraper *config.Config
	DB      *db.Config `mapstructure:"database"`
}

// New 讀取App 啟動程式設定檔
func New() (*Config, error) {
	viper.AutomaticEnv()

	configPath := viper.GetString("CONFIG_PATH")

	if configPath == "" {
		configPath = viper.GetString("PROJ_DIR") + "/deployment/config"
	}

	configName := viper.GetString("CONFIG_NAME")
	if configName == "" {
		configName = "app"
	}

	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)
	viper.SetConfigType("yaml")

	var config Config

	if err := viper.ReadInConfig(); err != nil {
		log.Error().Msgf("Error reading config file, %s", err)
		return &config, err
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		log.Error().Msgf("unable to decode into struct, %v", err)
		return &config, err
	}
	return &config, nil
}
