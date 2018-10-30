// Extract data from web pages using XPath and format into XML files.
// A Feed43 replacement you can run yourself.
package main

import (
	"fmt"
//	"log"
//	"flag" // Process command line args. http://golang.org/pkg/flag/

//	"net/http"
    "bufio"
    "os"

//	"github.com/PuerkitoBio/goquery"
    "launchpad.net/xmlpath"
)


// Stores the user's extraction rules for a given webpage.
type FeedConfig struct {
	SourceUrl string
	DestinationPath string
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
		"https://penny-arcade.com/comic/",
		"./html-cache/comic.html",
		map[string]string{
			"title": "//div[@id=\"comic\"]//h2",
			"imgurl": "//div[@id=\"comicFrame\"]//img/@src",
			"pageurl": "//meta[@property=\"og:url\"]/@content",
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


    // Open the raw file for reading.
    rawFile, err := os.Open(config.DestinationPath)
    if err != nil {
        panic(fmt.Sprintf("Input file `%s` not read: %s\n", config.DestinationPath, err.Error()))
    }

    // Buffer the file input.
    // Parse the file as HTML.
    bufferedFile := bufio.NewReader(rawFile)
    parsedXml, err := xmlpath.ParseHTML(bufferedFile)
    if err != nil {
        panic(fmt.Sprintf("Input file `%s` not read: %s\n", config.DestinationPath, err.Error()))
    }

	fmt.Printf("Scanning HTML file: %s\n", config.DestinationPath)
    for token, xpath := range config.Tokens {
        fmt.Printf("Looking for: %s => %s\n", token, xpath)

        path := xmlpath.MustCompile(xpath)
        if value, ok := path.String(parsedXml); ok {
            fmt.Println("Found: ", value)
        }
    }
}
