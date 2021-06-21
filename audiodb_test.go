package audiodb

import (
	"fmt"
	"testing"
)

func TestCreate(t *testing.T) {
	db := NewAudioDB()
	if db == nil {
		t.Fail()
	}
}

func TestSearchTitle(t *testing.T) {
	db := NewAudioDB()
	data := [][]string{
		{"linkin park","numb"},
		{"eminem", "rap god"},
		{"blink-182","down"},
		{"green day","american idiot"},
		{"iron maiden","powerslave"},
		{"ensiferum","token of time"},
	}

	for _,v := range(data) {
		t.Run(v[0] + " - " +v[1],func(t *testing.T) {
			tracks,err := db.SearchTitle(v[0],v[1])
			if err != nil {
				t.Error("search failed: ",err)
			} else if len(tracks.Tracks) == 0 {
				t.Error(fmt.Sprintf("no result for %s - %s",v[0],v[1]))
			}
		})
	}
}

func TestSearchArtist(t *testing.T) {
	db := NewAudioDB()
	data := []string{
		"linkin park",
		"eminem",
		"blink 182",
		"green day",
		"iron maiden",
		"ensiferum",
	}

	for _,v := range(data) {
		t.Run(fmt.Sprintf("Searching for %v",v),func(t *testing.T) {
			artists,err := db.SearchArtist(v)
			if err != nil {
				t.Error("search failed: ",err)
			} else if len(artists.Artists) == 0 {
				t.Error("no result for artist ",v)
			}
		})
	}
}