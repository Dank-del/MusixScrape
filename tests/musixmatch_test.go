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
	"log"
	"testing"

	"github.com/Dank-del/MusixScrape/musixScrape"
)

var (
	_songNames = []string{
		"麻婆豆腐 - シロガネ",
		"Yuu Miyashita - Koufukuron",
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
	c := musixScrape.New(nil)
	res, err := c.GetLyricsFromLink("https://www.musixmatch.com/lyrics/BAND-MAID/Sense")
	if err != nil {
		t.Error(err)
	}
	if res != nil {
		t.Log(res)
	}

}

func TestSearch(t *testing.T) {
	c := musixScrape.New(nil)
	var res []musixScrape.LyricResult
	var err error

	for _, current := range _songNames {
		res, err = c.Search(current)
		if err != nil {
			t.Error("Error while searching for", current, err)
			return
		}

		if len(res) == 0 || res[0].Song == "" {
			t.Error("No results found for", current)
			return
		}

		log.Println("Found lyrics for " + current + ":" + res[0].Lyrics)
	}

	for _, current := range _aimerSongNames {
		res, err = c.Search(current)
		if err != nil {
			t.Error("Error while searching for aimer's songs:", current, err)
			return
		}

		if len(res) == 0 || res[0].Song == "" {
			t.Error("No results found for:", current)
			return
		}
	}
}
