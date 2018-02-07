package controllers

import (
	"bufio"
	"fmt"
	"models"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

// CrawlerController c
type CrawlerController struct {
}

// NewCrawlerController c
func NewCrawlerController() *CrawlerController {
	return &CrawlerController{}
}

// Craw 爬原價屋
func (crawler *CrawlerController) Craw() (items []models.Item) {
	// var items []models.Item
	c := colly.NewCollector()

	c.OnHTML("tr[bgcolor='efefe0']", func(e *colly.HTMLElement) {

		count := 0
		scanner := bufio.NewScanner(strings.NewReader(e.ChildText("optgroup")))
		for scanner.Scan() {
			item := models.Item{}

			optgroup := e.ChildAttrs("optgroup", "label")[count]

			if len(scanner.Text()) < 3 {
				count++
			} else {
				category := e.ChildText("td:nth-child(2)")
				subCate := strings.Split(scanner.Text(), ` `)[0]
				itemName := strings.Split(scanner.Text(), ",")[0]
				oriPrice := regexp.MustCompile(`\$\d{1,}`).FindString(scanner.Text())
				oriPrice = regexp.MustCompile(`\d{1,}`).FindString(oriPrice)

				specialPrice := regexp.MustCompile(`↘\$\d{1,}\s`).FindString(scanner.Text())
				if len(specialPrice) > 1 {
					specialPrice = regexp.MustCompile(`\d{1,}`).FindString(specialPrice)
				}
				coolMoney := regexp.MustCompile(`↓酷幣\d{1,}↓`).FindString(scanner.Text())
				if len(coolMoney) > 1 {
					coolMoney = regexp.MustCompile(`\d{1,}`).FindString(coolMoney)
				}
				item.Category.Name = category
				item.SubCate.Name = subCate
				item.Name = itemName
				item.OriPrice, _ = strconv.Atoi(oriPrice)
				item.SpecialPrice, _ = strconv.Atoi(specialPrice)
				item.CoolMoney, _ = strconv.Atoi(coolMoney)
				item.Group = optgroup
				item.Tags = []string{category, subCate}
				items = append(items, item)
				if item.SpecialPrice != 0 {
					item.Tags = append(item.Tags, "onSale")
				}
				if item.CoolMoney != 0 {
					item.Tags = append(item.Tags, "coolpcDiscount")
				}
				// fmt.Println(item)
				// fmt.Println(oriPrice, specialPrice, coolMoney)
				// fmt.Println("category = ", category, "subCate = ", subCate, "Name = ", Item)
			}
		}
		// fmt.Println("===========")

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("爬", r.URL)
	})

	c.Visit("http://www.coolpc.com.tw/evaluate.php")

	return
}
