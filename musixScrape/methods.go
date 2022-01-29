package musixScrape

import (
	"strings"

	"github.com/gocolly/colly/v2"
)

func (c *Client) GetLyricsFromLink(url string) (*LyricResult, error) {
	res := new(LyricResult)
	// Instantiate default collector
	// On every an element which has href attribute call callback
	c.Opts.Collector.OnHTML(c.Opts.LyricCssSelector, func(e *colly.HTMLElement) {
		// link := e.Attr("href")
		res.Lyrics = e.Text
	})
	c.Opts.Collector.OnHTML(c.Opts.SongNameCssSelector, func(e *colly.HTMLElement) {
		res.Song = strings.ReplaceAll(e.Text, "Lyrics", "")
	})

	c.Opts.Collector.OnHTML(c.Opts.ArtistNameCssSelector, func(e *colly.HTMLElement) {
		res.Artist = e.Text
	})
	// Before making a request print "Visiting ..."
	c.Opts.Collector.OnRequest(func(r *colly.Request) {
		if c.loggerFunc != nil {
			c.loggerFunc("[MusixScrape Debug] Visiting", r.URL.String())
		}
	})

	err := c.Opts.Collector.Visit(url)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) Search(query string) ([]LyricResult, error) {
	var res []LyricResult
	var err error
	url := c.Opts.SearchUrl + query
	/*
		c.OnHTML(`div.box:nth-child(1) > div:nth-child(2) > div:nth-child(1) > ul:nth-child(1) > li:nth-child(1) > div:nth-child(1) > div:nth-child(2) > div:nth-child(1) > h2:nth-child(1) > a:nth-child(1)`,
			func(e *colly.HTMLElement) {
				if l := e.Attr("href"); l != "" {
					link := "https://" + Domain + l
					ly, err := GetLyricsFromLink(link)
					if err != nil {
						return
					}
					res = append(res, *ly)
				}
			})
	*/
	c.Opts.Collector.OnHTML(c.Opts.SearchOuterCssSelector, func(element *colly.HTMLElement) {
		element.ForEach(c.Opts.SearchInnerCssSelector,
			func(i int, element *colly.HTMLElement) {
				if l := element.Attr("href"); l != "" {
					link := "https://" + c.Opts.Domain + l
					ly, err := c.GetLyricsFromLink(link)
					if err != nil {
						return
					}

					res = append(res, *ly)
				}
			})
	})
	// Before making a request print "Visiting ..."
	c.Opts.Collector.OnRequest(func(r *colly.Request) {
		if c.loggerFunc != nil {
			c.loggerFunc("[MusixScrape Debug] Visiting", r.URL.String())
		}
	})

	err = c.Opts.Collector.Visit(url)
	if err != nil {
		return nil, err
	}

	return res, nil
}
