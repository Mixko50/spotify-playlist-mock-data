package main

import (
	"mock_spotify_data/database"
	"mock_spotify_data/procedures"
	"sync"
)

func main() {
	database.Init()
	procedures.RandomData()

	var wg sync.WaitGroup
	wg.Add(3)
	go procedures.Level1Operation(&wg)
	go procedures.Level2Operation(&wg)
	go procedures.Level3Operation(&wg)
	wg.Wait()
}
