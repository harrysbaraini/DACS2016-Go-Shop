package system

import (
	"harrysbaraini/monoedit/models"

	"github.com/jinzhu/gorm"
	// Import sqlite dialect
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB ...
// Database instance
var DB *gorm.DB

// RegisterDatabase ...
// This function registers a GORM instance
func init() {
	// Open a connection to GORM
	db, err := gorm.Open("sqlite3", "shop.db")
	if err != nil {
		panic("Failed to connect database")
	}

	DB = db

	DB.AutoMigrate(models.Product{})
}
