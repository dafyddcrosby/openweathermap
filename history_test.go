// Copyright 2015 Brian J. Downs
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package openweathermap

import (
	"net/http"
	"reflect"
	"testing"
	"time"
)

// TestNewHistory verifies NewHistorical does as advertised
func TestNewHistory(t *testing.T) {
	t.Parallel()

	for d := range DataUnits {
		t.Logf("Data unit: %s", d)

		if ValidDataUnit(d) {
			c, err := NewHistorical(d)
			if err != nil {
				t.Error(err)
			}
			if reflect.TypeOf(c).String() != "*openweathermap.HistoricalWeatherData" {
				t.Error("incorrect data type returned")
			}
		} else {
			t.Errorf("unusable data unit - %s", d)
		}
	}

	_, err := NewHistorical("asdf")
	if err == nil {
		t.Error("created instance when it shouldn't have")
	}
}

// TestNewHistoryWithCustomHttpClient will verify that a new instance of HistoricalWeatherData
// is created with custom http client
func TestNewHistoryWithCustomHttpClient(t *testing.T) {

	hc := http.DefaultClient
	hc.Timeout = time.Duration(1) * time.Second
	h, err := NewHistorical("c", WithHttpClient(hc))
	if err != nil {
		t.Error(err)
	}

	if reflect.TypeOf(h).String() != "*openweathermap.HistoricalWeatherData" {
		t.Error("incorrect data type returned")
	}

	expected := time.Duration(1) * time.Second
	if h.client.Timeout != expected {
		t.Errorf("Expected Duration %v, but got %v", expected, h.client.Timeout)
	}
}

// TestNewHistoryWithInvalidHttpClient will verify that returns an error with
// invalid http client
func TestNewHistoryWithInvalidHttpClient(t *testing.T) {

	h, err := NewHistorical("c", WithHttpClient(nil))
	if err != nil {
		t.Logf("Received expected bad client error. message: %s", err.Error())
	}
	if h != nil {
		t.Log("Expected nil, but got %v", h)
	}
}

// TestHistoryByName
func TestHistoryByName(t *testing.T) {
	t.Parallel()
	h, err := NewHistorical("F")
	if err != nil {
		t.Error(err)
	}
	if err := h.HistoryByName("Vancouver"); err != nil {
		t.Error(err)
	}
}

// TestHistoryByID
func TestHistoryByID(t *testing.T) {
	t.Parallel()
	h, err := NewHistorical("F")
	if err != nil {
		t.Error(err)
	}
	hp := &HistoricalParameters{
		Start: 1461598510,
		End:   1461588510,
		Cnt:   1,
	}
	if err := h.HistoryByID(5344157, hp); err != nil {
		t.Error(err)
	}
}

// TestHistoryByCoord
func TestHistoryByCoord(t *testing.T) {
	t.Parallel()
	h, err := NewHistorical("F")
	if err != nil {
		t.Error(err)
	}
	coords := &Coordinates{
		Longitude: -112.07,
		Latitude:  33.45,
	}
	hp := &HistoricalParameters{
		Start: 1461598510,
		End:   1461588510,
		Cnt:   1,
	}
	if err := h.HistoryByCoord(coords, hp); err != nil {
		t.Error(err)
	}
}
