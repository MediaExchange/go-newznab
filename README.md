# go-newznab

[![GoDoc](https://godoc.org/github.com/mediaexchange/nazbaz/github?status.svg)](https://godoc.org/github.com/mediaexchange/nazbaz)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue.svg)](https://www.apache.org/licenses/LICENSE-2.0)
[![Go version](https://img.shields.io/badge/go-~%3E1.15-green.svg)](https://golang.org/doc/devel/release.html#go1.15)

`go-newznab` is an HTTP client library for NZB indexers.

## Usage

This library was purposely designed to be stateless so that it could more
easily by used inside a web service. There are other implementation of NZB
and Newznab clients that are initialized once with the API URL and access
key. That works well if many calls will be made to the same service, but in
a web service which is typically stateless, this pattern is difficult to
work with.

Instead, the API URL and access key are passed into each call, then the
various parameters used to run a search are added as needed. The pattern used
is called "functional options" and was first publicized for Go by Dave
Cheney in an presentation titled "Functional options for friendly APIs."
This has since been [published on his blog
](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis)

Newznab servers support running generalized and specific searches. Various
parameters may be added to each to further restrict the results returned
from the service.

Search for "The Terminator" and restrict to HD (2k) and UHD (4k) media.

```go
res, err := newznab.Search("http://example.com/api", "my-api-key",
	newznab.Query("The Terminator"),
	newznab.Categories(newznab.Movies_HD, newznab.Movies_UHD))
```

Search for Season 2, Episode 22 of "The Office" in HD.

```go
res, err := newznab.TvSearch("http://example.com/api", "my-api-key",
	newznab.Query("The Office"),
	newznab.Categories(newznab.TV_HD),
	newznab.Season(2),
	newznab.Episode(22))
```

Search for "The Great Gatsby", 2013 release, by its IMDB ID, in UHD.

```go
res, err := newznab.MovieSearch("http://example.com/api", "my-api-key",
	newznab.ImdbId(1343092),
	newznab.Categories(newznab.Movies_UHD))
```

Download an NZB file:

```go
res, err := newznab.GetNzb("http://example.com/api", "my-api-key", "nzb-id")
```

## Contributing

 1.  Fork it
 2.  Create a feature branch (`git checkout -b new-feature`)
 3.  Commit changes (`git commit -am "Added new feature xyz"`)
 4.  Push the branch (`git push origin new-feature`)
 5.  Create a new pull request.

## Maintainers

* [Media Exchange](http://github.com/MediaExchange)

## License

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
