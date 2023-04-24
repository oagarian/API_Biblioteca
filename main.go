package main

import (
	"fmt"
  "github.com/gofiber/fiber/v2"
)

func handler(ctx *fiber.Ctx) error {
  fmt.Fprint(ctx, "API Biblioteca")
  return nil
}

func main() {
	fmt.Println("Hello, World!")
}
