package main

import (
	"fmt"
	"log"
	"time"

	_ "modernc.org/sqlite"
	"xorm.io/xorm"
)

type user struct {
	ID        int64     `xorm:"pk autoincr"`
	Name      string    `xorm:"varchar(50) notnull"`
	Email     string    `xorm:"varchar(100) unique"`
	CreatedAt time.Time `xorm:"created"`
}

func main() {
	engine, err := xorm.NewEngine("sqlite", "file::memory:?cache=shared")
	if err != nil {
		log.Fatal(err)
	}
	defer engine.Close()

	if err := engine.Sync2(new(user)); err != nil {
		log.Fatal(err)
	}

	if _, err := engine.Insert(&user{Name: "Carol", Email: "carol@example.com"}); err != nil {
		log.Fatal(err)
	}
	if _, err := engine.Insert(&user{Name: "David", Email: "david@example.com"}); err != nil {
		log.Fatal(err)
	}

	var users []user
	if err := engine.Asc("id").Find(&users); err != nil {
		log.Fatal(err)
	}

	fmt.Println("xorm inserted rows:", len(users))
	for _, item := range users {
		fmt.Printf("id=%d name=%s email=%s\n", item.ID, item.Name, item.Email)
	}
}
