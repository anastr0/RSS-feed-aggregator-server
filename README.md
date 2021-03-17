# RSS-feed-aggregator-server
REST Server in GO that aggregates RSS feed 


# Setup

```

// clone repo 
git clone git@github.com:anaswaratrajan/RSS-feed-aggregator-server.git 

cd RSS-feed-aggregator-server
```

#  Run

```
PORT=:5000 go run server go
```

# Endpoints 

* POST `/v1/rss-feed/` 
* POST `/v1/aggregate-feed/` 

# Usage 

```
// curl RSS-feed
curl -X POST http://localhost:5000/v1/rss-feed/ \
-d '{"feedURL":"https://home.cern/api/news/news/feed.rss"}' 

// curl aggregated RSS-feed
curl -X POST http://localhost:5000/v1/aggregate-feed/ \
-d '{"feedURLs":["https://home.cern/api/news/news/feed.rss", "https://www.nasa.gov/rss/dyn/breaking_news.rss"]}' 
```
