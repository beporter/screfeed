# screfeed: a Feed43 Replacement you run yourself

Extract data from web pages using XPath and format into XML files.


## Development

### Requirements

* [Git](https://git-scm.com/downloads) v2.0+
* [Go](https://golang.org/dl/) v1.11+


### Setup

* `cd $GOPATH`
* `git clone git@github.com:beporter/screfeed.git src/github.com/beporter/screfeed`
* `cd src/github.com/beporter/screfeed`


### Compiling

* (Run once): `cd $GOPATH/src/github.com/beporter/screfeed`
* Run: `go install && screfeed`


### Tests

TODO



## Design Notes:

The goal is to be able to:

1. Scrape a webpage (or many) on a recurring basis (scheduled on your own via cron or whatever).
1. Parse the scraped pages using xpaths to capture snippets into a set of tokens.
1. Use those tokens to build an RSS xml feed file (made web accessible on your own) containing content that includes a mix of the token values and your own HTML formatting.
1. You can then subscribe to your RSS feed using the web accessible URL to the generated xml files.


### Rough Idea Outline/Notes

* Define the following configs for each "feed":
	* Source URL. The URL from which to scrape the HTML.
	* Extraction `[token => //xpath]` pairs. Each pair consists of a name to identify it, and an xpath used to extract data from the fetched HTML content.
	* Output format. (Must interpolate tokens into the provided formatting strings):
		* Local xml filename and location to write to.
		* Feed URL. (Typically the "source" URL.)
		* Feed title.
		* Feed description.
		* Item title format.
		* Item link format.
		* Item body format.
* Tool will be designed to run via cron, so no built-in scheduling is necessary.
* If fetching a page fails, tool must not replace an existing xml file. (Poor person's caching is to leave the old generated file in place.)
* Should require as few external dependencies as possible.
* If at all possible, should be distributable as a single executable binary file. (not _that_ important for myself though.)


### Architecture

* Wrapper/executor: Main entry point for the program, executed by the command line. Responsible for coordinating config loading, argument parsing, looping over requested feeds?
* Option parser: Read command line args, set internal config appropriately.
* Config object?: Provides access to runtime information, command line args, imported config file settings.
* Fetcher: Takes a source URL and a list of token:xpath pairs. Downloads the source HTML file, extracts tokens using provided rules, returns an object containing the extracted values (each of which may be an indexed array of matches if there were multiple.)
* Formatter: Takes the object containing extracted properties and the formatting rules from the config and generates the appropriate rss compatible XML document. Must provode a "toString()" method to pass to a file writer.
* File writer: Writes data to the filesystem using a configured path and provided data string.


### Questions

* How do we handle xpaths that match multiple items? Ideally, this should create a multiple `<item>` xml file, by pairing the indexed entries for each token. I.e.: The first `title` match with the first `link` match and first `description` match. Could get complicated/messy if the count of elements returned from the html document for each xpath does not match for all tokens.


### License

For the time being, this is mine: **All rights are reserved.**

Granted, it's public on Github so I can't really expect to enforce that and nor would I try, but using Github is for my own convenience.


### TODO

* Pick a language. Requirements:
	* Compilable?
	* HTTP GET support.
	* Excellent DOM processing.
	* XML building, rss2-specific would be helpful?
	* Unit testable.
* Pick a config structure.
* Pick a template engine?
* git init
* Build the main() wrapper.
* Set up test suite.
* _Build/test the rest._
* Write docs as I go.


### Language Evaluations

| Language  | Compile?  | HTTP GET?  | DOM/XPATH?  | XML/RSS lib?  | Testing?  |
|-----------|-----------|------------|-------------|---------------|-----------|
| PHP  | No (.phar)  | Yes  | DOMXpath extension  | built-in  | phpunit  |
| Go  | Yes  | net/http  | Poor, 3rdP. [xmlpath](http://godoc.org/launchpad.net/xmlpath), [goquery](https://github.com/PuerkitoBio/goquery), [gokogiri](https://github.com/moovweb/gokogiri)  | encoding/xml  | built-in  |
| Perl  | No (unnecessary)  | LWP::Simple  | XML::XPath  | XML::RSS [ref](http://stackoverflow.com/a/14617203/70876)  | Test::Simple, Test::More, [Devel:Cover](http://blogs.perl.org/users/neilb/2014/08/check-your-test-coverage-with-develcover.html)  |
| C  | Yes  | libCurl  | libXml [ref](http://www.xmlsoft.org/examples/#xpath2.c)  | VTD-XML [ref](http://vtd-xml.sourceforge.net/codeSample/cs2.html)  | Many options? [ref](http://stackoverflow.com/questions/65820/unit-testing-c-code)  |
