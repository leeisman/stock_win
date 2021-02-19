package config

type Config struct {
	ScrapeModel string `mapstructure:"scrape_mode"`
	YahooURL    string `mapstructure:"yahoo_url"`
}
