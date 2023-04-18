package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	YOUTUBE_API           string
	SPOTIFY_CLIENT_SECRET string
	SPOTIFY_CLIENT_ID     string
}

func NewConfig() *Config {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error occured while loading .env file: ", err)
	}

	return &Config{
		YOUTUBE_API:           os.Getenv("YOUTUBE_API"),
		SPOTIFY_CLIENT_SECRET: os.Getenv("SPOTIFY_CLIENT_SECRET"),
		SPOTIFY_CLIENT_ID:     os.Getenv("SPOTIFY_CLIENT_ID"),
	}
}
