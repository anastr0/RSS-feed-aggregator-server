package main

import (
	"net/http"
	
	"github.com/labstack/echo/v4"

	"github.com/anaswaratrajan/RSS-feed-aggregator-server/handler"
)

func main() {
	e := echo.New()

	e.GET("/v1", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/v1/rss-feed", handler.RSSFeedHandler)
	e.GET("/v1/aggregate-feed", handler.RSSAggregateHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
