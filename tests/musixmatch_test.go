package tests_test

import (
	"log"
	"testing"

	"github.com/Dank-del/MusixScrape/musixScrape"
)

var (
	_songNames = []string{
		"麻婆豆腐 - シロガネ",
	}
	_aimerSongNames = []string{
		//"Aimer – Open The Doors",
		//"Aimer – Asa ga kuru",
		//"Aimer – I beg you",
		//"Aimer – Closer",
		"Aimer – Falling Alone",
		"Aimer – Stars in the Rain",
	}
)

func TestSearchByLink(t *testing.T) {
	c := musixScrape.New(nil, nil)
	res, err := c.GetLyricsFromLink("https://www.musixmatch.com/lyrics/BAND-MAID/Sense")
	if err != nil {
		t.Error(err)
	}
	if res != nil {
		t.Log(res)
	}

}

func TestSearch(t *testing.T) {
	c := musixScrape.New(nil, nil)
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
