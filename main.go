package main

import (
	"AmazonScrapper/utils"
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"os"
)

type product struct {
	ProductName string `selector:"span.a-size-medium.a-color-base.a-text-normal"`
	Stars       string `selector:"span.a-icon-alt"`
	Price       string `selector:"span.a-price > span.a-offscreen"`
}

func main() {
	searchTerm := "nintendo+switch"
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	products := make([]*product, 0)

	c.OnHTML("div.s-result-list.s-search-results.sg-row", func(e *colly.HTMLElement) {
		e.ForEach("div.a-section.a-spacing-medium", func(_ int, e *colly.HTMLElement) {
			c := &product{}

			e.Unmarshal(c)

			if len(c.ProductName) == 0 || len(c.Price) == 0 {
				return
			}

			utils.FormatStars(&c.Stars)
			utils.FormatPrice(&c.Price)

			products = append(products, c)
		})
	})

	c.Visit("https://www.amazon.com/s?k=" + searchTerm + "&ref=nb_sb_noss_1")
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")

	// Dump json to the standard output
	enc.Encode(products)
}