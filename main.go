package main

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func main() {

	c := colly.NewCollector()

	// c.OnHTML("select[option]", func(e *colly.HTMLElement) {
	// 	e.Request.Visit(e.Attr("text"))
	// })

	// str := ""

	// c.OnHTML("select", func(e *colly.HTMLElement) {
	// e.Request.Visit(e.Attr("text"))
	// str += e.Text
	// })

	c.OnHTML("tr[bgcolor='efefe0']", func(e *colly.HTMLElement) {
		// e.Request.Visit(e.Attr("text"))

		count := 0
		fmt.Println(e.ChildText("td:nth-child(2)"))

		scanner := bufio.NewScanner(strings.NewReader(e.ChildText("optgroup")))

		for scanner.Scan() {
			if len(scanner.Text()) < 3 {
				fmt.Println("====", e.ChildAttrs("optgroup", "label")[count], "===")
				count++
			} else {
				category := e.ChildText("td:nth-child(2)")
				subCate := strings.Split(scanner.Text(), ` `)[0]
				fmt.Println("category = ", category, "subCate = ", subCate)
				fmt.Println(scanner.Text())
			}
		}
		fmt.Println("===========")

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("http://www.coolpc.com.tw/evaluate.php")

	// err := ioutil.WriteFile("./abc.txt", []byte(str), 0666)
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
