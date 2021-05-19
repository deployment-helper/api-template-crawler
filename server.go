package main

import (
	"github.com/deployment-helper/api-template-crawler/github"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"encoding/json"
	"log"
	"os"
)

type RepoReq struct {
	Owner   string
	Name    string
	ReqType string
}

func main() {
	app := fiber.New()
	// Default middleware config
	app.Use(logger.New())

	app.Get("/health", func(c *fiber.Ctx) error {
		c.Type("json")
		return c.SendString("Ok Github")
	})

	app.Post("/v1/crawl", func(c *fiber.Ctx) error {
		// TODO: API Request schema validation
		// TODO: API request jwt token validation
		log.Printf("%s started", "/v1/crawl")
		var templates []github.Template
		var repoRequests = &[]RepoReq{}
		if err := c.BodyParser(repoRequests); err != nil {
			return err
		}
		for _, r := range *repoRequests {
			var templatePointer, err = github.GetRepository("", r.Owner, r.Name)
			if err != nil {
				log.Print(err)
				continue
			}
			template := *templatePointer
			templates = append(templates, template)
		}
		resp, _ := json.Marshal(templates)
		c.Type("json", "utf-8")
		return c.SendString(string(resp))
	})

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
		log.Printf("defaulting to port %s", port)
	}

	app.Listen(":" + port)
}
