package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

const (
	BREAKFAST = "Breakfast"
	LUNCH     = "Lunch"
	DINNER    = "Dinner"

	CROSSROADS        = "Crossroads"
	CAFE_3            = "Cafe_3"
	CLARK_KERR_CAMPUS = "Clark_Kerr_Campus"
	FOOTHILL          = "Foothill"
)

func formatParam(param string) string {
	return strings.ReplaceAll(strings.Title(strings.ToLower(param)), "%20", "_")
}

func main() {
	go cleanCache()

	app := fiber.New()

	app.Use(compress.New())
	app.Use(etag.New())
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]interface{}{
			"message": "bueno",
		})
	})

	app.Get("/:restaurant", func(c *fiber.Ctx) error {
		return c.JSON(map[string][]MealData{
			"data": {
				getData(formatParam(c.Params("restaurant", CROSSROADS)), BREAKFAST),
				getData(formatParam(c.Params("restaurant", CROSSROADS)), LUNCH),
				getData(formatParam(c.Params("restaurant", CROSSROADS)), DINNER),
			},
		})
	})
	app.Get("/:restaurant/:meal", func(c *fiber.Ctx) error {
		return c.JSON(
			getData(
				formatParam(c.Params("restaurant", CROSSROADS)),
				formatParam(c.Params("meal", LUNCH)),
			),
		)
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

	c.OnHTML(".cafe-location ."+restaurant+" .meal-period ."+meal+" .cat-name", func(e *colly.HTMLElement) {
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

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("error:", e, r.Request.URL, r.StatusCode)
	})

	c.DisableCookies()
	c.CacheDir = "./cache"
	c.Visit("https://caldining.berkeley.edu/menus/")

	return d
}

func cleanCache() {
	for range time.Tick(3 * 60 * time.Minute) {
		err := os.Remove("./cache")
		if err != nil {
			fmt.Println(err)
		}

		c := colly.NewCollector()
		c.Visit("https://caldining.berkeley.edu/menus/")
	}
}
