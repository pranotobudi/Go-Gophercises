package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Options struct {
	Text string
	Arc  string
}
type StoryArc struct {
	Title   string
	Story   []string
	Options []Options
}

// type StoryArc struct {
// 	Title   string   `json:"title"`
// 	Story   []string `json:"story"`
// 	Options []struct {
// 		Text string `json:"text"`
// 		Arc  string `json:"arc"`
// 	} `json:"options"`
// }
type jsonContent map[string]StoryArc

func main() {
	// Open our jsonFile
	filename := flag.String("file", "story.json", "file with CYOA story")
	flag.Parse()
	jsonFile, err := os.Open(*filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened story.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	// byteValue, _ := ioutil.ReadAll(jsonFile)
	// var result map[string]interface{}
	var result jsonContent
	d := json.NewDecoder(jsonFile)
	d.Decode(&result)

	// json.Unmarshal([]byte(byteValue), &result)
	var page StoryArc
	// fmt.Println(page.Title)

	tmpl := template.Must(template.ParseFiles("layout.html"))
	// if err != nil {
	// 	fmt.Println("failed to parse layout.html ", err)
	// }
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		page = result["intro"]
		tmpl.Execute(w, page)
	})
	http.HandleFunc("/new-york", func(w http.ResponseWriter, r *http.Request) {
		page = result["new-york"]
		tmpl.Execute(w, page)
	})
	http.HandleFunc("/debate", func(w http.ResponseWriter, r *http.Request) {
		page = result["debate"]
		tmpl.Execute(w, page)
	})
	http.HandleFunc("/sean-kelly", func(w http.ResponseWriter, r *http.Request) {
		page = result["sean-kelly"]
		tmpl.Execute(w, page)
	})
	http.HandleFunc("/mark-bates", func(w http.ResponseWriter, r *http.Request) {
		page = result["mark-bates"]
		tmpl.Execute(w, page)
	})
	http.HandleFunc("/denver", func(w http.ResponseWriter, r *http.Request) {
		page = result["denver"]
		tmpl.Execute(w, page)
	})
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		page = result["home"]
		tmpl.Execute(w, page)
	})
	http.ListenAndServe(":80", nil)

	// newYork := result["new-york"]
	// debate := result["debate"]
	// seanKelly := result["sean-kelly"]
	// markBates := result["mark-bates"]
	// denver := result["denver"]
	// home := result["home"]

}
