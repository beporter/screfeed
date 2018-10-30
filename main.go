// Extract data from web pages using XPath and format into XML files.
// A Feed43 replacement you can run yourself.
package main

import (
	"fmt"
//	"log"
//	"flag" // Process command line args. http://golang.org/pkg/flag/

//	"net/http"
//	"io/ioutil"

//	"github.com/PuerkitoBio/goquery"
)


// Stores the user's extraction rules for a given webpage.
type FeedConfig struct {
	SourceUrl string
	Tokens map[string]string
}


// Defines the output format for a given feed.
// The *Format elements will have `%token%` values interpolated with the
// corresponding extracted value from FeedConfig::Tokens[token].
type FeedFormat struct {
	XmlFilePath string
	FeedUrl string
	FeedTitle string
	FeedDescription string
	ItemTitleFormat string
	ItemLinkFormat string
	ItemBodyFormat string
}


// Experimenting. Will eventually be the main entry for fetching pages,
// extracting tokens from them, and writing xml files using the
// extracted values.
func main() {
	fmt.Printf("Starting up.\n")
	config := FeedConfig{
		"http://penny-arcade.com/comic/",
		map[string]string{
			"title": "//xpath1",
			"imgurl": "//xpath2",
			"pageurl": "//xpath3",
		},
	}
	format := FeedFormat{
		"./xml-cache/pa-comic.xml",
		"http://penny-arcade.com/comic/",
		"Penny Arcade: Comic",
		"PA is in the jpeg business.",
		"%title%",
		"%pageurl%",
		"<p><a href=\"%pageurl%\"><img src=\"%imgurl%\" alt=\"%title%\" style=\"width: 100%;\"/></a></p>",
	}
	fmt.Println(config)
	fmt.Println(format)
}
