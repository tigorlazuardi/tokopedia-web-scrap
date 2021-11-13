package service

import (
	"context"
	"fmt"

	"github.com/tigorlazuardi/tokopedia-web-scrap/csvwriter"
	"github.com/tigorlazuardi/tokopedia-web-scrap/scraper"
)

type ScrapService struct {
	scraper  scraper.Scraper
	writer   *csvwriter.CSVWriter
	startURL string
}

func NewScrapService(s scraper.Scraper, filename string, startURL string) ScrapService {
	if s == nil {
		panic("scraper is nil")
	}
	if startURL == "" {
		panic("start url is empty")
	}
	if filename == "" {
		panic("filename must be given")
	}
	w, err := csvwriter.NewCSVWriterToNewFile(filename)
	if err != nil {
		panic(err)
	}
	return ScrapService{s, w, startURL}
}

// start html scraping and save to file
func (s ScrapService) Scrap() error {
	data, err := s.scraper.Scrap(context.Background(), s.startURL)
	if err != nil {
		// TODO: Implement logger
		fmt.Println(err)
		return err
	}

	_ = data
	// err = s.writer.WriteFull(data)
	// if err != nil {
	// 	// TODO: Implement logger
	// 	fmt.Println(err)
	// 	return err
	// }
	return err
}
