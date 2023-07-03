package youtube

import (
	"context"
	"log"

	"github.com/alyumi/music_searcher/config"
	"google.golang.org/api/option"
	yt "google.golang.org/api/youtube/v3"
)

type YoutubeClient struct {
	Client *yt.Service
	config *Config
}

type Config struct {
	YOUTUBE_API string
}

const (
	VideoURL = "https://www.youtube.com/watch/?v="
)

func NewClient(conf config.Config) *YoutubeClient {
	var (
		ctx    = context.Background()
		apiKey = conf.YOUTUBE_API
	)
	client, err := yt.NewService(ctx, option.WithAPIKey(apiKey))

	if err != nil {
		log.Println("couldn't create youtube service", err)
	}

	return &YoutubeClient{
		Client: client,
		config: &Config{
			YOUTUBE_API: apiKey,
		},
	}
}
