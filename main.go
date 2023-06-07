package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("HI, this is a THEA API endpoint, go to /fonts to access the font list")
	})

	app.Get("/fonts", func(c *fiber.Ctx) error {
		return c.SendFile("./dist.json")
	})

	app.Get("/dl/:value", func(c *fiber.Ctx) error {
		return c.SendFile("./fonts/" + c.Params("value"))
	})

	log.Fatal(app.Listen(":3000"))
}
