package sitemap

import (
	"fmt"
	"net/http"

	"github.com/pranotobudi/Go-Gophercises/tree/main/Go-Sitemap-Builder/link"
)

func Build(url string) map[string]bool {
	// var checkedURL map[string]bool
	checkedURL := make(map[string]bool)

	//base
	url = link.NormalizeURL(url, parentUrl)
	fmt.Printf("normalized url: %v \n", url)
	if !checkedURL[url] {
		res, err := http.Get(url)
		if err != nil {
			fmt.Println("can't open the URL")
		}
		defer res.Body.Close()

		URLs, err := link.ParseURL(res.Body)
		if err != nil {
			panic(err)
		}
		checkedURL[url] = true
		fmt.Printf("checkedURL[%v] = %v \n", url, checkedURL[url])
		for _, childUrl := range URLs {
			fmt.Printf("childURL: %v \n", childUrl)
			newCheckedURL := Build(childUrl, parentUrl)
			for key, value := range newCheckedURL {
				checkedURL[key] = value
			}
			fmt.Printf("%v \n", url)
		}
	}
	return checkedURL
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func convertToXML(checkedURL []string) string {
	return ""
}
