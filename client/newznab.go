/*
   Copyright 2020 MediaExchange.io

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

// Package client provides HTTP clients for Newznab servers.
package client

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/MediaExchange/nazbaz/model"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Caps returns the server capabilities.
func Caps(capabilities model.Capabilities) (string, error) {
	// Build the URL to request from.
	u, err := EncodeUrl(capabilities.Api.Url, "t", "caps", "o", "json")
	if err != nil {
		return "", err
	}

	return Execute(u)
}

func GetNzb(m model.Get) (string, error) {
	u, err := EncodeUrl(m.Api.Url,
		"apikey", m.Api.Key,
		"t", "get",
		"o", "json",
		"id", m.Id)
	if err != nil {
		return "", err
	}

	// Retrieve the NZB file.
	body, err := Execute(u)
	if err != nil {
		return "", err
	}

	// Unmarshal the NZB.
	var nzb model.Nzb
	err = xml.Unmarshal([]byte(body), &nzb)
	if err != nil {
		return "", err
	}

	// Marshal to JSON for the client.
	b, err := json.Marshal(nzb)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// MovieSearch runs a search for movies
func MovieSearch(m model.MovieSearch) (string, error) {
	// Build the URL to request from.
	cat := ""
	if len(m.Categories) > 0 {
		cat = strings.Join(m.Categories, ",")
	}
	u, err := EncodeUrl(m.Api.Url,
		"apikey", m.Api.Key,
		"t", "movie",
		"o", "json",
		"extended", "1",
		"q", m.Query,
		"cat", cat,
		"imdbid", m.IMDBID)
	if err != nil {
		return "", err
	}

	return Execute(u)
}

// Search runs a general search.
func Search(m model.Search) (string, error) {
	// Build the URL to request from.
	cat := ""
	if len(m.Categories) > 0 {
		cat = strings.Join(m.Categories, ",")
	}
	u, err := EncodeUrl(m.Api.Url,
		"apikey", m.Api.Key,
		"t", "search",
		"o", "json",
		"extended", "1",
		"q", m.Query,
		"cat", cat)
	if err != nil {
		return "", err
	}
	fmt.Println(u)

	return Execute(u)
}

// TvSearch runs a search for TV shows.
func TvSearch(m model.TvSearch) (string, error) {
	// Build the URL to request from.
	cat := ""
	if len(m.Categories) > 0 {
		cat = strings.Join(m.Categories, ",")
	}
	u, err := EncodeUrl(m.Api.Url,
		"apikey", m.Api.Key,
		"t", "tvsearch",
		"o", "json",
		"extended", "1",
		"q", m.Query,
		"cat", cat,
		"season", m.Season,
		"ep", m.Episode,
		"rid", m.TVRageId,
		"tvdbid", m.TVDBID,
		"traktid", m.TraktId,
		"tvmazeid", m.TVMazeId,
		"imdbid", m.IMDBID,
		"tmdbid", m.TMDBID)
	if err != nil {
		return "", err
	}

	return Execute(u)
}

// Execute accepts the constructed URL and performs a GET operation.
func Execute(u *url.URL) (string, error) {
	// Run the request
	res, err := http.Get(u.String())
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	// Bail out now if the status isn't OK.
	if res.StatusCode != http.StatusOK {
		return "", errors.New(res.Status)
	}

	// Return the response body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// EncodeUrl returns a URL with a properly encoded query string.
func EncodeUrl(base string, params ...string) (*url.URL, error) {
	// Must have an even number of query parameters for q=s form.
	if len(params) % 2 != 0 {
		return nil, errors.New("incorrect number of parameters to EncodeUrl")
	}

	// Parse the base URL.
	u, err := url.Parse(base)
	if err != nil {
		return nil, err
	}

	// Build the query parameter hash.
	q := url.Values{}
	for i := 0; i < len(params); i += 2 {
		key := params[i]
		val := params[i+1]
		if len(val) > 0 {
			q.Add(key, val)
		}
	}

	// Add the encoded query string to the URL and return it.
	u.RawQuery = q.Encode()
	return u, nil
}
