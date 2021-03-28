package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/pranotobudi/Go-Gophercises/tree/main/Go-HTML-Link-Parser/link"
)

func main() {
	// var s = `
	// <html>
	// 	<body>
	// 		<h1>Hello</h1>
	// 		<a href="/other-page">link to another page</a>
	// 		<a href="/other-page2">
	// 			link to page 2
	// 			<span>
	// 				another additional note
	// 				<p> paragraph </p>
	// 			</span>
	// 		</a>
	// 	</body>
	// </html>
	// `
	// r := strings.NewReader(s)

	fileName := flag.String("filename", "ex4.html", "html file name")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}
	links, err := link.Parse(file)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v \n", links)

	// htmlTokenizer(file)
	//htmlParser(file)
	// buf, err := ioutil.ReadFile(*fileName)
	// if err != nil {
	// 	fmt.Println("file not found")
	// }

}

// func htmlParser(file io.Reader) {
// 	node, err := html.Parse(file)
// 	if err != nil {
// 		fmt.Println("can't open file")
// 	}
// 	var f func(*html.Node)
// 	f = func(n *html.Node) {
// 		fmt.Printf("data: %v \n", n.Data)
// 		fmt.Printf("node type: %v \n", n.Type)
// 		// fmt.Printf("firstChild: %v \n", n.FirstChild)
// 		// fmt.Printf("NextSibling: %v \n", n.NextSibling)
// 		if n.Type == html.ElementNode && n.Data == "a" {

// 			for _, a := range n.Attr {
// 				fmt.Printf("Attr: %v \n", a)
// 				if a.Key == "href" {
// 					fmt.Printf("url: %v \n", a.Val)
// 					break
// 				}
// 			}
// 		}
// 		for c := n.FirstChild; c != nil; c = c.NextSibling {
// 			f(c)
// 		}
// 	}
// 	f(node)

// 	// fmt.Printf("Node type: %v \n", (*node).Type)
// 	// fmt.Printf("Node Data: %v \n", (*node).Data)
// 	// fmt.Printf("Node DataAtom: %v \n", (*node).DataAtom)
// 	// fmt.Printf("Node FirstChild: %v \n", (*node).FirstChild)
// 	// fmt.Printf("Node LastChild: %v \n", (*node).LastChild)
// 	// fmt.Printf("Node Attrib: %v \n", (*node).Attr)
// }

// func htmlTokenizer(file io.Reader) {
// 	z := html.NewTokenizer(file)
// 	fmt.Println(z)
// 	for {
// 		tt := z.Next()
// 		fmt.Printf("tt: %v \n", tt)

// 		switch {
// 		case tt == html.ErrorToken:
// 			// End of the document, we're done
// 			return
// 		case tt == html.StartTagToken:
// 			t := z.Token()
// 			fmt.Printf("t: %v t.Data: %v \n", t, t.Data)
// 			isAnchor := t.Data == "a"
// 			fmt.Printf("t attrib: %v \n", t.Attr)
// 			if isAnchor {
// 				fmt.Println("We found a link!")
// 				for _, a := range t.Attr {
// 					fmt.Printf("a: %v \n", a)
// 					if a.Key == "href" {
// 						fmt.Println("Found href:", a.Val)

// 						break
// 					}
// 				}

// 			}
// 		}
// 	}

// }
