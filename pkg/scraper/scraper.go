package scraper

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
	"stock_win/pkg/db"
	"stock_win/pkg/scraper/config"
	"stock_win/pkg/scraper/internal"
	"stock_win/pkg/scraper/internal/yahoo"
)

func Scrape(lc fx.Lifecycle, cfg *config.Config, connections *db.Connections) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Info().Msgf("Starting scrape, mode on %s", cfg.ScrapeModel)
				scraper := newScraper(ctx, cfg, connections)
				scraper.Scrape()
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}

func newScraper(ctx context.Context, cfg *config.Config, connections *db.Connections) internal.ScraperInterface {
	switch cfg.ScrapeModel {
	case "yahoo":
		return yahoo.NewScraper(ctx, cfg,connections)
	}
	return nil
}
