package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"os"
	"stock_win/config"
	"stock_win/pkg/db"
	"stock_win/pkg/scraper"
)

var Scrape = &cobra.Command{
	Run: run,
	Use: "scrape",
}

func run(cmd *cobra.Command, args []string) {
	cfg, err := config.New()
	if err != nil {
		os.Exit(1)
		return
	}

	app := fx.New(
		fx.Supply(*cfg),
		fx.Provide(
			db.InitDatabases,
		),
		fx.Invoke(
			scraper.Scrape,
		),
	)

	app.Run()
}
