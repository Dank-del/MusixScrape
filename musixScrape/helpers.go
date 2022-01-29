package musixScrape

import (
	"github.com/gocolly/colly/v2"
)

func New(collector *colly.Collector, opts *ScrapeOpts) *Client {
	if opts == nil {
		opts = &ScrapeOpts{
			Domain:                Domain,
			ArtistNameCssSelector: ArtistNameCssSelector,
			SongNameCssSelector:   SongNameCssSelector,
			SearchUrl:             SearchUrl,
			LyricCssSelector:      LyricCssSelector,
		}
	}

	if collector == nil {
		collector = colly.NewCollector(
			colly.AllowedDomains(opts.Domain),
		)
	}

	return &Client{
		Collector:  collector,
		Opts:       opts,
		loggerFunc: opts.Logger,
	}
}
