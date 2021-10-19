package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/gocolly/colly"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(cache.New(cache.Config{
		Expiration: 2 * 60 * time.Minute,
	}))
	app.Use(compress.New())
	app.Use(etag.New())
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]interface{}{
			"message": "bueno",
		})
	})

	app.Get("/crossroads", func(c *fiber.Ctx) error {
		return c.JSON(getData("crossroads", ""))
	})
	app.Get("/crossroads/lunch", func(c *fiber.Ctx) error {
		return c.JSON(getData("crossroads", "lunch"))
	})
	app.Get("/crossroads/dinner", func(c *fiber.Ctx) error {
		return c.JSON(getData("crossroads", "dinner"))
	})

	if runtime.GOOS == "darwin" {
		app.Listen("localhost:3000")
	} else {
		app.Listen(":3000")
	}
}

func getData(restaurant string, meal string) map[string]interface{} {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML(".cafe-location .Crossroads .meal-period ."+meal+" ", func(e *colly.HTMLElement) {
		// e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://caldining.berkeley.edu/menus/")

	return map[string]interface{}{
		"stautus": "ok",
	}
}
