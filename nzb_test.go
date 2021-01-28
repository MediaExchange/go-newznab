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
	"github.com/MediaExchange/assert"
	"io/ioutil"
	"testing"
)

func TestShortFromXml(t *testing.T) {
	original, err := ioutil.ReadFile("testdata/nzb-short.xml")
	if err != nil {
		t.Error(err)
	}

	nzb, err := fromXml(original)
	assert.With(t).That(err).IsNil()
	assert.With(t).That(len(nzb.File)).IsEqualTo(2)
	assert.With(t).That(len(nzb.File[0].Groups.Group)).IsEqualTo(1)
	assert.With(t).That(len(nzb.File[1].Groups.Group)).IsEqualTo(1)
	assert.With(t).That(len(nzb.File[0].Segments.Segment)).IsEqualTo(1)
	assert.With(t).That(len(nzb.File[1].Segments.Segment)).IsEqualTo(1)

	expected, err := ioutil.ReadFile("testdata/nzb-short.json")
	if err != nil {
		t.Error(err)
	}

	actual, _ := json.MarshalIndent(nzb, "", "  ")
	assert.With(t).That(string(actual)).IsEqualTo(string(expected))
}

func TestLongFromXml(t *testing.T) {
	original, err := ioutil.ReadFile("testdata/nzb-long.xml")
	if err != nil {
		t.Error(err)
	}

	nzb, err := fromXml(original)

	assert.With(t).That(err).IsNil()
	assert.With(t).That(len(nzb.File)).IsEqualTo(55)
	assert.With(t).That(len(nzb.File[0].Groups.Group)).IsEqualTo(1)
	assert.With(t).That(len(nzb.File[54].Groups.Group)).IsEqualTo(1)
	assert.With(t).That(len(nzb.File[0].Segments.Segment)).IsEqualTo(1)
	assert.With(t).That(len(nzb.File[54].Segments.Segment)).IsEqualTo(42)

	expected, err := ioutil.ReadFile("testdata/nzb-long.json")
	if err != nil {
		t.Error(err)
	}

	actual, _ := json.MarshalIndent(nzb, "", "  ")
	assert.With(t).That(string(actual)).IsEqualTo(string(expected))
}
