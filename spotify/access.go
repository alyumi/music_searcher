package spotify

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

const (
	accounts_url = "https://accounts.spotify.com/api/token"
)

// Добавить проверку по времени получения данных
type AccessData struct {
	AccessToken string    `json:"access_token"`
	TokenType   string    `json:"token_type"`
	ExpiresIn   int       `json:"expires_in"`
	ReceiveTime time.Time `json:"receive_time"`
}

type ErrorData struct {
	Err Err `json:"error"`
}

type Err struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (ad *AccessData) GetAccessData() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

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
	q := req.URL.Query()
	q.Add("grant_type", grant_type)
	q.Add("client_id", client_id)
	q.Add("client_secret", client_secret)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if resp.StatusCode != http.StatusOK {
		log.Fatal("ERRORRRRR")
	}
	if err != nil {
		log.Println("cannot send request to gain access data:", err)
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("client: could not read response body: %s\n", err)
	}

	data := &AccessData{}
	data.ReceiveTime = time.Now()
	err = json.Unmarshal(responseBody, data)
	if err != nil {
		log.Println(err)
	}

	ad.AccessToken = data.AccessToken
	ad.TokenType = data.TokenType
	ad.ExpiresIn = data.ExpiresIn
	ad.ReceiveTime = time.Now()

	ioutil.WriteFile("temp_data.json", responseBody, 0600)

}

func readTempFile() (*AccessData, error) {
	accessData := &AccessData{}
	errorData := &ErrorData{}
	data, err := os.ReadFile("temp_data.json")
	if err != nil {
		log.Println("Cannot read temp_data file:", err)
		return accessData, err
	}

	json.Unmarshal(data, errorData)
	status, err := strconv.Atoi(errorData.Err.Status)
	if err != nil {
		if status != 0 && status != 200 {
			log.Println("request error:", errorData.Err.Status)
			return accessData, errors.New("status code: " + errorData.Err.Status)
		}
	}

	json.Unmarshal(data, accessData)

	return accessData, nil
}

func UseTempData() *AccessData {

	accessData, err := readTempFile()
	if err != nil {
		log.Fatal("Error occured:", err)
	}

	return accessData
}

func IsAccessDataValid() bool {

	accessData, err := readTempFile()
	if err != nil {
		log.Println("accessData error:", err)
		return false
	}

	deadlineTime := accessData.ReceiveTime.Add(time.Duration(accessData.ExpiresIn))
	deltaDeadline := time.Since(deadlineTime)

	return deltaDeadline > time.Duration(0)
}
