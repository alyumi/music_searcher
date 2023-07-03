package app

import (
	"context"
	"log"

	"github.com/alyumi/music_searcher/config"
	"github.com/alyumi/music_searcher/spotify"
	yt "github.com/alyumi/music_searcher/youtube"
)

type App struct {
	s *spotify.Client
	y *yt.YoutubeClient
}

func InitApp() *App {
	conf := config.NewConfig()
	spoti := spotify.NewClient(*conf)
	ytClient := yt.NewClient(*conf)

	return &App{
		s: spoti,
		y: ytClient,
	}

}

func (a App) FindLinks(URL string) []string {
	// URL := "https://open.spotify.com/track/4r13d29427UZ9lyGrhKjxJ?si=1536726afa7142fc"
	log.Print("searching links...")
	searchForYoutube := a.s.FormSearch(URL, "get")
	log.Println("\n")
	searchResult := a.y.YoutubeSearchService(context.Background(), searchForYoutube)
	links := yt.GetSearchLinks(searchResult)

	return links
}
