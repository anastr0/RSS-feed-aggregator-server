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


// all functions

func parseDate(dateString string) time.Time {
	// lots of date formats to handle, so used a package.
	t, err := dateparse.ParseLocal(dateString)
	if err != nil {
		fmt.Printf("Error parsing date")
		return time.Time{} // return zero value, (shows up in bottom of feed )
	}
	return t
}

func fetchRSS(feedURL string) []RSSFeedItem {

	// fetch content from given feedURL and decodes into rss struct, returns list of items
	resp, err := http.Get(feedURL) // fetch from feedURL
	if err != nil {
		log.Fatalln(err)
	}
	rss := RSS{} // new RSS model
	reqBody, err := ioutil.ReadAll(resp.Body)
    xml.Unmarshal(reqBody, &rss)

	RSSFeed := make([]RSSFeedItem, 0)

	// return feed items list
	for _, item := range rss.Channel.Items {
		pubDate := parseDate(item.PubDate) // parse date from dateString
		rssFeedItem := RSSFeedItem{ 
			PubDate:        pubDate,
			Title:     item.Title,   
			Link:      item.Link,
			Desc:      item.Desc, 
			Guid:      item.Guid,
			Source:  item.Source, 
			Author:   item.Author,  
		}
		RSSFeed = append(RSSFeed, rssFeedItem)
	}
	return RSSFeed 
}

func getRSSFeed(feedURL string) []RSSFeedItem {
	// called by getRSSFeedhandler
	// returns list of feed items from given feedURL

	rssFeed := fetchRSS(feedURL) // fetch and get feeditems
	// sort RSS feed items
	sort.Slice(rssFeed, func(i, j int) bool {
		return rssFeed[j].PubDate.Before(rssFeed[i].PubDate)
	})
	return rssFeed
}
func RSSFeedHandler(c echo.Context) error {
	// : Get RSS feed to given feedURL and return to client

	// Read json from request body and get feedURL
	json_map := make(map[string]interface{}) 
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	feedURL := json_map["feedURL"].(string)
	if err!=nil {
		return c.String(http.StatusOK, "Provide valid Json")
	}

	// Get RSS feed for the given feedURL and return as JSON
	rssFeed := &RSSFeed{
		Updated: time.Now(),
		RSSFeedList: getRSSFeed(feedURL),
	}
	return c.JSON(http.StatusOK, rssFeed)
}

