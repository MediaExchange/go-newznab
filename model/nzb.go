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

package model

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"golang.org/x/net/html/charset"
)

// Nzb files contain an optional header and one or more File entries.
type Nzb struct {
	XMLName xml.Name	`xml:"nzb" json:"-"`
	Head	Head		`xml:"head" json:"head,omitempty"`
	File 	[]File		`xml:"file" json:"file"`
}

// Head contains zero or more Meta structs.
type Head struct {
	XMLName	xml.Name 	`xml:"head" json:"-"`
	Meta    []Meta      `xml:"meta" json:"meta,omitempty"`
}

// Meta structs are optional, but contain key-value pairs that describe the content of the Nzb.
type Meta struct {
	XMLName xml.Name	`xml:"meta" json:"-"`
	Type 	string		`xml:"type,attr" json:"type,omitempty"`
	Value 	string		`xml:",chardata" json:"value,omitempty"`
}

// File describes a single file available for download.
type File struct {
	XMLName 	xml.Name	`xml:"file" json:"-"`
	Poster  	string		`xml:"poster,attr" json:"poster"`
	Date    	string		`xml:"date,attr" json:"date"`
	Subject 	string		`xml:"subject,attr" json:"subject"`
	Groups  	Groups		`xml:"groups" json:"group"`
	Segments	Segments	`xml:"segments" json:"segment"`
}

// Groups contains one or more Group structs.
type Groups struct {
	XMLName xml.Name	`xml:"groups" json:"-"`
	Group   []Group		`xml:"group" json:"group"`
}

// Group contains the name of a Usenet new group that the File is available from.
type Group struct {
	XMLName xml.Name	`xml:"group" json:"-"`
	Value   string		`xml:",chardata" json:"value"`
}

// Segments contains each Segment that may be downloaded.
type Segments struct {
	XMLName xml.Name	`xml:"segments" json:"-"`
	Segment []Segment	`xml:"segment" json:"segment"`
}

// Segment contains information about a piece of the File that can be downloaded from a news group.
type Segment struct {
	XMLName   xml.Name `xml:"segment" json:"-"`
	Bytes     string   `xml:"bytes,attr" json:"bytes"`
	Number    string   `xml:"number,attr" json:"number"`
	MessageId string   `xml:",chardata" json:"name"`
}

// MarshalJSON converts a Groups struct to an array of group names for better presentation in JSON format.
func (g Groups) MarshalJSON() ([]byte, error) {
	numGroups := len(g.Group)
	groupVals := make([]string, numGroups)
	for i, group := range g.Group {
		groupVals[i] = group.Value
	}
	return json.Marshal(groupVals)
}

// MarshalJSON converts a Segments struct to an array of unnamed Segment objects for better presentation in JSON format.
func (s Segments) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Segment)
}

// fromXML decodes NZB XML content to an Nzb struct.
func fromXml(data []byte) (nzb Nzb, err error) {
	// Some NZB files use iso-8859-1 encoding instead of UTF-8. This
	// implementation was taken from https://stackoverflow.com/a/32224438
	reader := bytes.NewReader(data)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&nzb)
	return
}
