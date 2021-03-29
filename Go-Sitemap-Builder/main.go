package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/pranotobudi/Go-Gophercises/tree/main/Go-Sitemap-Builder/link"
)

func main() {
	fmt.Println("bismillah")
	urlFlag := flag.String("url", "https://gophercises.com", "URL for sitemap to build on")
	maxDepth := flag.Int("depth", 3, "maximumd depth to traverse the URL")
	flag.Parse()

	links := bsf(*urlFlag, *maxDepth)

	turnToXML(links)

}

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

type loc struct {
	Value string `xml:"loc"`
}
type urlset struct {
	Xmlns string `xml:"xmlns,attr"`
	Urls  []loc  `xml:"url"`
}

func turnToXML(links []string) []string {
	var result []string
	toXML := urlset{
		Xmlns: xmlns,
	}
	for _, link := range links {
		toXML.Urls = append(toXML.Urls, loc{link})
	}
	fmt.Printf("%v", xml.Header)
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("  ", "    ")
	if err := enc.Encode(toXML); err != nil {
		fmt.Printf("error: %v\n", err)
	}
	// fmt.Printf("%v \n", enc)
	return result
}

func bsf(urlStr string, maxDepth int) []string {

	finalURLs := make(map[string]struct{})
	nextURLs := make(map[string]struct{})
	visitedURLs := make(map[string]struct{})
	pageURLs := get(urlStr)
	for _, pageURL := range pageURLs {
		nextURLs[pageURL] = struct{}{}
	}

	for i := 0; i <= maxDepth; i++ {
		finalURLs, nextURLs = nextURLs, make(map[string]struct{})
		for page, _ := range finalURLs {
			if _, ok := visitedURLs[page]; ok {
				continue
			}
			visitedURLs[page] = struct{}{}
			links := get(page)
			for _, link := range links {
				nextURLs[link] = struct{}{}
			}
		}
	}
	result := make([]string, len(finalURLs))
	counter := 0
	for link, _ := range finalURLs {
		// fmt.Printf("finalURL: %v \n", link)
		result[counter] = link
		counter++
	}
	return result
}
func get(urlStr string) []string {
	res, _ := http.Get(urlStr)
	links, _ := link.ParseURL(res.Body)

	reqUrl := res.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	defer res.Body.Close()

	return filter(links, baseUrl.String())

}
func filter(links []string, baseURL string) []string {
	var pageURLs []string
	for _, link := range links {
		// fmt.Printf("url: %v \n", link)
		switch {
		case strings.HasPrefix(link, "/"):
			pageURLs = append(pageURLs, baseURL+link)
		case strings.HasPrefix(link, "http"): //https://gophercises.com
			if strings.Contains(link, baseURL) {
				// fmt.Printf("baseURL: %v \n", baseURL)
				pageURLs = append(pageURLs, link)
			}
		}
	}
	return pageURLs
}
