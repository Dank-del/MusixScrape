package musixScrape

import (
	"github.com/gocolly/colly/v2"
)

type LyricResult struct {
	Song   string
	Artist string
	Lyrics string
}

type Client struct {
	Collector  *colly.Collector
	Opts       *ScrapeOpts
	loggerFunc func(v ...interface{})
}

type ScrapeOpts struct {
	Domain                string
	SearchUrl             string
	LyricCssSelector      string
	SongNameCssSelector   string
	ArtistNameCssSelector string
	Logger                func(v ...interface{})
}
