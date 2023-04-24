package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func handler(ctx *fiber.Ctx) error {
  fmt.Fprint(ctx, "API Biblioteca")
  return nil
}

//Functions 
func registerBook(c fiber.Ctx) error {
  return nil
}

func main() {
	app := fiber.New
  app.Get("/", handler)
  app.Listen(":8080")
}
