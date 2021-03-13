package handler

import (
  "net/http"

  "github.com/labstack/echo"
)

func RSSAggregateHandler(c echo.Context) error {
  // Take feedURL as formData
  // TODO: Get RSS feed and return to client
  
  return c.String(http.StatusOK, "Aggregated daily news feed")
}