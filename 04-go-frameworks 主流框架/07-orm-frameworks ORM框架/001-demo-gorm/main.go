package main

import (
	"fmt"
	"log"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type user struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"uniqueIndex"`
	CreatedAt time.Time
}

func main() {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(&user{}); err != nil {
		log.Fatal(err)
	}

	users := []user{
		{Name: "Alice", Email: "alice@example.com"},
		{Name: "Bob", Email: "bob@example.com"},
	}
	if err := db.Create(&users).Error; err != nil {
		log.Fatal(err)
	}

	var result []user
	if err := db.Order("id asc").Find(&result).Error; err != nil {
		log.Fatal(err)
	}

	if err := db.Model(&user{}).Where("email = ?", "alice@example.com").Update("name", "Alice Chen").Error; err != nil {
		log.Fatal(err)
	}

	fmt.Println("gorm inserted rows:", len(result))
	for _, item := range result {
		fmt.Printf("id=%d name=%s email=%s\n", item.ID, item.Name, item.Email)
	}
}
