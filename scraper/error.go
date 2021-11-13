package scraper

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/tigorlazuardi/tokopedia-web-scrap/pkg"
)

type Error struct {
	Err      error           `json:"error,omitempty"`
	Request  RequestInfo     `json:"request"`
	Response ResponseInfo    `json:"response"`
	Location pkg.CallerTrace `json:"location"`
}

func (e Error) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return "<nil>"
}

type RequestInfo struct {
	Method string      `json:"method"`
	Url    string      `json:"url"`
	Body   interface{} `json:"body"`
	Header http.Header `json:"header"`
}

type ResponseInfo struct {
	Status             int           `json:"status"`
	Body               interface{}   `json:"body"`
	Header             http.Header   `json:"header"`
	ConnectionDuration time.Duration `json:"connection_duration"`
}

func requestInfoFromCollyRequest(req *colly.Request) RequestInfo {
	var body []byte
	if req.Body != nil {
		body, _ = ioutil.ReadAll(req.Body)
	} else {
		body = []byte(``)
	}
	return RequestInfo{
		Method: req.Method,
		Url:    req.URL.String(),
		Body:   string(body),
		Header: req.Headers.Clone(),
	}
}

func responseInfoFromCollyResponse(res *colly.Response) ResponseInfo {
	ri := ResponseInfo{
		Status: res.StatusCode,
		Body:   res.Body,
		Header: res.Headers.Clone(),
	}
	if res.Trace != nil {
		ri.ConnectionDuration = res.Trace.ConnectDuration
	}
	return ri
}

func newError(err error, req RequestInfo, res ResponseInfo) Error {
	loc := pkg.GetCallerInfo(2)
	return Error{err, req, res, loc}
}
