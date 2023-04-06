package main

import (
	"log"

	"github.com/alyumi/music_searcher/spotify"
)

func main() {

	s := spotify.NewClient()
	s.ReceiveURL("https://open.spotify.com/track/4r13d29427UZ9lyGrhKjxJ?si=1536726afa7142fc")
	s.ChooseMethod("get")
	searchForYoutube := s.FormSearch()
	log.Println(searchForYoutube)
}
