package tests_test

import (
	"testing"

	"github.com/Dank-del/MusixScrape/musixScrape"
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
	res, err := c.Search("麻婆豆腐 - シロガネ")
	if err != nil {
		t.Error(err)
		return
	}

	if res != nil {
		t.Log(res)
	}
}
