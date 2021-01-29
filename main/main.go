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

package main

import (
	"errors"
	"flag"
	"fmt"
	newznab "github.com/MediaExchange/nazbaz"
	"os"
	"strings"
)

var (
	Url   string
	Key   string
	Query string
)

func main() {
	flag.StringVar(&Url, "url", "", "Full URL of the newznab API")
	flag.StringVar(&Key, "key", "", "API key for the server")
	flag.StringVar(&Query, "query", "", "Query to send")
	flag.Parse()

	if len(os.Args) == 1 {
		help()
		return
	}

	switch strings.ToLower(os.Args[1]) {
	case "help":
		help()
	case "search":
		search()
	default:
		help()
	}
}

// help displays the program's sub-commands and arguments.
func help() {
	fmt.Println("Usage:")
	fmt.Print('\n')
	fmt.Println("    newznab <command> [arguments]")
	fmt.Print('\n')
	fmt.Println("The commands are:")
	fmt.Print('\n')
	fmt.Println("    search    Run a general search")
	fmt.Print("\n")
	flag.PrintDefaults()
}

// search runs an unrestricted search for anything, but parameters can be used to narrow down the returned results.
func search() {
	params, err := validateAndParse()
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := newznab.Search(Url, Key, params...)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}

func validateAndParse() ([]newznab.Param, error) {
	// Parse the arguments
	err := flag.CommandLine.Parse(os.Args[2:])
	if err != nil {
		return nil, err
	}

	if len(Url) == 0 {
		return nil, errors.New("-url is required")
	}

	if len(Key) == 0 {
		return nil, errors.New("-key is required")
	}

	params := make([]newznab.Param, 0)
	if len(Query) > 0 {
		params = append(params, newznab.Query(Query))
	}

	return params, nil
}
