package procedures

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"mock_spotify_data/types/payload"
	"mock_spotify_data/utils/text"
	"os"
	"path/filepath"
)

var CurrentData *payload.SpotifyPlaylist

func RandomData() {

	// File name
	fileArray := [15]string{"all_out_2010_playlist", "daily_playlist", "greatest_1_playlist", "greatest_2_playlist", "greatest_3_playlist", "greatest_4_playlist", "greatest_5_playlist", "kpop_playlist", "lofi_1_playlist", "lofi_2_playlist", "nohguties_2_playlist", "noughties_1_playlist", "soft_10_s_playlist", "store_appropiate_playlist", "top_150_from_2010_2020_playlist"}

	// Random file number
	fileIndex := text.RandomNumber(15)

	// Show selected file
	fmt.Println("Selected playlist is " + fileArray[fileIndex])

	// Open our jsonFile
	absPath, _ := filepath.Abs("./data/" + fileArray[fileIndex] + ".json")

	jsonFile, err := os.Open(absPath)

	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully Opened users.json")

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result *payload.SpotifyPlaylist
	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Validating data")

	// Validate data
	if result.Href == nil {
		log.Fatal("Playlist ID is nil")
	}

	for _, item := range *result.Items {
		fmt.Println("Validation track id : " + *item.Track.Id)

		// Artist

		if &item.Track.Artists[0].Id == nil {
			log.Fatal("Track artist id is nil")
		}

		if &item.Track.Artists[0].Name == nil {
			log.Fatal("Track artist Name is nil")
		}

		if &item.Track.Artists[0].Href == nil {
			log.Fatal("Track artist link is nil")
		}

		// Album
		if &item.Track.Album.Id == nil {
			log.Fatal("Track album id is nil")
		}

		if &item.Track.Album.Name == nil {
			log.Fatal("Track album name is nil")
		}

		if len(item.Track.Album.Images) == 0 {
			log.Fatal("Track album image is nil")
		}

		// Track
		if item.Track.Id == nil {
			log.Fatal("Track ID is nil")
		}

		if item.Track.Name == nil {
			log.Fatal("Track Name is nil")
		}

		if item.Track.DurationMs == nil {
			log.Fatal("Track duration is nil")
		}

		if item.Track.Popularity == nil {
			log.Fatal("Track popularity is nil")
		}

		if item.Track.Explicit == nil {
			log.Fatal("Track explicit is nil")
		}

		// Random explicit
		explicit := rand.Intn(2) == 0
		item.Track.Explicit = &explicit

		if item.Track.PreviewUrl == nil {
			item.Track.PreviewUrl = nil
		}
	}

	fmt.Println("Validate data complete")

	CurrentData = result

	fmt.Println("Successfully Changed random data")
}
