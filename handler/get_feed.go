package handler

import (
  "net/http"

  "github.com/labstack/echo/v4"
)

func RSSFeedHandler(c echo.Context) error {
  // Take feedURL as formData
  // TODO: Get RSS feed and return to client

  feedURL := c.FormValue("feedURL") 
  return c.String(http.StatusOK, feedURL) 

}