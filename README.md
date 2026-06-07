# <img src="https://s.mxmcdn.net/site/images/logo_icon.svg" width="45px" align="left"></img>MusixScrape

[![Go Reference](https://pkg.go.dev/badge/github.com/Dank-del/MusixScrape.svg)](https://pkg.go.dev/github.com/Dank-del/MusixScrape)
[![Go-linux](https://github.com/Dank-del/MusixScrape/actions/workflows/go-linux.yml/badge.svg)](https://github.com/Dank-del/MusixScrape/actions/workflows/go-linux.yml)
[![Go-macos](https://github.com/Dank-del/MusixScrape/actions/workflows/go-macos.yml/badge.svg)](https://github.com/Dank-del/MusixScrape/actions/workflows/go-macos.yml)
[![Go-windows](https://github.com/Dank-del/MusixScrape/actions/workflows/go-windows.yml/badge.svg)](https://github.com/Dank-del/MusixScrape/actions/workflows/go-windows.yml)

A small Go library for scraping song lyrics and search results from [musixmatch.com](https://www.musixmatch.com).

## Features

- Fetch lyrics from a Musixmatch song URL.
- Search Musixmatch by free-text query and get a list of `Title / URL` matches.
- Zero third-party scraping framework dependencies, just `net/http` and `goquery`.
- CI-tested on Linux, macOS, and Windows.

## Install

```sh
go get github.com/Dank-del/MusixScrape/musixScrape
```

## Usage

```go
package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/Dank-del/MusixScrape/musixScrape"
)

func main() {
	c := musixScrape.New()

	// Fetch lyrics from a song URL.
	songURL, err := url.Parse("https://www.musixmatch.com/lyrics/BAND-MAID/Sense")
	if err != nil {
		log.Fatal(err)
	}
	lyrics, err := c.GetLyrics(songURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lyrics.Artist, lyrics.Title)

	// Search by query.
	results, err := c.Search("Yuu Miyashita - Koufukuron")
	if err != nil {
		log.Fatal(err)
	}
	for _, r := range results {
		fmt.Println(r)
	}
}
```

## License

[![GPLv3](https://www.gnu.org/graphics/gplv3-127x51.png)](https://www.gnu.org/licenses/gpl-3.0.en.html)

Licensed under the [GNU General Public License v3](https://www.gnu.org/licenses/gpl-3.0.en.html).
