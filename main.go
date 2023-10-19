package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/google/uuid"
)

type User struct {
	Id    string
	Name  string
	Email string
}

func createUser(c *fiber.Ctx) error {
	user := User{}
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	user.Id = uuid.NewString()
	return c.Status(fiber.StatusOK).JSON(user)
}

func ahri(c *fiber.Ctx) error {
	resp, err := http.Get("https://api.consumet.org/news/ann/recent-feeds")
	if err != nil {
		log.Fatalln(err)
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	sb := string(body)
	log.Println(sb)
	return c.Status(fiber.StatusOK).SendString(sb)
}

func main() {
	app := fiber.New()

	//MIDDLESWARES
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/createUser", createUser)
	app.Get("/ahri", ahri)

	app.Listen(":3000")
}
