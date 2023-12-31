package main

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/api", func(c *fiber.Ctx) error {
		slack_name := c.Query("slack_name")
		track := c.Query("track")
		res := fiber.Map{
			"slack_name":      slack_name,
			"current_day":     time.Now().Weekday().String(),
			"utc_time":        time.Now().UTC().Format("2006-01-02T15:04:05Z"),
			"track":           track,
			"github_file_url": "https://github.com/codewithed/hng_ix_01/blob/main/main.go",
			"github_repo_url": "https://github.com/codewithed/hng_ix_01",
			"status_code":     200,
		}

		return c.JSON(res)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
	log.Fatal(app.Listen("0.0.0.0:" + port))
}
