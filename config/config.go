package config

import "github.com/wctang723/KoreMitai/database"

type ApiConfig struct {
	Myqu           *database.Queries
	Platform       string
	Tokensecretkey string
}
