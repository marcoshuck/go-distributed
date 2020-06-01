package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/marcoshuck/go-distributed/challenge/pkg/sum"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&sum.Sum{})
	fmt.Println("The database has been migrated.")
}
