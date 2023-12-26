package main

import (
	"html/template"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Create a new engine
	engine := html.New("./views", ".html")

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		}, "layouts/main")
	})

	app.Post("/submit", func(c *fiber.Ctx) error {
		tmpl, err := template.ParseFiles("./views/partials/result.html")
		if err != nil {
			return err
		}
		return tmpl.Execute(c.Response().BodyWriter(), fiber.Map{
			"Name":  c.FormValue("name"),
			"Email": c.FormValue("email"),
		})
	})

	log.Fatal(app.Listen(":3000"))
}
