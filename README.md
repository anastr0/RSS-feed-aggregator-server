# RSS-feed-aggregator-server
REST Server api in GO, aggregates RSS feed, 
hosted [here](https://rss-feed-aggregater-server.herokuapp.com/v1/)!  


Check out [client](https://github.com/anaswaratrajan/RSS-feed-aggregator-client)) repository.
## Table of contents

- [Prerequisites](#prerequisites)
- [Setup and run](#Setup)
- [Usage](#Usage)
- [Features](#Features)
- [Repository-structure ](#Repository-structure )

## Prerequisites 

* Go 1.4+

## Setup

```
git clone git@github.com:anaswaratrajan/RSS-feed-aggregator-server.git 

cd RSS-feed-aggregator-server

PORT=5000 go run server go
```

http://localhost:5000

## Endpoints 

* POST `/v1/rss-feed/` 

* POST `/v1/aggregate-feed/`

## Usage 

Examples - 

curl RSS-feed
```
curl -X POST http://localhost:5000/v1/rss-feed/ \
-d '{"feedURL":"https://home.cern/api/news/news/feed.rss"}' 
```
curl aggregated RSS-feed
```
curl -X POST http://localhost:5000/v1/aggregate-feed/ \
-d '{"feedURLs":["https://home.cern/api/news/news/feed.rss", "https://www.nasa.gov/rss/dyn/breaking_news.rss"]}' 
```

## Repository 

```
.
├── bin
│   └── .. 
├── go.mod
├── go.sum
├── handler
│   └── get_feed.go
├── LICENSE
├── Procfile
├── README.md
├── server.go
. 
. 
```