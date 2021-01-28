/*
   Copyright 2021 MediaExchange.io

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

package newznab

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	Execute = execute
)

// BookSearch performs a search restricted to e-books.
func BookSearch(url string, key string, params ...Param) (string, error) {
	p := append(params, extended(), Apikey(key), Type("book"))
	u, err := EncodeUrl(url, p...)
	if err != nil {
		return "", err
	}

	return Execute(u)
}

// GetCapabilities returns the capabilities of the server.
func GetCapabilities(url string) (string, error) {
	// Build the URL to request from.
	u, err := EncodeUrl(url, Type("caps"))

	if err != nil {
		return "", err
	}

	return Execute(u)
}

func GetNzb(url string, key string, id string) (string, error) {
	u, err := EncodeUrl(url, Apikey(key), nzbid(id), Type("get"))
	if err != nil {
		return "", err
	}

	// Retrieve the NZB file.
	body, err := Execute(u)
	if err != nil {
		return "", err
	}

	// Unmarshal the NZB.
	var nzb Nzb
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

// MovieSearch performs a search restricted to movies.
func MovieSearch(url string, key string, params ...Param) (string, error) {
	p := append(params, extended(), Apikey(key), Type("movie"))
	u, err := EncodeUrl(url, p...)
	if err != nil {
		return "", err
	}

	return Execute(u)
}

// MusicSearch performs a search restricted to music.
func MusicSearch(url string, key string, params ...Param) (string, error) {
	p := append(params, extended(), Apikey(key), Type("music"))
	u, err := EncodeUrl(url, p...)
	if err != nil {
		return "", err
	}

	return Execute(u)
}

// Search performs a general search which can include any of media.
func Search(url string, key string, params ...Param) (string, error) {
	p := append(params, extended(), Apikey(key), Type("search"))
	u, err := EncodeUrl(url, p...)
	if err != nil {
		return "", err
	}

	return Execute(u)
}

// TvSearch performs a search restricted to TV shows.
func TvSearch(url string, key string, params ...Param) (string, error) {
	p := append(params, extended(), Apikey(key), Type("tvsearch"))
	u, err := EncodeUrl(url, p...)
	if err != nil {
		return "", err
	}

	return Execute(u)
}

// EncodeUrl returns a URL with a properly encoded query string.
func EncodeUrl(base string, params ...Param) (*url.URL, error) {
	// Parse the base URL.
	u, err := url.Parse(base)
	if err != nil {
		return nil, err
	}

	q := url.Values{}
	for _, param := range params {
		q.Add(param.Name, param.Value)
	}

	u.RawQuery = q.Encode()
	return u, nil
}

// Execute accepts the constructed URL and performs a GET operation.
func execute(u *url.URL) (string, error) {
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

// extended returns a Param that directs the service to produce all extended attributes.
// This function is private because it's included with every call. There is no need for the consumer of this library to specify it.
func extended() Param {
	return Param{
		Name:  "extended",
		Value: "1",
	}
}

func nzbid(id string) Param {
	return Param{
		Name:  "id",
		Value: id,
	}
}
