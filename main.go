package main

import (
	"github.com/Iuptec/tupa"
)

func main() {
	server := tupa.NewAPIServer(":6969")
	server.RegisterRoutes(managerRoutes())

	server.New()

}

func managerRoutes() []tupa.RouteInfo {
	return []tupa.RouteInfo{
		{
			Path:    "/api/v1/tickers",
			Handler: HandleListAll,
			Method:  "GET",
		},
		{
			Path:    "/api/v1/{ticker}",
			Handler: HandleTicker,
			Method:  "GET",
		},
	}
}
