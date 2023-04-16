package database

import (
	"log"
	model "mock_spotify_data/model/db"
)

func DropAllTables() {
	if result := Gorm.Migrator().DropTable(
		&model.Lv1Track{},
		&model.Lv2Artist{},
		&model.Lv2Album{},
		&model.Lv2Track{},
		&model.Lv3Album{},
		&model.Lv3Artist{},
		&model.Lv3AlbumArtist{},
		&model.Lv3Track{},
		&model.Lv3Activity{}); result != nil {
		log.Fatal("Unable to drop tables")
	}
}
