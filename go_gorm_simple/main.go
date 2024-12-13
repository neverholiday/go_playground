package main

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Books struct {
	Id   string `gorm:"column:id"`
	Name string `gorm:"column:name"`
}

type Authors struct {
	Id   string `gorm:"column:id"`
	Name string `gorm:"column:name"`
}

type BookInfo struct {
	BookName   string `gorm:"book_name"`
	AuthorName string `gorm:"author_name"`
}

func main() {
	dsn := "/Users/neverholiday/sqlitedbs/test.db"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// book := Books{
	// 	Id:   uuid.New().String(),
	// 	Name: "One Piece",
	// }

	// err = db.Table("books").Create(&book).Error
	// if err != nil {
	// 	panic(err)
	// }

	// author := Authors{
	// 	Id:   uuid.New().String(),
	// 	Name: "Eichiro Oda",
	// }

	// err = db.Table("authors").Create(&author).Error
	// if err != nil {
	// 	panic(err)
	// }

	_ = uuid.New()

	var bookInfos []BookInfo

	err = db.Raw(` SELECT 	
							a.name AS author_name, 
							b.name AS book_name
					FROM 
							books_authors ba 
					LEFT JOIN 
							authors a on a.id = ba.author_id 
					LEFT JOIN  
							books b on b.id = ba.book_id`).
		Scan(&bookInfos).Error

	if err != nil {
		panic(err)
	}

	for _, book := range bookInfos {
		fmt.Println(book.AuthorName, "|", book.BookName)
	}

}
