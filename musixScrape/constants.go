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

const (
	Domain                = "www.musixmatch.com"
	SearchUrl             = "https://" + Domain + "/search/"
	LyricCssSelector      = `p.mxm-lyrics__content:nth-child(2) > span:nth-child(1)`
	SongNameCssSelector   = ".mxm-track-title__track"
	ArtistNameCssSelector = `.mxm-track-title__artist`
	SearchOuterSelector   = `div.box:nth-child(2) > div:nth-child(2) > div:nth-child(1) > ul:nth-child(1)`
	SearchInnerSelector   = `div.box:nth-child(2) > div:nth-child(2) > div:nth-child(1) > ul:nth-child(1) > li:nth-child(1) > div:nth-child(1) > div:nth-child(2) > div:nth-child(1) > h2:nth-child(1) > a:nth-child(1)`
)
