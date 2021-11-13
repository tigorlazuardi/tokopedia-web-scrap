package cmd

import (
	"os"
	"time"

	"github.com/tigorlazuardi/tokopedia-web-scrap/scraper"
	"github.com/tigorlazuardi/tokopedia-web-scrap/service"
)

func Execute() int {
	scrap := scraper.NewTokopediaProductScraper(
		`div[data-testid="lstCL2ProductList"]`,
		`a[href]`,
		"https://ts.tokopedia.com/promo",
		10*time.Second,
		100,
	)
	filename := "tokopedia_" + time.Now().Format("2006_01_02T15_04_05") + ".csv"
	service := service.NewScrapService(scrap, filename, "https://www.tokopedia.com/p/handphone-tablet/handphone?ob=5")
	err := service.Scrap()
	if err != nil {
		return 1
	}
	return 0
}

func getArgs() []string {
	return os.Args[1:]
}
