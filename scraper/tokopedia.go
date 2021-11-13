package scraper

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
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
	limit      int
}

// implements Scraper interface
func (ts TokopediaProductScraper) Scrap(ctx context.Context, url string) (data ScrapeDataList, err error) {
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:83.0) Gecko/20100101 Firefox/83.0")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("tokopedia returned status code of %d", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return
	}
	res.Body.Close()

	doc.Find(`div[data-testid="lstCL2ProductList"]>div>a`).Children().Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Text())
	})

	result := make(ScrapeDataList, len(ts.collection))
	for _, v := range ts.collection {
		result = append(result, *v)
	}
	return result, ts.err
}

func (ts TokopediaProductScraper) hasError() bool {
	ts.mu.RLock()
	err := ts.err
	ts.mu.RUnlock()
	return err != nil
}

// Create Tokopedia Product Scraper
func NewTokopediaProductScraper(productListQuerySelector, descriptionQuerySelector, blackListPrefix string, timeout time.Duration, limit int) *TokopediaProductScraper {
	return &TokopediaProductScraper{
		productListQuerySelector: productListQuerySelector,
		descriptionQuerySelector: descriptionQuerySelector,
		blackListPrefix:          blackListPrefix,
		timeout:                  timeout,
		collection:               make(map[string]*ScrapeData),
		mu:                       &sync.RWMutex{},
		limit:                    limit,
	}
}
