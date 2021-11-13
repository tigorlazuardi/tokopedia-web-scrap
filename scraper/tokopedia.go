package scraper

type TokopediaScraper struct {
}

// Gets ScrapeData from url.
// Visitor should not visit any url aside from this given url and next pagination.
func (ts TokopediaScraper) Scrap(url string) (data ScrapeDataList, err error) {
	panic("not implemented") // TODO: Implement
}

func NewTokopediaScraper(url string) *TokopediaScraper {
	return &TokopediaScraper{}
}
