package config

import "github.com/joho/godotenv"

func loadENV() (e error) {
	e = godotenv.Load()
	return
}
