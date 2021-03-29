package link

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

//Link represents a link (<a href="">) in html document
type Link struct {
	Href string
	Text string
}

// var r io.Reader

// links, err := link.Parse(r)

//Parse will take in an HTML document and will return a slice of links parsed from it
func ParseURL(r io.Reader) ([]string, error) {
	doc, err := html.Parse(r)
	if err != nil {
		panic(err)
	}
	nodes := LinkNodes(doc)
	var links []Link
	for _, node := range nodes {
		links = append(links, BuildLink(node))
		// fmt.Printf("%v \n", node)
	}
	// dfs(doc, "")
	var URLs []string
	for _, link := range links {
		URLs = append(URLs, link.Href)
	}
	return URLs, nil
}

func NormalizeURL(url string, parentUrl string) string {
	var result string
	if strings.Contains(url, "http://") { //remove http://
		result = "https://" + url[7:]
	}
	// fmt.Printf("inside normalize func, url: %v \n", url)
	if !strings.Contains(url, ".") {
		result = parentUrl + url //parentUrl in the form of https
		// fmt.Printf("inside if, result: %v \n", result)
	} else {
		result = url
	}
	return result
}

func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		panic(err)
	}
	nodes := LinkNodes(doc)
	var links []Link
	for _, node := range nodes {
		links = append(links, BuildLink(node))
		fmt.Printf("%v \n", node)
	}
	// dfs(doc, "")

	return links, nil
}

func BuildLink(n *html.Node) Link {
	var result Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			result.Href = attr.Val
			break
		}
	}
	result.Text = BuildText(n)

	return result
}

func BuildText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var result string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result += BuildText(c) + " "
	}

	temp_result_slice := strings.Fields(result) //trim all whitespace and return slice of word in strings
	result = strings.Join(temp_result_slice, " ")
	return result
}

func LinkNodes(n *html.Node) []*html.Node {
	var result []*html.Node
	if n.Type == html.ElementNode && n.Data == "a" {
		//this block is the base condition
		return append(result, n)
		// return []*html.Node{n}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result = append(result, LinkNodes(c)...)
	}
	return result

}

func dfs(n *html.Node, padding string) {
	msg := n.Data
	if n.Type == html.ElementNode {
		msg = "<" + msg + ">"
	}
	fmt.Println(padding, msg)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, padding+"  ")
	}
}
