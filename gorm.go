package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()
	db.AutoMigrate(&Product{})

	db.Create(&Product{Code: "123", Price: 1000})

	var product Product
	db.First(&product, 1)
	db.First(&product, "code = ?", "123")

	db.Model(&product).Update("Price", 2000)
}
