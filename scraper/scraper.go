package scraper

import "strconv"

type Scraper interface {
	// The filter to check if http destination should be scraped
	Filter(string) bool
	// TODO: Change to HTML Element later
	Parse(string /* should be html element */) error
}

type ScrapeData struct {
	ProductName  string `json:"product_name"`
	Description  string `json:"description"`
	ImageLink    string `json:"image_link"`
	Price        int    `json:"price"`
	Rating       int    `json:"rating"`
	MerchantName string `json:"merchant_name"`
}

func (s ScrapeData) CSVHeader() []string {
	return []string{
		"product_name", "description", "image_link",
		"price", "rating", "merchant_name",
	}
}

func (s ScrapeData) CSVRow() []string {
	return []string{
		s.ProductName, s.Description, s.ImageLink,
		strconv.Itoa(s.Price), strconv.Itoa(s.Rating), s.MerchantName,
	}
}

type ScrapeDataList []ScrapeData

func (sdl ScrapeDataList) CSVHeader() []string {
	return (ScrapeData{}).CSVHeader()
}

func (sdl ScrapeDataList) CSVRows() [][]string {
	list := make([][]string, len(sdl))
	for _, scrapeData := range sdl {
		list = append(list, scrapeData.CSVRow())
	}
	return list
}
