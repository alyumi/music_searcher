package methods

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type Track struct {
	ID string `json:"id"`

	Name      string   `json:"name"`
	AlbumName Album    `json:"album"`
	Artists   []Artist `json:"artists"`
}

func (t Track) Get(id string, h http.Header) *Track {
	return t.get(id, h)
}

func (t Track) get(id string, h http.Header) *Track {
	var (
		track    = &Track{}
		client   = &http.Client{}
		endpoint = "https://api.spotify.com/v1/tracks/" + id
	)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Fatal("error occured", err)
	}

	req.Header = h

	resp, err := client.Do(req)
	if resp.StatusCode != 200 {
		log.Fatal(resp.StatusCode)
	}
	if err != nil {
		log.Fatal(err)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("client: could not read response body: %s\n", err)
	}

	err = json.Unmarshal(responseBody, track)
	if err != nil {
		log.Println(err)
	}

	//t = track
	// t.AlbumName = track.AlbumName
	// t.Artists = track.Artists
	// t.ID = track.ID
	// t.Name = track.Name

	os.WriteFile("track_data.json", responseBody, 0600)
	return track
}
