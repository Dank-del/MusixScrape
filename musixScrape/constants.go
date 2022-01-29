package musixScrape

const (
	Domain                = "www.musixmatch.com"
	SearchUrl             = "https://" + Domain + "/search/"
	LyricCssSelector      = `p.mxm-lyrics__content:nth-child(2) > span:nth-child(1)`
	SongNameCssSelector   = ".mxm-track-title__track"
	ArtistNameCssSelector = `.mxm-track-title__artist`
)
