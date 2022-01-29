package musixScrape

const (
	Domain                = "www.musixmatch.com"
	SearchUrl             = "https://" + Domain + "/search/"
	LyricCssSelector      = `p.mxm-lyrics__content:nth-child(2) > span:nth-child(1)`
	SongNameCssSelector   = ".mxm-track-title__track"
	ArtistNameCssSelector = `.mxm-track-title__artist`
	SearchOuterSelector   = `div.box:nth-child(2) > div:nth-child(2) > div:nth-child(1) > ul:nth-child(1)`
	SearchInnerSelector   = `div.box:nth-child(2) > div:nth-child(2) > div:nth-child(1) > ul:nth-child(1) > li:nth-child(1) > div:nth-child(1) > div:nth-child(2) > div:nth-child(1) > h2:nth-child(1) > a:nth-child(1)`
)
