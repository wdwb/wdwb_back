//go:build wireinject
// +build wireinject

package di

import (
	"github.com/changchanghwang/wdwb_back/internal/libs/db"
	secClient "github.com/changchanghwang/wdwb_back/internal/libs/sec-client"
	"github.com/changchanghwang/wdwb_back/internal/server"
	"github.com/changchanghwang/wdwb_back/internal/services/filings"
	"github.com/changchanghwang/wdwb_back/internal/services/holdings"
	"github.com/changchanghwang/wdwb_back/internal/services/investors"
	"github.com/changchanghwang/wdwb_back/internal/services/stocks"
	"github.com/changchanghwang/wdwb_back/internal/services/sync"
	"github.com/google/wire"
)

func InitializeServer() (*server.Server, error) {
	wire.Build(
		db.Init,
		server.ServerSet,
		secClient.New,
		stocks.StockSet,
		filings.FilingSet,
		investors.InvestorSet,
		holdings.HoldingSet,
		sync.SyncSet,
	)
	return &server.Server{}, nil
}
