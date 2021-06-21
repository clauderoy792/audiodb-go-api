package main

import (
	"fmt"

	"gitlab.com/claude.roy790/audiodb"
)

func main() {
	db := audiodb.NewAudioDB()
	tracks, err := db.SearchTitle("eminem", "rap god")
	if err == nil {
		for _, t := range tracks.Tracks {
			fmt.Printf("a:%s, t:%s, alb:%s\n", t.Artist, t.Title, t.Album)
		}
	} else {
		fmt.Println("error searhcing title: ", err)
	}

	artists, err := db.SearchArtist("blink 182")
	if err == nil {
		for _, artist := range artists.Artists {
			fmt.Println("artist: ", artist.Name)
		}
	} else {
		fmt.Println("error: ", err)
	}
}
