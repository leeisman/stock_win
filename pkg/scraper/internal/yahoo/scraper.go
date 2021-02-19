package yahoo

import (
	"context"
	"stock_win/pkg/db"
	"stock_win/pkg/model"
	"stock_win/pkg/scraper/config"
)

type Scraper struct {
	ctx         context.Context
	cfg         *config.Config
	connections *db.Connections
}

func NewScraper(ctx context.Context, cfg *config.Config, connections *db.Connections) *Scraper {
	return &Scraper{
		ctx:         ctx,
		cfg:         cfg,
		connections: connections,
	}
}

func (s *Scraper) Scrape() {
	tx := s.connections.WriteDB
	stock := &model.Stock{
		SymbolID:   "1",
		SymbolName: "測試",
		Price:      100,
	}
	tx.Create(stock)
}
