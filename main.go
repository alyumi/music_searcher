package main

import (
	"context"
	"log"

	"github.com/alyumi/music_searcher/config"
	"github.com/alyumi/music_searcher/spotify"
	yt "github.com/alyumi/music_searcher/youtube"
)

func main() {

	conf := config.NewConfig()
	s := spotify.NewClient(*conf)

	URL := "https://open.spotify.com/track/4r13d29427UZ9lyGrhKjxJ?si=1536726afa7142fc"
	searchForYoutube := s.FormSearch(URL, "get")
	var ss []string
	ss = append(ss, searchForYoutube)

	searchCall := yt.YoutubeSearchService(context.Background(), *conf, []string(ss))
	links := yt.GetSearchLinks(searchCall)

	for _, link := range links {
		log.Println(link)
	}
}
