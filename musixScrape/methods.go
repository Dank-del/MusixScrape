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
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func (c *Client) GetLyrics(url *url.URL) (Lyrics, error) {
	lyrics := Lyrics{
		Artist: "",
		Song:   "",
		Lyrics: "",
	}
	startParsingSong := false
	startParsingArtist := false
	startParsingLyrics := false
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return lyrics, err
	}
	req.Header.Set("User-Agent", c.UserAgent)
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return lyrics, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return lyrics, fmt.Errorf("musixmatch returned %d", resp.StatusCode)
	}
	z := html.NewTokenizer(resp.Body)
	for {
		z.Next()
		token := z.Token()
		switch token.Type {
		case html.ErrorToken:
			if errors.Is(z.Err(), io.EOF) {
				return lyrics, nil
			}
			return lyrics, z.Err()
		case html.EndTagToken:
			// the song title starts after <small>Lyrics</small>
			startParsingSong = token.DataAtom == atom.Small
			startParsingArtist = false
			startParsingLyrics = false
			break
		case html.TextToken:
			if startParsingSong {
				lyrics.Song += token.Data
			} else if startParsingArtist {
				lyrics.Artist += token.Data
			} else if startParsingLyrics {
				lyrics.Lyrics += token.Data
			}
			break
		case html.StartTagToken:
			if token.DataAtom == atom.A {
				for _, attr := range token.Attr {
					if attr.Key == "class" {
						startParsingArtist = stringInArray(strings.Split(attr.Val, " "), "mxm-track-title__artist")
						if startParsingArtist {
							break
						}
					}
				}
			} else if token.DataAtom == atom.Span {
				for _, attr := range token.Attr {
					if attr.Key == "class" && attr.Val == "lyrics__content__ok" {
						startParsingLyrics = true
						break
					}
				}
			}
		}
	}
}

func (c *Client) Search(query string) ([]SearchResult, error) {
	var results []SearchResult
	result := SearchResult{
		Artist: "",
		Song:   "",
		Url:    nil,
	}
	startParsingResults := false
	startParsingSong := false
	startParsingArtist := false
	// used for song and artist
	temporaryString := ""
	req, err := http.NewRequest("GET", c.SearchUrl+url.PathEscape(query), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", c.UserAgent)
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("musixmatch returned %d", resp.StatusCode)
	}
	z := html.NewTokenizer(resp.Body)
	for {
		z.Next()
		token := z.Token()
		switch token.Type {
		case html.ErrorToken:
			if errors.Is(z.Err(), io.EOF) {
				return results, nil
			}
			return nil, z.Err()
		case html.StartTagToken:
			if token.DataAtom == atom.Ul && !startParsingResults {
				for _, attr := range token.Attr {
					if attr.Key == "class" {
						startParsingResults = stringInArray(strings.Split(attr.Val, " "), "tracks")
						if startParsingResults {
							results = nil
							break
						}
					}
				}
			} else if token.DataAtom == atom.A && startParsingResults {
				href := ""
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						href = attr.Val
					} else if attr.Key == "class" {
						if attr.Val == "title" {
							startParsingSong = true
						} else if attr.Val == "artist" {
							startParsingArtist = true
						}
					}
				}
				if startParsingSong {
					songUrl, err := resp.Request.URL.Parse(href)
					if err != nil {
						return nil, fmt.Errorf("failed to parse song url: %s", err)
					}
					result.Url = songUrl
				}
			}
			break
		case html.TextToken:
			if startParsingSong || startParsingArtist {
				temporaryString += token.Data
			}
			break
		case html.EndTagToken:
			if startParsingSong {
				result.Song = strings.TrimSpace(temporaryString)
				temporaryString = ""
				startParsingSong = false
			} else if startParsingArtist {
				result.Artist = strings.TrimSpace(temporaryString)
				temporaryString = ""
				startParsingArtist = false
			} else if token.DataAtom == atom.Li && startParsingResults {
				results = append(results, result)
				result = SearchResult{
					Artist: "",
					Song:   "",
					Url:    nil,
				}
			} else if token.DataAtom == atom.Ul {
				startParsingResults = false
			}
		}
	}
}
