package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func titulo(urls ...string) <-chan string {
	c := make(chan string)

	for _, url := range urls {
		go func(cUrl string) {
			resp, _ := http.Get(cUrl)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)</title>")
			aRetorno := r.FindStringSubmatch(string(html))

			// teste para evitar erro
			if cap(aRetorno) == 0 {
				c <- "Erro ao ler página " + cUrl
				return
			}

			c <- aRetorno[1]
		}(url)
	}

	return c
}

func main() {
	t1 := titulo("https://www.instagram.com", "https://www.twitch.tv")
	t2 := titulo("https://www.youtube.com", "https://github.com")
	fmt.Println("Primeiros: ", <-t1, <-t2)
	fmt.Println("Segundos: ", <-t1, <-t2)
}
