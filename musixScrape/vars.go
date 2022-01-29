package musixScrape

import "github.com/gocolly/colly/v2"

var DefaultOpts = &ScrapeOpts{
	Domain:                 Domain,
	ArtistNameCssSelector:  ArtistNameCssSelector,
	SongNameCssSelector:    SongNameCssSelector,
	SearchUrl:              SearchUrl,
	LyricCssSelector:       LyricCssSelector,
	SearchOuterCssSelector: SearchOuterSelector,
	SearchInnerCssSelector: SearchInnerSelector,
	Collector: colly.NewCollector(
		colly.AllowedDomains(Domain),
		colly.AllowURLRevisit()),
}
