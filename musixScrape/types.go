/*
   MusixScrape - A Golang library to scrape lyrics from musixmatch.com
   Copyright (C) 2021  Sayan Biswas, ALiwoto

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

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
	Opts       *ScrapeOpts
	loggerFunc func(v ...interface{})
}

type ScrapeOpts struct {
	Domain                 string
	SearchUrl              string
	LyricCssSelector       string
	SongNameCssSelector    string
	ArtistNameCssSelector  string
	SearchOuterCssSelector string
	SearchInnerCssSelector string
	Collector              *colly.Collector
	Logger                 func(v ...interface{})
}
