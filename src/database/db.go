package database

import (
	"ca-server/src/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:truong123456@tcp(localhost:3306)/sys?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	//using the database 'sys'
	useSysDB := "USE sys"
	if err := DB.Exec(useSysDB).Error; err != nil {
		log.Fatal("Failed to select database 'sys':", err)
	}

	// Automatically migrate schema
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Message{})
	DB.AutoMigrate(&models.Group{})
	DB.AutoMigrate(&models.GroupMember{})
	DB.AutoMigrate(&models.MediaMessage{})
	DB.AutoMigrate(&models.Notification{})
	DB.AutoMigrate(&models.Relationship{})
}
