package youtube

import (
	"context"
	"log"
	"net/http"

	yt "google.golang.org/api/youtube/v3"
)

func (ytClient YoutubeClient) YoutubeSearchService(ctx context.Context, searchQuery string) []*yt.SearchResult {
	log.Print("Searching in youtube...")
	searchService := ytClient.Client.Search
	call := searchService.List([]string{"id,snippet"}).Q(searchQuery)
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
