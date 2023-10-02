package database

import (
	"github.com/samsunil1999/ad_space_auction_service/configs"
	"github.com/samsunil1999/ad_space_auction_service/models/entities"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

// Connect to the MySql database
func Connect() {
	dsn := configs.App.DbUser + ":" +
		configs.App.DbPassword + "@tcp" +
		"(" + configs.App.DbHost + ":" + configs.App.DbPort + ")/" +
		configs.App.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
}

func SyncDatabase() {
	Db.AutoMigrate(&entities.AdSpaces{})
	Db.AutoMigrate(&entities.Bids{})
	Db.AutoMigrate(&entities.Bidders{})
}
