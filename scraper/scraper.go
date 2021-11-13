package scraper

import "strconv"

type Scraper interface {
	// Gets ScrapeData from url.
	// Visitor should not visit any url aside from this given url and next pagination.
	Scrap(url string) (data ScrapeDataList, err error)
}

type ScrapeData struct {
	ProductName  string `json:"product_name"`
	Description  string `json:"description"`
	ImageLink    string `json:"image_link"`
	Price        int    `json:"price"`
	Rating       int    `json:"rating"`
	MerchantName string `json:"merchant_name"`
}

// implements csvwriter.RowHeaderGetter
func (s ScrapeData) CSVHeader() []string {
	return []string{
		"product_name", "description", "image_link",
		"price", "rating", "merchant_name",
	}
}

// implements csvwriter.RowBodyGetter
func (s ScrapeData) CSVRow() []string {
	return []string{
		s.ProductName, s.Description, s.ImageLink,
		strconv.Itoa(s.Price), strconv.Itoa(s.Rating), s.MerchantName,
	}
}

type ScrapeDataList []ScrapeData

// implements csvwriter.RowHeaderGetter
func (sdl ScrapeDataList) CSVHeader() []string {
	return (ScrapeData{}).CSVHeader()
}

// implements csvwriter.MultiRowBodyGetter
func (sdl ScrapeDataList) CSVRows() [][]string {
	list := make([][]string, len(sdl))
	for _, scrapeData := range sdl {
		list = append(list, scrapeData.CSVRow())
	}
	return list
}
