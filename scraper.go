package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getPage(url string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error getting page:", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return
	}

	fmt.Println(string(body))
}

func main() {
	getPage("http://www.iamjason.me")
}
