package service

import (
	"ocg.com/train/model"
	"ocg.com/train/repository"
)

func FindBookById(Id int64) (*model.Book, error) {
	return repository.Books.FindBookById(Id)
}

func UpdateBookRating(book *model.Book) *model.Book {
	sum := 0
	count := 0
	reviews := repository.Reviews.FindAllReview()
	for _, review := range reviews {
		if book.Id == review.BookId {
			sum += review.Rating
			count++
		}
	}
	rating := float32(sum) / float32(count)
	book.Rating = rating
	return book
}
