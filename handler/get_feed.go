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

// models

type reqFeedURL struct {
	feedURL string `json:"feedURL"`
}

type RSS struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
	Desc  string `xml:"description"`
	Items []Item `xml:"item"`
}

type Item struct {
	Title     string    `xml:"title"`
	Link      string    `xml:"link"`
	Desc      string    `xml:"description"`
	Guid      string    `xml:"guid"`
	Source   string    `xml:"source"`
	Author    string    `xml:"author"`
   PubDate    string `xml:"pubDate"`
 }

type RSSFeedItem struct {
	PubDate    time.Time `json:"pubDate"`
	Title     string    `json:"title"`
	Link      string   `json:"link"`
	Desc      string   `json:"description"`
	Guid      string   `json:"guid,omitempty"`
	Source   string   `json:"source,omitempty"`
	Author    string   `json:"author,omitempty"`
	
}

type RSSFeed struct {
	Updated			time.Time //`json:"LastUpdated"`
	RSSFeedList []RSSFeedItem //`json:"RSSFeedList"`
}


func RSSFeedHandler(c echo.Context) error {
  // Take feedURL as formData
  // TODO: Get RSS feed and return to client

  feedURL := c.FormValue("feedURL") 
  return c.String(http.StatusOK, feedURL) 

}