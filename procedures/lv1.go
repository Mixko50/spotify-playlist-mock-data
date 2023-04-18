package procedures

import (
	"log"
	"mock_spotify_data/database"
	model "mock_spotify_data/model/db"
	"mock_spotify_data/utils/text"
	"strconv"
	"sync"
)

func Level1Operation(wg *sync.WaitGroup) {
	for _, item := range *CurrentData.Items {
		releaseYear, err := strconv.ParseInt(item.Track.Album.ReleaseDate[0:4], 10, 64)
		if err != nil {
			database.DropAllTables()
			log.Fatal(err)
		}
		randomCount := int64(text.RandomNumber(200)) + 1
		if result := database.Gorm.Create(&model.Lv1Track{
			SpotifyId:  item.Track.Id,
			Name:       item.Track.Name,
			Album:      &item.Track.Album.Name,
			Artist:     &item.Track.Artists[0].Name,
			ArtworkUrl: item.Track.Album.Images[0].Url,
			Duration:   item.Track.DurationMs,
			Popularity: item.Track.Popularity,
			Explicit:   item.Track.Explicit,
			Year:       &releaseYear,
			Count:      &randomCount,
		}); result.Error != nil {
			database.DropAllTables()
			log.Fatal(result.Error)
		}
	}
	wg.Done()
}
