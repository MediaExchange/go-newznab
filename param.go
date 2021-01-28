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
	"fmt"
	"strconv"
	"strings"
)

// Parameters added to the URL to make a query.
type Param struct {
	Name string
	Value string
}

// Album returns a Param that restricts the search of music to a specific album title.
func Album(a string) Param {
	return Param{
		Name:  "album",
		Value: a,
	}
}

// Apikey returns a Param that defines the key used to access the API.
func Apikey(k string) Param {
	return Param{
		Name:  "apikey",
		Value: k,
	}
}

// Artist returns a Param that restricts the search of music to a specific artist.
func Artist(a string) Param {
	return Param{
		Name:  "artist",
		Value: a,
	}
}

// Author returns a Param that restricts the search of e-books to a specific author.
func Author(a string) Param {
	return Param{
		Name:  "author",
		Value: a,
	}
}

// Categories returns a Param that defines the media categories to restrict the search within.
func Categories(cats ...int) Param {
	// Convert the integers to strings.
	c := make([]string, len(cats))
	for i := range cats {
		c[i] = strconv.Itoa(cats[i])
	}

	return Param{
		Name:  "cat",
		Value: strings.Join(c, ","),
	}
}

// Episode returns a Param that restricts the search of TV shows to a specific episode
func Episode(e int) Param {
	return Param{
		Name:  "episode",
		Value: fmt.Sprintf("E%02d", e),
	}
}

// Genre returns a Param that restricts the search to a specific genre of media
func Genre(g string) Param {
	return Param{
		Name:  "genre",
		Value: g,
	}
}

// ImdbId returns a Param that contains the IMDB ID of the media to search for.
func ImdbId(i int) Param {
	return Param{
		Name:  "imdbid",
		Value: strconv.Itoa(i),
	}
}

// Json returns a Param that directs the service to produce JSON formatted output.
func Json() Param {
	return Param{
		Name:  "o",
		Value: "json",
	}
}

// Label returns a Param that restricts the search of music to a specific publisher or label name.
func Label(l string) Param {
	return Param{
		Name:  "label",
		Value: l,
	}
}

// Limit returns a Param that defines the maximum number of results to return.
func Limit(l int) Param {
	return Param{
		Name:  "limit",
		Value: strconv.Itoa(l),
	}
}

// MaxAge returns a Param that directs the service to return only results that were uploaded in the last "m" days.
func MaxAge(m int) Param {
	return Param{
		Name:  "maxage",
		Value: strconv.Itoa(m),
	}
}

// Offset returns a Param that directs the service to return results starting a the specified offset. This is useful
// when a query would return more results than the service is able to provide in a single response. The consumer of
// this library can retrieve the next batch by re-running the same query, but with an offset.
func Offset(o int) Param {
	return Param{
		Name:  "offset",
		Value: strconv.Itoa(o),
	}
}

// Query returns a Param with the term to search for.
func Query(q string) Param {
	return Param{
		Name:  "q",
		Value: q,
	}
}

// Season returns a Param that restricts the search of TV shows to a specific season
func Season(s int) Param {
	return Param{
		Name:  "season",
		Value: fmt.Sprintf("S%02d", s),
	}
}

// Title returns a Param that restricts the search of e-books to a specific title.
func Title(t string) Param {
	return Param{
		Name:  "title",
		Value: t,
	}
}

// Track returns a Param that restricts the search of music to a specific track name.
func Track(t string) Param {
	return Param{
		Name:  "track",
		Value: t,
	}
}

// Type returns a Param that defines the type of request being made.
func Type(t string) Param {
	return Param{
		Name:  "t",
		Value: t,
	}
}

// Xml returns a Param that directs the service to produce XML formatted output.
func Xml() Param {
	return Param{
		Name:  "o",
		Value: "xml",
	}
}

// Year returns a Param that restricts the search of music to a specific year.
func Year(y string) Param {
	return Param{
		Name:  "year",
		Value: y,
	}
}
