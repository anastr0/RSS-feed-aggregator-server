package main

import (
	"github.com/anaswaratrajan/RSS-feed-aggregator-server/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"fmt"
	"os"
)

func main() {
	//port := ":5000"
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// enable CORs to any origin
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/v1/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Healthcheck") // ping check
	})

	e.POST("/v1/rss-feed/", handler.RSSFeedHandler)            // get RSS feed to given feedURL
	e.POST("/v1/aggregate-feed/", handler.RSSAggregateHandler) // get aggregated RSS feed to list of feedURLs

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s", os.Getenv("PORT")))) 
}
