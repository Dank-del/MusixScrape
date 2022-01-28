package MusixScrape

import "testing"

func TestSearchByLink(t *testing.T) {
	res, err := GetLyricsFromLink("https://www.musixmatch.com/lyrics/BAND-MAID/Sense")
	if err != nil {
		t.Error(err)
	}
	if res != nil {
		t.Log(res)
	}

}

func TestSearch(t *testing.T) {
	res, err := Search("Band Maid - Sense")
	if err != nil {
		t.Error(err)
	}
	if res != nil {
		t.Log(res)
	}
}
