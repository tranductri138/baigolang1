package controller

import (
	"github.com/gofiber/fiber/v2"
	"ocg.com/train/model"
	"ocg.com/train/service"
)

func GetAllReview(c *fiber.Ctx) error {
	return c.JSON(service.GetAllReview())
}

func CreateNewReView(c *fiber.Ctx) error {
	review := new(model.Review)
	err := c.BodyParser(&review)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	newReview := service.CreateReview(review)
	return c.JSON(newReview)
}

func UpdateReviewById(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	review := new(model.Review)
	err := c.BodyParser(&review)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	updateReview := service.UpdateReView(int64(id), review)
	return c.JSON(updateReview)

}

func GetReviewByBookId(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	return c.JSON(service.GetReviewByBookId(int64(id)))
}

func DeleteReviewById(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	return c.JSON(service.DeleteReviewById(int64(id)))
}
