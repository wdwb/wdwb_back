package server

import (
	_ "github.com/changchanghwang/wdwb_back/docs"
	"github.com/changchanghwang/wdwb_back/internal/config"
	holdings "github.com/changchanghwang/wdwb_back/internal/services/holdings/presentation"
	investors "github.com/changchanghwang/wdwb_back/internal/services/investors/presentation"
	stocks "github.com/changchanghwang/wdwb_back/internal/services/stocks/presentation"
	sync "github.com/changchanghwang/wdwb_back/internal/services/sync/presentation"
	"github.com/gofiber/fiber/v2"
	swagger "github.com/gofiber/swagger"
)

type Route struct {
	stocks   *stocks.StockController
	sync     *sync.SyncController
	investor *investors.InvestorController
	holding  *holdings.HoldingController
}

func NewRoute(
	stocks *stocks.StockController,
	sync *sync.SyncController,
	investor *investors.InvestorController,
	holding *holdings.HoldingController,
) *Route {
	return &Route{
		stocks:   stocks,
		sync:     sync,
		investor: investor,
		holding:  holding,
	}
}

// @title			wdwb API
// @version		1.0
// @description	API Server for wdwb
// @contact.name	API Support
// @contact.email	window95pill@gmail.com
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath		/
func (r *Route) Route(app *fiber.App) {
	// swagger
	app.Get("/swagger/*", swagger.HandlerDefault, swagger.New(swagger.Config{ // custom
		URL:         config.Origin + "/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
	}))
	// stocks
	stocksGroup := app.Group("/stocks")
	r.stocks.Route(stocksGroup)

	// sync
	syncGroup := app.Group("/sync")
	r.sync.Route(syncGroup)

	// investors
	investorsGroup := app.Group("/investors")
	r.investor.Route(investorsGroup)

	// holdings
	holdingsGroup := app.Group("/holdings")
	r.holding.Route(holdingsGroup)
}
