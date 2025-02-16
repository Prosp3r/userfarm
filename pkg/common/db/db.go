package db

import (
	"log"

	"github.com/Prosp3r/deepseek/userfarm/pkg/common/models"
	"github.com/Prosp3r/userfarm/pkg/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	// db.AutoMigrate(&models.Url{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Task{})
	db.AutoMigrate(&models.Referral{})
	// db.AutoMigrate(&models.UserTask{})

	return db
}
