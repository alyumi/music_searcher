package spotify

import (
	"errors"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/alyumi/music_searcher/config"
	"github.com/alyumi/music_searcher/spotify/methods"
)

type Client struct {
	ad     AccessData
	pd     PathData
	td     TrackData
	config Config
}

type Config struct {
	SPOTIFY_CLIENT_ID     string
	SPOTIFY_CLIENT_SECRET string
}

type TrackData struct {
	Name   string
	Artist string
	Album  string
}

func NewClient(config config.Config) *Client {

	client := &Client{
		config: Config{
			SPOTIFY_CLIENT_ID:     config.SPOTIFY_CLIENT_ID,
			SPOTIFY_CLIENT_SECRET: config.SPOTIFY_CLIENT_SECRET,
		},
	}

	if IsAccessDataValid() {
		log.Println("Access data valid and exists")
		client.ad = *UseTempData()
	} else {
		log.Println("Access data not valid or not exists")
		client.ad.GetAccessData(client.config)
	}

	return client
}

func (c Client) GetHeader() http.Header {
	header := http.Header{}
	value := c.ad.TokenType + "  " + c.ad.AccessToken
	header.Add("Authorization", value)
	return header
}

func (c Client) parseURL(rawURL string) (*PathData, error) {
	raw, err := url.Parse(rawURL)
	if err != nil {
		log.Println(err)
	}
	pd := &PathData{}

	if checkURL(raw) {
		pd.getPathData(raw.Path)
		return pd, nil
	} else {
		return pd, errors.New("wrong URL")
	}

}

func (c *Client) ReceiveURL(rawURL string) {

	ans, err := c.parseURL(rawURL)

	if err != nil {
		log.Fatal("Fatal error occured:", err)
	}

	c.pd.ID = ans.ID
	c.pd.Name = ans.Name

}

func checkURL(URL *url.URL) bool {
	var (
		host   = "open.spotify.com"
		scheme = "https"
	)

	if scheme != URL.Scheme {
		return false
	}
	if host != URL.Host {
		return false
	}

	return true
}

type Getter interface {
	Get(id string, c http.Header)
}

// Implement
func (c *Client) ChooseMethod(method string) {

	t := &methods.Track{}

	if strings.ToLower(method) == methods.Get {
		header := c.GetHeader()

		id := c.pd.ID
		t = t.Get(id, header)
	}

	artist := formArtists(t.Artists)
	c.td.Artist = artist
	c.td.Name = t.Name
	c.td.Album = t.AlbumName.Name
}

func (c Client) FormSearch(URL string, method string) string {
	c.ReceiveURL(URL)
	c.ChooseMethod(method)

	return c.td.Name + " " + c.td.Artist
}

func formArtists(artists []methods.Artist) string {
	var ans string
	for _, artist := range artists {
		if ans != "" {
			ans = ans + ", " + artist.Name
		} else {
			ans = artist.Name
		}
	}

	return ans
}

type Variant struct {
	Name   string
	Method string
}

func NewVariant(name, method string) *Variant {
	return &Variant{
		Name:   name,
		Method: method,
	}
}

func (v Variant) NewTrack() *methods.Track {
	return &methods.Track{}
}

func (v Variant) NewGetter() Getter {
	var g Getter
	return g
}
