package html

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

// Titulo obtem o título de uma página html
func Titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)
			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			matches := r.FindStringSubmatch(string(html))
			if len(matches) > 1 {
				c <- strings.TrimSpace(matches[1])
			} else {
				c <- "Sem título: " + url
			}
		}(url) // invocando a função automaticamente "go func(url string) --> função anonima"
	}
	return c
}
