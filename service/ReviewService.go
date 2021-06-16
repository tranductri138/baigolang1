package service

import (
	"ocg.com/train/model"
	"ocg.com/train/repository"
)

func GetAllReview() []*model.Review {
	return repository.Reviews.FindAllReview()
}

func CreateReview(review *model.Review) *model.Review {
	newReview := repository.Reviews.CreateNewReview(review)
	bookId := newReview.BookId
	book, _ := FindBookById(bookId)
	UpdateBookRating(book)
	return newReview
}

func DeleteReviewById(Id int64) error {
	review, _ := repository.Reviews.FindReviewById(Id)
	bookID := review.BookId
	book, _ := FindBookById(bookID)
	UpdateBookRating(book)
	return repository.Reviews.DeleteReviewById(Id)
}

func UpdateReView(Id int64, newReview *model.Review) *model.Review {
	err := DeleteReviewById(Id)
	if err != nil {
		return nil
	}
	review := CreateReview(newReview)
	review.Id = Id
	book, _ := FindBookById(review.BookId)
	UpdateBookRating(book)
	return review
}

func GetReviewByBookId(bookId int64) map[int64][]*model.Review {
	return repository.Reviews.GetReviewByBookId(bookId)
}
