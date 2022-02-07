package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
	"strings"
)

// 练习5.5：实现函数countWordsAndImages。
func main() {
	words, images, err := CountWordsAndImages(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("words: %d, imges: %d", words, images)
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}

	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n == nil {
		return
	}

	if n.Type == html.TextNode {
		words = len(strings.Fields(n.Data))
	}
	if n.Type == html.ElementNode && n.Data == "img" {
		images = 1
	}

	var w, i int
	w, i = countWordsAndImages(n.FirstChild)
	words += w
	images += i
	w, i = countWordsAndImages(n.NextSibling)
	words += w
	images += i

	return
}
