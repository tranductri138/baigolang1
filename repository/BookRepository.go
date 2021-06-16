package repository

import (
	"errors"
	"fmt"
	"ocg.com/train/model"
)

type BookRepo struct {
	books  map[int64]*model.Book
	autoID int64 //đây là biến đếm tự tăng gán giá trị cho id của Book
}

var Books BookRepo

func init() { //func init luôn chạy đầu tiên khi chúng ta import package
	Books = BookRepo{autoID: 0}
	Books.books = make(map[int64]*model.Book)
	Books.InitData("sql:45312")
}

//Pointer receiver ~ method trong Java. Đối tượng chủ thể là *BookRepo
func (r *BookRepo) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}

func (r *BookRepo) CreateNewBook(book *model.Book) int64 {
	nextID := r.getAutoID()
	book.Id = nextID
	r.books[nextID] = book
	return nextID
}

func (r *BookRepo) InitData(connection string) {
	fmt.Println("Connect to ", connection)

	r.CreateNewBook(&model.Book{
		Title: "Dế Mèn Phiêu Lưu Ký",
		Authors: []model.Author{
			{FullName: "Tô Hoài", Country: "Vietnam"},
			{FullName: "Hames", Country: "Turkey"},
		},
		Rating: 0})

	r.CreateNewBook(&model.Book{
		Title: "100 năm cô đơn",
		Authors: []model.Author{
			{FullName: "Gabriel Garcia Marquez", Country: "Columbia"},
			{FullName: "Ivan", Country: "Russia"},
		},
		Rating: 0})
}

func (r *BookRepo) GetAllBooks() map[int64]*model.Book {
	return r.books
}

func (r *BookRepo) FindBookById(Id int64) (*model.Book, error) {
	if book, ok := r.books[Id]; ok {
		return book, nil
	} else {
		return nil, errors.New("book not found")
	}
}

func (r *BookRepo) DeleteBookById(Id int64) error {
	if _, ok := r.books[Id]; ok {
		delete(r.books, Id)
		return nil
	} else {
		return errors.New("book not found")
	}
}

func (r *BookRepo) UpdateBookById(book *model.Book) error {
	if _, ok := r.books[book.Id]; ok {
		r.books[book.Id] = book
		return nil
	} else {
		return errors.New("book not found")
	}
}

func (r *BookRepo) Upsert(book *model.Book) int64 {
	if _, ok := r.books[book.Id]; ok {
		r.books[book.Id] = book //tìm thấy thì update
		return book.Id
	} else { //không tìm thấy thì tạo mới
		return r.CreateNewBook(book)
	}
}
