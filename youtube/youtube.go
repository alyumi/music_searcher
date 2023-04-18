package youtube

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/alyumi/music_searcher/config"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	yt "google.golang.org/api/youtube/v3"
)

type Service struct {
	config *Config
}

type Config struct {
	API string
}

func NewConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	var (
		apiKey = os.Getenv("YOUTUBE_API")
	)

	return &Config{
		API: apiKey,
	}
}

func NewService() *Service {
	return &Service{
		config: NewConfig(),
	}
}

func YoutubeSearchService(ctx context.Context, conf config.Config, searchQuery []string) []*yt.SearchResult {

	apiKey := conf.YOUTUBE_API
	client, err := yt.NewService(ctx, option.WithAPIKey(apiKey))

	if err != nil {
		log.Println("Youtube service error:", err)
	}

	call := client.Search.List([]string{"id,snippet"}).Q(searchQuery[0])
	response, err := call.Do()
	if err != nil {
		log.Println("error while making call to youtube search api")
	}

	return response.Items
}

func GetSearchLinks(searchList []*yt.SearchResult) []string {
	var workingLinks, unchekedLinks []string

	for _, v := range searchList {
		unchekedLinks = append(unchekedLinks, v.Id.VideoId)
	}

	links := make(chan []string, 1)
	defer close(links)
	log.Println("I'm here")

	go findWorkingLinks(unchekedLinks, links)

	for _, v := range <-links {
		workingLinks = append(workingLinks, formYtLink(v))
	}

	return workingLinks
}

func findWorkingLinks(unchekedLinks []string, links chan []string) {

	var checkedLinks []string

	for _, link := range unchekedLinks {
		resp, err := http.Get(formYtLink(link))
		if err != nil {
			log.Printf("Error occured %s", err)
		}
		if resp.StatusCode == http.StatusOK {
			log.Println(link)
			checkedLinks = append(checkedLinks, link)
		} else {
			continue
		}
	}

	links <- checkedLinks

}

func formYtLink(id string) string {
	URL := "https://www.youtube.com/watch/?v="
	return URL + id
}
