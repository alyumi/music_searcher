package spotify

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

const (
	accounts_url = "https://accounts.spotify.com/api/token"
)

type AccessData struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func (ad *AccessData) GetAccessData() {

	var (
		client        = &http.Client{}
		client_id     = os.Getenv("SPOTIFY_CLIENT_ID")
		client_secret = os.Getenv("SPOTIFY_CLIENT_SECRET")
		grant_type    = "client_credentials"
	)
	req, err := http.NewRequest("POST", accounts_url, nil)
	if err != nil {
		log.Println(err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.URL.Query().Add("grant_type", grant_type)
	req.URL.Query().Add("client_id", client_id)
	req.URL.Query().Add("client_secret", client_secret)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("cannot send request to gain access data:", err)
	}
	defer resp.Body.Close()

	data := &AccessData{}
	err = json.Unmarshal(resp.Body, data)
}
