package routes

import (
	"github.com/gofiber/fiber/v2"
	"ocg.com/train/controller"
)

func ConfigBookRouter(router *fiber.Router) {
	(*router).Get("/books", controller.GetAllBooks)
	(*router).Post("/books", controller.CreateNewBook)
	(*router).Get("/books/:id", controller.GetBookById)
	(*router).Delete("/books/:id", controller.DeleteBookById)
	(*router).Patch("/books", controller.UpdateBookById)
	(*router).Put("/books/:id", controller.UpsertBook)
}

func ConfigReviewRouter(router *fiber.Router) {
	(*router).Get("/reviews/books/:id", controller.GetReviewByBookId)
	(*router).Get("/reviews", controller.GetAllReview)
	(*router).Patch("/reviews/:id", controller.UpdateReviewById)
	(*router).Post("/reviews", controller.CreateNewReView)
	(*router).Delete("/reviews/:id", controller.DeleteReviewById)
}
