package main

import (
	"github.com/budisetionugroho123/money_manegement/migrations"
	"github.com/gofiber/fiber/v2"
)

func main() {
	migrations.RunMigration()

	app := fiber.New()
	app.Listen(":5000")

}
