package procedures

import (
	"log"
	"mock_spotify_data/database"
	model "mock_spotify_data/model/db"
	"mock_spotify_data/utils/text"
	"strconv"
	"time"
)

func Level3Operation() {
	nowTime := time.Now()

	for _, item := range *CurrentData.Items {
		// --------------- Artist ---------------

		var artistModel *model.Lv3Artist

		// Find artist first, if found the record then we don't need to insert it again
		if result := database.Gorm.First(&artistModel, "spotify_id = ?", item.Track.Album.Artists[0].Id); result.RowsAffected == 0 {
			artist := &model.Lv3Artist{
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

		var albumModel *model.Lv3Album

		// Find album first, if found the record then we don't need to insert it again
		if result := database.Gorm.First(&albumModel, "spotify_id = ?", &item.Track.Album.Id); result.RowsAffected == 0 {
			releaseYear, _ := strconv.ParseInt(item.Track.Album.ReleaseDate[0:4], 10, 64)

			// Album
			album := &model.Lv3Album{
				SpotifyId:  &item.Track.Album.Id,
				Name:       &item.Track.Album.Name,
				ArtworkUrl: item.Track.Album.Images[0].Url,
				Year:       &releaseYear,
			}

			albumModel = album

			// Insert Album
			if insertResult := database.Gorm.Create(album); insertResult.Error != nil {
				database.DropAllTables()
				log.Fatal(insertResult.Error)
			}
		}

		// --------------- Album & Artist ---------------

		var albumAndArtistModel *model.Lv3AlbumArtist

		if result := database.Gorm.First(&albumAndArtistModel, "album_id = ? AND artist_id = ?", albumModel.Id, artistModel.Id); result.RowsAffected == 0 {
			albumArtist := &model.Lv3AlbumArtist{
				AlbumId:  albumModel.Id,
				ArtistId: artistModel.Id,
			}

			// Insert Album & Artist
			if insertResult := database.Gorm.Create(albumArtist); insertResult.Error != nil {
				database.DropAllTables()
				log.Fatal(insertResult.Error)
			}
		}

		// --------------- Track ---------------

		// Track
		track := &model.Lv3Track{
			SpotifyId:  item.Track.Id,
			Name:       item.Track.Name,
			AlbumId:    albumModel.Id,
			Duration:   item.Track.DurationMs,
			Popularity: item.Track.Popularity,
			Explicit:   item.Track.Explicit,
			PreviewUrl: item.Track.PreviewUrl,
		}

		// Insert Track
		if result := database.Gorm.Create(track); result.Error != nil {
			database.DropAllTables()
			log.Fatal(result.Error)
		}

		// --------------- Activity ---------------

		shuffleState := text.RandomShuffleState()
		repeatState := text.RandomRepeatState()
		deviceName := text.RandomDevice()
		context := text.GetPlaylistHref(*CurrentData.Href)

		for i := 0; i < text.RandomNumber(10); i++ {
			// Activity
			activity := &model.Lv3Activity{
				TrackId:      track.Id,
				ShuffleState: &shuffleState,
				RepeatState:  &repeatState,
				DeviceName:   &deviceName,
				Context:      &context,
				Timestamp:    &nowTime,
			}

			// Insert Activity
			if result := database.Gorm.Create(activity); result.Error != nil {
				database.DropAllTables()
				log.Fatal(result.Error)
			}

			nowTime = nowTime.Add(time.Minute * 1)
		}
	}
}
