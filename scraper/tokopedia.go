package scraper

import (
	"strconv"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
)

type TokopediaProductScraper struct {
	productListQuerySelector string
	blackListPrefix          string
	descriptionQuerySelector string
	timeout                  time.Duration
	// collection is collection for scrape data. The key is url to product description
	collection map[string]*ScrapeData
	err        error
	mu         *sync.RWMutex
}

// implements Scraper interface
func (ts TokopediaProductScraper) Scrap(url string) (data ScrapeDataList, err error) {
	c := colly.NewCollector()
	c.SetRequestTimeout(ts.timeout)
	c.OnHTML(ts.productListQuerySelector, ts.productListScraper)
	c.OnError(ts.onError)

	i := 1
	for len(ts.collection) < 100 {
		ts.mu.RLock()
		if ts.err != nil {
			ts.mu.RUnlock()
			return data, ts.err
		}
		ts.mu.RUnlock()
		url := url + "&page=" + strconv.Itoa(i)
		c.Visit(url)
	}
	result := make(ScrapeDataList, len(ts.collection))
	for _, v := range ts.collection {
		result = append(result, *v)
	}
	return result, ts.err
}

func (ts TokopediaProductScraper) productListScraper(element *colly.HTMLElement) {

}

func (ts TokopediaProductScraper) descriptionScraper(url string) colly.HTMLCallback {
	return func(element *colly.HTMLElement) {

	}
}

func (ts TokopediaProductScraper) onError(res *colly.Response, e error) {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	ts.err = e
}

// Create Tokopedia Product Scraper
func NewTokopediaProductScraper(productListQuerySelector, descriptionQuerySelector, blackListPrefix string, timeout time.Duration) *TokopediaProductScraper {
	return &TokopediaProductScraper{
		productListQuerySelector: productListQuerySelector,
		descriptionQuerySelector: descriptionQuerySelector,
		blackListPrefix:          blackListPrefix,
		timeout:                  timeout,
		collection:               make(map[string]*ScrapeData),
		mu:                       &sync.RWMutex{},
	}
}
