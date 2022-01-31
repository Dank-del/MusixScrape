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

import "net/http"

// New builds a Client with default settings. If you want to change them,
// change them after calling New or construct the Client manually
func New() *Client {
	return &Client{
		SearchUrl: DefaultSearchUrl,
		UserAgent: DefaultUserAgent,
		HttpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
	}
}

func stringInArray(array []string, str string) bool {
	for _, i := range array {
		if i == str {
			return true
		}
	}
	return false
}
