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

package tests_test

import (
	"net/url"
	"testing"

	"github.com/Dank-del/MusixScrape/musixScrape"
)

var (
	_songNames = []string{
		"麻婆豆腐 - シロガネ",
		"Yuu Miyashita - Koufukuron",
		"bo en - My Time",
	}
	_aimerSongNames = []string{
		"Aimer – Open The Doors",
		"Aimer – Asa ga kuru",
		"Aimer – I beg you",
		"Aimer – Closer",
		"Aimer – Falling Alone",
		"Aimer – Stars in the Rain",
	}
)

func TestSearchByLink(t *testing.T) {
	c := musixScrape.New()
	url, err := url.Parse("https://www.musixmatch.com/lyrics/BAND-MAID/Sense")
	if err != nil {
		t.Error(err)
		return
	}
	res, err := c.GetLyrics(url)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%s - %s:\n%s", res.Artist, res.Song, res.Lyrics)
}

func TestSearch(t *testing.T) {
	c := musixScrape.New()
	var res []musixScrape.SearchResult
	var lyrics musixScrape.Lyrics
	var err error

	for _, current := range _songNames {
		res, err = c.Search(current)
		if err != nil {
			t.Error("Error while searching for", current, err)
			return
		}

		if len(res) == 0 {
			t.Error("No results found for", current)
			return
		}
		lyrics, err = c.GetLyrics(res[0].Url)
		if err != nil {
			t.Error("Error while getting lyrics for", current, err)
			return
		}
		t.Logf("%s - %s:\n%s", res[0].Artist, res[0].Song, lyrics.Lyrics)
	}

	for _, current := range _aimerSongNames {
		res, err = c.Search(current)
		if err != nil {
			t.Error("Error while searching for aimer's songs:", current, err)
			return
		}

		if len(res) == 0 {
			t.Error("No results found for:", current)
			return
		}
	}
}
