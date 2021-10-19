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

const (
	BREAKFAST = "Breakfast"
	LUNCH     = "Lunch"
	DINNER    = "Dinner"
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
		return c.JSON(map[string][]MealData{
			"data": {
				getData("Crossroads", BREAKFAST),
				getData("crossroads", LUNCH),
				getData("crossroads", DINNER),
			},
		})
	})
	app.Get("/crossroads/lunch", func(c *fiber.Ctx) error {
		return c.JSON(getData("crossroads", LUNCH))
	})
	app.Get("/crossroads/dinner", func(c *fiber.Ctx) error {
		return c.JSON(getData("crossroads", DINNER))
	})

	if runtime.GOOS == "darwin" {
		app.Listen("localhost:3000")
	} else {
		app.Listen(":3000")
	}
}

type MealData struct {
	Name     string        `json:"name"`
	Sections []MealSection `json:"sections"`
}

type MealSection struct {
	Name  string     `json:"name"`
	Items []MealItem `json:"items"`
}

type MealItem struct {
	Name string `json:"name"`
}

func getData(restaurant string, meal string) MealData {
	d := MealData{
		Name: meal,
	}
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML(".cafe-location .Crossroads .meal-period ."+meal+" .cat-name", func(e *colly.HTMLElement) {
		s := MealSection{}
		childNodes := e.DOM.Children()

		s.Name = childNodes.First().Text()
		e.ForEach(".recip", func(_ int, recipeElement *colly.HTMLElement) {
			item := MealItem{
				Name: recipeElement.DOM.Children().First().Text(),
			}

			s.Items = append(s.Items, item)
		})

		d.Sections = append(d.Sections, s)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://caldining.berkeley.edu/menus/")

	return d
}
