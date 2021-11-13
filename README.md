# Web Scraper Tokopedia

Scraps tokopedia mobile handphone page for the top 100 entries into csv file.

## Requirements

- golang 1.13

## Installation

### 1st Method

Easiest way is to use `go install`. But it depends on your setup. ensure `$GOPATH/bin` is in your `$PATH`, or if you set `$GOBIN` env var, it should also be in `$PATH`.

```
go install github.com/tigorlazuardi/tokopedia-web-scrap
```

Run the command using:

```
tokopedia-web-scrap
```

### 2nd Method

Clone this repo. Then run `go build`.

Run the app using:
```
./tokopedia-web-scrap
```
