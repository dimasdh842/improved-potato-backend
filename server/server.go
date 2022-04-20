package server

import (
	"improved_potato/helper"
	"improved_potato/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func getAllRedirects(c *fiber.Ctx) error {
	imtos, err := model.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting all links " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(imtos)
}

func getOne(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not parse id " + err.Error(),
		})
	}

	imto, err := model.GetOne(int(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not retrive url from db " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(imto)
}

func createOne(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var imto model.Imto
	err := c.BodyParser(&imto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not retrive url from db " + err.Error(),
		})
	}

	if imto.Random {
		imto.Imto = helper.RandomURL(8)
	}

	imto, err = model.CreateOne(imto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not retrive url from db " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(imto)
}

func updateOne(c *fiber.Ctx) error {
	c.Accepts("application/json")
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)

	var imto model.Imto
	imto, err = model.GetOne(int(id))

	var input model.Imto
	err = c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not retrive url from db " + err.Error(),
		})
	}

	imto.Imto = input.Imto
	imto.Random = input.Random
	imto.Redirect = input.Redirect

	imto, err = model.UpdateOne(imto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not retrive url from db " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(imto)

}

func deleteOne(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not retrive url from db " + err.Error(),
		})
	}

	err = model.DeleteOne(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not retrive url from db " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "delete successfully",
	})

}

func redirect(c *fiber.Ctx) error {
	uniqURL := c.Params("redirect")
	imto, err := model.FindByUniqueUrl(uniqURL)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not retrive url from db " + err.Error(),
		})
	}

	imto.Clicked += 1
	imto, err = model.UpdateOne(imto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error updating " + err.Error(),
		})
	}

	return c.Redirect(imto.Redirect, fiber.StatusTemporaryRedirect)

}

func SetupAndListen() {
	router := fiber.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin,Content-Type,Accept",
	}))
	router.Get("/r/:redirect", redirect)

	router.Get("/all", getAllRedirects)
	router.Get("/all/:id", getOne)
	router.Post("/all", createOne)
	router.Patch("/all/:id", updateOne)
	router.Delete("/all/:id", deleteOne)

	router.Listen(":3000")

}
