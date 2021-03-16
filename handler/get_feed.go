package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/araddon/dateparse"
	"net/http"
	"io/ioutil"
	"encoding/xml"
	"encoding/json"
	"fmt"
	"log"
	"time"
	"sort"
)

func RSSFeedHandler(c echo.Context) error {
  // Take feedURL as formData
  // TODO: Get RSS feed and return to client

  feedURL := c.FormValue("feedURL") 
  return c.String(http.StatusOK, feedURL) 

}