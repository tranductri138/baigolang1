package repository

import (
	"errors"
	"ocg.com/train/model"
)

type ReviewRepo struct {
	reviews map[int64]*model.Review
	autoID  int64
}

var Reviews ReviewRepo

func init() {
	Reviews = ReviewRepo{autoID: 0}
	Reviews.reviews = make(map[int64]*model.Review)
}

func (r *ReviewRepo) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}

func (r *ReviewRepo) CreateNewReview(review *model.Review) *model.Review {
	nextID := r.getAutoID()
	review.Id = nextID
	r.reviews[nextID] = review
	return review

}

func (r *ReviewRepo) FindAllReview() (result []*model.Review) {
	mapReviews := r.reviews
	for _, review := range mapReviews {
		result = append(result, review)
	}
	return result
}

func (r *ReviewRepo) FindReviewById(Id int64) (*model.Review, error) {
	review, ok := r.reviews[Id]
	if ok {
		return review, nil
	} else {
		return nil, errors.New("review not fount")
	}
}

func (r *ReviewRepo) UpdateReviewById(review *model.Review) error {
	if _, ok := r.reviews[review.Id]; ok {
		r.reviews[review.Id] = review
		return nil
	} else {
		return errors.New("review not found")
	}
}

func (r *ReviewRepo) GetReviewByBookId(bookId int64) (result map[int64][]*model.Review) {
	result = make(map[int64][]*model.Review)
	reviews := r.FindAllReview()
	for _, review := range reviews {
		if bookId == review.BookId {
			result[bookId] = append(result[bookId], review)
		}
	}
	return result
}

func (r *ReviewRepo) DeleteReviewById(Id int64) error {
	if _, ok := r.reviews[Id]; ok {
		delete(r.reviews, Id)
		return nil
	} else {
		return errors.New("book not found")
	}
}
