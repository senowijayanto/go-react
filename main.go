package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	fmt.Println("Go React Application")
	app := fiber.New()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	todos := []*Todo{}

	// Get Todos
	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"data": todos})
	})

	// Create Todos
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Todo body is required"})
		}

		todo.ID = len(todos) + 1
		todos = append(todos, todo)

		return c.Status(201).JSON(todo)
	})

	// Update Todos
	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		intID, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		for i, todo := range todos {
			if todo.ID == intID {
				todos[i].Completed = true
				return c.Status(200).JSON(todos[i])
			}
		}

		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	})

	// Delete Todos
	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		intID, err := strconv.Atoi(id)

		if err != nil {
			return err
		}

		for i, todo := range todos {
			if todo.ID == intID {
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"success": true})
			}
		}
		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"message": "Hello World"})
	})

	log.Fatal(app.Listen(":" + PORT))
}