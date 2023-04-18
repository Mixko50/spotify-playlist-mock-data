package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	model "mock_spotify_data/model/db"
	"os"
)

var Gorm *gorm.DB

func Init() {

	db, err := gorm.Open(mysql.Open(os.Getenv("DATABASE_DSN")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Unable to connect the database", err)
	}

	Gorm = db

	log.Println("Initializing the database")

	DropAllTables()

	databaseInitErr := Gorm.AutoMigrate(
		&model.Lv1Track{},
		&model.Lv2Artist{},
		&model.Lv2Album{},
		&model.Lv2Track{},
		&model.Lv3Album{},
		&model.Lv3Artist{},
		&model.Lv3AlbumArtist{},
		&model.Lv3Track{},
		&model.Lv3Activity{},
	)
	if databaseInitErr != nil {
		log.Fatal("Unable to migrate database", err)
	}

	log.Println("Initialized mysql connection")
}
