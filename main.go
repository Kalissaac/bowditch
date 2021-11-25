package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
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

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		url := strings.TrimSuffix(r.URL.Path, "/")
		portions := strings.Split(url, "/")
		fmt.Println(url, portions, len(portions))

		if len(portions) == 1 { // /
			w.Write([]byte(`{"message": "bowditch is alive"}`))
			return
		} else if len(portions) == 2 { // /restaurant
			restaurant := portions[1]
			data, err := json.Marshal(map[string][]MealData{
				"data": {
					getData(formatParam(restaurant), BREAKFAST),
					getData(formatParam(restaurant), LUNCH),
					getData(formatParam(restaurant), DINNER),
				},
			})

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf(`{"message": "%s"}`, err)))
			} else {
				w.Write(data)
			}
			return
		} else if len(portions) == 3 { // /restaurant/meal
			restaurant := portions[1]
			meal := portions[2]
			data, err := json.Marshal(getData(
				formatParam(restaurant),
				formatParam(meal),
			))

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf(`{"message": "%s"}`, err)))
			} else {
				w.Write(data)
			}
			return
		}

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Not found"}`))
	})

	addr := ":3000"
	if runtime.GOOS == "darwin" {
		addr = "localhost:3000"
	}
	log.Fatal(http.ListenAndServe(addr, mux))
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
	c.CacheDir = "./.cache"
	c.Visit("https://caldining.berkeley.edu/menus/")

	return d
}

func cleanCache() {
	for range time.Tick(3 * 60 * time.Minute) {
		err := os.RemoveAll("./.cache")
		if err != nil {
			fmt.Println(err)
		}

		c := colly.NewCollector()
		c.DisableCookies()
		c.CacheDir = "./.cache"
		c.Visit("https://caldining.berkeley.edu/menus/")
	}
}
