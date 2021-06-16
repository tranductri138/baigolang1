package controller

import (
	"fmt"
	"ocg.com/train/model"
	"ocg.com/train/repository"

	"github.com/gofiber/fiber/v2"
)

func GetAllBooks(c *fiber.Ctx) error {
	return c.JSON(repository.Books.GetAllBooks())
}

func CreateNewBook(c *fiber.Ctx) error {
	book := new(model.Book)
	err := c.BodyParser(&book)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	bookId := repository.Books.CreateNewBook(book)
	return c.SendString(fmt.Sprintf("New book is created successfully with id = %d", bookId))
}

func GetBookById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	book, err := repository.Books.FindBookById(int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	return c.JSON(book)
}

func DeleteBookById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	err = repository.Books.DeleteBookById(int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	return c.SendString("delete successfully")
}

func UpdateBookById(c *fiber.Ctx) error {
	updatedBook := new(model.Book)
	err := c.BodyParser(&updatedBook)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	err = repository.Books.UpdateBookById(updatedBook)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	return c.SendString(fmt.Sprintf("Book with id = %d is successfully updated", updatedBook.Id))
}

func UpsertBook(c *fiber.Ctx) error {
	book := new(model.Book)

	err := c.BodyParser(&book)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	id := repository.Books.Upsert(book)
	return c.SendString(fmt.Sprintf("Book with id = %d is successfully upserted", id))
}
