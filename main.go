package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		slack_name := c.Query("slack_name")
		track := c.Query("track")
		res := &fiber.Map{
			"slack_name":      slack_name,
			"current_day":     time.Now().Weekday(),
			"utc_time":        time.Now().UTC(),
			"track":           track,
			"github_file_url": "https://github.com/codewithed/hng_ix_01/blob/main/main.go",
			"github_repo_url": "https://github.com/codewithed/hng_ix_01",
			"status_code":     200,
		}

		return c.JSON(res)
	})

	log.Fatal(app.Listen(":3000"))
}
