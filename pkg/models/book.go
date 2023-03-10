package models

import (
	"go-movies-crud/pkg/config"

	"github.com/PranavJayachandran/go-bookstore/config"
	"github.com/jinzhu/gorm"
)

var db=gorm.DB

type Book struct{
	gorm.Model
	Name string `gotm:""json:"name"`
	Author string `json:"author`
	Publication string `json:"publication:`
}

func init(){
	config.Connect()
	db=config.GetDB()
	db.AutoMigrate(&Book{})
}

func(b *Book) CreateBook() *Book{
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks()[]Book{
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBooksById(Id int64) (*Book,*gorm.db){
	var getBook Book
	db:=db.where("ID=?",Id).Find(&getBook)
	return &getBook,db
}

func DeleteBook(Id int64) Book{
	var book Book
	db.where("ID=?",Id).Delete(book)
	return book
}
