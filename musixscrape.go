package MusixScrape

import (
	"github.com/gocolly/colly/v2"
	"log"
	"strings"
)

func GetLyricsFromLink(url string) (*LyricResult, error) {
	res := LyricResult{}
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains(Domain),
	)
	// On every a element which has href attribute call callback
	c.OnHTML(LyricCssSelector, func(e *colly.HTMLElement) {
		// link := e.Attr("href")
		res.Lyrics = e.Text
	})
	c.OnHTML(SongNameCssSelector, func(e *colly.HTMLElement) {
		res.Song = strings.ReplaceAll(e.Text, "Lyrics", "")
	})

	c.OnHTML(ArtistNameCssSelector, func(e *colly.HTMLElement) {
		res.Artist = e.Text
	})
	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Println("[MusixScrape Debug] Visiting", r.URL.String())
	})
	err := c.Visit(url)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func Search(query string) ([]LyricResult, error) {
	var res []LyricResult
	var err error
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains(Domain),
	)
	url := SearchUrl + query
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
	c.OnHTML(`div.box:nth-child(2) > div:nth-child(2) > div:nth-child(1) > ul:nth-child(1)`, func(element *colly.HTMLElement) {
		element.ForEach(`div.box:nth-child(2) > div:nth-child(2) > div:nth-child(1) > ul:nth-child(1) > li:nth-child(1) > div:nth-child(1) > div:nth-child(2) > div:nth-child(1) > h2:nth-child(1) > a:nth-child(1)`, func(i int, element *colly.HTMLElement) {
			if l := element.Attr("href"); l != "" {
				link := "https://" + Domain + l
				ly, err := GetLyricsFromLink(link)
				if err != nil {
					return
				}
				res = append(res, *ly)
			}
		})
	})
	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Println("[MusixScrape Debug] Visiting", r.URL.String())
	})
	err = c.Visit(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}
