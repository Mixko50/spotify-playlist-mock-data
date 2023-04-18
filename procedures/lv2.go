package procedures

import (
	"log"
	"mock_spotify_data/database"
	model "mock_spotify_data/model/db"
	"mock_spotify_data/utils/text"
	"strconv"
	"sync"
)

func Level2Operation(wg *sync.WaitGroup) {
	for _, item := range *CurrentData.Items {

		// --------------- Artist ---------------

		var artistModel *model.Lv2Artist

		// Find artist first, if found the record then we don't need to insert it again
		if result := database.Gorm.First(&artistModel, "spotify_id = ?", item.Track.Album.Artists[0].Id); result.RowsAffected == 0 {
			artist := &model.Lv2Artist{
				SpotifyId: &item.Track.Album.Artists[0].Id,
				Name:      &item.Track.Album.Artists[0].Name,
				Href:      &item.Track.Album.Artists[0].Href,
			}

			artistModel = artist

			// Insert
			if insertResult := database.Gorm.Create(artist); insertResult.Error != nil {
				database.DropAllTables()
				log.Fatal(insertResult.Error)
			}
		}

		// --------------- Album ---------------

		var albumModel *model.Lv2Album

		// Find album first, if found the record then we don't need to insert it again
		if result := database.Gorm.First(&albumModel, "spotify_id = ?", &item.Track.Album.Id); result.RowsAffected == 0 {
			releaseYear, _ := strconv.ParseInt(item.Track.Album.ReleaseDate[0:4], 10, 64)

			// Album
			album := &model.Lv2Album{
				SpotifyId:  &item.Track.Album.Id,
				Name:       &item.Track.Album.Name,
				ArtistId:   artistModel.Id,
				ArtworkUrl: item.Track.Album.Images[0].Url,
				Year:       &releaseYear,
			}

			albumModel = album

			// Insert Album
			if insertResult := database.Gorm.Create(album); insertResult.Error != nil {
				database.DropAllTables()
				panic(insertResult.Error)
			}
		}

		// --------------- Track ---------------

		randomCount := int64(text.RandomNumber(200)) + 1

		// Track
		track := &model.Lv2Track{
			SpotifyId:  item.Track.Id,
			Name:       item.Track.Name,
			AlbumId:    albumModel.Id,
			Duration:   item.Track.DurationMs,
			Popularity: item.Track.Popularity,
			Explicit:   item.Track.Explicit,
			Count:      &randomCount,
		}

		// Insert Track
		if result := database.Gorm.Create(track); result.Error != nil {
			database.DropAllTables()
			panic(result.Error)
		}
	}

	wg.Done()
}
