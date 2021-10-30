package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
type ArcStory struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

func main() {
	http.HandleFunc("/", loadStory())
	fmt.Println("starting server on port 8080")
	http.ListenAndServe(":8080", nil)
}
func loadJSON() map[string]ArcStory {
	filename := flag.String("file", "gopher.json", "the JSON file with CYOA story")
	jsonFile, err := os.Open(*filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// var users Users
	story := make(map[string]ArcStory)
	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &story)
	// fmt.Println("JSON FILE:", story["intro"])
	return story
}
func loadStory() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("layout.html"))
		story := loadJSON()
		path := r.URL.Path
		fmt.Println("path:", path)
		var arc string
		if path == "/" {
			arc = "intro"
		} else {
			arc = path
		}
		arc = strings.Trim(arc, "/")
		fmt.Println("arc:", arc)
		tmpl.Execute(w, story[arc])
	}

}
