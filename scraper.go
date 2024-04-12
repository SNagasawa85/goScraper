package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func getPage(url string) string {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error getting page:", err)
		return ""
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("unable to close response body")
		}
	}(response.Body)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return ""
	}
	return string(body)
}

func getLinks(body string, baseUrl string) string {
	reader := strings.NewReader(body)
	document, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		fmt.Println("Error loading HTTP response body into goquery reader:", err)
		return ""
	}
	var links []string
	document.Find("a").Each(func(index int, element *goquery.Selection) {
		link, exists := element.Attr("href")
		if exists && !(strings.HasPrefix(link, "http") || strings.HasPrefix(link, "https")) {
			cleanedLink := strings.TrimPrefix(link, "/")
			combinedLink := fmt.Sprintf("%s%s", strings.TrimSuffix(baseUrl, "/"), cleanedLink)
			links = append(links, combinedLink)
		}
	})

	if len(links) > 0 {
		rand.Seed(time.Now().UnixNano())
		randIndex := rand.Intn(len(links))
		return links[randIndex]
	}
	return ""
}

func getSecondPage(url string) {
	if url == "" {
		fmt.Println("No URL GIVEN")
		return
	}
	fmt.Println("Getting html for:", url)
	secondPage := getPage(url)
	fmt.Println(secondPage)
}

func main() {
	var baseUrl string = "https://zerotomastery.io/blog/golang-practice-projects/"
	body := getPage(baseUrl)
	secondPageUrl := getLinks(body, baseUrl)
	getSecondPage(secondPageUrl)
}
