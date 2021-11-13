package scraper

import "io"

type Scraper interface {
	// The filter to check if http destination should be scraped
	Filter(string) bool
	// TODO: Change to HTML Element later
	Parse(string /* should be html element */) error
	// Gets the parsed data stream
	Stream() io.ReadCloser
}
