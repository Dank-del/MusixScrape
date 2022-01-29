package musixScrape

func New(opts *ScrapeOpts) *Client {
	if opts == nil {
		opts = DefaultOpts
	}
	return &Client{
		Opts:       opts,
		loggerFunc: opts.Logger,
	}
}
