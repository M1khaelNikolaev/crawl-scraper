package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	links, err := GetPageLinks("https://gcheb.cap.ru/")
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Ссылки на странице:")
	for _, link := range links {
		fmt.Println(link)
	}
}

func GetPageLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	var links []string
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, exists := s.Attr("href")
		if exists {
			links = append(links, link)
		}
	})
	return links, nil
}
