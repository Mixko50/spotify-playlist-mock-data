package main

import (
	"mock_spotify_data/database"
	"mock_spotify_data/procedures"
)

func main() {
	database.Init()
	procedures.RandomData()

	procedures.Level1Operation()
	procedures.Level2Operation()
	procedures.Level3Operation()
}
